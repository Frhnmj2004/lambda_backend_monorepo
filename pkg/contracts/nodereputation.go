// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// NodeReputationProvider is an auto generated low-level Go binding around an user-defined struct.
type NodeReputationProvider struct {
	GpuModel     string
	Vram         *big.Int
	LastSeen     *big.Int
	IsRegistered bool
}

// NodeReputationMetaData contains all meta data concerning the NodeReputation contract.
var NodeReputationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"}],\"name\":\"JobCountIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"NodeHeartbeat\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"gpuModel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vram\",\"type\":\"uint256\"}],\"name\":\"NodeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getActiveProvidersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getProviderInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"gpuModel\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"vram\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastSeen\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"internalType\":\"structNodeReputation.Provider\",\"name\":\"providerInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"jobCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"incrementJobs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"isProviderActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providers\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"gpuModel\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"vram\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastSeen\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"gpuModel\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"vram\",\"type\":\"uint256\"}],\"name\":\"registerNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sendHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"successfulJobsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalProviders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NodeReputationABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeReputationMetaData.ABI instead.
var NodeReputationABI = NodeReputationMetaData.ABI

// NodeReputation is an auto generated Go binding around an Ethereum contract.
type NodeReputation struct {
	NodeReputationCaller     // Read-only binding to the contract
	NodeReputationTransactor // Write-only binding to the contract
	NodeReputationFilterer   // Log filterer for contract events
}

// NodeReputationCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeReputationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeReputationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeReputationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeReputationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeReputationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeReputationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeReputationSession struct {
	Contract     *NodeReputation   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeReputationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeReputationCallerSession struct {
	Contract *NodeReputationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// NodeReputationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeReputationTransactorSession struct {
	Contract     *NodeReputationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// NodeReputationRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeReputationRaw struct {
	Contract *NodeReputation // Generic contract binding to access the raw methods on
}

// NodeReputationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeReputationCallerRaw struct {
	Contract *NodeReputationCaller // Generic read-only contract binding to access the raw methods on
}

// NodeReputationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeReputationTransactorRaw struct {
	Contract *NodeReputationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeReputation creates a new instance of NodeReputation, bound to a specific deployed contract.
func NewNodeReputation(address common.Address, backend bind.ContractBackend) (*NodeReputation, error) {
	contract, err := bindNodeReputation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeReputation{NodeReputationCaller: NodeReputationCaller{contract: contract}, NodeReputationTransactor: NodeReputationTransactor{contract: contract}, NodeReputationFilterer: NodeReputationFilterer{contract: contract}}, nil
}

// NewNodeReputationCaller creates a new read-only instance of NodeReputation, bound to a specific deployed contract.
func NewNodeReputationCaller(address common.Address, caller bind.ContractCaller) (*NodeReputationCaller, error) {
	contract, err := bindNodeReputation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeReputationCaller{contract: contract}, nil
}

// NewNodeReputationTransactor creates a new write-only instance of NodeReputation, bound to a specific deployed contract.
func NewNodeReputationTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeReputationTransactor, error) {
	contract, err := bindNodeReputation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeReputationTransactor{contract: contract}, nil
}

// NewNodeReputationFilterer creates a new log filterer instance of NodeReputation, bound to a specific deployed contract.
func NewNodeReputationFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeReputationFilterer, error) {
	contract, err := bindNodeReputation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeReputationFilterer{contract: contract}, nil
}

// bindNodeReputation binds a generic wrapper to an already deployed contract.
func bindNodeReputation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NodeReputationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeReputation *NodeReputationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeReputation.Contract.NodeReputationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeReputation *NodeReputationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeReputation.Contract.NodeReputationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeReputation *NodeReputationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeReputation.Contract.NodeReputationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeReputation *NodeReputationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeReputation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeReputation *NodeReputationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeReputation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeReputation *NodeReputationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeReputation.Contract.contract.Transact(opts, method, params...)
}

