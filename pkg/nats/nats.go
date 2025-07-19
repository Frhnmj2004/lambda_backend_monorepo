package nats

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

// NATSClient wraps the NATS client for messaging
type NATSClient struct {
	conn *nats.Conn
	js   nats.JetStreamContext
}

// NewNATSConnection creates a new NATS connection
func NewNATSConnection(natsURL string) (*NATSClient, error) {
	// Connect to NATS
	conn, err := nats.Connect(natsURL,
		nats.Name("lamda-backend"),
		nats.ReconnectWait(time.Second),
		nats.MaxReconnects(-1),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			fmt.Printf("NATS disconnected: %v\n", err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			fmt.Printf("NATS reconnected to %s\n", nc.ConnectedUrl())
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	// Create JetStream context
	js, err := conn.JetStream()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to create JetStream context: %w", err)
	}

	return &NATSClient{
		conn: conn,
		js:   js,
	}, nil
}

// Publish publishes a message to a subject
func (n *NATSClient) Publish(subject string, data interface{}) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	return n.conn.Publish(subject, payload)
}

// PublishWithReply publishes a message and waits for a reply
func (n *NATSClient) PublishWithReply(subject string, data interface{}, timeout time.Duration) ([]byte, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal message: %w", err)
	}

	msg, err := n.conn.Request(subject, payload, timeout)
	if err != nil {
		return nil, fmt.Errorf("failed to publish request: %w", err)
	}

	return msg.Data, nil
}

// Subscribe subscribes to a subject
func (n *NATSClient) Subscribe(subject string, handler func([]byte)) (*nats.Subscription, error) {
	return n.conn.Subscribe(subject, func(msg *nats.Msg) {
		handler(msg.Data)
	})
}

// SubscribeWithQueue subscribes to a subject with queue group for load balancing
func (n *NATSClient) SubscribeWithQueue(subject, queue string, handler func([]byte)) (*nats.Subscription, error) {
	return n.conn.QueueSubscribe(subject, queue, func(msg *nats.Msg) {
		handler(msg.Data)
	})
}

// SubscribeWithReply subscribes to a subject and sends a reply
func (n *NATSClient) SubscribeWithReply(subject string, handler func([]byte) ([]byte, error)) (*nats.Subscription, error) {
	return n.conn.Subscribe(subject, func(msg *nats.Msg) {
		response, err := handler(msg.Data)
		if err != nil {
			// Log error but don't send error response to avoid breaking the protocol
			fmt.Printf("Error handling message: %v\n", err)
			return
		}
		msg.Respond(response)
	})
}

// PublishJetStream publishes a message to JetStream
func (n *NATSClient) PublishJetStream(subject string, data interface{}) (*nats.PubAck, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal message: %w", err)
	}

	return n.js.Publish(subject, payload)
}

// SubscribeJetStream subscribes to a JetStream subject
func (n *NATSClient) SubscribeJetStream(subject string, handler func([]byte)) (*nats.Subscription, error) {
	return n.js.Subscribe(subject, func(msg *nats.Msg) {
		handler(msg.Data)
		msg.Ack()
	})
}

// CreateStream creates a JetStream stream
func (n *NATSClient) CreateStream(name string, subjects []string) error {
	_, err := n.js.AddStream(&nats.StreamConfig{
		Name:     name,
		Subjects: subjects,
	})
	return err
}

// Close closes the NATS connection
func (n *NATSClient) Close() {
	if n.conn != nil {
		n.conn.Close()
	}
}

// IsConnected checks if the client is connected
func (n *NATSClient) IsConnected() bool {
	return n.conn != nil && n.conn.IsConnected()
}

// WaitForConnection waits for the client to be connected with timeout
func (n *NATSClient) WaitForConnection(ctx context.Context, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for NATS connection")
		default:
			if n.IsConnected() {
				return nil
			}
			time.Sleep(time.Second)
		}
	}
}
