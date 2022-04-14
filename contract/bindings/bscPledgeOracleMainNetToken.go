// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"math/big"
	"strings"
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

// BscPledgeOracleMainnetTokenMetaData contains all meta data concerning the BscPledgeOracleMainnetToken contract.
var BscPledgeOracleMainnetTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getAssetsAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"assets\",\"type\":\"uint256[]\"}],\"name\":\"getPrices\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlying\",\"type\":\"uint256\"}],\"name\":\"getUnderlyingAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlying\",\"type\":\"uint256\"}],\"name\":\"getUnderlyingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aggergator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"setAssetsAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDecimals\",\"type\":\"uint256\"}],\"name\":\"setDecimals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"assets\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlying\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"aggergator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"setUnderlyingAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlying\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setUnderlyingPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600160045534801561001557600080fd5b50600061002061006f565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350610073565b3390565b610d7f806100826000396000f3fe608060405234801561001057600080fd5b50600436106100e95760003560e01c80638c8885c81161008c578063cd9ffa0b11610066578063cd9ffa0b1461034b578063d05eaae014610381578063da663257146104a8578063f2fde38b146104c5576100e9565b80638c8885c8146102e45780638da5cb5b14610301578063b889a98914610325576100e9565b8063434b33ac116100c8578063434b33ac14610247578063715018a61461027957806375e443aa1461028157806383532667146102c1576100e9565b8062e4768b146100ee57806309cb3a4e1461011c57806341976e091461020f575b600080fd5b61011a6004803603604081101561010457600080fd5b506001600160a01b0381351690602001356104eb565b005b6101bf6004803603602081101561013257600080fd5b81019060208101813564010000000081111561014d57600080fd5b82018360208201111561015f57600080fd5b8035906020019184602083028401116401000000008311171561018157600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610569945050505050565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101fb5781810151838201526020016101e3565b505050509050019250505060405180910390f35b6102356004803603602081101561022557600080fd5b50356001600160a01b0316610603565b60408051918252519081900360200190f35b61011a6004803603606081101561025d57600080fd5b508035906001600160a01b03602082013516906040013561061d565b61011a610709565b61029e6004803603602081101561029757600080fd5b50356107b5565b604080516001600160a01b03909316835260208301919091528051918290030190f35b61011a600480360360408110156102d757600080fd5b50803590602001356107de565b61011a600480360360208110156102fa57600080fd5b50356108a3565b61030961090a565b604080516001600160a01b039092168252519081900360200190f35b61029e6004803603602081101561033b57600080fd5b50356001600160a01b0316610919565b61011a6004803603606081101561036157600080fd5b506001600160a01b03813581169160208101359091169060400135610944565b61011a6004803603604081101561039757600080fd5b8101906020810181356401000000008111156103b257600080fd5b8201836020820111156103c457600080fd5b803590602001918460208302840111640100000000831117156103e657600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929594936020810193503591505064010000000081111561043657600080fd5b82018360208201111561044857600080fd5b8035906020019184602083028401116401000000008311171561046a57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506109e3945050505050565b610235600480360360208110156104be57600080fd5b5035610ac8565b61011a600480360360208110156104db57600080fd5b50356001600160a01b0316610bdb565b6104f3610cdd565b6001600160a01b031661050461090a565b6001600160a01b03161461054d576040805162461bcd60e51b81526020600482018190526024820152600080516020610d08833981519152604482015290519081900360640190fd5b6001600160a01b03909116600090815260036020526040902055565b8051606090818167ffffffffffffffff8111801561058657600080fd5b506040519080825280602002602001820160405280156105b0578160200160208202803683370190505b50905060005b828110156105f9576105da8582815181106105cd57fe5b6020026020010151610ac8565b8282815181106105e657fe5b60209081029190910101526001016105b6565b509150505b919050565b6000610617826001600160a01b0316610ac8565b92915050565b610625610cdd565b6001600160a01b031661063661090a565b6001600160a01b03161461067f576040805162461bcd60e51b81526020600482018190526024820152600080516020610d08833981519152604482015290519081900360640190fd5b600083116106d0576040805162461bcd60e51b8152602060048201526019602482015278756e6465726c79696e672063616e6e6f74206265207a65726f60381b604482015290519081900360640190fd5b600092835260016020908152604080852080546001600160a01b0319166001600160a01b03959095169490941790935560029052912055565b610711610cdd565b6001600160a01b031661072261090a565b6001600160a01b03161461076b576040805162461bcd60e51b81526020600482018190526024820152600080516020610d08833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000908152600160209081526040808320546002909252909120546001600160a01b0390911691565b6107e6610cdd565b6001600160a01b03166107f761090a565b6001600160a01b031614610840576040805162461bcd60e51b81526020600482018190526024820152600080516020610d08833981519152604482015290519081900360640190fd5b60008211610891576040805162461bcd60e51b8152602060048201526019602482015278756e6465726c79696e672063616e6e6f74206265207a65726f60381b604482015290519081900360640190fd5b60009182526003602052604090912055565b6108ab610cdd565b6001600160a01b03166108bc61090a565b6001600160a01b031614610905576040805162461bcd60e51b81526020600482018190526024820152600080516020610d08833981519152604482015290519081900360640190fd5b600455565b6000546001600160a01b031690565b6001600160a01b03908116600090815260016020908152604080832054600290925290912054911691565b61094c610cdd565b6001600160a01b031661095d61090a565b6001600160a01b0316146109a6576040805162461bcd60e51b81526020600482018190526024820152600080516020610d08833981519152604482015290519081900360640190fd5b6001600160a01b03928316600090815260016020908152604080832080546001600160a01b03191695909616949094179094556002909352912055565b6109eb610cdd565b6001600160a01b03166109fc61090a565b6001600160a01b031614610a45576040805162461bcd60e51b81526020600482018190526024820152600080516020610d08833981519152604482015290519081900360640190fd5b8051825114610a855760405162461bcd60e51b8152600401808060200182810382526022815260200180610d286022913960400191505060405180910390fd5b815160005b81811015610ac257828181518110610a9e57fe5b60209081029190910181015160008381526003909252604090912055600101610a8a565b50505050565b6000818152600160205260408120546001600160a01b03168015610bc5576000816001600160a01b031663feaf968c6040518163ffffffff1660e01b815260040160a06040518083038186803b158015610b2157600080fd5b505afa158015610b35573d6000803e3d6000fd5b505050506040513d60a0811015610b4b57600080fd5b50602090810151600086815260029092526040909120549091506012811015610b8c5780601203600a0a6004548381610b8057fe5b040293505050506105fe565b6012811115610bba5780601203600a0a6004548381610ba757fe5b0481610baf57fe5b0493505050506105fe565b6004548281610baf57fe5b50506000818152600360205260409020546105fe565b610be3610cdd565b6001600160a01b0316610bf461090a565b6001600160a01b031614610c3d576040805162461bcd60e51b81526020600482018190526024820152600080516020610d08833981519152604482015290519081900360640190fd5b6001600160a01b038116610c825760405162461bcd60e51b8152600401808060200182810382526026815260200180610ce26026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b339056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572696e7075742061727261797327206c656e67746820617265206e6f7420657175616ca2646970667358221220982b7cfec6b25f4d94b203d539ae44765e503bab1ed99f20cf313cc09f77474c64736f6c634300060c0033",
}

// BscPledgeOracleMainnetTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use BscPledgeOracleMainnetTokenMetaData.ABI instead.
var BscPledgeOracleMainnetTokenABI = BscPledgeOracleMainnetTokenMetaData.ABI

// BscPledgeOracleMainnetTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BscPledgeOracleMainnetTokenMetaData.Bin instead.
var BscPledgeOracleMainnetTokenBin = BscPledgeOracleMainnetTokenMetaData.Bin

// DeployBscPledgeOracleMainnetToken deploys a new Ethereum contract, binding an instance of BscPledgeOracleMainnetToken to it.
func DeployBscPledgeOracleMainnetToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BscPledgeOracleMainnetToken, error) {
	parsed, err := BscPledgeOracleMainnetTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BscPledgeOracleMainnetTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BscPledgeOracleMainnetToken{BscPledgeOracleMainnetTokenCaller: BscPledgeOracleMainnetTokenCaller{contract: contract}, BscPledgeOracleMainnetTokenTransactor: BscPledgeOracleMainnetTokenTransactor{contract: contract}, BscPledgeOracleMainnetTokenFilterer: BscPledgeOracleMainnetTokenFilterer{contract: contract}}, nil
}

// BscPledgeOracleMainnetToken is an auto generated Go binding around an Ethereum contract.
type BscPledgeOracleMainnetToken struct {
	BscPledgeOracleMainnetTokenCaller     // Read-only binding to the contract
	BscPledgeOracleMainnetTokenTransactor // Write-only binding to the contract
	BscPledgeOracleMainnetTokenFilterer   // Log filterer for contract events
}

// BscPledgeOracleMainnetTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BscPledgeOracleMainnetTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscPledgeOracleMainnetTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BscPledgeOracleMainnetTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscPledgeOracleMainnetTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BscPledgeOracleMainnetTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscPledgeOracleMainnetTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BscPledgeOracleMainnetTokenSession struct {
	Contract     *BscPledgeOracleMainnetToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BscPledgeOracleMainnetTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BscPledgeOracleMainnetTokenCallerSession struct {
	Contract *BscPledgeOracleMainnetTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// BscPledgeOracleMainnetTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BscPledgeOracleMainnetTokenTransactorSession struct {
	Contract     *BscPledgeOracleMainnetTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// BscPledgeOracleMainnetTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BscPledgeOracleMainnetTokenRaw struct {
	Contract *BscPledgeOracleMainnetToken // Generic contract binding to access the raw methods on
}

// BscPledgeOracleMainnetTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BscPledgeOracleMainnetTokenCallerRaw struct {
	Contract *BscPledgeOracleMainnetTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BscPledgeOracleMainnetTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BscPledgeOracleMainnetTokenTransactorRaw struct {
	Contract *BscPledgeOracleMainnetTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBscPledgeOracleMainnetToken creates a new instance of BscPledgeOracleMainnetToken, bound to a specific deployed contract.
func NewBscPledgeOracleMainnetToken(address common.Address, backend bind.ContractBackend) (*BscPledgeOracleMainnetToken, error) {
	contract, err := bindBscPledgeOracleMainnetToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleMainnetToken{BscPledgeOracleMainnetTokenCaller: BscPledgeOracleMainnetTokenCaller{contract: contract}, BscPledgeOracleMainnetTokenTransactor: BscPledgeOracleMainnetTokenTransactor{contract: contract}, BscPledgeOracleMainnetTokenFilterer: BscPledgeOracleMainnetTokenFilterer{contract: contract}}, nil
}

// NewBscPledgeOracleMainnetTokenCaller creates a new read-only instance of BscPledgeOracleMainnetToken, bound to a specific deployed contract.
func NewBscPledgeOracleMainnetTokenCaller(address common.Address, caller bind.ContractCaller) (*BscPledgeOracleMainnetTokenCaller, error) {
	contract, err := bindBscPledgeOracleMainnetToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleMainnetTokenCaller{contract: contract}, nil
}

// NewBscPledgeOracleMainnetTokenTransactor creates a new write-only instance of BscPledgeOracleMainnetToken, bound to a specific deployed contract.
func NewBscPledgeOracleMainnetTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BscPledgeOracleMainnetTokenTransactor, error) {
	contract, err := bindBscPledgeOracleMainnetToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleMainnetTokenTransactor{contract: contract}, nil
}

