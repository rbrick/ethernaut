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

// FallbackMetaData contains all meta data concerning the Fallback contract.
var FallbackMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"contribute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"contributions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContribution\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"d7bb99ba": "contribute()",
		"42e94c90": "contributions(address)",
		"f10fdf5c": "getContribution()",
		"8da5cb5b": "owner()",
		"3ccfd60b": "withdraw()",
	},
	Bin: "0x608060405234801561000f575f80fd5b50600180546001600160a01b031916339081179091555f908152602081905260409020683635c9adc5dea0000090556102d38061004b5f395ff3fe60806040526004361061004c575f3560e01c80633ccfd60b1461009057806342e94c90146100a65780638da5cb5b146100e4578063d7bb99ba1461011b578063f10fdf5c14610123575f80fd5b3661008c575f3411801561006d5750335f9081526020819052604090205415155b610075575f80fd5b600180546001600160a01b03191633908117909155005b5f80fd5b34801561009b575f80fd5b506100a4610143565b005b3480156100b1575f80fd5b506100d16100c036600461024b565b5f6020819052908152604090205481565b6040519081526020015b60405180910390f35b3480156100ef575f80fd5b50600154610103906001600160a01b031681565b6040516001600160a01b0390911681526020016100db565b6100a46101da565b34801561012e575f80fd5b50335f908152602081905260409020546100d1565b6001546001600160a01b031633146101a15760405162461bcd60e51b815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e6572000000000000000000604482015260640160405180910390fd5b6001546040516001600160a01b03909116904780156108fc02915f818181858888f193505050501580156101d7573d5f803e3d5ffd5b50565b66038d7ea4c6800034106101ec575f80fd5b335f908152602081905260408120805434929061020a908490610278565b90915550506001546001600160a01b03165f9081526020819052604080822054338352912054111561024957600180546001600160a01b031916331790555b565b5f6020828403121561025b575f80fd5b81356001600160a01b0381168114610271575f80fd5b9392505050565b8082018082111561029757634e487b7160e01b5f52601160045260245ffd5b9291505056fea264697066735822122070e70ad0755618280f0fecad85fd340c71736b14f481e18a9c0cd9839684d8d064736f6c63430008140033",
}

// FallbackABI is the input ABI used to generate the binding from.
// Deprecated: Use FallbackMetaData.ABI instead.
var FallbackABI = FallbackMetaData.ABI

// Deprecated: Use FallbackMetaData.Sigs instead.
// FallbackFuncSigs maps the 4-byte function signature to its string representation.
var FallbackFuncSigs = FallbackMetaData.Sigs

// FallbackBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FallbackMetaData.Bin instead.
var FallbackBin = FallbackMetaData.Bin

// DeployFallback deploys a new Ethereum contract, binding an instance of Fallback to it.
func DeployFallback(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Fallback, error) {
	parsed, err := FallbackMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FallbackBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Fallback{FallbackCaller: FallbackCaller{contract: contract}, FallbackTransactor: FallbackTransactor{contract: contract}, FallbackFilterer: FallbackFilterer{contract: contract}}, nil
}

// Fallback is an auto generated Go binding around an Ethereum contract.
type Fallback struct {
	FallbackCaller     // Read-only binding to the contract
	FallbackTransactor // Write-only binding to the contract
	FallbackFilterer   // Log filterer for contract events
}

