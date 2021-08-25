// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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

// TargetCoreExMetaData contains all meta data concerning the TargetCoreEx contract.
var TargetCoreExMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_RequestID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_result\",\"type\":\"bytes\"}],\"name\":\"CrossChainResponseSent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_RequestID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_endpointAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callData\",\"type\":\"bytes\"}],\"name\":\"callService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506105ec806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806375d219fa14610030575b600080fd5b61004a60048036038101906100459190610236565b61004c565b005b826000151560008083815260200190815260200160002060009054906101000a900460ff161515146100b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100aa90610388565b60405180910390fd5b6000808473ffffffffffffffffffffffffffffffffffffffff16846040516100db9190610341565b6000604051808303816000865af19150503d8060008114610118576040519150601f19603f3d011682016040523d82523d6000602084013e61011d565b606091505b5091509150600115158215151415610194577fdd9b7291a69fc50c9dbd1ee123efce8f2f5086ebe46cea5855a5331b2eb6df4c8682604051610160929190610358565b60405180910390a1600160008088815260200190815260200160002060006101000a81548160ff0219169083151502179055505b505050505050565b60006101af6101aa846103cd565b6103a8565b9050828152602081018484840111156101cb576101ca610519565b5b6101d6848285610472565b509392505050565b6000813590506101ed81610588565b92915050565b6000813590506102028161059f565b92915050565b600082601f83011261021d5761021c610514565b5b813561022d84826020860161019c565b91505092915050565b60008060006060848603121561024f5761024e610523565b5b600061025d868287016101f3565b935050602061026e868287016101de565b925050604084013567ffffffffffffffff81111561028f5761028e61051e565b5b61029b86828701610208565b9150509250925092565b6102ae81610448565b82525050565b60006102bf826103fe565b6102c98185610409565b93506102d9818560208601610481565b6102e281610528565b840191505092915050565b60006102f8826103fe565b610302818561041a565b9350610312818560208601610481565b80840191505092915050565b600061032b602383610425565b915061033682610539565b604082019050919050565b600061034d82846102ed565b915081905092915050565b600060408201905061036d60008301856102a5565b818103602083015261037f81846102b4565b90509392505050565b600060208201905081810360008301526103a18161031e565b9050919050565b60006103b26103c3565b90506103be82826104b4565b919050565b6000604051905090565b600067ffffffffffffffff8211156103e8576103e76104e5565b5b6103f182610528565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600061044182610452565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b82818337600083830152505050565b60005b8381101561049f578082015181840152602081019050610484565b838111156104ae576000848401525b50505050565b6104bd82610528565b810181811067ffffffffffffffff821117156104dc576104db6104e5565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f6953657276696365436f726545783a206475706c69636174656420726571756560008201527f7374210000000000000000000000000000000000000000000000000000000000602082015250565b61059181610436565b811461059c57600080fd5b50565b6105a881610448565b81146105b357600080fd5b5056fea2646970667358221220ac341e29c08e67a3a81417011b503b8369d152b35178e16f3cf690fc0def7ad364736f6c63430008070033",
}

// TargetCoreExABI is the input ABI used to generate the binding from.
// Deprecated: Use TargetCoreExMetaData.ABI instead.
var TargetCoreExABI = TargetCoreExMetaData.ABI

// TargetCoreExBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TargetCoreExMetaData.Bin instead.
var TargetCoreExBin = TargetCoreExMetaData.Bin

// DeployTargetCoreEx deploys a new Ethereum contract, binding an instance of TargetCoreEx to it.
func DeployTargetCoreEx(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TargetCoreEx, error) {
	parsed, err := TargetCoreExMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TargetCoreExBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TargetCoreEx{TargetCoreExCaller: TargetCoreExCaller{contract: contract}, TargetCoreExTransactor: TargetCoreExTransactor{contract: contract}, TargetCoreExFilterer: TargetCoreExFilterer{contract: contract}}, nil
}

