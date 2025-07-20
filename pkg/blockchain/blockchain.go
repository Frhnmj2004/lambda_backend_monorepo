package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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

// CreateTransactOpts creates transaction options for contract interactions
func (e *EVMClient) CreateTransactOpts(ctx context.Context, privateKeyHex string) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to get public key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	chainID, err := e.GetChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	// Get the current nonce
	nonce, err := e.GetNonce(ctx, fromAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))

	// Get gas price
	gasPrice, err := e.GetGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %w", err)
	}
	auth.GasPrice = gasPrice

	return auth, nil
}

// WaitForTransaction waits for a transaction to be mined
func (e *EVMClient) WaitForTransaction(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	return bind.WaitMined(ctx, e.client, tx)
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
	// Create a new context with timeout
	timeoutCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Try to get the latest block number as a connection test
	_, err := e.client.BlockNumber(timeoutCtx)
	if err != nil {
		return fmt.Errorf("failed to connect to blockchain: %w", err)
	}

	return nil
}

// GetClient returns the underlying ethclient for contract initialization
func (e *EVMClient) GetClient() *ethclient.Client {
	return e.client
}
