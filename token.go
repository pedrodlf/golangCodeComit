// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"linkedAtMeetings\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerFactory\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"meeting\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"returnCode\",\"type\":\"uint256\"}],\"name\":\"EventCreateLinkedAtMeeting\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"meeting\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"initialPoints\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"returnCode\",\"type\":\"uint256\"}],\"name\":\"EventRegisterUserAtMeetingFactory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"meeting\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"returnCode\",\"type\":\"uint256\"}],\"name\":\"EventDeregisterUserAtMeetingFactory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"meeting\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"points\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"returnCode\",\"type\":\"uint256\"}],\"name\":\"EventGetUserPointsFactory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"caller\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"action\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"points\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"returnCode\",\"type\":\"uint256\"}],\"name\":\"EventAddPointsFactory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"caller\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"action\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"points\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"returnCode\",\"type\":\"uint256\"}],\"name\":\"EventPayPrizeFactory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"caller\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"returnCode\",\"type\":\"uint256\"}],\"name\":\"EventCleanUserPointsFactory\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"meeting\",\"type\":\"uint32\"},{\"name\":\"startTime\",\"type\":\"uint256\"},{\"name\":\"endingTime\",\"type\":\"uint256\"}],\"name\":\"createLinkedAtMeeting\",\"outputs\":[{\"name\":\"code\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"meeting\",\"type\":\"uint32\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"initialPoints\",\"type\":\"uint32\"}],\"name\":\"registerUserAtMeeting\",\"outputs\":[{\"name\":\"code\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"meeting\",\"type\":\"uint32\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"deregisterUserAtMeeting\",\"outputs\":[{\"name\":\"code\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"meeting\",\"type\":\"uint32\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserPoints\",\"outputs\":[{\"name\":\"code\",\"type\":\"uint256\"},{\"name\":\"points\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"meeting\",\"type\":\"uint32\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"action\",\"type\":\"uint32\"},{\"name\":\"points\",\"type\":\"uint32\"}],\"name\":\"addPoints\",\"outputs\":[{\"name\":\"code\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"meeting\",\"type\":\"uint32\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"action\",\"type\":\"uint32\"},{\"name\":\"points\",\"type\":\"uint32\"}],\"name\":\"payPrize\",\"outputs\":[{\"name\":\"code\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"meeting\",\"type\":\"uint32\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"cleanUserPoints\",\"outputs\":[{\"name\":\"code\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// LinkedAtMeetings is a free data retrieval call binding the contract method 0x722465cf.
//
// Solidity: function linkedAtMeetings( uint32) constant returns(address)
func (_Token *TokenCaller) LinkedAtMeetings(opts *bind.CallOpts, arg0 uint32) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "linkedAtMeetings", arg0)
	return *ret0, err
}

// LinkedAtMeetings is a free data retrieval call binding the contract method 0x722465cf.
//
// Solidity: function linkedAtMeetings( uint32) constant returns(address)
func (_Token *TokenSession) LinkedAtMeetings(arg0 uint32) (common.Address, error) {
	return _Token.Contract.LinkedAtMeetings(&_Token.CallOpts, arg0)
}

// LinkedAtMeetings is a free data retrieval call binding the contract method 0x722465cf.
//
// Solidity: function linkedAtMeetings( uint32) constant returns(address)
func (_Token *TokenCallerSession) LinkedAtMeetings(arg0 uint32) (common.Address, error) {
	return _Token.Contract.LinkedAtMeetings(&_Token.CallOpts, arg0)
}

// OwnerFactory is a free data retrieval call binding the contract method 0xd2779e0e.
//
// Solidity: function ownerFactory() constant returns(address)
func (_Token *TokenCaller) OwnerFactory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "ownerFactory")
	return *ret0, err
}

// OwnerFactory is a free data retrieval call binding the contract method 0xd2779e0e.
//
// Solidity: function ownerFactory() constant returns(address)
func (_Token *TokenSession) OwnerFactory() (common.Address, error) {
	return _Token.Contract.OwnerFactory(&_Token.CallOpts)
}

// OwnerFactory is a free data retrieval call binding the contract method 0xd2779e0e.
//
// Solidity: function ownerFactory() constant returns(address)
func (_Token *TokenCallerSession) OwnerFactory() (common.Address, error) {
	return _Token.Contract.OwnerFactory(&_Token.CallOpts)
}