// GetActiveProvidersCount is a free data retrieval call binding the contract method 0x13c3e4b3.
//
// Solidity: function getActiveProvidersCount() view returns(uint256 count)
func (_NodeReputation *NodeReputationCaller) GetActiveProvidersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeReputation.contract.Call(opts, &out, "getActiveProvidersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveProvidersCount is a free data retrieval call binding the contract method 0x13c3e4b3.
//
// Solidity: function getActiveProvidersCount() view returns(uint256 count)
func (_NodeReputation *NodeReputationSession) GetActiveProvidersCount() (*big.Int, error) {
	return _NodeReputation.Contract.GetActiveProvidersCount(&_NodeReputation.CallOpts)
}

// GetActiveProvidersCount is a free data retrieval call binding the contract method 0x13c3e4b3.
//
// Solidity: function getActiveProvidersCount() view returns(uint256 count)
func (_NodeReputation *NodeReputationCallerSession) GetActiveProvidersCount() (*big.Int, error) {
	return _NodeReputation.Contract.GetActiveProvidersCount(&_NodeReputation.CallOpts)
}

// GetProviderInfo is a free data retrieval call binding the contract method 0x7583902f.
//
// Solidity: function getProviderInfo(address provider) view returns((string,uint256,uint256,bool) providerInfo, uint256 jobCount)
func (_NodeReputation *NodeReputationCaller) GetProviderInfo(opts *bind.CallOpts, provider common.Address) (struct {
	ProviderInfo NodeReputationProvider
	JobCount     *big.Int
}, error) {
	var out []interface{}
	err := _NodeReputation.contract.Call(opts, &out, "getProviderInfo", provider)

	outstruct := new(struct {
		ProviderInfo NodeReputationProvider
		JobCount     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProviderInfo = *abi.ConvertType(out[0], new(NodeReputationProvider)).(*NodeReputationProvider)
	outstruct.JobCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetProviderInfo is a free data retrieval call binding the contract method 0x7583902f.
//
// Solidity: function getProviderInfo(address provider) view returns((string,uint256,uint256,bool) providerInfo, uint256 jobCount)
func (_NodeReputation *NodeReputationSession) GetProviderInfo(provider common.Address) (struct {
	ProviderInfo NodeReputationProvider
	JobCount     *big.Int
}, error) {
	return _NodeReputation.Contract.GetProviderInfo(&_NodeReputation.CallOpts, provider)
}

// GetProviderInfo is a free data retrieval call binding the contract method 0x7583902f.
//
// Solidity: function getProviderInfo(address provider) view returns((string,uint256,uint256,bool) providerInfo, uint256 jobCount)
func (_NodeReputation *NodeReputationCallerSession) GetProviderInfo(provider common.Address) (struct {
	ProviderInfo NodeReputationProvider
	JobCount     *big.Int
}, error) {
	return _NodeReputation.Contract.GetProviderInfo(&_NodeReputation.CallOpts, provider)
}

// IsProviderActive is a free data retrieval call binding the contract method 0x696724cb.
//
// Solidity: function isProviderActive(address provider) view returns(bool)
func (_NodeReputation *NodeReputationCaller) IsProviderActive(opts *bind.CallOpts, provider common.Address) (bool, error) {
	var out []interface{}
	err := _NodeReputation.contract.Call(opts, &out, "isProviderActive", provider)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProviderActive is a free data retrieval call binding the contract method 0x696724cb.
//
// Solidity: function isProviderActive(address provider) view returns(bool)
func (_NodeReputation *NodeReputationSession) IsProviderActive(provider common.Address) (bool, error) {
	return _NodeReputation.Contract.IsProviderActive(&_NodeReputation.CallOpts, provider)
}

// IsProviderActive is a free data retrieval call binding the contract method 0x696724cb.
//
// Solidity: function isProviderActive(address provider) view returns(bool)
func (_NodeReputation *NodeReputationCallerSession) IsProviderActive(provider common.Address) (bool, error) {
	return _NodeReputation.Contract.IsProviderActive(&_NodeReputation.CallOpts, provider)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeReputation *NodeReputationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeReputation.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeReputation *NodeReputationSession) Owner() (common.Address, error) {
	return _NodeReputation.Contract.Owner(&_NodeReputation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeReputation *NodeReputationCallerSession) Owner() (common.Address, error) {
	return _NodeReputation.Contract.Owner(&_NodeReputation.CallOpts)
}

// Providers is a free data retrieval call binding the contract method 0x0787bc27.
//
// Solidity: function providers(address ) view returns(string gpuModel, uint256 vram, uint256 lastSeen, bool isRegistered)
func (_NodeReputation *NodeReputationCaller) Providers(opts *bind.CallOpts, arg0 common.Address) (struct {
	GpuModel     string
	Vram         *big.Int
	LastSeen     *big.Int
	IsRegistered bool
}, error) {
	var out []interface{}
	err := _NodeReputation.contract.Call(opts, &out, "providers", arg0)

	outstruct := new(struct {
		GpuModel     string
		Vram         *big.Int
		LastSeen     *big.Int
		IsRegistered bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GpuModel = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Vram = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastSeen = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsRegistered = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Providers is a free data retrieval call binding the contract method 0x0787bc27.
//
// Solidity: function providers(address ) view returns(string gpuModel, uint256 vram, uint256 lastSeen, bool isRegistered)
func (_NodeReputation *NodeReputationSession) Providers(arg0 common.Address) (struct {
	GpuModel     string
	Vram         *big.Int
	LastSeen     *big.Int
	IsRegistered bool
}, error) {
	return _NodeReputation.Contract.Providers(&_NodeReputation.CallOpts, arg0)
}

// Providers is a free data retrieval call binding the contract method 0x0787bc27.
//
// Solidity: function providers(address ) view returns(string gpuModel, uint256 vram, uint256 lastSeen, bool isRegistered)
func (_NodeReputation *NodeReputationCallerSession) Providers(arg0 common.Address) (struct {
	GpuModel     string
	Vram         *big.Int
	LastSeen     *big.Int
	IsRegistered bool
}, error) {
	return _NodeReputation.Contract.Providers(&_NodeReputation.CallOpts, arg0)
}

// SuccessfulJobsCount is a free data retrieval call binding the contract method 0x3c097e3e.
//
// Solidity: function successfulJobsCount(address ) view returns(uint256)
func (_NodeReputation *NodeReputationCaller) SuccessfulJobsCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NodeReputation.contract.Call(opts, &out, "successfulJobsCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SuccessfulJobsCount is a free data retrieval call binding the contract method 0x3c097e3e.
//
// Solidity: function successfulJobsCount(address ) view returns(uint256)
func (_NodeReputation *NodeReputationSession) SuccessfulJobsCount(arg0 common.Address) (*big.Int, error) {
	return _NodeReputation.Contract.SuccessfulJobsCount(&_NodeReputation.CallOpts, arg0)
}

// SuccessfulJobsCount is a free data retrieval call binding the contract method 0x3c097e3e.
//
// Solidity: function successfulJobsCount(address ) view returns(uint256)
func (_NodeReputation *NodeReputationCallerSession) SuccessfulJobsCount(arg0 common.Address) (*big.Int, error) {
	return _NodeReputation.Contract.SuccessfulJobsCount(&_NodeReputation.CallOpts, arg0)
}

// TotalProviders is a free data retrieval call binding the contract method 0x96616b1f.
//
// Solidity: function totalProviders() view returns(uint256)
func (_NodeReputation *NodeReputationCaller) TotalProviders(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeReputation.contract.Call(opts, &out, "totalProviders")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalProviders is a free data retrieval call binding the contract method 0x96616b1f.
//
// Solidity: function totalProviders() view returns(uint256)
func (_NodeReputation *NodeReputationSession) TotalProviders() (*big.Int, error) {
	return _NodeReputation.Contract.TotalProviders(&_NodeReputation.CallOpts)
}

// TotalProviders is a free data retrieval call binding the contract method 0x96616b1f.
//
// Solidity: function totalProviders() view returns(uint256)
func (_NodeReputation *NodeReputationCallerSession) TotalProviders() (*big.Int, error) {
	return _NodeReputation.Contract.TotalProviders(&_NodeReputation.CallOpts)
}

// IncrementJobs is a paid mutator transaction binding the contract method 0x7ebcd727.
//
// Solidity: function incrementJobs(address provider) returns()
func (_NodeReputation *NodeReputationTransactor) IncrementJobs(opts *bind.TransactOpts, provider common.Address) (*types.Transaction, error) {
	return _NodeReputation.contract.Transact(opts, "incrementJobs", provider)
}

// IncrementJobs is a paid mutator transaction binding the contract method 0x7ebcd727.
//
// Solidity: function incrementJobs(address provider) returns()
func (_NodeReputation *NodeReputationSession) IncrementJobs(provider common.Address) (*types.Transaction, error) {
	return _NodeReputation.Contract.IncrementJobs(&_NodeReputation.TransactOpts, provider)
}

// IncrementJobs is a paid mutator transaction binding the contract method 0x7ebcd727.
//
// Solidity: function incrementJobs(address provider) returns()
func (_NodeReputation *NodeReputationTransactorSession) IncrementJobs(provider common.Address) (*types.Transaction, error) {
	return _NodeReputation.Contract.IncrementJobs(&_NodeReputation.TransactOpts, provider)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xf3ef747f.
//
// Solidity: function registerNode(string gpuModel, uint256 vram) returns()
func (_NodeReputation *NodeReputationTransactor) RegisterNode(opts *bind.TransactOpts, gpuModel string, vram *big.Int) (*types.Transaction, error) {
	return _NodeReputation.contract.Transact(opts, "registerNode", gpuModel, vram)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xf3ef747f.
//
// Solidity: function registerNode(string gpuModel, uint256 vram) returns()
func (_NodeReputation *NodeReputationSession) RegisterNode(gpuModel string, vram *big.Int) (*types.Transaction, error) {
	return _NodeReputation.Contract.RegisterNode(&_NodeReputation.TransactOpts, gpuModel, vram)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xf3ef747f.
//
// Solidity: function registerNode(string gpuModel, uint256 vram) returns()
func (_NodeReputation *NodeReputationTransactorSession) RegisterNode(gpuModel string, vram *big.Int) (*types.Transaction, error) {
	return _NodeReputation.Contract.RegisterNode(&_NodeReputation.TransactOpts, gpuModel, vram)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NodeReputation *NodeReputationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeReputation.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NodeReputation *NodeReputationSession) RenounceOwnership() (*types.Transaction, error) {
	return _NodeReputation.Contract.RenounceOwnership(&_NodeReputation.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NodeReputation *NodeReputationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NodeReputation.Contract.RenounceOwnership(&_NodeReputation.TransactOpts)
}

// SendHeartbeat is a paid mutator transaction binding the contract method 0x34f167a9.
//
// Solidity: function sendHeartbeat() returns()
func (_NodeReputation *NodeReputationTransactor) SendHeartbeat(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeReputation.contract.Transact(opts, "sendHeartbeat")
}

// SendHeartbeat is a paid mutator transaction binding the contract method 0x34f167a9.
//
// Solidity: function sendHeartbeat() returns()
func (_NodeReputation *NodeReputationSession) SendHeartbeat() (*types.Transaction, error) {
	return _NodeReputation.Contract.SendHeartbeat(&_NodeReputation.TransactOpts)
}

// SendHeartbeat is a paid mutator transaction binding the contract method 0x34f167a9.
//
// Solidity: function sendHeartbeat() returns()
func (_NodeReputation *NodeReputationTransactorSession) SendHeartbeat() (*types.Transaction, error) {
	return _NodeReputation.Contract.SendHeartbeat(&_NodeReputation.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeReputation *NodeReputationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NodeReputation.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeReputation *NodeReputationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NodeReputation.Contract.TransferOwnership(&_NodeReputation.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeReputation *NodeReputationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NodeReputation.Contract.TransferOwnership(&_NodeReputation.TransactOpts, newOwner)
}

// NodeReputationJobCountIncrementedIterator is returned from FilterJobCountIncremented and is used to iterate over the raw logs and unpacked data for JobCountIncremented events raised by the NodeReputation contract.
type NodeReputationJobCountIncrementedIterator struct {
	Event *NodeReputationJobCountIncremented // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodeReputationJobCountIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeReputationJobCountIncremented)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodeReputationJobCountIncremented)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodeReputationJobCountIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeReputationJobCountIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeReputationJobCountIncremented represents a JobCountIncremented event raised by the NodeReputation contract.
type NodeReputationJobCountIncremented struct {
	Provider common.Address
	NewCount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterJobCountIncremented is a free log retrieval operation binding the contract event 0xf0b433b19371a783f4bfbaf7c118079db6c086603cd0c6ba2b6478c955b8b963.
//
// Solidity: event JobCountIncremented(address indexed provider, uint256 newCount)
func (_NodeReputation *NodeReputationFilterer) FilterJobCountIncremented(opts *bind.FilterOpts, provider []common.Address) (*NodeReputationJobCountIncrementedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _NodeReputation.contract.FilterLogs(opts, "JobCountIncremented", providerRule)
	if err != nil {
		return nil, err
	}
	return &NodeReputationJobCountIncrementedIterator{contract: _NodeReputation.contract, event: "JobCountIncremented", logs: logs, sub: sub}, nil
}

// WatchJobCountIncremented is a free log subscription operation binding the contract event 0xf0b433b19371a783f4bfbaf7c118079db6c086603cd0c6ba2b6478c955b8b963.
//
// Solidity: event JobCountIncremented(address indexed provider, uint256 newCount)
func (_NodeReputation *NodeReputationFilterer) WatchJobCountIncremented(opts *bind.WatchOpts, sink chan<- *NodeReputationJobCountIncremented, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _NodeReputation.contract.WatchLogs(opts, "JobCountIncremented", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeReputationJobCountIncremented)
				if err := _NodeReputation.contract.UnpackLog(event, "JobCountIncremented", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseJobCountIncremented is a log parse operation binding the contract event 0xf0b433b19371a783f4bfbaf7c118079db6c086603cd0c6ba2b6478c955b8b963.
//
// Solidity: event JobCountIncremented(address indexed provider, uint256 newCount)
func (_NodeReputation *NodeReputationFilterer) ParseJobCountIncremented(log types.Log) (*NodeReputationJobCountIncremented, error) {
	event := new(NodeReputationJobCountIncremented)
	if err := _NodeReputation.contract.UnpackLog(event, "JobCountIncremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeReputationNodeHeartbeatIterator is returned from FilterNodeHeartbeat and is used to iterate over the raw logs and unpacked data for NodeHeartbeat events raised by the NodeReputation contract.
type NodeReputationNodeHeartbeatIterator struct {
	Event *NodeReputationNodeHeartbeat // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodeReputationNodeHeartbeatIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeReputationNodeHeartbeat)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodeReputationNodeHeartbeat)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodeReputationNodeHeartbeatIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeReputationNodeHeartbeatIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeReputationNodeHeartbeat represents a NodeHeartbeat event raised by the NodeReputation contract.
type NodeReputationNodeHeartbeat struct {
	Provider  common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNodeHeartbeat is a free log retrieval operation binding the contract event 0x0f1ef8be219b2bb5e8b630cb970a8ac2b44c0f8a3a3db371d8a653e50af49a11.
//
// Solidity: event NodeHeartbeat(address indexed provider, uint256 timestamp)
func (_NodeReputation *NodeReputationFilterer) FilterNodeHeartbeat(opts *bind.FilterOpts, provider []common.Address) (*NodeReputationNodeHeartbeatIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _NodeReputation.contract.FilterLogs(opts, "NodeHeartbeat", providerRule)
	if err != nil {
		return nil, err
	}
	return &NodeReputationNodeHeartbeatIterator{contract: _NodeReputation.contract, event: "NodeHeartbeat", logs: logs, sub: sub}, nil
}

// WatchNodeHeartbeat is a free log subscription operation binding the contract event 0x0f1ef8be219b2bb5e8b630cb970a8ac2b44c0f8a3a3db371d8a653e50af49a11.
//
// Solidity: event NodeHeartbeat(address indexed provider, uint256 timestamp)
func (_NodeReputation *NodeReputationFilterer) WatchNodeHeartbeat(opts *bind.WatchOpts, sink chan<- *NodeReputationNodeHeartbeat, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _NodeReputation.contract.WatchLogs(opts, "NodeHeartbeat", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeReputationNodeHeartbeat)
				if err := _NodeReputation.contract.UnpackLog(event, "NodeHeartbeat", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodeHeartbeat is a log parse operation binding the contract event 0x0f1ef8be219b2bb5e8b630cb970a8ac2b44c0f8a3a3db371d8a653e50af49a11.
//
// Solidity: event NodeHeartbeat(address indexed provider, uint256 timestamp)
func (_NodeReputation *NodeReputationFilterer) ParseNodeHeartbeat(log types.Log) (*NodeReputationNodeHeartbeat, error) {
	event := new(NodeReputationNodeHeartbeat)
	if err := _NodeReputation.contract.UnpackLog(event, "NodeHeartbeat", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeReputationNodeRegisteredIterator is returned from FilterNodeRegistered and is used to iterate over the raw logs and unpacked data for NodeRegistered events raised by the NodeReputation contract.
type NodeReputationNodeRegisteredIterator struct {
	Event *NodeReputationNodeRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodeReputationNodeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeReputationNodeRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodeReputationNodeRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodeReputationNodeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeReputationNodeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeReputationNodeRegistered represents a NodeRegistered event raised by the NodeReputation contract.
type NodeReputationNodeRegistered struct {
	Provider common.Address
	GpuModel string
	Vram     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNodeRegistered is a free log retrieval operation binding the contract event 0x6f3c990a2f665f94ed1db0af40275bedd4d3d025f541e2c77112281fbee63600.
//
// Solidity: event NodeRegistered(address indexed provider, string gpuModel, uint256 vram)
func (_NodeReputation *NodeReputationFilterer) FilterNodeRegistered(opts *bind.FilterOpts, provider []common.Address) (*NodeReputationNodeRegisteredIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _NodeReputation.contract.FilterLogs(opts, "NodeRegistered", providerRule)
	if err != nil {
		return nil, err
	}
	return &NodeReputationNodeRegisteredIterator{contract: _NodeReputation.contract, event: "NodeRegistered", logs: logs, sub: sub}, nil
}

// WatchNodeRegistered is a free log subscription operation binding the contract event 0x6f3c990a2f665f94ed1db0af40275bedd4d3d025f541e2c77112281fbee63600.
//
// Solidity: event NodeRegistered(address indexed provider, string gpuModel, uint256 vram)
func (_NodeReputation *NodeReputationFilterer) WatchNodeRegistered(opts *bind.WatchOpts, sink chan<- *NodeReputationNodeRegistered, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _NodeReputation.contract.WatchLogs(opts, "NodeRegistered", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeReputationNodeRegistered)
				if err := _NodeReputation.contract.UnpackLog(event, "NodeRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodeRegistered is a log parse operation binding the contract event 0x6f3c990a2f665f94ed1db0af40275bedd4d3d025f541e2c77112281fbee63600.
//
// Solidity: event NodeRegistered(address indexed provider, string gpuModel, uint256 vram)
func (_NodeReputation *NodeReputationFilterer) ParseNodeRegistered(log types.Log) (*NodeReputationNodeRegistered, error) {
	event := new(NodeReputationNodeRegistered)
	if err := _NodeReputation.contract.UnpackLog(event, "NodeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeReputationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NodeReputation contract.
type NodeReputationOwnershipTransferredIterator struct {
	Event *NodeReputationOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodeReputationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeReputationOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodeReputationOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodeReputationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeReputationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeReputationOwnershipTransferred represents a OwnershipTransferred event raised by the NodeReputation contract.
type NodeReputationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NodeReputation *NodeReputationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NodeReputationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NodeReputation.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NodeReputationOwnershipTransferredIterator{contract: _NodeReputation.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NodeReputation *NodeReputationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NodeReputationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NodeReputation.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeReputationOwnershipTransferred)
				if err := _NodeReputation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NodeReputation *NodeReputationFilterer) ParseOwnershipTransferred(log types.Log) (*NodeReputationOwnershipTransferred, error) {
	event := new(NodeReputationOwnershipTransferred)
	if err := _NodeReputation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
