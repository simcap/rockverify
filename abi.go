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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RockVerifyABI is the input ABI used to generate the binding from.
const RockVerifyABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"urlShasum\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"fileShasum\",\"type\":\"bytes32\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"downloadables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"urlShasum\",\"type\":\"bytes32\"}],\"name\":\"lookup\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"urlShasum\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"fileShasum\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RockVerifyBin is the compiled bytecode used for deploying new contracts.
var RockVerifyBin = "0x608060405234801561001057600080fd5b50610165806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80632f92673214610046578063c5d76f391461006b578063f39ec1f71461009a575b600080fd5b6100696004803603604081101561005c57600080fd5b50803590602001356100b7565b005b6100886004803603602081101561008157600080fd5b503561010b565b60408051918252519081900360200190f35b610088600480360360208110156100b057600080fd5b503561011d565b60008281526020818152604091829020839055815133815290810184905280820183905290517f6e7788094390b17970960cf310096f0de10956a8bcc4ec0cabe0bbfeb24d06f79181900360600190a15050565b60006020819052908152604090205481565b6000908152602081905260409020549056fea2646970667358221220697016fdb5c1e96023e6224fc119a64ada6ccd6f0c97d78d6442d7e6c2813ddf64736f6c63430006010033"

// DeployRockVerify deploys a new Ethereum contract, binding an instance of RockVerify to it.
func DeployRockVerify(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RockVerify, error) {
	parsed, err := abi.JSON(strings.NewReader(RockVerifyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RockVerifyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RockVerify{RockVerifyCaller: RockVerifyCaller{contract: contract}, RockVerifyTransactor: RockVerifyTransactor{contract: contract}, RockVerifyFilterer: RockVerifyFilterer{contract: contract}}, nil
}

// RockVerify is an auto generated Go binding around an Ethereum contract.
type RockVerify struct {
	RockVerifyCaller     // Read-only binding to the contract
	RockVerifyTransactor // Write-only binding to the contract
	RockVerifyFilterer   // Log filterer for contract events
}

// RockVerifyCaller is an auto generated read-only Go binding around an Ethereum contract.
type RockVerifyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RockVerifyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RockVerifyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RockVerifyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RockVerifyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RockVerifySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RockVerifySession struct {
	Contract     *RockVerify       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RockVerifyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RockVerifyCallerSession struct {
	Contract *RockVerifyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RockVerifyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RockVerifyTransactorSession struct {
	Contract     *RockVerifyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RockVerifyRaw is an auto generated low-level Go binding around an Ethereum contract.
type RockVerifyRaw struct {
	Contract *RockVerify // Generic contract binding to access the raw methods on
}

// RockVerifyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RockVerifyCallerRaw struct {
	Contract *RockVerifyCaller // Generic read-only contract binding to access the raw methods on
}

// RockVerifyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RockVerifyTransactorRaw struct {
	Contract *RockVerifyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRockVerify creates a new instance of RockVerify, bound to a specific deployed contract.
func NewRockVerify(address common.Address, backend bind.ContractBackend) (*RockVerify, error) {
	contract, err := bindRockVerify(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RockVerify{RockVerifyCaller: RockVerifyCaller{contract: contract}, RockVerifyTransactor: RockVerifyTransactor{contract: contract}, RockVerifyFilterer: RockVerifyFilterer{contract: contract}}, nil
}

// NewRockVerifyCaller creates a new read-only instance of RockVerify, bound to a specific deployed contract.
func NewRockVerifyCaller(address common.Address, caller bind.ContractCaller) (*RockVerifyCaller, error) {
	contract, err := bindRockVerify(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RockVerifyCaller{contract: contract}, nil
}

// NewRockVerifyTransactor creates a new write-only instance of RockVerify, bound to a specific deployed contract.
func NewRockVerifyTransactor(address common.Address, transactor bind.ContractTransactor) (*RockVerifyTransactor, error) {
	contract, err := bindRockVerify(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RockVerifyTransactor{contract: contract}, nil
}

// NewRockVerifyFilterer creates a new log filterer instance of RockVerify, bound to a specific deployed contract.
func NewRockVerifyFilterer(address common.Address, filterer bind.ContractFilterer) (*RockVerifyFilterer, error) {
	contract, err := bindRockVerify(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RockVerifyFilterer{contract: contract}, nil
}

// bindRockVerify binds a generic wrapper to an already deployed contract.
func bindRockVerify(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RockVerifyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RockVerify *RockVerifyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RockVerify.Contract.RockVerifyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RockVerify *RockVerifyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RockVerify.Contract.RockVerifyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RockVerify *RockVerifyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RockVerify.Contract.RockVerifyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RockVerify *RockVerifyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RockVerify.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RockVerify *RockVerifyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RockVerify.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RockVerify *RockVerifyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RockVerify.Contract.contract.Transact(opts, method, params...)
}

// Downloadables is a free data retrieval call binding the contract method 0xc5d76f39.
//
// Solidity: function downloadables(bytes32 ) constant returns(bytes32)
func (_RockVerify *RockVerifyCaller) Downloadables(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RockVerify.contract.Call(opts, out, "downloadables", arg0)
	return *ret0, err
}

// Downloadables is a free data retrieval call binding the contract method 0xc5d76f39.
//
// Solidity: function downloadables(bytes32 ) constant returns(bytes32)
func (_RockVerify *RockVerifySession) Downloadables(arg0 [32]byte) ([32]byte, error) {
	return _RockVerify.Contract.Downloadables(&_RockVerify.CallOpts, arg0)
}

// Downloadables is a free data retrieval call binding the contract method 0xc5d76f39.
//
// Solidity: function downloadables(bytes32 ) constant returns(bytes32)
func (_RockVerify *RockVerifyCallerSession) Downloadables(arg0 [32]byte) ([32]byte, error) {
	return _RockVerify.Contract.Downloadables(&_RockVerify.CallOpts, arg0)
}

// Lookup is a free data retrieval call binding the contract method 0xf39ec1f7.
//
// Solidity: function lookup(bytes32 urlShasum) constant returns(bytes32)
func (_RockVerify *RockVerifyCaller) Lookup(opts *bind.CallOpts, urlShasum [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RockVerify.contract.Call(opts, out, "lookup", urlShasum)
	return *ret0, err
}

// Lookup is a free data retrieval call binding the contract method 0xf39ec1f7.
//
// Solidity: function lookup(bytes32 urlShasum) constant returns(bytes32)
func (_RockVerify *RockVerifySession) Lookup(urlShasum [32]byte) ([32]byte, error) {
	return _RockVerify.Contract.Lookup(&_RockVerify.CallOpts, urlShasum)
}

// Lookup is a free data retrieval call binding the contract method 0xf39ec1f7.
//
// Solidity: function lookup(bytes32 urlShasum) constant returns(bytes32)
func (_RockVerify *RockVerifyCallerSession) Lookup(urlShasum [32]byte) ([32]byte, error) {
	return _RockVerify.Contract.Lookup(&_RockVerify.CallOpts, urlShasum)
}

// Register is a paid mutator transaction binding the contract method 0x2f926732.
//
// Solidity: function register(bytes32 urlShasum, bytes32 fileShasum) returns()
func (_RockVerify *RockVerifyTransactor) Register(opts *bind.TransactOpts, urlShasum [32]byte, fileShasum [32]byte) (*types.Transaction, error) {
	return _RockVerify.contract.Transact(opts, "register", urlShasum, fileShasum)
}

// Register is a paid mutator transaction binding the contract method 0x2f926732.
//
// Solidity: function register(bytes32 urlShasum, bytes32 fileShasum) returns()
func (_RockVerify *RockVerifySession) Register(urlShasum [32]byte, fileShasum [32]byte) (*types.Transaction, error) {
	return _RockVerify.Contract.Register(&_RockVerify.TransactOpts, urlShasum, fileShasum)
}

// Register is a paid mutator transaction binding the contract method 0x2f926732.
//
// Solidity: function register(bytes32 urlShasum, bytes32 fileShasum) returns()
func (_RockVerify *RockVerifyTransactorSession) Register(urlShasum [32]byte, fileShasum [32]byte) (*types.Transaction, error) {
	return _RockVerify.Contract.Register(&_RockVerify.TransactOpts, urlShasum, fileShasum)
}

// RockVerifyRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the RockVerify contract.
type RockVerifyRegisteredIterator struct {
	Event *RockVerifyRegistered // Event containing the contract specifics and raw log

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
func (it *RockVerifyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RockVerifyRegistered)
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
		it.Event = new(RockVerifyRegistered)
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
func (it *RockVerifyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RockVerifyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RockVerifyRegistered represents a Registered event raised by the RockVerify contract.
type RockVerifyRegistered struct {
	From       common.Address
	UrlShasum  [32]byte
	FileShasum [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x6e7788094390b17970960cf310096f0de10956a8bcc4ec0cabe0bbfeb24d06f7.
//
// Solidity: event Registered(address from, bytes32 urlShasum, bytes32 fileShasum)
func (_RockVerify *RockVerifyFilterer) FilterRegistered(opts *bind.FilterOpts) (*RockVerifyRegisteredIterator, error) {

	logs, sub, err := _RockVerify.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &RockVerifyRegisteredIterator{contract: _RockVerify.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x6e7788094390b17970960cf310096f0de10956a8bcc4ec0cabe0bbfeb24d06f7.
//
// Solidity: event Registered(address from, bytes32 urlShasum, bytes32 fileShasum)
func (_RockVerify *RockVerifyFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *RockVerifyRegistered) (event.Subscription, error) {

	logs, sub, err := _RockVerify.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RockVerifyRegistered)
				if err := _RockVerify.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0x6e7788094390b17970960cf310096f0de10956a8bcc4ec0cabe0bbfeb24d06f7.
//
// Solidity: event Registered(address from, bytes32 urlShasum, bytes32 fileShasum)
func (_RockVerify *RockVerifyFilterer) ParseRegistered(log types.Log) (*RockVerifyRegistered, error) {
	event := new(RockVerifyRegistered)
	if err := _RockVerify.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	return event, nil
}