// AddPoints is a paid mutator transaction binding the contract method 0x5cb0e01a.
//
// Solidity: function addPoints(meeting uint32, user address, action uint32, points uint32) returns(code uint256)
func (_Token *TokenTransactor) AddPoints(opts *bind.TransactOpts, meeting uint32, user common.Address, action uint32, points uint32) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "addPoints", meeting, user, action, points)
}

// AddPoints is a paid mutator transaction binding the contract method 0x5cb0e01a.
//
// Solidity: function addPoints(meeting uint32, user address, action uint32, points uint32) returns(code uint256)
func (_Token *TokenSession) AddPoints(meeting uint32, user common.Address, action uint32, points uint32) (*types.Transaction, error) {
	return _Token.Contract.AddPoints(&_Token.TransactOpts, meeting, user, action, points)
}

// AddPoints is a paid mutator transaction binding the contract method 0x5cb0e01a.
//
// Solidity: function addPoints(meeting uint32, user address, action uint32, points uint32) returns(code uint256)
func (_Token *TokenTransactorSession) AddPoints(meeting uint32, user common.Address, action uint32, points uint32) (*types.Transaction, error) {
	return _Token.Contract.AddPoints(&_Token.TransactOpts, meeting, user, action, points)
}

// CleanUserPoints is a paid mutator transaction binding the contract method 0xb0b9d224.
//
// Solidity: function cleanUserPoints(meeting uint32, user address) returns(code uint256)
func (_Token *TokenTransactor) CleanUserPoints(opts *bind.TransactOpts, meeting uint32, user common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "cleanUserPoints", meeting, user)
}

// CleanUserPoints is a paid mutator transaction binding the contract method 0xb0b9d224.
//
// Solidity: function cleanUserPoints(meeting uint32, user address) returns(code uint256)
func (_Token *TokenSession) CleanUserPoints(meeting uint32, user common.Address) (*types.Transaction, error) {
	return _Token.Contract.CleanUserPoints(&_Token.TransactOpts, meeting, user)
}

// CleanUserPoints is a paid mutator transaction binding the contract method 0xb0b9d224.
//
// Solidity: function cleanUserPoints(meeting uint32, user address) returns(code uint256)
func (_Token *TokenTransactorSession) CleanUserPoints(meeting uint32, user common.Address) (*types.Transaction, error) {
	return _Token.Contract.CleanUserPoints(&_Token.TransactOpts, meeting, user)
}

// CreateLinkedAtMeeting is a paid mutator transaction binding the contract method 0xb0ba0537.
//
// Solidity: function createLinkedAtMeeting(meeting uint32, startTime uint256, endingTime uint256) returns(code uint256)
func (_Token *TokenTransactor) CreateLinkedAtMeeting(opts *bind.TransactOpts, meeting uint32, startTime *big.Int, endingTime *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "createLinkedAtMeeting", meeting, startTime, endingTime)
}

// CreateLinkedAtMeeting is a paid mutator transaction binding the contract method 0xb0ba0537.
//
// Solidity: function createLinkedAtMeeting(meeting uint32, startTime uint256, endingTime uint256) returns(code uint256)
func (_Token *TokenSession) CreateLinkedAtMeeting(meeting uint32, startTime *big.Int, endingTime *big.Int) (*types.Transaction, error) {
	return _Token.Contract.CreateLinkedAtMeeting(&_Token.TransactOpts, meeting, startTime, endingTime)
}

// CreateLinkedAtMeeting is a paid mutator transaction binding the contract method 0xb0ba0537.
//
// Solidity: function createLinkedAtMeeting(meeting uint32, startTime uint256, endingTime uint256) returns(code uint256)
func (_Token *TokenTransactorSession) CreateLinkedAtMeeting(meeting uint32, startTime *big.Int, endingTime *big.Int) (*types.Transaction, error) {
	return _Token.Contract.CreateLinkedAtMeeting(&_Token.TransactOpts, meeting, startTime, endingTime)
}

// DeregisterUserAtMeeting is a paid mutator transaction binding the contract method 0x74ebcd49.
//
// Solidity: function deregisterUserAtMeeting(meeting uint32, user address) returns(code uint256)
func (_Token *TokenTransactor) DeregisterUserAtMeeting(opts *bind.TransactOpts, meeting uint32, user common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "deregisterUserAtMeeting", meeting, user)
}

