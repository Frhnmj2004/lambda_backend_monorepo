package blockchain

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EVMClient wraps the Ethereum client for blockchain interactions
type EVMClient struct {
	client *ethclient.Client
	url    string
}

// NewEVMClient creates a new Ethereum client
func NewEVMClient(rpcURL string) (*EVMClient, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}

	return &EVMClient{
		client: client,
		url:    rpcURL,
	}, nil
}

// GetLatestBlockNumber returns the latest block number
func (e *EVMClient) GetLatestBlockNumber(ctx context.Context) (uint64, error) {
	return e.client.BlockNumber(ctx)
}

// GetBlockByNumber returns a block by its number
func (e *EVMClient) GetBlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	return e.client.BlockByNumber(ctx, number)
}

// GetLogs retrieves logs from the blockchain
func (e *EVMClient) GetLogs(ctx context.Context, query interface{}) ([]types.Log, error) {
	filterQuery, ok := query.(ethereum.FilterQuery)
	if !ok {
		return nil, fmt.Errorf("query must be of type ethereum.FilterQuery")
	}
	return e.client.FilterLogs(ctx, filterQuery)
}

// GetTransactionReceipt returns the receipt of a transaction
func (e *EVMClient) GetTransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return e.client.TransactionReceipt(ctx, txHash)
}

// GetBalance returns the balance of an account
func (e *EVMClient) GetBalance(ctx context.Context, address common.Address) (*big.Int, error) {
	return e.client.BalanceAt(ctx, address, nil)
}

// GetNonce returns the nonce of an account
func (e *EVMClient) GetNonce(ctx context.Context, address common.Address) (uint64, error) {
	return e.client.PendingNonceAt(ctx, address)
}

// GetGasPrice returns the current gas price
func (e *EVMClient) GetGasPrice(ctx context.Context) (*big.Int, error) {
	return e.client.SuggestGasPrice(ctx)
}

// GetChainID returns the chain ID
func (e *EVMClient) GetChainID(ctx context.Context) (*big.Int, error) {
	return e.client.ChainID(ctx)
}

// Close closes the client connection
func (e *EVMClient) Close() {
	if e.client != nil {
		e.client.Close()
	}
}

// IsConnected checks if the client is connected
func (e *EVMClient) IsConnected(ctx context.Context) bool {
	_, err := e.client.BlockNumber(ctx)
	return err == nil
}

// WaitForConnection waits for the client to be connected with timeout
func (e *EVMClient) WaitForConnection(ctx context.Context, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for blockchain connection")
		default:
			if e.IsConnected(ctx) {
				return nil
			}
			time.Sleep(time.Second)
		}
	}
}
