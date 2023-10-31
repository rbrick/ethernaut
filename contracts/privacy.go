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

// PrivacyMetaData contains all meta data concerning the Privacy contract.
var PrivacyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"_data\",\"type\":\"bytes32[3]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"_key\",\"type\":\"bytes16\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b3cea217": "ID()",
		"cf309012": "locked()",
		"e1afb08c": "unlock(bytes16)",
	},
	Bin: "0x60806040525f805460ff1916600190811790915542908190556002805461ffff909216620100000263ffffffff199092169190911761ff0a179055348015610045575f80fd5b50604051610266380380610266833981016040819052610064916100c9565b61007060038281610077565b505061014e565b82600381019282156100a5579160200282015b828111156100a557825182559160200191906001019061008a565b506100b19291506100b5565b5090565b5b808211156100b1575f81556001016100b6565b5f606082840312156100d9575f80fd5b82601f8301126100e7575f80fd5b604051606081016001600160401b038111828210171561011557634e487b7160e01b5f52604160045260245ffd5b604052806060840185811115610129575f80fd5b845b8181101561014357805183526020928301920161012b565b509195945050505050565b61010b8061015b5f395ff3fe6080604052348015600e575f80fd5b5060043610603a575f3560e01c8063b3cea21714603e578063cf309012146059578063e1afb08c146073575b5f80fd5b604660015481565b6040519081526020015b60405180910390f35b5f5460649060ff1681565b60405190151581526020016050565b6082607e36600460a9565b6084565b005b6005546001600160801b0319828116911614609d575f80fd5b505f805460ff19169055565b5f6020828403121560b8575f80fd5b81356001600160801b03198116811460ce575f80fd5b939250505056fea2646970667358221220eb55af9f8eed73c312eedc81f99f4dde7da7f1bede97d19a9199f2659b06c76e64736f6c63430008140033",
}

// PrivacyABI is the input ABI used to generate the binding from.
// Deprecated: Use PrivacyMetaData.ABI instead.
var PrivacyABI = PrivacyMetaData.ABI

// Deprecated: Use PrivacyMetaData.Sigs instead.
// PrivacyFuncSigs maps the 4-byte function signature to its string representation.
var PrivacyFuncSigs = PrivacyMetaData.Sigs

// PrivacyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PrivacyMetaData.Bin instead.
var PrivacyBin = PrivacyMetaData.Bin

// DeployPrivacy deploys a new Ethereum contract, binding an instance of Privacy to it.
func DeployPrivacy(auth *bind.TransactOpts, backend bind.ContractBackend, _data [3][32]byte) (common.Address, *types.Transaction, *Privacy, error) {
	parsed, err := PrivacyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PrivacyBin), backend, _data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Privacy{PrivacyCaller: PrivacyCaller{contract: contract}, PrivacyTransactor: PrivacyTransactor{contract: contract}, PrivacyFilterer: PrivacyFilterer{contract: contract}}, nil
}

// Privacy is an auto generated Go binding around an Ethereum contract.
type Privacy struct {
	PrivacyCaller     // Read-only binding to the contract
	PrivacyTransactor // Write-only binding to the contract
	PrivacyFilterer   // Log filterer for contract events
}

// PrivacyCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrivacyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivacyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrivacyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivacyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrivacyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivacySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrivacySession struct {
	Contract     *Privacy          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PrivacyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrivacyCallerSession struct {
	Contract *PrivacyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PrivacyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrivacyTransactorSession struct {
	Contract     *PrivacyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PrivacyRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrivacyRaw struct {
	Contract *Privacy // Generic contract binding to access the raw methods on
}

// PrivacyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrivacyCallerRaw struct {
	Contract *PrivacyCaller // Generic read-only contract binding to access the raw methods on
}

// PrivacyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrivacyTransactorRaw struct {
	Contract *PrivacyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrivacy creates a new instance of Privacy, bound to a specific deployed contract.
func NewPrivacy(address common.Address, backend bind.ContractBackend) (*Privacy, error) {
	contract, err := bindPrivacy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Privacy{PrivacyCaller: PrivacyCaller{contract: contract}, PrivacyTransactor: PrivacyTransactor{contract: contract}, PrivacyFilterer: PrivacyFilterer{contract: contract}}, nil
}

// NewPrivacyCaller creates a new read-only instance of Privacy, bound to a specific deployed contract.
func NewPrivacyCaller(address common.Address, caller bind.ContractCaller) (*PrivacyCaller, error) {
	contract, err := bindPrivacy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrivacyCaller{contract: contract}, nil
}

// NewPrivacyTransactor creates a new write-only instance of Privacy, bound to a specific deployed contract.
func NewPrivacyTransactor(address common.Address, transactor bind.ContractTransactor) (*PrivacyTransactor, error) {
	contract, err := bindPrivacy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrivacyTransactor{contract: contract}, nil
}

// NewPrivacyFilterer creates a new log filterer instance of Privacy, bound to a specific deployed contract.
func NewPrivacyFilterer(address common.Address, filterer bind.ContractFilterer) (*PrivacyFilterer, error) {
	contract, err := bindPrivacy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrivacyFilterer{contract: contract}, nil
}

// bindPrivacy binds a generic wrapper to an already deployed contract.
func bindPrivacy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PrivacyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Privacy *PrivacyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Privacy.Contract.PrivacyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Privacy *PrivacyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Privacy.Contract.PrivacyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Privacy *PrivacyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Privacy.Contract.PrivacyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Privacy *PrivacyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Privacy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Privacy *PrivacyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Privacy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Privacy *PrivacyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Privacy.Contract.contract.Transact(opts, method, params...)
}

// ID is a free data retrieval call binding the contract method 0xb3cea217.
//
// Solidity: function ID() view returns(uint256)
func (_Privacy *PrivacyCaller) ID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Privacy.contract.Call(opts, &out, "ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ID is a free data retrieval call binding the contract method 0xb3cea217.
//
// Solidity: function ID() view returns(uint256)
func (_Privacy *PrivacySession) ID() (*big.Int, error) {
	return _Privacy.Contract.ID(&_Privacy.CallOpts)
}

// ID is a free data retrieval call binding the contract method 0xb3cea217.
//
// Solidity: function ID() view returns(uint256)
func (_Privacy *PrivacyCallerSession) ID() (*big.Int, error) {
	return _Privacy.Contract.ID(&_Privacy.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_Privacy *PrivacyCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Privacy.contract.Call(opts, &out, "locked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_Privacy *PrivacySession) Locked() (bool, error) {
	return _Privacy.Contract.Locked(&_Privacy.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_Privacy *PrivacyCallerSession) Locked() (bool, error) {
	return _Privacy.Contract.Locked(&_Privacy.CallOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xe1afb08c.
//
// Solidity: function unlock(bytes16 _key) returns()
func (_Privacy *PrivacyTransactor) Unlock(opts *bind.TransactOpts, _key [16]byte) (*types.Transaction, error) {
	return _Privacy.contract.Transact(opts, "unlock", _key)
}

// Unlock is a paid mutator transaction binding the contract method 0xe1afb08c.
//
// Solidity: function unlock(bytes16 _key) returns()
func (_Privacy *PrivacySession) Unlock(_key [16]byte) (*types.Transaction, error) {
	return _Privacy.Contract.Unlock(&_Privacy.TransactOpts, _key)
}

// Unlock is a paid mutator transaction binding the contract method 0xe1afb08c.
//
// Solidity: function unlock(bytes16 _key) returns()
func (_Privacy *PrivacyTransactorSession) Unlock(_key [16]byte) (*types.Transaction, error) {
	return _Privacy.Contract.Unlock(&_Privacy.TransactOpts, _key)
}