// NewBscPledgeOracleMainnetTokenFilterer creates a new log filterer instance of BscPledgeOracleMainnetToken, bound to a specific deployed contract.
func NewBscPledgeOracleMainnetTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BscPledgeOracleMainnetTokenFilterer, error) {
	contract, err := bindBscPledgeOracleMainnetToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleMainnetTokenFilterer{contract: contract}, nil
}

// bindBscPledgeOracleMainnetToken binds a generic wrapper to an already deployed contract.
func bindBscPledgeOracleMainnetToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BscPledgeOracleMainnetTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscPledgeOracleMainnetToken.Contract.BscPledgeOracleMainnetTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.BscPledgeOracleMainnetTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.BscPledgeOracleMainnetTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscPledgeOracleMainnetToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.contract.Transact(opts, method, params...)
}

// GetAssetsAggregator is a free data retrieval call binding the contract method 0xb889a989.
//
// Solidity: function getAssetsAggregator(address asset) view returns(address, uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCaller) GetAssetsAggregator(opts *bind.CallOpts, asset common.Address) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleMainnetToken.contract.Call(opts, &out, "getAssetsAggregator", asset)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetAssetsAggregator is a free data retrieval call binding the contract method 0xb889a989.
//
// Solidity: function getAssetsAggregator(address asset) view returns(address, uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) GetAssetsAggregator(asset common.Address) (common.Address, *big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetAssetsAggregator(&_BscPledgeOracleMainnetToken.CallOpts, asset)
}

// GetAssetsAggregator is a free data retrieval call binding the contract method 0xb889a989.
//
// Solidity: function getAssetsAggregator(address asset) view returns(address, uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCallerSession) GetAssetsAggregator(asset common.Address) (common.Address, *big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetAssetsAggregator(&_BscPledgeOracleMainnetToken.CallOpts, asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCaller) GetPrice(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleMainnetToken.contract.Call(opts, &out, "getPrice", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) GetPrice(asset common.Address) (*big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetPrice(&_BscPledgeOracleMainnetToken.CallOpts, asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCallerSession) GetPrice(asset common.Address) (*big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetPrice(&_BscPledgeOracleMainnetToken.CallOpts, asset)
}

// GetPrices is a free data retrieval call binding the contract method 0x09cb3a4e.
//
// Solidity: function getPrices(uint256[] assets) view returns(uint256[])
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCaller) GetPrices(opts *bind.CallOpts, assets []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleMainnetToken.contract.Call(opts, &out, "getPrices", assets)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPrices is a free data retrieval call binding the contract method 0x09cb3a4e.
//
// Solidity: function getPrices(uint256[] assets) view returns(uint256[])
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) GetPrices(assets []*big.Int) ([]*big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetPrices(&_BscPledgeOracleMainnetToken.CallOpts, assets)
}

// GetPrices is a free data retrieval call binding the contract method 0x09cb3a4e.
//
// Solidity: function getPrices(uint256[] assets) view returns(uint256[])
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCallerSession) GetPrices(assets []*big.Int) ([]*big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetPrices(&_BscPledgeOracleMainnetToken.CallOpts, assets)
}

// GetUnderlyingAggregator is a free data retrieval call binding the contract method 0x75e443aa.
//
// Solidity: function getUnderlyingAggregator(uint256 underlying) view returns(address, uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCaller) GetUnderlyingAggregator(opts *bind.CallOpts, underlying *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleMainnetToken.contract.Call(opts, &out, "getUnderlyingAggregator", underlying)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetUnderlyingAggregator is a free data retrieval call binding the contract method 0x75e443aa.
//
// Solidity: function getUnderlyingAggregator(uint256 underlying) view returns(address, uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) GetUnderlyingAggregator(underlying *big.Int) (common.Address, *big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetUnderlyingAggregator(&_BscPledgeOracleMainnetToken.CallOpts, underlying)
}