// FallbackCaller is an auto generated read-only Go binding around an Ethereum contract.
type FallbackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FallbackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FallbackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FallbackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FallbackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FallbackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FallbackSession struct {
	Contract     *Fallback         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FallbackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FallbackCallerSession struct {
	Contract *FallbackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FallbackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FallbackTransactorSession struct {
	Contract     *FallbackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FallbackRaw is an auto generated low-level Go binding around an Ethereum contract.
type FallbackRaw struct {
	Contract *Fallback // Generic contract binding to access the raw methods on
}

// FallbackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FallbackCallerRaw struct {
	Contract *FallbackCaller // Generic read-only contract binding to access the raw methods on
}

// FallbackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FallbackTransactorRaw struct {
	Contract *FallbackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFallback creates a new instance of Fallback, bound to a specific deployed contract.
func NewFallback(address common.Address, backend bind.ContractBackend) (*Fallback, error) {
	contract, err := bindFallback(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fallback{FallbackCaller: FallbackCaller{contract: contract}, FallbackTransactor: FallbackTransactor{contract: contract}, FallbackFilterer: FallbackFilterer{contract: contract}}, nil
}

// NewFallbackCaller creates a new read-only instance of Fallback, bound to a specific deployed contract.
func NewFallbackCaller(address common.Address, caller bind.ContractCaller) (*FallbackCaller, error) {
	contract, err := bindFallback(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FallbackCaller{contract: contract}, nil
}

// NewFallbackTransactor creates a new write-only instance of Fallback, bound to a specific deployed contract.
func NewFallbackTransactor(address common.Address, transactor bind.ContractTransactor) (*FallbackTransactor, error) {
	contract, err := bindFallback(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FallbackTransactor{contract: contract}, nil
}

// NewFallbackFilterer creates a new log filterer instance of Fallback, bound to a specific deployed contract.
func NewFallbackFilterer(address common.Address, filterer bind.ContractFilterer) (*FallbackFilterer, error) {
	contract, err := bindFallback(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FallbackFilterer{contract: contract}, nil
}

// bindFallback binds a generic wrapper to an already deployed contract.
func bindFallback(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FallbackABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fallback *FallbackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fallback.Contract.FallbackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fallback *FallbackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fallback.Contract.FallbackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fallback *FallbackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fallback.Contract.FallbackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fallback *FallbackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fallback.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fallback *FallbackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fallback.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fallback *FallbackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fallback.Contract.contract.Transact(opts, method, params...)
}

// Contributions is a free data retrieval call binding the contract method 0x42e94c90.
//
// Solidity: function contributions(address ) view returns(uint256)
func (_Fallback *FallbackCaller) Contributions(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Fallback.contract.Call(opts, &out, "contributions", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Contributions is a free data retrieval call binding the contract method 0x42e94c90.
//
// Solidity: function contributions(address ) view returns(uint256)
func (_Fallback *FallbackSession) Contributions(arg0 common.Address) (*big.Int, error) {
	return _Fallback.Contract.Contributions(&_Fallback.CallOpts, arg0)
}

// Contributions is a free data retrieval call binding the contract method 0x42e94c90.
//
// Solidity: function contributions(address ) view returns(uint256)
func (_Fallback *FallbackCallerSession) Contributions(arg0 common.Address) (*big.Int, error) {
	return _Fallback.Contract.Contributions(&_Fallback.CallOpts, arg0)
}

// GetContribution is a free data retrieval call binding the contract method 0xf10fdf5c.
//
// Solidity: function getContribution() view returns(uint256)
func (_Fallback *FallbackCaller) GetContribution(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fallback.contract.Call(opts, &out, "getContribution")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContribution is a free data retrieval call binding the contract method 0xf10fdf5c.
//
// Solidity: function getContribution() view returns(uint256)
func (_Fallback *FallbackSession) GetContribution() (*big.Int, error) {
	return _Fallback.Contract.GetContribution(&_Fallback.CallOpts)
}

// GetContribution is a free data retrieval call binding the contract method 0xf10fdf5c.
//
// Solidity: function getContribution() view returns(uint256)
func (_Fallback *FallbackCallerSession) GetContribution() (*big.Int, error) {
	return _Fallback.Contract.GetContribution(&_Fallback.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fallback *FallbackCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fallback.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fallback *FallbackSession) Owner() (common.Address, error) {
	return _Fallback.Contract.Owner(&_Fallback.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fallback *FallbackCallerSession) Owner() (common.Address, error) {
	return _Fallback.Contract.Owner(&_Fallback.CallOpts)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_Fallback *FallbackTransactor) Contribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fallback.contract.Transact(opts, "contribute")
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_Fallback *FallbackSession) Contribute() (*types.Transaction, error) {
	return _Fallback.Contract.Contribute(&_Fallback.TransactOpts)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_Fallback *FallbackTransactorSession) Contribute() (*types.Transaction, error) {
	return _Fallback.Contract.Contribute(&_Fallback.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Fallback *FallbackTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fallback.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Fallback *FallbackSession) Withdraw() (*types.Transaction, error) {
	return _Fallback.Contract.Withdraw(&_Fallback.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Fallback *FallbackTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Fallback.Contract.Withdraw(&_Fallback.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Fallback *FallbackTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fallback.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Fallback *FallbackSession) Receive() (*types.Transaction, error) {
	return _Fallback.Contract.Receive(&_Fallback.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Fallback *FallbackTransactorSession) Receive() (*types.Transaction, error) {
	return _Fallback.Contract.Receive(&_Fallback.TransactOpts)
}