// DeregisterUserAtMeeting is a paid mutator transaction binding the contract method 0x74ebcd49.
//
// Solidity: function deregisterUserAtMeeting(meeting uint32, user address) returns(code uint256)
func (_Token *TokenSession) DeregisterUserAtMeeting(meeting uint32, user common.Address) (*types.Transaction, error) {
	return _Token.Contract.DeregisterUserAtMeeting(&_Token.TransactOpts, meeting, user)
}

// DeregisterUserAtMeeting is a paid mutator transaction binding the contract method 0x74ebcd49.
//
// Solidity: function deregisterUserAtMeeting(meeting uint32, user address) returns(code uint256)
func (_Token *TokenTransactorSession) DeregisterUserAtMeeting(meeting uint32, user common.Address) (*types.Transaction, error) {
	return _Token.Contract.DeregisterUserAtMeeting(&_Token.TransactOpts, meeting, user)
}

// GetUserPoints is a paid mutator transaction binding the contract method 0x8c3a96a3.
//
// Solidity: function getUserPoints(meeting uint32, user address) returns(code uint256, points uint32)
func (_Token *TokenTransactor) GetUserPoints(opts *bind.TransactOpts, meeting uint32, user common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "getUserPoints", meeting, user)
}

// GetUserPoints is a paid mutator transaction binding the contract method 0x8c3a96a3.
//
// Solidity: function getUserPoints(meeting uint32, user address) returns(code uint256, points uint32)
func (_Token *TokenSession) GetUserPoints(meeting uint32, user common.Address) (*types.Transaction, error) {
	return _Token.Contract.GetUserPoints(&_Token.TransactOpts, meeting, user)
}

// GetUserPoints is a paid mutator transaction binding the contract method 0x8c3a96a3.
//
// Solidity: function getUserPoints(meeting uint32, user address) returns(code uint256, points uint32)
func (_Token *TokenTransactorSession) GetUserPoints(meeting uint32, user common.Address) (*types.Transaction, error) {
	return _Token.Contract.GetUserPoints(&_Token.TransactOpts, meeting, user)
}

// PayPrize is a paid mutator transaction binding the contract method 0x4cff35b7.
//
// Solidity: function payPrize(meeting uint32, user address, action uint32, points uint32) returns(code uint256)
func (_Token *TokenTransactor) PayPrize(opts *bind.TransactOpts, meeting uint32, user common.Address, action uint32, points uint32) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "payPrize", meeting, user, action, points)
}

// PayPrize is a paid mutator transaction binding the contract method 0x4cff35b7.
//
// Solidity: function payPrize(meeting uint32, user address, action uint32, points uint32) returns(code uint256)
func (_Token *TokenSession) PayPrize(meeting uint32, user common.Address, action uint32, points uint32) (*types.Transaction, error) {
	return _Token.Contract.PayPrize(&_Token.TransactOpts, meeting, user, action, points)
}

// PayPrize is a paid mutator transaction binding the contract method 0x4cff35b7.
//
// Solidity: function payPrize(meeting uint32, user address, action uint32, points uint32) returns(code uint256)
func (_Token *TokenTransactorSession) PayPrize(meeting uint32, user common.Address, action uint32, points uint32) (*types.Transaction, error) {
	return _Token.Contract.PayPrize(&_Token.TransactOpts, meeting, user, action, points)
}

// RegisterUserAtMeeting is a paid mutator transaction binding the contract method 0xb5e9e929.
//
// Solidity: function registerUserAtMeeting(meeting uint32, user address, initialPoints uint32) returns(code uint256)
func (_Token *TokenTransactor) RegisterUserAtMeeting(opts *bind.TransactOpts, meeting uint32, user common.Address, initialPoints uint32) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "registerUserAtMeeting", meeting, user, initialPoints)
}

// RegisterUserAtMeeting is a paid mutator transaction binding the contract method 0xb5e9e929.
//
// Solidity: function registerUserAtMeeting(meeting uint32, user address, initialPoints uint32) returns(code uint256)
func (_Token *TokenSession) RegisterUserAtMeeting(meeting uint32, user common.Address, initialPoints uint32) (*types.Transaction, error) {
	return _Token.Contract.RegisterUserAtMeeting(&_Token.TransactOpts, meeting, user, initialPoints)
}