// GetUnderlyingAggregator is a free data retrieval call binding the contract method 0x75e443aa.
//
// Solidity: function getUnderlyingAggregator(uint256 underlying) view returns(address, uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCallerSession) GetUnderlyingAggregator(underlying *big.Int) (common.Address, *big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetUnderlyingAggregator(&_BscPledgeOracleMainnetToken.CallOpts, underlying)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xda663257.
//
// Solidity: function getUnderlyingPrice(uint256 underlying) view returns(uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCaller) GetUnderlyingPrice(opts *bind.CallOpts, underlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleMainnetToken.contract.Call(opts, &out, "getUnderlyingPrice", underlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xda663257.
//
// Solidity: function getUnderlyingPrice(uint256 underlying) view returns(uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) GetUnderlyingPrice(underlying *big.Int) (*big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetUnderlyingPrice(&_BscPledgeOracleMainnetToken.CallOpts, underlying)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xda663257.
//
// Solidity: function getUnderlyingPrice(uint256 underlying) view returns(uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCallerSession) GetUnderlyingPrice(underlying *big.Int) (*big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetUnderlyingPrice(&_BscPledgeOracleMainnetToken.CallOpts, underlying)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscPledgeOracleMainnetToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) Owner() (common.Address, error) {
	return _BscPledgeOracleMainnetToken.Contract.Owner(&_BscPledgeOracleMainnetToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCallerSession) Owner() (common.Address, error) {
	return _BscPledgeOracleMainnetToken.Contract.Owner(&_BscPledgeOracleMainnetToken.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.RenounceOwnership(&_BscPledgeOracleMainnetToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.RenounceOwnership(&_BscPledgeOracleMainnetToken.TransactOpts)
}

// SetAssetsAggregator is a paid mutator transaction binding the contract method 0xcd9ffa0b.
//
// Solidity: function setAssetsAggregator(address asset, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) SetAssetsAggregator(opts *bind.TransactOpts, asset common.Address, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "setAssetsAggregator", asset, aggergator, _decimals)
}

// SetAssetsAggregator is a paid mutator transaction binding the contract method 0xcd9ffa0b.
//
// Solidity: function setAssetsAggregator(address asset, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) SetAssetsAggregator(asset common.Address, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetAssetsAggregator(&_BscPledgeOracleMainnetToken.TransactOpts, asset, aggergator, _decimals)
}

// SetAssetsAggregator is a paid mutator transaction binding the contract method 0xcd9ffa0b.
//
// Solidity: function setAssetsAggregator(address asset, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) SetAssetsAggregator(asset common.Address, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetAssetsAggregator(&_BscPledgeOracleMainnetToken.TransactOpts, asset, aggergator, _decimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 newDecimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) SetDecimals(opts *bind.TransactOpts, newDecimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "setDecimals", newDecimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 newDecimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) SetDecimals(newDecimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetDecimals(&_BscPledgeOracleMainnetToken.TransactOpts, newDecimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 newDecimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) SetDecimals(newDecimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetDecimals(&_BscPledgeOracleMainnetToken.TransactOpts, newDecimals)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address asset, uint256 price) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) SetPrice(opts *bind.TransactOpts, asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "setPrice", asset, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address asset, uint256 price) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) SetPrice(asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetPrice(&_BscPledgeOracleMainnetToken.TransactOpts, asset, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address asset, uint256 price) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) SetPrice(asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetPrice(&_BscPledgeOracleMainnetToken.TransactOpts, asset, price)
}

// SetPrices is a paid mutator transaction binding the contract method 0xd05eaae0.
//
// Solidity: function setPrices(uint256[] assets, uint256[] prices) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) SetPrices(opts *bind.TransactOpts, assets []*big.Int, prices []*big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "setPrices", assets, prices)
}

