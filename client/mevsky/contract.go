// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mevsky

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MevskyABI is the input ABI used to generate the binding from.
const MevskyABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minBounty\",\"type\":\"uint256\"}],\"name\":\"MinBountySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"TurnedOff\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bounty\",\"type\":\"uint256\"}],\"name\":\"TurnedOn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"minBounty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"on\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMinBounty\",\"type\":\"uint256\"}],\"name\":\"setMinBounty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"turnOff\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"turnOn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Mevsky is an auto generated Go binding around an Ethereum contract.
type Mevsky struct {
	MevskyCaller     // Read-only binding to the contract
	MevskyTransactor // Write-only binding to the contract
	MevskyFilterer   // Log filterer for contract events
}

// MevskyCaller is an auto generated read-only Go binding around an Ethereum contract.
type MevskyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MevskyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MevskyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MevskyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MevskyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MevskySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MevskySession struct {
	Contract     *Mevsky           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MevskyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MevskyCallerSession struct {
	Contract *MevskyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MevskyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MevskyTransactorSession struct {
	Contract     *MevskyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MevskyRaw is an auto generated low-level Go binding around an Ethereum contract.
type MevskyRaw struct {
	Contract *Mevsky // Generic contract binding to access the raw methods on
}

// MevskyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MevskyCallerRaw struct {
	Contract *MevskyCaller // Generic read-only contract binding to access the raw methods on
}

// MevskyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MevskyTransactorRaw struct {
	Contract *MevskyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMevsky creates a new instance of Mevsky, bound to a specific deployed contract.
func NewMevsky(address common.Address, backend bind.ContractBackend) (*Mevsky, error) {
	contract, err := bindMevsky(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mevsky{MevskyCaller: MevskyCaller{contract: contract}, MevskyTransactor: MevskyTransactor{contract: contract}, MevskyFilterer: MevskyFilterer{contract: contract}}, nil
}

// NewMevskyCaller creates a new read-only instance of Mevsky, bound to a specific deployed contract.
func NewMevskyCaller(address common.Address, caller bind.ContractCaller) (*MevskyCaller, error) {
	contract, err := bindMevsky(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MevskyCaller{contract: contract}, nil
}

// NewMevskyTransactor creates a new write-only instance of Mevsky, bound to a specific deployed contract.
func NewMevskyTransactor(address common.Address, transactor bind.ContractTransactor) (*MevskyTransactor, error) {
	contract, err := bindMevsky(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MevskyTransactor{contract: contract}, nil
}

// NewMevskyFilterer creates a new log filterer instance of Mevsky, bound to a specific deployed contract.
func NewMevskyFilterer(address common.Address, filterer bind.ContractFilterer) (*MevskyFilterer, error) {
	contract, err := bindMevsky(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MevskyFilterer{contract: contract}, nil
}

// bindMevsky binds a generic wrapper to an already deployed contract.
func bindMevsky(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MevskyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mevsky *MevskyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mevsky.Contract.MevskyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mevsky *MevskyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mevsky.Contract.MevskyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mevsky *MevskyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mevsky.Contract.MevskyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mevsky *MevskyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mevsky.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mevsky *MevskyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mevsky.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mevsky *MevskyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mevsky.Contract.contract.Transact(opts, method, params...)
}

// MinBounty is a free data retrieval call binding the contract method 0x89b8db55.
//
// Solidity: function minBounty() view returns(uint256)
func (_Mevsky *MevskyCaller) MinBounty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mevsky.contract.Call(opts, &out, "minBounty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBounty is a free data retrieval call binding the contract method 0x89b8db55.
//
// Solidity: function minBounty() view returns(uint256)
func (_Mevsky *MevskySession) MinBounty() (*big.Int, error) {
	return _Mevsky.Contract.MinBounty(&_Mevsky.CallOpts)
}

// MinBounty is a free data retrieval call binding the contract method 0x89b8db55.
//
// Solidity: function minBounty() view returns(uint256)
func (_Mevsky *MevskyCallerSession) MinBounty() (*big.Int, error) {
	return _Mevsky.Contract.MinBounty(&_Mevsky.CallOpts)
}

// On is a free data retrieval call binding the contract method 0x67b7c034.
//
// Solidity: function on() view returns(bool)
func (_Mevsky *MevskyCaller) On(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Mevsky.contract.Call(opts, &out, "on")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// On is a free data retrieval call binding the contract method 0x67b7c034.
//
// Solidity: function on() view returns(bool)
func (_Mevsky *MevskySession) On() (bool, error) {
	return _Mevsky.Contract.On(&_Mevsky.CallOpts)
}

// On is a free data retrieval call binding the contract method 0x67b7c034.
//
// Solidity: function on() view returns(bool)
func (_Mevsky *MevskyCallerSession) On() (bool, error) {
	return _Mevsky.Contract.On(&_Mevsky.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mevsky *MevskyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mevsky.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mevsky *MevskySession) Owner() (common.Address, error) {
	return _Mevsky.Contract.Owner(&_Mevsky.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mevsky *MevskyCallerSession) Owner() (common.Address, error) {
	return _Mevsky.Contract.Owner(&_Mevsky.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mevsky *MevskyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mevsky.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mevsky *MevskySession) RenounceOwnership() (*types.Transaction, error) {
	return _Mevsky.Contract.RenounceOwnership(&_Mevsky.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mevsky *MevskyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mevsky.Contract.RenounceOwnership(&_Mevsky.TransactOpts)
}

// SetMinBounty is a paid mutator transaction binding the contract method 0x3b58f49a.
//
// Solidity: function setMinBounty(uint256 newMinBounty) returns()
func (_Mevsky *MevskyTransactor) SetMinBounty(opts *bind.TransactOpts, newMinBounty *big.Int) (*types.Transaction, error) {
	return _Mevsky.contract.Transact(opts, "setMinBounty", newMinBounty)
}

// SetMinBounty is a paid mutator transaction binding the contract method 0x3b58f49a.
//
// Solidity: function setMinBounty(uint256 newMinBounty) returns()
func (_Mevsky *MevskySession) SetMinBounty(newMinBounty *big.Int) (*types.Transaction, error) {
	return _Mevsky.Contract.SetMinBounty(&_Mevsky.TransactOpts, newMinBounty)
}

// SetMinBounty is a paid mutator transaction binding the contract method 0x3b58f49a.
//
// Solidity: function setMinBounty(uint256 newMinBounty) returns()
func (_Mevsky *MevskyTransactorSession) SetMinBounty(newMinBounty *big.Int) (*types.Transaction, error) {
	return _Mevsky.Contract.SetMinBounty(&_Mevsky.TransactOpts, newMinBounty)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mevsky *MevskyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mevsky.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mevsky *MevskySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mevsky.Contract.TransferOwnership(&_Mevsky.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mevsky *MevskyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mevsky.Contract.TransferOwnership(&_Mevsky.TransactOpts, newOwner)
}

// TurnOff is a paid mutator transaction binding the contract method 0xb547a726.
//
// Solidity: function turnOff(address receiver) returns()
func (_Mevsky *MevskyTransactor) TurnOff(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _Mevsky.contract.Transact(opts, "turnOff", receiver)
}

// TurnOff is a paid mutator transaction binding the contract method 0xb547a726.
//
// Solidity: function turnOff(address receiver) returns()
func (_Mevsky *MevskySession) TurnOff(receiver common.Address) (*types.Transaction, error) {
	return _Mevsky.Contract.TurnOff(&_Mevsky.TransactOpts, receiver)
}

// TurnOff is a paid mutator transaction binding the contract method 0xb547a726.
//
// Solidity: function turnOff(address receiver) returns()
func (_Mevsky *MevskyTransactorSession) TurnOff(receiver common.Address) (*types.Transaction, error) {
	return _Mevsky.Contract.TurnOff(&_Mevsky.TransactOpts, receiver)
}

// TurnOn is a paid mutator transaction binding the contract method 0xba33ed7d.
//
// Solidity: function turnOn() payable returns()
func (_Mevsky *MevskyTransactor) TurnOn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mevsky.contract.Transact(opts, "turnOn")
}

// TurnOn is a paid mutator transaction binding the contract method 0xba33ed7d.
//
// Solidity: function turnOn() payable returns()
func (_Mevsky *MevskySession) TurnOn() (*types.Transaction, error) {
	return _Mevsky.Contract.TurnOn(&_Mevsky.TransactOpts)
}

// TurnOn is a paid mutator transaction binding the contract method 0xba33ed7d.
//
// Solidity: function turnOn() payable returns()
func (_Mevsky *MevskyTransactorSession) TurnOn() (*types.Transaction, error) {
	return _Mevsky.Contract.TurnOn(&_Mevsky.TransactOpts)
}

// MevskyMinBountySetIterator is returned from FilterMinBountySet and is used to iterate over the raw logs and unpacked data for MinBountySet events raised by the Mevsky contract.
type MevskyMinBountySetIterator struct {
	Event *MevskyMinBountySet // Event containing the contract specifics and raw log

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
func (it *MevskyMinBountySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MevskyMinBountySet)
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
		it.Event = new(MevskyMinBountySet)
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
func (it *MevskyMinBountySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MevskyMinBountySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MevskyMinBountySet represents a MinBountySet event raised by the Mevsky contract.
type MevskyMinBountySet struct {
	MinBounty *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinBountySet is a free log retrieval operation binding the contract event 0xa5324cc7956c4d8e6e433de4d6a5e735c98fa2f3af113a918f33a0f94e22a967.
//
// Solidity: event MinBountySet(uint256 minBounty)
func (_Mevsky *MevskyFilterer) FilterMinBountySet(opts *bind.FilterOpts) (*MevskyMinBountySetIterator, error) {

	logs, sub, err := _Mevsky.contract.FilterLogs(opts, "MinBountySet")
	if err != nil {
		return nil, err
	}
	return &MevskyMinBountySetIterator{contract: _Mevsky.contract, event: "MinBountySet", logs: logs, sub: sub}, nil
}

// WatchMinBountySet is a free log subscription operation binding the contract event 0xa5324cc7956c4d8e6e433de4d6a5e735c98fa2f3af113a918f33a0f94e22a967.
//
// Solidity: event MinBountySet(uint256 minBounty)
func (_Mevsky *MevskyFilterer) WatchMinBountySet(opts *bind.WatchOpts, sink chan<- *MevskyMinBountySet) (event.Subscription, error) {

	logs, sub, err := _Mevsky.contract.WatchLogs(opts, "MinBountySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MevskyMinBountySet)
				if err := _Mevsky.contract.UnpackLog(event, "MinBountySet", log); err != nil {
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

// ParseMinBountySet is a log parse operation binding the contract event 0xa5324cc7956c4d8e6e433de4d6a5e735c98fa2f3af113a918f33a0f94e22a967.
//
// Solidity: event MinBountySet(uint256 minBounty)
func (_Mevsky *MevskyFilterer) ParseMinBountySet(log types.Log) (*MevskyMinBountySet, error) {
	event := new(MevskyMinBountySet)
	if err := _Mevsky.contract.UnpackLog(event, "MinBountySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MevskyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mevsky contract.
type MevskyOwnershipTransferredIterator struct {
	Event *MevskyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MevskyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MevskyOwnershipTransferred)
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
		it.Event = new(MevskyOwnershipTransferred)
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
func (it *MevskyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MevskyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MevskyOwnershipTransferred represents a OwnershipTransferred event raised by the Mevsky contract.
type MevskyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mevsky *MevskyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MevskyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mevsky.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MevskyOwnershipTransferredIterator{contract: _Mevsky.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mevsky *MevskyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MevskyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mevsky.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MevskyOwnershipTransferred)
				if err := _Mevsky.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Mevsky *MevskyFilterer) ParseOwnershipTransferred(log types.Log) (*MevskyOwnershipTransferred, error) {
	event := new(MevskyOwnershipTransferred)
	if err := _Mevsky.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MevskyTurnedOffIterator is returned from FilterTurnedOff and is used to iterate over the raw logs and unpacked data for TurnedOff events raised by the Mevsky contract.
type MevskyTurnedOffIterator struct {
	Event *MevskyTurnedOff // Event containing the contract specifics and raw log

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
func (it *MevskyTurnedOffIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MevskyTurnedOff)
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
		it.Event = new(MevskyTurnedOff)
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
func (it *MevskyTurnedOffIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MevskyTurnedOffIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MevskyTurnedOff represents a TurnedOff event raised by the Mevsky contract.
type MevskyTurnedOff struct {
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTurnedOff is a free log retrieval operation binding the contract event 0x5bd0ded7f3a6710e01c75062c0eb193fc026930098b14f9fd962e1317067a157.
//
// Solidity: event TurnedOff(address receiver)
func (_Mevsky *MevskyFilterer) FilterTurnedOff(opts *bind.FilterOpts) (*MevskyTurnedOffIterator, error) {

	logs, sub, err := _Mevsky.contract.FilterLogs(opts, "TurnedOff")
	if err != nil {
		return nil, err
	}
	return &MevskyTurnedOffIterator{contract: _Mevsky.contract, event: "TurnedOff", logs: logs, sub: sub}, nil
}

// WatchTurnedOff is a free log subscription operation binding the contract event 0x5bd0ded7f3a6710e01c75062c0eb193fc026930098b14f9fd962e1317067a157.
//
// Solidity: event TurnedOff(address receiver)
func (_Mevsky *MevskyFilterer) WatchTurnedOff(opts *bind.WatchOpts, sink chan<- *MevskyTurnedOff) (event.Subscription, error) {

	logs, sub, err := _Mevsky.contract.WatchLogs(opts, "TurnedOff")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MevskyTurnedOff)
				if err := _Mevsky.contract.UnpackLog(event, "TurnedOff", log); err != nil {
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

// ParseTurnedOff is a log parse operation binding the contract event 0x5bd0ded7f3a6710e01c75062c0eb193fc026930098b14f9fd962e1317067a157.
//
// Solidity: event TurnedOff(address receiver)
func (_Mevsky *MevskyFilterer) ParseTurnedOff(log types.Log) (*MevskyTurnedOff, error) {
	event := new(MevskyTurnedOff)
	if err := _Mevsky.contract.UnpackLog(event, "TurnedOff", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MevskyTurnedOnIterator is returned from FilterTurnedOn and is used to iterate over the raw logs and unpacked data for TurnedOn events raised by the Mevsky contract.
type MevskyTurnedOnIterator struct {
	Event *MevskyTurnedOn // Event containing the contract specifics and raw log

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
func (it *MevskyTurnedOnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MevskyTurnedOn)
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
		it.Event = new(MevskyTurnedOn)
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
func (it *MevskyTurnedOnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MevskyTurnedOnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MevskyTurnedOn represents a TurnedOn event raised by the Mevsky contract.
type MevskyTurnedOn struct {
	Bounty *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTurnedOn is a free log retrieval operation binding the contract event 0x1e4ee824371beb9d98ccaa8792e4e4e73012b547e6b63e34e191fa970fb8d088.
//
// Solidity: event TurnedOn(uint256 bounty)
func (_Mevsky *MevskyFilterer) FilterTurnedOn(opts *bind.FilterOpts) (*MevskyTurnedOnIterator, error) {

	logs, sub, err := _Mevsky.contract.FilterLogs(opts, "TurnedOn")
	if err != nil {
		return nil, err
	}
	return &MevskyTurnedOnIterator{contract: _Mevsky.contract, event: "TurnedOn", logs: logs, sub: sub}, nil
}

// WatchTurnedOn is a free log subscription operation binding the contract event 0x1e4ee824371beb9d98ccaa8792e4e4e73012b547e6b63e34e191fa970fb8d088.
//
// Solidity: event TurnedOn(uint256 bounty)
func (_Mevsky *MevskyFilterer) WatchTurnedOn(opts *bind.WatchOpts, sink chan<- *MevskyTurnedOn) (event.Subscription, error) {

	logs, sub, err := _Mevsky.contract.WatchLogs(opts, "TurnedOn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MevskyTurnedOn)
				if err := _Mevsky.contract.UnpackLog(event, "TurnedOn", log); err != nil {
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

// ParseTurnedOn is a log parse operation binding the contract event 0x1e4ee824371beb9d98ccaa8792e4e4e73012b547e6b63e34e191fa970fb8d088.
//
// Solidity: event TurnedOn(uint256 bounty)
func (_Mevsky *MevskyFilterer) ParseTurnedOn(log types.Log) (*MevskyTurnedOn, error) {
	event := new(MevskyTurnedOn)
	if err := _Mevsky.contract.UnpackLog(event, "TurnedOn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
