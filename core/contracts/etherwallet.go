// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package etherwallet

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

// EtherwalletMetaData contains all meta data concerning the Etherwallet contract.
var EtherwalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AmountMustBeGreaterThanZeroError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NoFundsAvailableForWithdrawalError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailedError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// EtherwalletABI is the input ABI used to generate the binding from.
// Deprecated: Use EtherwalletMetaData.ABI instead.
var EtherwalletABI = EtherwalletMetaData.ABI

// Etherwallet is an auto generated Go binding around an Ethereum contract.
type Etherwallet struct {
	EtherwalletCaller     // Read-only binding to the contract
	EtherwalletTransactor // Write-only binding to the contract
	EtherwalletFilterer   // Log filterer for contract events
}

// EtherwalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherwalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherwalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherwalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherwalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherwalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherwalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherwalletSession struct {
	Contract     *Etherwallet      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EtherwalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherwalletCallerSession struct {
	Contract *EtherwalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// EtherwalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherwalletTransactorSession struct {
	Contract     *EtherwalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EtherwalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherwalletRaw struct {
	Contract *Etherwallet // Generic contract binding to access the raw methods on
}

// EtherwalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherwalletCallerRaw struct {
	Contract *EtherwalletCaller // Generic read-only contract binding to access the raw methods on
}

// EtherwalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherwalletTransactorRaw struct {
	Contract *EtherwalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherwallet creates a new instance of Etherwallet, bound to a specific deployed contract.
func NewEtherwallet(address common.Address, backend bind.ContractBackend) (*Etherwallet, error) {
	contract, err := bindEtherwallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Etherwallet{EtherwalletCaller: EtherwalletCaller{contract: contract}, EtherwalletTransactor: EtherwalletTransactor{contract: contract}, EtherwalletFilterer: EtherwalletFilterer{contract: contract}}, nil
}

// NewEtherwalletCaller creates a new read-only instance of Etherwallet, bound to a specific deployed contract.
func NewEtherwalletCaller(address common.Address, caller bind.ContractCaller) (*EtherwalletCaller, error) {
	contract, err := bindEtherwallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherwalletCaller{contract: contract}, nil
}

// NewEtherwalletTransactor creates a new write-only instance of Etherwallet, bound to a specific deployed contract.
func NewEtherwalletTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherwalletTransactor, error) {
	contract, err := bindEtherwallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherwalletTransactor{contract: contract}, nil
}

// NewEtherwalletFilterer creates a new log filterer instance of Etherwallet, bound to a specific deployed contract.
func NewEtherwalletFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherwalletFilterer, error) {
	contract, err := bindEtherwallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherwalletFilterer{contract: contract}, nil
}

// bindEtherwallet binds a generic wrapper to an already deployed contract.
func bindEtherwallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EtherwalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Etherwallet *EtherwalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Etherwallet.Contract.EtherwalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Etherwallet *EtherwalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Etherwallet.Contract.EtherwalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Etherwallet *EtherwalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Etherwallet.Contract.EtherwalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Etherwallet *EtherwalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Etherwallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Etherwallet *EtherwalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Etherwallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Etherwallet *EtherwalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Etherwallet.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Etherwallet *EtherwalletTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Etherwallet.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Etherwallet *EtherwalletSession) Deposit() (*types.Transaction, error) {
	return _Etherwallet.Contract.Deposit(&_Etherwallet.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Etherwallet *EtherwalletTransactorSession) Deposit() (*types.Transaction, error) {
	return _Etherwallet.Contract.Deposit(&_Etherwallet.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Etherwallet *EtherwalletTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Etherwallet.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Etherwallet *EtherwalletSession) Withdraw() (*types.Transaction, error) {
	return _Etherwallet.Contract.Withdraw(&_Etherwallet.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Etherwallet *EtherwalletTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Etherwallet.Contract.Withdraw(&_Etherwallet.TransactOpts)
}
