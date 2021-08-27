// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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

// TargetCoreExABI is the input ABI used to generate the binding from.
const TargetCoreExABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_RequestID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_result\",\"type\":\"bytes\"}],\"name\":\"CrossChainResponseSent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_RequestID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_endpointAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callData\",\"type\":\"bytes\"}],\"name\":\"callService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TargetCoreExBin is the compiled bytecode used for deploying new contracts.
var TargetCoreExBin = "0x608060405234801561001057600080fd5b5061032a806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806375d219fa14610030575b600080fd5b6101136004803603606081101561004657600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561008d57600080fd5b82018360208201111561009f57600080fd5b803590602001918460018302840111640100000000831117156100c157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610115565b005b826000151560008083815260200190815260200160002060009054906101000a900460ff16151514610192576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806102d26023913960400191505060405180910390fd5b60008251905060606000602085016000808583348b5af1915081600181146101c157600081146101e8576101ed565b3d6040519450601f19601f6020830101168501604052808552806000602087013e506101ed565b600080fd5b505060018114156102c8577fdd9b7291a69fc50c9dbd1ee123efce8f2f5086ebe46cea5855a5331b2eb6df4c87836040518083815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610261578082015181840152602081019050610246565b50505050905090810190601f16801561028e5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a1600160008089815260200190815260200160002060006101000a81548160ff0219169083151502179055505b5050505050505056fe6953657276696365436f726545783a206475706c696361746564207265717565737421a26469706673582212208f366c84f2b149964686de45b6ed4c0bd95634a5e3cd72e0cb5b749a5fac37c664736f6c634300060c0033"

// DeployTargetCoreEx deploys a new Ethereum contract, binding an instance of TargetCoreEx to it.
func DeployTargetCoreEx(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TargetCoreEx, error) {
	parsed, err := abi.JSON(strings.NewReader(TargetCoreExABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TargetCoreExBin), backend)
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
func (_TargetCoreEx *TargetCoreExRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_TargetCoreEx *TargetCoreExCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
	return event, nil
}
