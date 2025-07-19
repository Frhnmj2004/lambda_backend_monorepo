package auth

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// SIWEMessage represents a Sign-In With Ethereum message
type SIWEMessage struct {
	Domain         string `json:"domain"`
	Address        string `json:"address"`
	Statement      string `json:"statement"`
	URI            string `json:"uri"`
	Version        string `json:"version"`
	ChainID        int    `json:"chainId"`
	Nonce          string `json:"nonce"`
	IssuedAt       string `json:"issuedAt"`
	ExpirationTime string `json:"expirationTime,omitempty"`
	NotBefore      string `json:"notBefore,omitempty"`
	RequestID      string `json:"requestId,omitempty"`
	Resources      string `json:"resources,omitempty"`
}

// SIWESignature represents a SIWE signature
type SIWESignature struct {
	Message   SIWEMessage `json:"message"`
	Signature string      `json:"signature"`
}

// GenerateNonce generates a random nonce for SIWE
func GenerateNonce() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate nonce: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}

// CreateSIWEMessage creates a SIWE message for authentication
func CreateSIWEMessage(address, domain, uri string, chainID int) (*SIWEMessage, error) {
	nonce, err := GenerateNonce()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	issuedAt := now.Format("2006-01-02T15:04:05.000Z")
	expirationTime := now.Add(24 * time.Hour).Format("2006-01-02T15:04:05.000Z")

	return &SIWEMessage{
		Domain:         domain,
		Address:        address,
		Statement:      "Sign in to Lamda Network",
		URI:            uri,
		Version:        "1",
		ChainID:        chainID,
		Nonce:          nonce,
		IssuedAt:       issuedAt,
		ExpirationTime: expirationTime,
	}, nil
}

// FormatSIWEMessage formats a SIWE message for signing
func FormatSIWEMessage(message SIWEMessage) string {
	parts := []string{
		message.Domain + " wants you to sign in with your Ethereum account:",
		message.Address,
		"",
		message.Statement,
		"",
		"URI: " + message.URI,
		"Version: " + message.Version,
		"Chain ID: " + fmt.Sprintf("%d", message.ChainID),
		"Nonce: " + message.Nonce,
		"Issued At: " + message.IssuedAt,
	}

	if message.ExpirationTime != "" {
		parts = append(parts, "Expiration Time: "+message.ExpirationTime)
	}

	if message.NotBefore != "" {
		parts = append(parts, "Not Before: "+message.NotBefore)
	}

	if message.RequestID != "" {
		parts = append(parts, "Request ID: "+message.RequestID)
	}

	if message.Resources != "" {
		parts = append(parts, "Resources:")
		parts = append(parts, message.Resources)
	}

	return strings.Join(parts, "\n")
}

// VerifySIWESignature verifies a SIWE signature
func VerifySIWESignature(signature SIWESignature) (bool, error) {
	// Format the message
	formattedMessage := FormatSIWEMessage(signature.Message)

	// Hash the message
	messageHash := crypto.Keccak256Hash([]byte(formattedMessage))

	// Decode the signature
	sigBytes, err := hex.DecodeString(strings.TrimPrefix(signature.Signature, "0x"))
	if err != nil {
		return false, fmt.Errorf("failed to decode signature: %w", err)
	}

	// Recover the public key
	pubKey, err := crypto.SigToPub(messageHash.Bytes(), sigBytes)
	if err != nil {
		return false, fmt.Errorf("failed to recover public key: %w", err)
	}

	// Get the address from the public key
	recoveredAddress := crypto.PubkeyToAddress(*pubKey)

	// Compare with the claimed address
	claimedAddress := common.HexToAddress(signature.Message.Address)

	if recoveredAddress != claimedAddress {
		return false, fmt.Errorf("signature verification failed: address mismatch")
	}

	// Check if the message has expired
	expirationTime, err := time.Parse("2006-01-02T15:04:05.000Z", signature.Message.ExpirationTime)
	if err != nil {
		return false, fmt.Errorf("failed to parse expiration time: %w", err)
	}

	if time.Now().After(expirationTime) {
		return false, fmt.Errorf("message has expired")
	}

	return true, nil
}

// SignSIWEMessage signs a SIWE message with a private key
func SignSIWEMessage(message SIWEMessage, privateKey *ecdsa.PrivateKey) (string, error) {
	// Format the message
	formattedMessage := FormatSIWEMessage(message)

	// Hash the message
	messageHash := crypto.Keccak256Hash([]byte(formattedMessage))

	// Sign the hash
	signature, err := crypto.Sign(messageHash.Bytes(), privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign message: %w", err)
	}

	// Convert to hex
	return hex.EncodeToString(signature), nil
}

// ParseSIWEMessage parses a SIWE message from a formatted string
func ParseSIWEMessage(formattedMessage string) (*SIWEMessage, error) {
	lines := strings.Split(formattedMessage, "\n")
	if len(lines) < 10 {
		return nil, fmt.Errorf("invalid SIWE message format")
	}

	message := &SIWEMessage{}

	// Parse domain
	if strings.HasSuffix(lines[0], " wants you to sign in with your Ethereum account:") {
		message.Domain = strings.TrimSuffix(lines[0], " wants you to sign in with your Ethereum account:")
	}

	// Parse address
	message.Address = lines[1]

	// Parse statement (skip empty lines)
	statementIndex := 3
	for i := 2; i < len(lines); i++ {
		if lines[i] == "" {
			statementIndex = i + 1
			break
		}
	}
	if statementIndex < len(lines) {
		message.Statement = lines[statementIndex]
	}

	// Parse other fields
	for i := statementIndex + 2; i < len(lines); i++ {
		line := lines[i]
		if strings.HasPrefix(line, "URI: ") {
			message.URI = strings.TrimPrefix(line, "URI: ")
		} else if strings.HasPrefix(line, "Version: ") {
			message.Version = strings.TrimPrefix(line, "Version: ")
		} else if strings.HasPrefix(line, "Chain ID: ") {
			chainIDStr := strings.TrimPrefix(line, "Chain ID: ")
			if _, err := fmt.Sscanf(chainIDStr, "%d", &message.ChainID); err != nil {
				return nil, fmt.Errorf("failed to parse chain ID: %w", err)
			}
		} else if strings.HasPrefix(line, "Nonce: ") {
			message.Nonce = strings.TrimPrefix(line, "Nonce: ")
		} else if strings.HasPrefix(line, "Issued At: ") {
			message.IssuedAt = strings.TrimPrefix(line, "Issued At: ")
		} else if strings.HasPrefix(line, "Expiration Time: ") {
			message.ExpirationTime = strings.TrimPrefix(line, "Expiration Time: ")
		} else if strings.HasPrefix(line, "Not Before: ") {
			message.NotBefore = strings.TrimPrefix(line, "Not Before: ")
		} else if strings.HasPrefix(line, "Request ID: ") {
			message.RequestID = strings.TrimPrefix(line, "Request ID: ")
		}
	}

	return message, nil
}