// TargetCoreEx is an auto generated Go binding around an Ethereum contract.
type TargetCoreEx struct {
	TargetCoreExCaller     // Read-only binding to the contract
	TargetCoreExTransactor // Write-only binding to the contract
	TargetCoreExFilterer   // Log filterer for contract events
}

// TargetCoreExCaller is an auto generated read-only Go binding around an Ethereum contract.
type TargetCoreExCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TargetCoreExTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TargetCoreExTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TargetCoreExFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TargetCoreExFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TargetCoreExSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TargetCoreExSession struct {
	Contract     *TargetCoreEx     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TargetCoreExCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TargetCoreExCallerSession struct {
	Contract *TargetCoreExCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TargetCoreExTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TargetCoreExTransactorSession struct {
	Contract     *TargetCoreExTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TargetCoreExRaw is an auto generated low-level Go binding around an Ethereum contract.
type TargetCoreExRaw struct {
	Contract *TargetCoreEx // Generic contract binding to access the raw methods on
}

// TargetCoreExCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TargetCoreExCallerRaw struct {
	Contract *TargetCoreExCaller // Generic read-only contract binding to access the raw methods on
}

// TargetCoreExTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TargetCoreExTransactorRaw struct {
	Contract *TargetCoreExTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTargetCoreEx creates a new instance of TargetCoreEx, bound to a specific deployed contract.
func NewTargetCoreEx(address common.Address, backend bind.ContractBackend) (*TargetCoreEx, error) {
	contract, err := bindTargetCoreEx(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TargetCoreEx{TargetCoreExCaller: TargetCoreExCaller{contract: contract}, TargetCoreExTransactor: TargetCoreExTransactor{contract: contract}, TargetCoreExFilterer: TargetCoreExFilterer{contract: contract}}, nil
}

// NewTargetCoreExCaller creates a new read-only instance of TargetCoreEx, bound to a specific deployed contract.
func NewTargetCoreExCaller(address common.Address, caller bind.ContractCaller) (*TargetCoreExCaller, error) {
	contract, err := bindTargetCoreEx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TargetCoreExCaller{contract: contract}, nil
}

// NewTargetCoreExTransactor creates a new write-only instance of TargetCoreEx, bound to a specific deployed contract.
func NewTargetCoreExTransactor(address common.Address, transactor bind.ContractTransactor) (*TargetCoreExTransactor, error) {
	contract, err := bindTargetCoreEx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TargetCoreExTransactor{contract: contract}, nil
}

// NewTargetCoreExFilterer creates a new log filterer instance of TargetCoreEx, bound to a specific deployed contract.
func NewTargetCoreExFilterer(address common.Address, filterer bind.ContractFilterer) (*TargetCoreExFilterer, error) {
	contract, err := bindTargetCoreEx(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TargetCoreExFilterer{contract: contract}, nil
}

// bindTargetCoreEx binds a generic wrapper to an already deployed contract.
func bindTargetCoreEx(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TargetCoreExABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TargetCoreEx *TargetCoreExRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TargetCoreEx.Contract.TargetCoreExCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TargetCoreEx *TargetCoreExRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TargetCoreEx.Contract.TargetCoreExTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TargetCoreEx *TargetCoreExRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TargetCoreEx.Contract.TargetCoreExTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TargetCoreEx *TargetCoreExCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TargetCoreEx.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TargetCoreEx *TargetCoreExTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TargetCoreEx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TargetCoreEx *TargetCoreExTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TargetCoreEx.Contract.contract.Transact(opts, method, params...)
}

// CallService is a paid mutator transaction binding the contract method 0x75d219fa.
//
// Solidity: function callService(bytes32 _RequestID, address _endpointAddress, bytes _callData) returns()
func (_TargetCoreEx *TargetCoreExTransactor) CallService(opts *bind.TransactOpts, _RequestID [32]byte, _endpointAddress common.Address, _callData []byte) (*types.Transaction, error) {
	return _TargetCoreEx.contract.Transact(opts, "callService", _RequestID, _endpointAddress, _callData)
}

// CallService is a paid mutator transaction binding the contract method 0x75d219fa.
//
// Solidity: function callService(bytes32 _RequestID, address _endpointAddress, bytes _callData) returns()
func (_TargetCoreEx *TargetCoreExSession) CallService(_RequestID [32]byte, _endpointAddress common.Address, _callData []byte) (*types.Transaction, error) {
	return _TargetCoreEx.Contract.CallService(&_TargetCoreEx.TransactOpts, _RequestID, _endpointAddress, _callData)
}

// CallService is a paid mutator transaction binding the contract method 0x75d219fa.
//
// Solidity: function callService(bytes32 _RequestID, address _endpointAddress, bytes _callData) returns()
func (_TargetCoreEx *TargetCoreExTransactorSession) CallService(_RequestID [32]byte, _endpointAddress common.Address, _callData []byte) (*types.Transaction, error) {
	return _TargetCoreEx.Contract.CallService(&_TargetCoreEx.TransactOpts, _RequestID, _endpointAddress, _callData)
}

// TargetCoreExCrossChainResponseSentIterator is returned from FilterCrossChainResponseSent and is used to iterate over the raw logs and unpacked data for CrossChainResponseSent events raised by the TargetCoreEx contract.
type TargetCoreExCrossChainResponseSentIterator struct {
	Event *TargetCoreExCrossChainResponseSent // Event containing the contract specifics and raw log

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
func (it *TargetCoreExCrossChainResponseSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TargetCoreExCrossChainResponseSent)
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
		it.Event = new(TargetCoreExCrossChainResponseSent)
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
func (it *TargetCoreExCrossChainResponseSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TargetCoreExCrossChainResponseSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TargetCoreExCrossChainResponseSent represents a CrossChainResponseSent event raised by the TargetCoreEx contract.
type TargetCoreExCrossChainResponseSent struct {
	RequestID [32]byte
	Result    []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCrossChainResponseSent is a free log retrieval operation binding the contract event 0xdd9b7291a69fc50c9dbd1ee123efce8f2f5086ebe46cea5855a5331b2eb6df4c.
//
// Solidity: event CrossChainResponseSent(bytes32 _RequestID, bytes _result)
func (_TargetCoreEx *TargetCoreExFilterer) FilterCrossChainResponseSent(opts *bind.FilterOpts) (*TargetCoreExCrossChainResponseSentIterator, error) {

	logs, sub, err := _TargetCoreEx.contract.FilterLogs(opts, "CrossChainResponseSent")
	if err != nil {
		return nil, err
	}
	return &TargetCoreExCrossChainResponseSentIterator{contract: _TargetCoreEx.contract, event: "CrossChainResponseSent", logs: logs, sub: sub}, nil
}

// WatchCrossChainResponseSent is a free log subscription operation binding the contract event 0xdd9b7291a69fc50c9dbd1ee123efce8f2f5086ebe46cea5855a5331b2eb6df4c.
//
// Solidity: event CrossChainResponseSent(bytes32 _RequestID, bytes _result)
func (_TargetCoreEx *TargetCoreExFilterer) WatchCrossChainResponseSent(opts *bind.WatchOpts, sink chan<- *TargetCoreExCrossChainResponseSent) (event.Subscription, error) {

	logs, sub, err := _TargetCoreEx.contract.WatchLogs(opts, "CrossChainResponseSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TargetCoreExCrossChainResponseSent)
				if err := _TargetCoreEx.contract.UnpackLog(event, "CrossChainResponseSent", log); err != nil {
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

// ParseCrossChainResponseSent is a log parse operation binding the contract event 0xdd9b7291a69fc50c9dbd1ee123efce8f2f5086ebe46cea5855a5331b2eb6df4c.
//
// Solidity: event CrossChainResponseSent(bytes32 _RequestID, bytes _result)
func (_TargetCoreEx *TargetCoreExFilterer) ParseCrossChainResponseSent(log types.Log) (*TargetCoreExCrossChainResponseSent, error) {
	event := new(TargetCoreExCrossChainResponseSent)
	if err := _TargetCoreEx.contract.UnpackLog(event, "CrossChainResponseSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
