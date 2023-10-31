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
)

// FalloutMetaData contains all meta data concerning the Fallout contract.
var FalloutMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"Fal1out\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"allocator\",\"type\":\"address\"}],\"name\":\"allocatorBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectAllocations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"allocator\",\"type\":\"address\"}],\"name\":\"sendAllocation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FalloutABI is the input ABI used to generate the binding from.
// Deprecated: Use FalloutMetaData.ABI instead.
var FalloutABI = FalloutMetaData.ABI

// Fallout is an auto generated Go binding around an Ethereum contract.
type Fallout struct {
	FalloutCaller     // Read-only binding to the contract
	FalloutTransactor // Write-only binding to the contract
	FalloutFilterer   // Log filterer for contract events
}

// FalloutCaller is an auto generated read-only Go binding around an Ethereum contract.
type FalloutCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FalloutTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FalloutTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FalloutFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FalloutFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FalloutSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FalloutSession struct {
	Contract     *Fallout          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FalloutCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FalloutCallerSession struct {
	Contract *FalloutCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FalloutTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FalloutTransactorSession struct {
	Contract     *FalloutTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FalloutRaw is an auto generated low-level Go binding around an Ethereum contract.
type FalloutRaw struct {
	Contract *Fallout // Generic contract binding to access the raw methods on
}

// FalloutCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FalloutCallerRaw struct {
	Contract *FalloutCaller // Generic read-only contract binding to access the raw methods on
}

// FalloutTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FalloutTransactorRaw struct {
	Contract *FalloutTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFallout creates a new instance of Fallout, bound to a specific deployed contract.
func NewFallout(address common.Address, backend bind.ContractBackend) (*Fallout, error) {
	contract, err := bindFallout(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fallout{FalloutCaller: FalloutCaller{contract: contract}, FalloutTransactor: FalloutTransactor{contract: contract}, FalloutFilterer: FalloutFilterer{contract: contract}}, nil
}

// NewFalloutCaller creates a new read-only instance of Fallout, bound to a specific deployed contract.
func NewFalloutCaller(address common.Address, caller bind.ContractCaller) (*FalloutCaller, error) {
	contract, err := bindFallout(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FalloutCaller{contract: contract}, nil
}

// NewFalloutTransactor creates a new write-only instance of Fallout, bound to a specific deployed contract.
func NewFalloutTransactor(address common.Address, transactor bind.ContractTransactor) (*FalloutTransactor, error) {
	contract, err := bindFallout(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FalloutTransactor{contract: contract}, nil
}

// NewFalloutFilterer creates a new log filterer instance of Fallout, bound to a specific deployed contract.
func NewFalloutFilterer(address common.Address, filterer bind.ContractFilterer) (*FalloutFilterer, error) {
	contract, err := bindFallout(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FalloutFilterer{contract: contract}, nil
}

// bindFallout binds a generic wrapper to an already deployed contract.
func bindFallout(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FalloutABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fallout *FalloutRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fallout.Contract.FalloutCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fallout *FalloutRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fallout.Contract.FalloutTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fallout *FalloutRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fallout.Contract.FalloutTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fallout *FalloutCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fallout.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fallout *FalloutTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fallout.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fallout *FalloutTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fallout.Contract.contract.Transact(opts, method, params...)
}

// AllocatorBalance is a free data retrieval call binding the contract method 0xffd40b56.
//
// Solidity: function allocatorBalance(address allocator) view returns(uint256)
func (_Fallout *FalloutCaller) AllocatorBalance(opts *bind.CallOpts, allocator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Fallout.contract.Call(opts, &out, "allocatorBalance", allocator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllocatorBalance is a free data retrieval call binding the contract method 0xffd40b56.
//
// Solidity: function allocatorBalance(address allocator) view returns(uint256)
func (_Fallout *FalloutSession) AllocatorBalance(allocator common.Address) (*big.Int, error) {
	return _Fallout.Contract.AllocatorBalance(&_Fallout.CallOpts, allocator)
}

// AllocatorBalance is a free data retrieval call binding the contract method 0xffd40b56.
//
// Solidity: function allocatorBalance(address allocator) view returns(uint256)
func (_Fallout *FalloutCallerSession) AllocatorBalance(allocator common.Address) (*big.Int, error) {
	return _Fallout.Contract.AllocatorBalance(&_Fallout.CallOpts, allocator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fallout *FalloutCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fallout.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fallout *FalloutSession) Owner() (common.Address, error) {
	return _Fallout.Contract.Owner(&_Fallout.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fallout *FalloutCallerSession) Owner() (common.Address, error) {
	return _Fallout.Contract.Owner(&_Fallout.CallOpts)
}

// Fal1out is a paid mutator transaction binding the contract method 0x6fab5ddf.
//
// Solidity: function Fal1out() payable returns()
func (_Fallout *FalloutTransactor) Fal1out(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fallout.contract.Transact(opts, "Fal1out")
}

// Fal1out is a paid mutator transaction binding the contract method 0x6fab5ddf.
//
// Solidity: function Fal1out() payable returns()
func (_Fallout *FalloutSession) Fal1out() (*types.Transaction, error) {
	return _Fallout.Contract.Fal1out(&_Fallout.TransactOpts)
}

// Fal1out is a paid mutator transaction binding the contract method 0x6fab5ddf.
//
// Solidity: function Fal1out() payable returns()
func (_Fallout *FalloutTransactorSession) Fal1out() (*types.Transaction, error) {
	return _Fallout.Contract.Fal1out(&_Fallout.TransactOpts)
}

// Allocate is a paid mutator transaction binding the contract method 0xabaa9916.
//
// Solidity: function allocate() payable returns()
func (_Fallout *FalloutTransactor) Allocate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fallout.contract.Transact(opts, "allocate")
}

// Allocate is a paid mutator transaction binding the contract method 0xabaa9916.
//
// Solidity: function allocate() payable returns()
func (_Fallout *FalloutSession) Allocate() (*types.Transaction, error) {
	return _Fallout.Contract.Allocate(&_Fallout.TransactOpts)
}

// Allocate is a paid mutator transaction binding the contract method 0xabaa9916.
//
// Solidity: function allocate() payable returns()
func (_Fallout *FalloutTransactorSession) Allocate() (*types.Transaction, error) {
	return _Fallout.Contract.Allocate(&_Fallout.TransactOpts)
}

// CollectAllocations is a paid mutator transaction binding the contract method 0x8aa96f38.
//
// Solidity: function collectAllocations() returns()
func (_Fallout *FalloutTransactor) CollectAllocations(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fallout.contract.Transact(opts, "collectAllocations")
}

// CollectAllocations is a paid mutator transaction binding the contract method 0x8aa96f38.
//
// Solidity: function collectAllocations() returns()
func (_Fallout *FalloutSession) CollectAllocations() (*types.Transaction, error) {
	return _Fallout.Contract.CollectAllocations(&_Fallout.TransactOpts)
}

// CollectAllocations is a paid mutator transaction binding the contract method 0x8aa96f38.
//
// Solidity: function collectAllocations() returns()
func (_Fallout *FalloutTransactorSession) CollectAllocations() (*types.Transaction, error) {
	return _Fallout.Contract.CollectAllocations(&_Fallout.TransactOpts)
}

// SendAllocation is a paid mutator transaction binding the contract method 0xa2dea26f.
//
// Solidity: function sendAllocation(address allocator) returns()
func (_Fallout *FalloutTransactor) SendAllocation(opts *bind.TransactOpts, allocator common.Address) (*types.Transaction, error) {
	return _Fallout.contract.Transact(opts, "sendAllocation", allocator)
}

// SendAllocation is a paid mutator transaction binding the contract method 0xa2dea26f.
//
// Solidity: function sendAllocation(address allocator) returns()
func (_Fallout *FalloutSession) SendAllocation(allocator common.Address) (*types.Transaction, error) {
	return _Fallout.Contract.SendAllocation(&_Fallout.TransactOpts, allocator)
}

// SendAllocation is a paid mutator transaction binding the contract method 0xa2dea26f.
//
// Solidity: function sendAllocation(address allocator) returns()
func (_Fallout *FalloutTransactorSession) SendAllocation(allocator common.Address) (*types.Transaction, error) {
	return _Fallout.Contract.SendAllocation(&_Fallout.TransactOpts, allocator)
}