// RegisterUserAtMeeting is a paid mutator transaction binding the contract method 0xb5e9e929.
//
// Solidity: function registerUserAtMeeting(meeting uint32, user address, initialPoints uint32) returns(code uint256)
func (_Token *TokenTransactorSession) RegisterUserAtMeeting(meeting uint32, user common.Address, initialPoints uint32) (*types.Transaction, error) {
	return _Token.Contract.RegisterUserAtMeeting(&_Token.TransactOpts, meeting, user, initialPoints)
}

// TokenEventAddPointsFactoryIterator is returned from FilterEventAddPointsFactory and is used to iterate over the raw logs and unpacked data for EventAddPointsFactory events raised by the Token contract.
type TokenEventAddPointsFactoryIterator struct {
	Event *TokenEventAddPointsFactory // Event containing the contract specifics and raw log

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
func (it *TokenEventAddPointsFactoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenEventAddPointsFactory)
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
		it.Event = new(TokenEventAddPointsFactory)
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
func (it *TokenEventAddPointsFactoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenEventAddPointsFactoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenEventAddPointsFactory represents a EventAddPointsFactory event raised by the Token contract.
type TokenEventAddPointsFactory struct {
	Caller     uint32
	User       common.Address
	Action     uint32
	Points     uint32
	ReturnCode *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEventAddPointsFactory is a free log retrieval operation binding the contract event 0x1265a2e57d260c08034c1641161ac849fc65d94c6563dfb89ebfbe32a68d6205.
//
// Solidity: e EventAddPointsFactory(caller indexed uint32, user indexed address, action uint32, points uint32, returnCode indexed uint256)
func (_Token *TokenFilterer) FilterEventAddPointsFactory(opts *bind.FilterOpts, caller []uint32, user []common.Address, returnCode []*big.Int) (*TokenEventAddPointsFactoryIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var returnCodeRule []interface{}
	for _, returnCodeItem := range returnCode {
		returnCodeRule = append(returnCodeRule, returnCodeItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "EventAddPointsFactory", callerRule, userRule, returnCodeRule)
	if err != nil {
		return nil, err
	}
	return &TokenEventAddPointsFactoryIterator{contract: _Token.contract, event: "EventAddPointsFactory", logs: logs, sub: sub}, nil
}

// WatchEventAddPointsFactory is a free log subscription operation binding the contract event 0x1265a2e57d260c08034c1641161ac849fc65d94c6563dfb89ebfbe32a68d6205.
//
// Solidity: e EventAddPointsFactory(caller indexed uint32, user indexed address, action uint32, points uint32, returnCode indexed uint256)
func (_Token *TokenFilterer) WatchEventAddPointsFactory(opts *bind.WatchOpts, sink chan<- *TokenEventAddPointsFactory, caller []uint32, user []common.Address, returnCode []*big.Int) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var returnCodeRule []interface{}
	for _, returnCodeItem := range returnCode {
		returnCodeRule = append(returnCodeRule, returnCodeItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "EventAddPointsFactory", callerRule, userRule, returnCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenEventAddPointsFactory)
				if err := _Token.contract.UnpackLog(event, "EventAddPointsFactory", log); err != nil {
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

// TokenEventCleanUserPointsFactoryIterator is returned from FilterEventCleanUserPointsFactory and is used to iterate over the raw logs and unpacked data for EventCleanUserPointsFactory events raised by the Token contract.
type TokenEventCleanUserPointsFactoryIterator struct {
	Event *TokenEventCleanUserPointsFactory // Event containing the contract specifics and raw log

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
func (it *TokenEventCleanUserPointsFactoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenEventCleanUserPointsFactory)
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
		it.Event = new(TokenEventCleanUserPointsFactory)
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
func (it *TokenEventCleanUserPointsFactoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenEventCleanUserPointsFactoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenEventCleanUserPointsFactory represents a EventCleanUserPointsFactory event raised by the Token contract.
type TokenEventCleanUserPointsFactory struct {
	Caller     uint32
	User       common.Address
	ReturnCode *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEventCleanUserPointsFactory is a free log retrieval operation binding the contract event 0xa39804f33f39f2d6cbe7d3a2d3ae168b71fafe703602735645ebf2f4d19d259a.
//
// Solidity: e EventCleanUserPointsFactory(caller indexed uint32, user indexed address, returnCode indexed uint256)
func (_Token *TokenFilterer) FilterEventCleanUserPointsFactory(opts *bind.FilterOpts, caller []uint32, user []common.Address, returnCode []*big.Int) (*TokenEventCleanUserPointsFactoryIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var returnCodeRule []interface{}
	for _, returnCodeItem := range returnCode {
		returnCodeRule = append(returnCodeRule, returnCodeItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "EventCleanUserPointsFactory", callerRule, userRule, returnCodeRule)
	if err != nil {
		return nil, err
	}
	return &TokenEventCleanUserPointsFactoryIterator{contract: _Token.contract, event: "EventCleanUserPointsFactory", logs: logs, sub: sub}, nil
}

// WatchEventCleanUserPointsFactory is a free log subscription operation binding the contract event 0xa39804f33f39f2d6cbe7d3a2d3ae168b71fafe703602735645ebf2f4d19d259a.
//
// Solidity: e EventCleanUserPointsFactory(caller indexed uint32, user indexed address, returnCode indexed uint256)
func (_Token *TokenFilterer) WatchEventCleanUserPointsFactory(opts *bind.WatchOpts, sink chan<- *TokenEventCleanUserPointsFactory, caller []uint32, user []common.Address, returnCode []*big.Int) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var returnCodeRule []interface{}
	for _, returnCodeItem := range returnCode {
		returnCodeRule = append(returnCodeRule, returnCodeItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "EventCleanUserPointsFactory", callerRule, userRule, returnCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenEventCleanUserPointsFactory)
				if err := _Token.contract.UnpackLog(event, "EventCleanUserPointsFactory", log); err != nil {
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

// TokenEventCreateLinkedAtMeetingIterator is returned from FilterEventCreateLinkedAtMeeting and is used to iterate over the raw logs and unpacked data for EventCreateLinkedAtMeeting events raised by the Token contract.
type TokenEventCreateLinkedAtMeetingIterator struct {
	Event *TokenEventCreateLinkedAtMeeting // Event containing the contract specifics and raw log

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
func (it *TokenEventCreateLinkedAtMeetingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenEventCreateLinkedAtMeeting)
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
		it.Event = new(TokenEventCreateLinkedAtMeeting)
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
func (it *TokenEventCreateLinkedAtMeetingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenEventCreateLinkedAtMeetingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenEventCreateLinkedAtMeeting represents a EventCreateLinkedAtMeeting event raised by the Token contract.
type TokenEventCreateLinkedAtMeeting struct {
	Meeting    uint32
	ReturnCode *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEventCreateLinkedAtMeeting is a free log retrieval operation binding the contract event 0x28de32d3bc24cfbe5227a6c5dffa462b612f0f3b21f77c55cd6be19a7199fb3c.
//
// Solidity: e EventCreateLinkedAtMeeting(meeting indexed uint32, returnCode indexed uint256)
func (_Token *TokenFilterer) FilterEventCreateLinkedAtMeeting(opts *bind.FilterOpts, meeting []uint32, returnCode []*big.Int) (*TokenEventCreateLinkedAtMeetingIterator, error) {

	var meetingRule []interface{}
	for _, meetingItem := range meeting {
		meetingRule = append(meetingRule, meetingItem)
	}
	var returnCodeRule []interface{}
	for _, returnCodeItem := range returnCode {
		returnCodeRule = append(returnCodeRule, returnCodeItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "EventCreateLinkedAtMeeting", meetingRule, returnCodeRule)
	if err != nil {
		return nil, err
	}
	return &TokenEventCreateLinkedAtMeetingIterator{contract: _Token.contract, event: "EventCreateLinkedAtMeeting", logs: logs, sub: sub}, nil
}

// WatchEventCreateLinkedAtMeeting is a free log subscription operation binding the contract event 0x28de32d3bc24cfbe5227a6c5dffa462b612f0f3b21f77c55cd6be19a7199fb3c.
//
// Solidity: e EventCreateLinkedAtMeeting(meeting indexed uint32, returnCode indexed uint256)
func (_Token *TokenFilterer) WatchEventCreateLinkedAtMeeting(opts *bind.WatchOpts, sink chan<- *TokenEventCreateLinkedAtMeeting, meeting []uint32, returnCode []*big.Int) (event.Subscription, error) {

	var meetingRule []interface{}
	for _, meetingItem := range meeting {
		meetingRule = append(meetingRule, meetingItem)
	}
	var returnCodeRule []interface{}
	for _, returnCodeItem := range returnCode {
		returnCodeRule = append(returnCodeRule, returnCodeItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "EventCreateLinkedAtMeeting", meetingRule, returnCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenEventCreateLinkedAtMeeting)
				if err := _Token.contract.UnpackLog(event, "EventCreateLinkedAtMeeting", log); err != nil {
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

// TokenEventDeregisterUserAtMeetingFactoryIterator is returned from FilterEventDeregisterUserAtMeetingFactory and is used to iterate over the raw logs and unpacked data for EventDeregisterUserAtMeetingFactory events raised by the Token contract.
type TokenEventDeregisterUserAtMeetingFactoryIterator struct {
	Event *TokenEventDeregisterUserAtMeetingFactory // Event containing the contract specifics and raw log

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
func (it *TokenEventDeregisterUserAtMeetingFactoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenEventDeregisterUserAtMeetingFactory)
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
		it.Event = new(TokenEventDeregisterUserAtMeetingFactory)
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
func (it *TokenEventDeregisterUserAtMeetingFactoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenEventDeregisterUserAtMeetingFactoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenEventDeregisterUserAtMeetingFactory represents a EventDeregisterUserAtMeetingFactory event raised by the Token contract.
type TokenEventDeregisterUserAtMeetingFactory struct {
	Meeting    uint32
	User       common.Address
	ReturnCode *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEventDeregisterUserAtMeetingFactory is a free log retrieval operation binding the contract event 0x329ea9412c577d822aa3ba67ef83047e6356487cc23b1667b838086678dcf9ed.
//
// Solidity: e EventDeregisterUserAtMeetingFactory(meeting indexed uint32, user indexed address, returnCode indexed uint256)
func (_Token *TokenFilterer) FilterEventDeregisterUserAtMeetingFactory(opts *bind.FilterOpts, meeting []uint32, user []common.Address, returnCode []*big.Int) (*TokenEventDeregisterUserAtMeetingFactoryIterator, error) {

	var meetingRule []interface{}
	for _, meetingItem := range meeting {
		meetingRule = append(meetingRule, meetingItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var returnCodeRule []interface{}
	for _, returnCodeItem := range returnCode {
		returnCodeRule = append(returnCodeRule, returnCodeItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "EventDeregisterUserAtMeetingFactory", meetingRule, userRule, returnCodeRule)
	if err != nil {
		return nil, err
	}
	return &TokenEventDeregisterUserAtMeetingFactoryIterator{contract: _Token.contract, event: "EventDeregisterUserAtMeetingFactory", logs: logs, sub: sub}, nil
}

// WatchEventDeregisterUserAtMeetingFactory is a free log subscription operation binding the contract event 0x329ea9412c577d822aa3ba67ef83047e6356487cc23b1667b838086678dcf9ed.
//
// Solidity: e EventDeregisterUserAtMeetingFactory(meeting indexed uint32, user indexed address, returnCode indexed uint256)
func (_Token *TokenFilterer) WatchEventDeregisterUserAtMeetingFactory(opts *bind.WatchOpts, sink chan<- *TokenEventDeregisterUserAtMeetingFactory, meeting []uint32, user []common.Address, returnCode []*big.Int) (event.Subscription, error) {

	var meetingRule []interface{}
	for _, meetingItem := range meeting {
		meetingRule = append(meetingRule, meetingItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var returnCodeRule []interface{}
	for _, returnCodeItem := range returnCode {
		returnCodeRule = append(returnCodeRule, returnCodeItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "EventDeregisterUserAtMeetingFactory", meetingRule, userRule, returnCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenEventDeregisterUserAtMeetingFactory)
				if err := _Token.contract.UnpackLog(event, "EventDeregisterUserAtMeetingFactory", log); err != nil {
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

// TokenEventGetUserPointsFactoryIterator is returned from FilterEventGetUserPointsFactory and is used to iterate over the raw logs and unpacked data for EventGetUserPointsFactory events raised by the Token contract.
type TokenEventGetUserPointsFactoryIterator struct {
	Event *TokenEventGetUserPointsFactory // Event containing the contract specifics and raw log

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
func (it *TokenEventGetUserPointsFactoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenEventGetUserPointsFactory)
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
		it.Event = new(TokenEventGetUserPointsFactory)
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
func (it *TokenEventGetUserPointsFactoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenEventGetUserPointsFactoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenEventGetUserPointsFactory represents a EventGetUserPointsFactory event raised by the Token contract.
type TokenEventGetUserPointsFactory struct {
	Meeting    uint32
	User       common.Address
	Points     uint32
	ReturnCode *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEventGetUserPointsFactory is a free log retrieval operation binding the contract event 0x97685ae83b5b99a6fe01340065e7c578cdb72b9d484fc1d18bb9cf05216187f7.
//
// Solidity: e EventGetUserPointsFactory(meeting uint32, user indexed address, points indexed uint32, returnCode indexed uint256)
func (_Token *TokenFilterer) FilterEventGetUserPointsFactory(opts *bind.FilterOpts, user []common.Address, points []uint32, returnCode []*big.Int) (*TokenEventGetUserPointsFactoryIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pointsRule []interface{}
	for _, pointsItem := range points {
		pointsRule = append(pointsRule, pointsItem)
	}
	var returnCodeRule []interface{}
	for _, returnCodeItem := range returnCode {
		returnCodeRule = append(returnCodeRule, returnCodeItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "EventGetUserPointsFactory", userRule, pointsRule, returnCodeRule)
	if err != nil {
		return nil, err
	}
	return &TokenEventGetUserPointsFactoryIterator{contract: _Token.contract, event: "EventGetUserPointsFactory", logs: logs, sub: sub}, nil
}

// WatchEventGetUserPointsFactory is a free log subscription operation binding the contract event 0x97685ae83b5b99a6fe01340065e7c578cdb72b9d484fc1d18bb9cf05216187f7.
//
// Solidity: e EventGetUserPointsFactory(meeting uint32, user indexed address, points indexed uint32, returnCode indexed uint256)
func (_Token *TokenFilterer) WatchEventGetUserPointsFactory(opts *bind.WatchOpts, sink chan<- *TokenEventGetUserPointsFactory, user []common.Address, points []uint32, returnCode []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pointsRule []interface{}
	for _, pointsItem := range points {
		pointsRule = append(pointsRule, pointsItem)
	}
	var returnCodeRule []interface{}
	for _, returnCodeItem := range returnCode {
		returnCodeRule = append(returnCodeRule, returnCodeItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "EventGetUserPointsFactory", userRule, pointsRule, returnCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenEventGetUserPointsFactory)
				if err := _Token.contract.UnpackLog(event, "EventGetUserPointsFactory", log); err != nil {
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

// TokenEventPayPrizeFactoryIterator is returned from FilterEventPayPrizeFactory and is used to iterate over the raw logs and unpacked data for EventPayPrizeFactory events raised by the Token contract.
type TokenEventPayPrizeFactoryIterator struct {
	Event *TokenEventPayPrizeFactory // Event containing the contract specifics and raw log

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
func (it *TokenEventPayPrizeFactoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenEventPayPrizeFactory)
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
		it.Event = new(TokenEventPayPrizeFactory)
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
func (it *TokenEventPayPrizeFactoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenEventPayPrizeFactoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenEventPayPrizeFactory represents a EventPayPrizeFactory event raised by the Token contract.
type TokenEventPayPrizeFactory struct {
	Caller     uint32
	User       common.Address
	Action     uint32
	Points     uint32
	ReturnCode *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEventPayPrizeFactory is a free log retrieval operation binding the contract event 0x16c15b6cdc61e2fbea04a1bd03283dfb194f7768dec4298f0d3ffce6a70b84de.
//
// Solidity: e EventPayPrizeFactory(caller indexed uint32, user indexed address, action uint32, points uint32, returnCode indexed uint256)
func (_Token *TokenFilterer) FilterEventPayPrizeFactory(opts *bind.FilterOpts, caller []uint32, user []common.Address, returnCode []*big.Int) (*TokenEventPayPrizeFactoryIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var returnCodeRule []interface{}
	for _, returnCodeItem := range returnCode {
		returnCodeRule = append(returnCodeRule, returnCodeItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "EventPayPrizeFactory", callerRule, userRule, returnCodeRule)
	if err != nil {
		return nil, err
	}
	return &TokenEventPayPrizeFactoryIterator{contract: _Token.contract, event: "EventPayPrizeFactory", logs: logs, sub: sub}, nil
}

// WatchEventPayPrizeFactory is a free log subscription operation binding the contract event 0x16c15b6cdc61e2fbea04a1bd03283dfb194f7768dec4298f0d3ffce6a70b84de.
//
// Solidity: e EventPayPrizeFactory(caller indexed uint32, user indexed address, action uint32, points uint32, returnCode indexed uint256)
func (_Token *TokenFilterer) WatchEventPayPrizeFactory(opts *bind.WatchOpts, sink chan<- *TokenEventPayPrizeFactory, caller []uint32, user []common.Address, returnCode []*big.Int) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var returnCodeRule []interface{}
	for _, returnCodeItem := range returnCode {
		returnCodeRule = append(returnCodeRule, returnCodeItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "EventPayPrizeFactory", callerRule, userRule, returnCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenEventPayPrizeFactory)
				if err := _Token.contract.UnpackLog(event, "EventPayPrizeFactory", log); err != nil {
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

// TokenEventRegisterUserAtMeetingFactoryIterator is returned from FilterEventRegisterUserAtMeetingFactory and is used to iterate over the raw logs and unpacked data for EventRegisterUserAtMeetingFactory events raised by the Token contract.
type TokenEventRegisterUserAtMeetingFactoryIterator struct {
	Event *TokenEventRegisterUserAtMeetingFactory // Event containing the contract specifics and raw log

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
func (it *TokenEventRegisterUserAtMeetingFactoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenEventRegisterUserAtMeetingFactory)
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
		it.Event = new(TokenEventRegisterUserAtMeetingFactory)
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
func (it *TokenEventRegisterUserAtMeetingFactoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenEventRegisterUserAtMeetingFactoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenEventRegisterUserAtMeetingFactory represents a EventRegisterUserAtMeetingFactory event raised by the Token contract.
type TokenEventRegisterUserAtMeetingFactory struct {
	Meeting       uint32
	User          common.Address
	InitialPoints uint32
	ReturnCode    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEventRegisterUserAtMeetingFactory is a free log retrieval operation binding the contract event 0x6ec83203ea51e56b38f14171a77bfb644b673c0a0d9c38c7f3f34f09c79cc993.
//
// Solidity: e EventRegisterUserAtMeetingFactory(meeting indexed uint32, user indexed address, initialPoints uint32, returnCode indexed uint256)
func (_Token *TokenFilterer) FilterEventRegisterUserAtMeetingFactory(opts *bind.FilterOpts, meeting []uint32, user []common.Address, returnCode []*big.Int) (*TokenEventRegisterUserAtMeetingFactoryIterator, error) {

	var meetingRule []interface{}
	for _, meetingItem := range meeting {
		meetingRule = append(meetingRule, meetingItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var returnCodeRule []interface{}
	for _, returnCodeItem := range returnCode {
		returnCodeRule = append(returnCodeRule, returnCodeItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "EventRegisterUserAtMeetingFactory", meetingRule, userRule, returnCodeRule)
	if err != nil {
		return nil, err
	}
	return &TokenEventRegisterUserAtMeetingFactoryIterator{contract: _Token.contract, event: "EventRegisterUserAtMeetingFactory", logs: logs, sub: sub}, nil
}

// WatchEventRegisterUserAtMeetingFactory is a free log subscription operation binding the contract event 0x6ec83203ea51e56b38f14171a77bfb644b673c0a0d9c38c7f3f34f09c79cc993.
//
// Solidity: e EventRegisterUserAtMeetingFactory(meeting indexed uint32, user indexed address, initialPoints uint32, returnCode indexed uint256)
func (_Token *TokenFilterer) WatchEventRegisterUserAtMeetingFactory(opts *bind.WatchOpts, sink chan<- *TokenEventRegisterUserAtMeetingFactory, meeting []uint32, user []common.Address, returnCode []*big.Int) (event.Subscription, error) {

	var meetingRule []interface{}
	for _, meetingItem := range meeting {
		meetingRule = append(meetingRule, meetingItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var returnCodeRule []interface{}
	for _, returnCodeItem := range returnCode {
		returnCodeRule = append(returnCodeRule, returnCodeItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "EventRegisterUserAtMeetingFactory", meetingRule, userRule, returnCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenEventRegisterUserAtMeetingFactory)
				if err := _Token.contract.UnpackLog(event, "EventRegisterUserAtMeetingFactory", log); err != nil {
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
