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

// JobManagerJob is an auto generated low-level Go binding around an user-defined struct.
type JobManagerJob struct {
	JobId       [32]byte
	Renter      common.Address
	Provider    common.Address
	Payment     *big.Int
	Status      uint8
	CreatedAt   *big.Int
	ConfirmedAt *big.Int
}

// JobManagerMetaData contains all meta data concerning the JobManager contract.
var JobManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"confirmedAt\",\"type\":\"uint256\"}],\"name\":\"JobConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"renter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"JobCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaymentClaimed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"}],\"name\":\"confirmResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"providerAddress\",\"type\":\"address\"}],\"name\":\"createJob\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"}],\"name\":\"getJobInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"renter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"enumJobManager.JobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"confirmedAt\",\"type\":\"uint256\"}],\"internalType\":\"structJobManager.Job\",\"name\":\"job\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getJobStatistics\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"created\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"completed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"escrowAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalEscrowAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobManager.JobStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"isJobInStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"jobs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"renter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"enumJobManager.JobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"confirmedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalEscrowAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalJobsCompleted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalJobsCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// JobManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use JobManagerMetaData.ABI instead.
var JobManagerABI = JobManagerMetaData.ABI

// JobManager is an auto generated Go binding around an Ethereum contract.
type JobManager struct {
	JobManagerCaller     // Read-only binding to the contract
	JobManagerTransactor // Write-only binding to the contract
	JobManagerFilterer   // Log filterer for contract events
}

// JobManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type JobManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JobManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JobManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JobManagerSession struct {
	Contract     *JobManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JobManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JobManagerCallerSession struct {
	Contract *JobManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// JobManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JobManagerTransactorSession struct {
	Contract     *JobManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// JobManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type JobManagerRaw struct {
	Contract *JobManager // Generic contract binding to access the raw methods on
}

// JobManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JobManagerCallerRaw struct {
	Contract *JobManagerCaller // Generic read-only contract binding to access the raw methods on
}

// JobManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JobManagerTransactorRaw struct {
	Contract *JobManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJobManager creates a new instance of JobManager, bound to a specific deployed contract.
func NewJobManager(address common.Address, backend bind.ContractBackend) (*JobManager, error) {
	contract, err := bindJobManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JobManager{JobManagerCaller: JobManagerCaller{contract: contract}, JobManagerTransactor: JobManagerTransactor{contract: contract}, JobManagerFilterer: JobManagerFilterer{contract: contract}}, nil
}

// NewJobManagerCaller creates a new read-only instance of JobManager, bound to a specific deployed contract.
func NewJobManagerCaller(address common.Address, caller bind.ContractCaller) (*JobManagerCaller, error) {
	contract, err := bindJobManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JobManagerCaller{contract: contract}, nil
}

// NewJobManagerTransactor creates a new write-only instance of JobManager, bound to a specific deployed contract.
func NewJobManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*JobManagerTransactor, error) {
	contract, err := bindJobManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JobManagerTransactor{contract: contract}, nil
}

// NewJobManagerFilterer creates a new log filterer instance of JobManager, bound to a specific deployed contract.
func NewJobManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*JobManagerFilterer, error) {
	contract, err := bindJobManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JobManagerFilterer{contract: contract}, nil
}

// bindJobManager binds a generic wrapper to an already deployed contract.
func bindJobManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := JobManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JobManager *JobManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JobManager.Contract.JobManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JobManager *JobManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobManager.Contract.JobManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JobManager *JobManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JobManager.Contract.JobManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JobManager *JobManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JobManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JobManager *JobManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JobManager *JobManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JobManager.Contract.contract.Transact(opts, method, params...)
}

// GetJobInfo is a free data retrieval call binding the contract method 0x6164626a.
//
// Solidity: function getJobInfo(bytes32 jobId) view returns((bytes32,address,address,uint256,uint8,uint256,uint256) job)
func (_JobManager *JobManagerCaller) GetJobInfo(opts *bind.CallOpts, jobId [32]byte) (JobManagerJob, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "getJobInfo", jobId)

	if err != nil {
		return *new(JobManagerJob), err
	}

	out0 := *abi.ConvertType(out[0], new(JobManagerJob)).(*JobManagerJob)

	return out0, err

}