// SetPrices is a paid mutator transaction binding the contract method 0xd05eaae0.
//
// Solidity: function setPrices(uint256[] assets, uint256[] prices) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) SetPrices(assets []*big.Int, prices []*big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetPrices(&_BscPledgeOracleMainnetToken.TransactOpts, assets, prices)
}

// SetPrices is a paid mutator transaction binding the contract method 0xd05eaae0.
//
// Solidity: function setPrices(uint256[] assets, uint256[] prices) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) SetPrices(assets []*big.Int, prices []*big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetPrices(&_BscPledgeOracleMainnetToken.TransactOpts, assets, prices)
}

// SetUnderlyingAggregator is a paid mutator transaction binding the contract method 0x434b33ac.
//
// Solidity: function setUnderlyingAggregator(uint256 underlying, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) SetUnderlyingAggregator(opts *bind.TransactOpts, underlying *big.Int, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "setUnderlyingAggregator", underlying, aggergator, _decimals)
}

// SetUnderlyingAggregator is a paid mutator transaction binding the contract method 0x434b33ac.
//
// Solidity: function setUnderlyingAggregator(uint256 underlying, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) SetUnderlyingAggregator(underlying *big.Int, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetUnderlyingAggregator(&_BscPledgeOracleMainnetToken.TransactOpts, underlying, aggergator, _decimals)
}

// SetUnderlyingAggregator is a paid mutator transaction binding the contract method 0x434b33ac.
//
// Solidity: function setUnderlyingAggregator(uint256 underlying, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) SetUnderlyingAggregator(underlying *big.Int, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetUnderlyingAggregator(&_BscPledgeOracleMainnetToken.TransactOpts, underlying, aggergator, _decimals)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x83532667.
//
// Solidity: function setUnderlyingPrice(uint256 underlying, uint256 price) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) SetUnderlyingPrice(opts *bind.TransactOpts, underlying *big.Int, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "setUnderlyingPrice", underlying, price)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x83532667.
//
// Solidity: function setUnderlyingPrice(uint256 underlying, uint256 price) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) SetUnderlyingPrice(underlying *big.Int, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetUnderlyingPrice(&_BscPledgeOracleMainnetToken.TransactOpts, underlying, price)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x83532667.
//
// Solidity: function setUnderlyingPrice(uint256 underlying, uint256 price) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) SetUnderlyingPrice(underlying *big.Int, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetUnderlyingPrice(&_BscPledgeOracleMainnetToken.TransactOpts, underlying, price)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.TransferOwnership(&_BscPledgeOracleMainnetToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.TransferOwnership(&_BscPledgeOracleMainnetToken.TransactOpts, newOwner)
}

// BscPledgeOracleMainnetTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BscPledgeOracleMainnetToken contract.
type BscPledgeOracleMainnetTokenOwnershipTransferredIterator struct {
	Event *BscPledgeOracleMainnetTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BscPledgeOracleMainnetTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscPledgeOracleMainnetTokenOwnershipTransferred)
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
		it.Event = new(BscPledgeOracleMainnetTokenOwnershipTransferred)
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
func (it *BscPledgeOracleMainnetTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscPledgeOracleMainnetTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscPledgeOracleMainnetTokenOwnershipTransferred represents a OwnershipTransferred event raised by the BscPledgeOracleMainnetToken contract.
type BscPledgeOracleMainnetTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BscPledgeOracleMainnetTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BscPledgeOracleMainnetToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleMainnetTokenOwnershipTransferredIterator{contract: _BscPledgeOracleMainnetToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BscPledgeOracleMainnetTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BscPledgeOracleMainnetToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscPledgeOracleMainnetTokenOwnershipTransferred)
				if err := _BscPledgeOracleMainnetToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenFilterer) ParseOwnershipTransferred(log types.Log) (*BscPledgeOracleMainnetTokenOwnershipTransferred, error) {
	event := new(BscPledgeOracleMainnetTokenOwnershipTransferred)
	if err := _BscPledgeOracleMainnetToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