// GetJobInfo is a free data retrieval call binding the contract method 0x6164626a.
//
// Solidity: function getJobInfo(bytes32 jobId) view returns((bytes32,address,address,uint256,uint8,uint256,uint256) job)
func (_JobManager *JobManagerSession) GetJobInfo(jobId [32]byte) (JobManagerJob, error) {
	return _JobManager.Contract.GetJobInfo(&_JobManager.CallOpts, jobId)
}

// GetJobInfo is a free data retrieval call binding the contract method 0x6164626a.
//
// Solidity: function getJobInfo(bytes32 jobId) view returns((bytes32,address,address,uint256,uint8,uint256,uint256) job)
func (_JobManager *JobManagerCallerSession) GetJobInfo(jobId [32]byte) (JobManagerJob, error) {
	return _JobManager.Contract.GetJobInfo(&_JobManager.CallOpts, jobId)
}

// GetJobStatistics is a free data retrieval call binding the contract method 0x3b71ef18.
//
// Solidity: function getJobStatistics() view returns(uint256 created, uint256 completed, uint256 escrowAmount)
func (_JobManager *JobManagerCaller) GetJobStatistics(opts *bind.CallOpts) (struct {
	Created      *big.Int
	Completed    *big.Int
	EscrowAmount *big.Int
}, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "getJobStatistics")

	outstruct := new(struct {
		Created      *big.Int
		Completed    *big.Int
		EscrowAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Created = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Completed = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EscrowAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetJobStatistics is a free data retrieval call binding the contract method 0x3b71ef18.
//
// Solidity: function getJobStatistics() view returns(uint256 created, uint256 completed, uint256 escrowAmount)
func (_JobManager *JobManagerSession) GetJobStatistics() (struct {
	Created      *big.Int
	Completed    *big.Int
	EscrowAmount *big.Int
}, error) {
	return _JobManager.Contract.GetJobStatistics(&_JobManager.CallOpts)
}

// GetJobStatistics is a free data retrieval call binding the contract method 0x3b71ef18.
//
// Solidity: function getJobStatistics() view returns(uint256 created, uint256 completed, uint256 escrowAmount)
func (_JobManager *JobManagerCallerSession) GetJobStatistics() (struct {
	Created      *big.Int
	Completed    *big.Int
	EscrowAmount *big.Int
}, error) {
	return _JobManager.Contract.GetJobStatistics(&_JobManager.CallOpts)
}

// GetTotalEscrowAmount is a free data retrieval call binding the contract method 0xb5b17c67.
//
// Solidity: function getTotalEscrowAmount() view returns(uint256)
func (_JobManager *JobManagerCaller) GetTotalEscrowAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "getTotalEscrowAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalEscrowAmount is a free data retrieval call binding the contract method 0xb5b17c67.
//
// Solidity: function getTotalEscrowAmount() view returns(uint256)
func (_JobManager *JobManagerSession) GetTotalEscrowAmount() (*big.Int, error) {
	return _JobManager.Contract.GetTotalEscrowAmount(&_JobManager.CallOpts)
}

// GetTotalEscrowAmount is a free data retrieval call binding the contract method 0xb5b17c67.
//
// Solidity: function getTotalEscrowAmount() view returns(uint256)
func (_JobManager *JobManagerCallerSession) GetTotalEscrowAmount() (*big.Int, error) {
	return _JobManager.Contract.GetTotalEscrowAmount(&_JobManager.CallOpts)
}

// IsJobInStatus is a free data retrieval call binding the contract method 0xecee149f.
//
// Solidity: function isJobInStatus(bytes32 jobId, uint8 status) view returns(bool)
func (_JobManager *JobManagerCaller) IsJobInStatus(opts *bind.CallOpts, jobId [32]byte, status uint8) (bool, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "isJobInStatus", jobId, status)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsJobInStatus is a free data retrieval call binding the contract method 0xecee149f.
//
// Solidity: function isJobInStatus(bytes32 jobId, uint8 status) view returns(bool)
func (_JobManager *JobManagerSession) IsJobInStatus(jobId [32]byte, status uint8) (bool, error) {
	return _JobManager.Contract.IsJobInStatus(&_JobManager.CallOpts, jobId, status)
}

// IsJobInStatus is a free data retrieval call binding the contract method 0xecee149f.
//
// Solidity: function isJobInStatus(bytes32 jobId, uint8 status) view returns(bool)
func (_JobManager *JobManagerCallerSession) IsJobInStatus(jobId [32]byte, status uint8) (bool, error) {
	return _JobManager.Contract.IsJobInStatus(&_JobManager.CallOpts, jobId, status)
}

// Jobs is a free data retrieval call binding the contract method 0x38ed7cfc.
//
// Solidity: function jobs(bytes32 ) view returns(bytes32 jobId, address renter, address provider, uint256 payment, uint8 status, uint256 createdAt, uint256 confirmedAt)
func (_JobManager *JobManagerCaller) Jobs(opts *bind.CallOpts, arg0 [32]byte) (struct {
	JobId       [32]byte
	Renter      common.Address
	Provider    common.Address
	Payment     *big.Int
	Status      uint8
	CreatedAt   *big.Int
	ConfirmedAt *big.Int
}, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "jobs", arg0)

	outstruct := new(struct {
		JobId       [32]byte
		Renter      common.Address
		Provider    common.Address
		Payment     *big.Int
		Status      uint8
		CreatedAt   *big.Int
		ConfirmedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.JobId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Renter = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Provider = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Payment = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.CreatedAt = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ConfirmedAt = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Jobs is a free data retrieval call binding the contract method 0x38ed7cfc.
//
// Solidity: function jobs(bytes32 ) view returns(bytes32 jobId, address renter, address provider, uint256 payment, uint8 status, uint256 createdAt, uint256 confirmedAt)
func (_JobManager *JobManagerSession) Jobs(arg0 [32]byte) (struct {
	JobId       [32]byte
	Renter      common.Address
	Provider    common.Address
	Payment     *big.Int
	Status      uint8
	CreatedAt   *big.Int
	ConfirmedAt *big.Int
}, error) {
	return _JobManager.Contract.Jobs(&_JobManager.CallOpts, arg0)
}

// Jobs is a free data retrieval call binding the contract method 0x38ed7cfc.
//
// Solidity: function jobs(bytes32 ) view returns(bytes32 jobId, address renter, address provider, uint256 payment, uint8 status, uint256 createdAt, uint256 confirmedAt)
func (_JobManager *JobManagerCallerSession) Jobs(arg0 [32]byte) (struct {
	JobId       [32]byte
	Renter      common.Address
	Provider    common.Address
	Payment     *big.Int
	Status      uint8
	CreatedAt   *big.Int
	ConfirmedAt *big.Int
}, error) {
	return _JobManager.Contract.Jobs(&_JobManager.CallOpts, arg0)
}

// TotalEscrowAmount is a free data retrieval call binding the contract method 0xae916397.
//
// Solidity: function totalEscrowAmount() view returns(uint256)
func (_JobManager *JobManagerCaller) TotalEscrowAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "totalEscrowAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalEscrowAmount is a free data retrieval call binding the contract method 0xae916397.
//
// Solidity: function totalEscrowAmount() view returns(uint256)
func (_JobManager *JobManagerSession) TotalEscrowAmount() (*big.Int, error) {
	return _JobManager.Contract.TotalEscrowAmount(&_JobManager.CallOpts)
}

// TotalEscrowAmount is a free data retrieval call binding the contract method 0xae916397.
//
// Solidity: function totalEscrowAmount() view returns(uint256)
func (_JobManager *JobManagerCallerSession) TotalEscrowAmount() (*big.Int, error) {
	return _JobManager.Contract.TotalEscrowAmount(&_JobManager.CallOpts)
}

// TotalJobsCompleted is a free data retrieval call binding the contract method 0x8ee0decf.
//
// Solidity: function totalJobsCompleted() view returns(uint256)
func (_JobManager *JobManagerCaller) TotalJobsCompleted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "totalJobsCompleted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalJobsCompleted is a free data retrieval call binding the contract method 0x8ee0decf.
//
// Solidity: function totalJobsCompleted() view returns(uint256)
func (_JobManager *JobManagerSession) TotalJobsCompleted() (*big.Int, error) {
	return _JobManager.Contract.TotalJobsCompleted(&_JobManager.CallOpts)
}

// TotalJobsCompleted is a free data retrieval call binding the contract method 0x8ee0decf.
//
// Solidity: function totalJobsCompleted() view returns(uint256)
func (_JobManager *JobManagerCallerSession) TotalJobsCompleted() (*big.Int, error) {
	return _JobManager.Contract.TotalJobsCompleted(&_JobManager.CallOpts)
}

// TotalJobsCreated is a free data retrieval call binding the contract method 0x22ef57bd.
//
// Solidity: function totalJobsCreated() view returns(uint256)
func (_JobManager *JobManagerCaller) TotalJobsCreated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "totalJobsCreated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalJobsCreated is a free data retrieval call binding the contract method 0x22ef57bd.
//
// Solidity: function totalJobsCreated() view returns(uint256)
func (_JobManager *JobManagerSession) TotalJobsCreated() (*big.Int, error) {
	return _JobManager.Contract.TotalJobsCreated(&_JobManager.CallOpts)
}

// TotalJobsCreated is a free data retrieval call binding the contract method 0x22ef57bd.
//
// Solidity: function totalJobsCreated() view returns(uint256)
func (_JobManager *JobManagerCallerSession) TotalJobsCreated() (*big.Int, error) {
	return _JobManager.Contract.TotalJobsCreated(&_JobManager.CallOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xf5414023.
//
// Solidity: function claimReward(bytes32 jobId) returns()
func (_JobManager *JobManagerTransactor) ClaimReward(opts *bind.TransactOpts, jobId [32]byte) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "claimReward", jobId)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xf5414023.
//
// Solidity: function claimReward(bytes32 jobId) returns()
func (_JobManager *JobManagerSession) ClaimReward(jobId [32]byte) (*types.Transaction, error) {
	return _JobManager.Contract.ClaimReward(&_JobManager.TransactOpts, jobId)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xf5414023.
//
// Solidity: function claimReward(bytes32 jobId) returns()
func (_JobManager *JobManagerTransactorSession) ClaimReward(jobId [32]byte) (*types.Transaction, error) {
	return _JobManager.Contract.ClaimReward(&_JobManager.TransactOpts, jobId)
}

// ConfirmResult is a paid mutator transaction binding the contract method 0xf30abb59.
//
// Solidity: function confirmResult(bytes32 jobId) returns()
func (_JobManager *JobManagerTransactor) ConfirmResult(opts *bind.TransactOpts, jobId [32]byte) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "confirmResult", jobId)
}

// ConfirmResult is a paid mutator transaction binding the contract method 0xf30abb59.
//
// Solidity: function confirmResult(bytes32 jobId) returns()
func (_JobManager *JobManagerSession) ConfirmResult(jobId [32]byte) (*types.Transaction, error) {
	return _JobManager.Contract.ConfirmResult(&_JobManager.TransactOpts, jobId)
}

// ConfirmResult is a paid mutator transaction binding the contract method 0xf30abb59.
//
// Solidity: function confirmResult(bytes32 jobId) returns()
func (_JobManager *JobManagerTransactorSession) ConfirmResult(jobId [32]byte) (*types.Transaction, error) {
	return _JobManager.Contract.ConfirmResult(&_JobManager.TransactOpts, jobId)
}

// CreateJob is a paid mutator transaction binding the contract method 0x7706ebc8.
//
// Solidity: function createJob(bytes32 jobId, address providerAddress) payable returns()
func (_JobManager *JobManagerTransactor) CreateJob(opts *bind.TransactOpts, jobId [32]byte, providerAddress common.Address) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "createJob", jobId, providerAddress)
}

// CreateJob is a paid mutator transaction binding the contract method 0x7706ebc8.
//
// Solidity: function createJob(bytes32 jobId, address providerAddress) payable returns()
func (_JobManager *JobManagerSession) CreateJob(jobId [32]byte, providerAddress common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.CreateJob(&_JobManager.TransactOpts, jobId, providerAddress)
}

// CreateJob is a paid mutator transaction binding the contract method 0x7706ebc8.
//
// Solidity: function createJob(bytes32 jobId, address providerAddress) payable returns()
func (_JobManager *JobManagerTransactorSession) CreateJob(jobId [32]byte, providerAddress common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.CreateJob(&_JobManager.TransactOpts, jobId, providerAddress)
}

// JobManagerJobConfirmedIterator is returned from FilterJobConfirmed and is used to iterate over the raw logs and unpacked data for JobConfirmed events raised by the JobManager contract.
type JobManagerJobConfirmedIterator struct {
	Event *JobManagerJobConfirmed // Event containing the contract specifics and raw log

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
func (it *JobManagerJobConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobManagerJobConfirmed)
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
		it.Event = new(JobManagerJobConfirmed)
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
func (it *JobManagerJobConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobManagerJobConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobManagerJobConfirmed represents a JobConfirmed event raised by the JobManager contract.
type JobManagerJobConfirmed struct {
	JobId       [32]byte
	ConfirmedAt *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterJobConfirmed is a free log retrieval operation binding the contract event 0x8cc5b2bb3811643c5e58483c7be26e57579a2dae8738acac65d28bfc9ef9342a.
//
// Solidity: event JobConfirmed(bytes32 indexed jobId, uint256 confirmedAt)
func (_JobManager *JobManagerFilterer) FilterJobConfirmed(opts *bind.FilterOpts, jobId [][32]byte) (*JobManagerJobConfirmedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _JobManager.contract.FilterLogs(opts, "JobConfirmed", jobIdRule)
	if err != nil {
		return nil, err
	}
	return &JobManagerJobConfirmedIterator{contract: _JobManager.contract, event: "JobConfirmed", logs: logs, sub: sub}, nil
}

// WatchJobConfirmed is a free log subscription operation binding the contract event 0x8cc5b2bb3811643c5e58483c7be26e57579a2dae8738acac65d28bfc9ef9342a.
//
// Solidity: event JobConfirmed(bytes32 indexed jobId, uint256 confirmedAt)
func (_JobManager *JobManagerFilterer) WatchJobConfirmed(opts *bind.WatchOpts, sink chan<- *JobManagerJobConfirmed, jobId [][32]byte) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _JobManager.contract.WatchLogs(opts, "JobConfirmed", jobIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobManagerJobConfirmed)
				if err := _JobManager.contract.UnpackLog(event, "JobConfirmed", log); err != nil {
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

// ParseJobConfirmed is a log parse operation binding the contract event 0x8cc5b2bb3811643c5e58483c7be26e57579a2dae8738acac65d28bfc9ef9342a.
//
// Solidity: event JobConfirmed(bytes32 indexed jobId, uint256 confirmedAt)
func (_JobManager *JobManagerFilterer) ParseJobConfirmed(log types.Log) (*JobManagerJobConfirmed, error) {
	event := new(JobManagerJobConfirmed)
	if err := _JobManager.contract.UnpackLog(event, "JobConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobManagerJobCreatedIterator is returned from FilterJobCreated and is used to iterate over the raw logs and unpacked data for JobCreated events raised by the JobManager contract.
type JobManagerJobCreatedIterator struct {
	Event *JobManagerJobCreated // Event containing the contract specifics and raw log

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
func (it *JobManagerJobCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobManagerJobCreated)
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
		it.Event = new(JobManagerJobCreated)
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
func (it *JobManagerJobCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobManagerJobCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobManagerJobCreated represents a JobCreated event raised by the JobManager contract.
type JobManagerJobCreated struct {
	JobId    [32]byte
	Renter   common.Address
	Provider common.Address
	Payment  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterJobCreated is a free log retrieval operation binding the contract event 0xd4308b56cdccee295e288de7c02585786d44894362fac012f3053072b5257c6a.
//
// Solidity: event JobCreated(bytes32 indexed jobId, address indexed renter, address indexed provider, uint256 payment)
func (_JobManager *JobManagerFilterer) FilterJobCreated(opts *bind.FilterOpts, jobId [][32]byte, renter []common.Address, provider []common.Address) (*JobManagerJobCreatedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var renterRule []interface{}
	for _, renterItem := range renter {
		renterRule = append(renterRule, renterItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _JobManager.contract.FilterLogs(opts, "JobCreated", jobIdRule, renterRule, providerRule)
	if err != nil {
		return nil, err
	}
	return &JobManagerJobCreatedIterator{contract: _JobManager.contract, event: "JobCreated", logs: logs, sub: sub}, nil
}

// WatchJobCreated is a free log subscription operation binding the contract event 0xd4308b56cdccee295e288de7c02585786d44894362fac012f3053072b5257c6a.
//
// Solidity: event JobCreated(bytes32 indexed jobId, address indexed renter, address indexed provider, uint256 payment)
func (_JobManager *JobManagerFilterer) WatchJobCreated(opts *bind.WatchOpts, sink chan<- *JobManagerJobCreated, jobId [][32]byte, renter []common.Address, provider []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var renterRule []interface{}
	for _, renterItem := range renter {
		renterRule = append(renterRule, renterItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _JobManager.contract.WatchLogs(opts, "JobCreated", jobIdRule, renterRule, providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobManagerJobCreated)
				if err := _JobManager.contract.UnpackLog(event, "JobCreated", log); err != nil {
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

// ParseJobCreated is a log parse operation binding the contract event 0xd4308b56cdccee295e288de7c02585786d44894362fac012f3053072b5257c6a.
//
// Solidity: event JobCreated(bytes32 indexed jobId, address indexed renter, address indexed provider, uint256 payment)
func (_JobManager *JobManagerFilterer) ParseJobCreated(log types.Log) (*JobManagerJobCreated, error) {
	event := new(JobManagerJobCreated)
	if err := _JobManager.contract.UnpackLog(event, "JobCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobManagerPaymentClaimedIterator is returned from FilterPaymentClaimed and is used to iterate over the raw logs and unpacked data for PaymentClaimed events raised by the JobManager contract.
type JobManagerPaymentClaimedIterator struct {
	Event *JobManagerPaymentClaimed // Event containing the contract specifics and raw log

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
func (it *JobManagerPaymentClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobManagerPaymentClaimed)
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
		it.Event = new(JobManagerPaymentClaimed)
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
func (it *JobManagerPaymentClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobManagerPaymentClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobManagerPaymentClaimed represents a PaymentClaimed event raised by the JobManager contract.
type JobManagerPaymentClaimed struct {
	JobId    [32]byte
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPaymentClaimed is a free log retrieval operation binding the contract event 0x009aa28285e2f2f93f61312a2c375d49710858b14618190f383029b03c6ec3f4.
//
// Solidity: event PaymentClaimed(bytes32 indexed jobId, address indexed provider, uint256 amount)
func (_JobManager *JobManagerFilterer) FilterPaymentClaimed(opts *bind.FilterOpts, jobId [][32]byte, provider []common.Address) (*JobManagerPaymentClaimedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _JobManager.contract.FilterLogs(opts, "PaymentClaimed", jobIdRule, providerRule)
	if err != nil {
		return nil, err
	}
	return &JobManagerPaymentClaimedIterator{contract: _JobManager.contract, event: "PaymentClaimed", logs: logs, sub: sub}, nil
}

// WatchPaymentClaimed is a free log subscription operation binding the contract event 0x009aa28285e2f2f93f61312a2c375d49710858b14618190f383029b03c6ec3f4.
//
// Solidity: event PaymentClaimed(bytes32 indexed jobId, address indexed provider, uint256 amount)
func (_JobManager *JobManagerFilterer) WatchPaymentClaimed(opts *bind.WatchOpts, sink chan<- *JobManagerPaymentClaimed, jobId [][32]byte, provider []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _JobManager.contract.WatchLogs(opts, "PaymentClaimed", jobIdRule, providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobManagerPaymentClaimed)
				if err := _JobManager.contract.UnpackLog(event, "PaymentClaimed", log); err != nil {
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

// ParsePaymentClaimed is a log parse operation binding the contract event 0x009aa28285e2f2f93f61312a2c375d49710858b14618190f383029b03c6ec3f4.
//
// Solidity: event PaymentClaimed(bytes32 indexed jobId, address indexed provider, uint256 amount)
func (_JobManager *JobManagerFilterer) ParsePaymentClaimed(log types.Log) (*JobManagerPaymentClaimed, error) {
	event := new(JobManagerPaymentClaimed)
	if err := _JobManager.contract.UnpackLog(event, "PaymentClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
