// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// LillyboxWallet is an auto generated low-level Go binding around an user-defined struct.
type LillyboxWallet struct {
	KlayBalance           *big.Int
	StakedBalance         *big.Int
	StakedBlock           *big.Int
	UnstakeRequestBlock   *big.Int
	UnstakePendingBalance *big.Int
	Lock                  bool
}

// LillyboxMetaData contains all meta data concerning the Lillybox contract.
var LillyboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_adsCostPer10Times\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unstakeProcessTime\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blocknumber\",\"type\":\"uint256\"}],\"name\":\"DepositKlay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blocknumber\",\"type\":\"uint256\"}],\"name\":\"FlushReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usedLIL\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blocknumber\",\"type\":\"uint256\"}],\"name\":\"NewDonation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blocknumber\",\"type\":\"uint256\"}],\"name\":\"NewStaking\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blocknumber\",\"type\":\"uint256\"}],\"name\":\"UnStaking\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usedLIL\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blocknumber\",\"type\":\"uint256\"}],\"name\":\"WithdrawKlay\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AdsCostPer10Times\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIL_DECIMALS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LKLAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RewardPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"toNickname\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_donationAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lilAmount\",\"type\":\"uint256\"}],\"name\":\"donation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flushReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumAdsCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"mintAd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"mintVod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"showRevenuesFromAds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"showRevenuesFromFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"showWallet\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"klayBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakedBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakedBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeRequestBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakePendingBalance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"lock\",\"type\":\"bool\"}],\"internalType\":\"structLillybox.Wallet\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"vacuumAd\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lilAmount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawRevenue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LillyboxABI is the input ABI used to generate the binding from.
// Deprecated: Use LillyboxMetaData.ABI instead.
var LillyboxABI = LillyboxMetaData.ABI

// Lillybox is an auto generated Go binding around an Ethereum contract.
type Lillybox struct {
	LillyboxCaller     // Read-only binding to the contract
	LillyboxTransactor // Write-only binding to the contract
	LillyboxFilterer   // Log filterer for contract events
}

// LillyboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type LillyboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LillyboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LillyboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LillyboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LillyboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LillyboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LillyboxSession struct {
	Contract     *Lillybox         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LillyboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LillyboxCallerSession struct {
	Contract *LillyboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// LillyboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LillyboxTransactorSession struct {
	Contract     *LillyboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LillyboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type LillyboxRaw struct {
	Contract *Lillybox // Generic contract binding to access the raw methods on
}

// LillyboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LillyboxCallerRaw struct {
	Contract *LillyboxCaller // Generic read-only contract binding to access the raw methods on
}

// LillyboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LillyboxTransactorRaw struct {
	Contract *LillyboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLillybox creates a new instance of Lillybox, bound to a specific deployed contract.
func NewLillybox(address common.Address, backend bind.ContractBackend) (*Lillybox, error) {
	contract, err := bindLillybox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lillybox{LillyboxCaller: LillyboxCaller{contract: contract}, LillyboxTransactor: LillyboxTransactor{contract: contract}, LillyboxFilterer: LillyboxFilterer{contract: contract}}, nil
}

// NewLillyboxCaller creates a new read-only instance of Lillybox, bound to a specific deployed contract.
func NewLillyboxCaller(address common.Address, caller bind.ContractCaller) (*LillyboxCaller, error) {
	contract, err := bindLillybox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LillyboxCaller{contract: contract}, nil
}

// NewLillyboxTransactor creates a new write-only instance of Lillybox, bound to a specific deployed contract.
func NewLillyboxTransactor(address common.Address, transactor bind.ContractTransactor) (*LillyboxTransactor, error) {
	contract, err := bindLillybox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LillyboxTransactor{contract: contract}, nil
}

// NewLillyboxFilterer creates a new log filterer instance of Lillybox, bound to a specific deployed contract.
func NewLillyboxFilterer(address common.Address, filterer bind.ContractFilterer) (*LillyboxFilterer, error) {
	contract, err := bindLillybox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LillyboxFilterer{contract: contract}, nil
}

// bindLillybox binds a generic wrapper to an already deployed contract.
func bindLillybox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LillyboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lillybox *LillyboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lillybox.Contract.LillyboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lillybox *LillyboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lillybox.Contract.LillyboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lillybox *LillyboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lillybox.Contract.LillyboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lillybox *LillyboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lillybox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lillybox *LillyboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lillybox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lillybox *LillyboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lillybox.Contract.contract.Transact(opts, method, params...)
}

// AdsCostPer10Times is a free data retrieval call binding the contract method 0xf4f17214.
//
// Solidity: function AdsCostPer10Times() view returns(uint256)
func (_Lillybox *LillyboxCaller) AdsCostPer10Times(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lillybox.contract.Call(opts, &out, "AdsCostPer10Times")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdsCostPer10Times is a free data retrieval call binding the contract method 0xf4f17214.
//
// Solidity: function AdsCostPer10Times() view returns(uint256)
func (_Lillybox *LillyboxSession) AdsCostPer10Times() (*big.Int, error) {
	return _Lillybox.Contract.AdsCostPer10Times(&_Lillybox.CallOpts)
}

// AdsCostPer10Times is a free data retrieval call binding the contract method 0xf4f17214.
//
// Solidity: function AdsCostPer10Times() view returns(uint256)
func (_Lillybox *LillyboxCallerSession) AdsCostPer10Times() (*big.Int, error) {
	return _Lillybox.Contract.AdsCostPer10Times(&_Lillybox.CallOpts)
}

// LIL is a free data retrieval call binding the contract method 0x357bbf82.
//
// Solidity: function LIL() view returns(uint256)
func (_Lillybox *LillyboxCaller) LIL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lillybox.contract.Call(opts, &out, "LIL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LIL is a free data retrieval call binding the contract method 0x357bbf82.
//
// Solidity: function LIL() view returns(uint256)
func (_Lillybox *LillyboxSession) LIL() (*big.Int, error) {
	return _Lillybox.Contract.LIL(&_Lillybox.CallOpts)
}

// LIL is a free data retrieval call binding the contract method 0x357bbf82.
//
// Solidity: function LIL() view returns(uint256)
func (_Lillybox *LillyboxCallerSession) LIL() (*big.Int, error) {
	return _Lillybox.Contract.LIL(&_Lillybox.CallOpts)
}

// LILDECIMALS is a free data retrieval call binding the contract method 0xc77c0bc8.
//
// Solidity: function LIL_DECIMALS() view returns(uint256)
func (_Lillybox *LillyboxCaller) LILDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lillybox.contract.Call(opts, &out, "LIL_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LILDECIMALS is a free data retrieval call binding the contract method 0xc77c0bc8.
//
// Solidity: function LIL_DECIMALS() view returns(uint256)
func (_Lillybox *LillyboxSession) LILDECIMALS() (*big.Int, error) {
	return _Lillybox.Contract.LILDECIMALS(&_Lillybox.CallOpts)
}

// LILDECIMALS is a free data retrieval call binding the contract method 0xc77c0bc8.
//
// Solidity: function LIL_DECIMALS() view returns(uint256)
func (_Lillybox *LillyboxCallerSession) LILDECIMALS() (*big.Int, error) {
	return _Lillybox.Contract.LILDECIMALS(&_Lillybox.CallOpts)
}

// LKLAY is a free data retrieval call binding the contract method 0x44ec85cb.
//
// Solidity: function LKLAY() view returns(uint256)
func (_Lillybox *LillyboxCaller) LKLAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lillybox.contract.Call(opts, &out, "LKLAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LKLAY is a free data retrieval call binding the contract method 0x44ec85cb.
//
// Solidity: function LKLAY() view returns(uint256)
func (_Lillybox *LillyboxSession) LKLAY() (*big.Int, error) {
	return _Lillybox.Contract.LKLAY(&_Lillybox.CallOpts)
}

// LKLAY is a free data retrieval call binding the contract method 0x44ec85cb.
//
// Solidity: function LKLAY() view returns(uint256)
func (_Lillybox *LillyboxCallerSession) LKLAY() (*big.Int, error) {
	return _Lillybox.Contract.LKLAY(&_Lillybox.CallOpts)
}

// RewardPerBlock is a free data retrieval call binding the contract method 0x462258d3.
//
// Solidity: function RewardPerBlock() view returns(uint256)
func (_Lillybox *LillyboxCaller) RewardPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lillybox.contract.Call(opts, &out, "RewardPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerBlock is a free data retrieval call binding the contract method 0x462258d3.
//
// Solidity: function RewardPerBlock() view returns(uint256)
func (_Lillybox *LillyboxSession) RewardPerBlock() (*big.Int, error) {
	return _Lillybox.Contract.RewardPerBlock(&_Lillybox.CallOpts)
}

// RewardPerBlock is a free data retrieval call binding the contract method 0x462258d3.
//
// Solidity: function RewardPerBlock() view returns(uint256)
func (_Lillybox *LillyboxCallerSession) RewardPerBlock() (*big.Int, error) {
	return _Lillybox.Contract.RewardPerBlock(&_Lillybox.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256)
func (_Lillybox *LillyboxCaller) BalanceOf(opts *bind.CallOpts, owner common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lillybox.contract.Call(opts, &out, "balanceOf", owner, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256)
func (_Lillybox *LillyboxSession) BalanceOf(owner common.Address, id *big.Int) (*big.Int, error) {
	return _Lillybox.Contract.BalanceOf(&_Lillybox.CallOpts, owner, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256)
func (_Lillybox *LillyboxCallerSession) BalanceOf(owner common.Address, id *big.Int) (*big.Int, error) {
	return _Lillybox.Contract.BalanceOf(&_Lillybox.CallOpts, owner, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[])
func (_Lillybox *LillyboxCaller) BalanceOfBatch(opts *bind.CallOpts, owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Lillybox.contract.Call(opts, &out, "balanceOfBatch", owners, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[])
func (_Lillybox *LillyboxSession) BalanceOfBatch(owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Lillybox.Contract.BalanceOfBatch(&_Lillybox.CallOpts, owners, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[])
func (_Lillybox *LillyboxCallerSession) BalanceOfBatch(owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Lillybox.Contract.BalanceOfBatch(&_Lillybox.CallOpts, owners, ids)
}

// GetRewards is a free data retrieval call binding the contract method 0x79ee54f7.
//
// Solidity: function getRewards(address user) view returns(uint256)
func (_Lillybox *LillyboxCaller) GetRewards(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lillybox.contract.Call(opts, &out, "getRewards", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewards is a free data retrieval call binding the contract method 0x79ee54f7.
//
// Solidity: function getRewards(address user) view returns(uint256)
func (_Lillybox *LillyboxSession) GetRewards(user common.Address) (*big.Int, error) {
	return _Lillybox.Contract.GetRewards(&_Lillybox.CallOpts, user)
}

// GetRewards is a free data retrieval call binding the contract method 0x79ee54f7.
//
// Solidity: function getRewards(address user) view returns(uint256)
func (_Lillybox *LillyboxCallerSession) GetRewards(user common.Address) (*big.Int, error) {
	return _Lillybox.Contract.GetRewards(&_Lillybox.CallOpts, user)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Lillybox *LillyboxCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Lillybox.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Lillybox *LillyboxSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Lillybox.Contract.IsApprovedForAll(&_Lillybox.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Lillybox *LillyboxCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Lillybox.Contract.IsApprovedForAll(&_Lillybox.CallOpts, owner, operator)
}

// MinimumAdsCost is a free data retrieval call binding the contract method 0x81814ded.
//
// Solidity: function minimumAdsCost() view returns(uint256)
func (_Lillybox *LillyboxCaller) MinimumAdsCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lillybox.contract.Call(opts, &out, "minimumAdsCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumAdsCost is a free data retrieval call binding the contract method 0x81814ded.
//
// Solidity: function minimumAdsCost() view returns(uint256)
func (_Lillybox *LillyboxSession) MinimumAdsCost() (*big.Int, error) {
	return _Lillybox.Contract.MinimumAdsCost(&_Lillybox.CallOpts)
}

// MinimumAdsCost is a free data retrieval call binding the contract method 0x81814ded.
//
// Solidity: function minimumAdsCost() view returns(uint256)
func (_Lillybox *LillyboxCallerSession) MinimumAdsCost() (*big.Int, error) {
	return _Lillybox.Contract.MinimumAdsCost(&_Lillybox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lillybox *LillyboxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lillybox.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lillybox *LillyboxSession) Owner() (common.Address, error) {
	return _Lillybox.Contract.Owner(&_Lillybox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lillybox *LillyboxCallerSession) Owner() (common.Address, error) {
	return _Lillybox.Contract.Owner(&_Lillybox.CallOpts)
}

// ShowRevenuesFromAds is a free data retrieval call binding the contract method 0xd9718ca4.
//
// Solidity: function showRevenuesFromAds() view returns(uint256)
func (_Lillybox *LillyboxCaller) ShowRevenuesFromAds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lillybox.contract.Call(opts, &out, "showRevenuesFromAds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ShowRevenuesFromAds is a free data retrieval call binding the contract method 0xd9718ca4.
//
// Solidity: function showRevenuesFromAds() view returns(uint256)
func (_Lillybox *LillyboxSession) ShowRevenuesFromAds() (*big.Int, error) {
	return _Lillybox.Contract.ShowRevenuesFromAds(&_Lillybox.CallOpts)
}

// ShowRevenuesFromAds is a free data retrieval call binding the contract method 0xd9718ca4.
//
// Solidity: function showRevenuesFromAds() view returns(uint256)
func (_Lillybox *LillyboxCallerSession) ShowRevenuesFromAds() (*big.Int, error) {
	return _Lillybox.Contract.ShowRevenuesFromAds(&_Lillybox.CallOpts)
}

// ShowRevenuesFromFee is a free data retrieval call binding the contract method 0x26615aa8.
//
// Solidity: function showRevenuesFromFee() view returns(uint256)
func (_Lillybox *LillyboxCaller) ShowRevenuesFromFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lillybox.contract.Call(opts, &out, "showRevenuesFromFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ShowRevenuesFromFee is a free data retrieval call binding the contract method 0x26615aa8.
//
// Solidity: function showRevenuesFromFee() view returns(uint256)
func (_Lillybox *LillyboxSession) ShowRevenuesFromFee() (*big.Int, error) {
	return _Lillybox.Contract.ShowRevenuesFromFee(&_Lillybox.CallOpts)
}

// ShowRevenuesFromFee is a free data retrieval call binding the contract method 0x26615aa8.
//
// Solidity: function showRevenuesFromFee() view returns(uint256)
func (_Lillybox *LillyboxCallerSession) ShowRevenuesFromFee() (*big.Int, error) {
	return _Lillybox.Contract.ShowRevenuesFromFee(&_Lillybox.CallOpts)
}

// ShowWallet is a free data retrieval call binding the contract method 0x262655bf.
//
// Solidity: function showWallet(address user) view returns((uint256,uint256,uint256,uint256,uint256,bool))
func (_Lillybox *LillyboxCaller) ShowWallet(opts *bind.CallOpts, user common.Address) (LillyboxWallet, error) {
	var out []interface{}
	err := _Lillybox.contract.Call(opts, &out, "showWallet", user)

	if err != nil {
		return *new(LillyboxWallet), err
	}

	out0 := *abi.ConvertType(out[0], new(LillyboxWallet)).(*LillyboxWallet)

	return out0, err

}

// ShowWallet is a free data retrieval call binding the contract method 0x262655bf.
//
// Solidity: function showWallet(address user) view returns((uint256,uint256,uint256,uint256,uint256,bool))
func (_Lillybox *LillyboxSession) ShowWallet(user common.Address) (LillyboxWallet, error) {
	return _Lillybox.Contract.ShowWallet(&_Lillybox.CallOpts, user)
}

// ShowWallet is a free data retrieval call binding the contract method 0x262655bf.
//
// Solidity: function showWallet(address user) view returns((uint256,uint256,uint256,uint256,uint256,bool))
func (_Lillybox *LillyboxCallerSession) ShowWallet(user common.Address) (LillyboxWallet, error) {
	return _Lillybox.Contract.ShowWallet(&_Lillybox.CallOpts, user)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Lillybox *LillyboxCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Lillybox.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Lillybox *LillyboxSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Lillybox.Contract.SupportsInterface(&_Lillybox.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Lillybox *LillyboxCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Lillybox.Contract.SupportsInterface(&_Lillybox.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_Lillybox *LillyboxCaller) Uri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Lillybox.contract.Call(opts, &out, "uri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_Lillybox *LillyboxSession) Uri(tokenId *big.Int) (string, error) {
	return _Lillybox.Contract.Uri(&_Lillybox.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_Lillybox *LillyboxCallerSession) Uri(tokenId *big.Int) (string, error) {
	return _Lillybox.Contract.Uri(&_Lillybox.CallOpts, tokenId)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(bool)
func (_Lillybox *LillyboxTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lillybox.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(bool)
func (_Lillybox *LillyboxSession) Deposit() (*types.Transaction, error) {
	return _Lillybox.Contract.Deposit(&_Lillybox.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(bool)
func (_Lillybox *LillyboxTransactorSession) Deposit() (*types.Transaction, error) {
	return _Lillybox.Contract.Deposit(&_Lillybox.TransactOpts)
}

// Donation is a paid mutator transaction binding the contract method 0x23e7f146.
//
// Solidity: function donation(string toNickname, uint256 _donationAmount, uint256 _lilAmount) returns(bool)
func (_Lillybox *LillyboxTransactor) Donation(opts *bind.TransactOpts, toNickname string, _donationAmount *big.Int, _lilAmount *big.Int) (*types.Transaction, error) {
	return _Lillybox.contract.Transact(opts, "donation", toNickname, _donationAmount, _lilAmount)
}

// Donation is a paid mutator transaction binding the contract method 0x23e7f146.
//
// Solidity: function donation(string toNickname, uint256 _donationAmount, uint256 _lilAmount) returns(bool)
func (_Lillybox *LillyboxSession) Donation(toNickname string, _donationAmount *big.Int, _lilAmount *big.Int) (*types.Transaction, error) {
	return _Lillybox.Contract.Donation(&_Lillybox.TransactOpts, toNickname, _donationAmount, _lilAmount)
}

// Donation is a paid mutator transaction binding the contract method 0x23e7f146.
//
// Solidity: function donation(string toNickname, uint256 _donationAmount, uint256 _lilAmount) returns(bool)
func (_Lillybox *LillyboxTransactorSession) Donation(toNickname string, _donationAmount *big.Int, _lilAmount *big.Int) (*types.Transaction, error) {
	return _Lillybox.Contract.Donation(&_Lillybox.TransactOpts, toNickname, _donationAmount, _lilAmount)
}

// FlushReward is a paid mutator transaction binding the contract method 0xc362defe.
//
// Solidity: function flushReward() returns(uint256)
func (_Lillybox *LillyboxTransactor) FlushReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lillybox.contract.Transact(opts, "flushReward")
}

// FlushReward is a paid mutator transaction binding the contract method 0xc362defe.
//
// Solidity: function flushReward() returns(uint256)
func (_Lillybox *LillyboxSession) FlushReward() (*types.Transaction, error) {
	return _Lillybox.Contract.FlushReward(&_Lillybox.TransactOpts)
}

// FlushReward is a paid mutator transaction binding the contract method 0xc362defe.
//
// Solidity: function flushReward() returns(uint256)
func (_Lillybox *LillyboxTransactorSession) FlushReward() (*types.Transaction, error) {
	return _Lillybox.Contract.FlushReward(&_Lillybox.TransactOpts)
}

// MintAd is a paid mutator transaction binding the contract method 0xe79d6544.
//
// Solidity: function mintAd(string tokenURI) payable returns(uint256)
func (_Lillybox *LillyboxTransactor) MintAd(opts *bind.TransactOpts, tokenURI string) (*types.Transaction, error) {
	return _Lillybox.contract.Transact(opts, "mintAd", tokenURI)
}

// MintAd is a paid mutator transaction binding the contract method 0xe79d6544.
//
// Solidity: function mintAd(string tokenURI) payable returns(uint256)
func (_Lillybox *LillyboxSession) MintAd(tokenURI string) (*types.Transaction, error) {
	return _Lillybox.Contract.MintAd(&_Lillybox.TransactOpts, tokenURI)
}

// MintAd is a paid mutator transaction binding the contract method 0xe79d6544.
//
// Solidity: function mintAd(string tokenURI) payable returns(uint256)
func (_Lillybox *LillyboxTransactorSession) MintAd(tokenURI string) (*types.Transaction, error) {
	return _Lillybox.Contract.MintAd(&_Lillybox.TransactOpts, tokenURI)
}

// MintVod is a paid mutator transaction binding the contract method 0xb04fb68f.
//
// Solidity: function mintVod(string tokenURI) payable returns(uint256)
func (_Lillybox *LillyboxTransactor) MintVod(opts *bind.TransactOpts, tokenURI string) (*types.Transaction, error) {
	return _Lillybox.contract.Transact(opts, "mintVod", tokenURI)
}

// MintVod is a paid mutator transaction binding the contract method 0xb04fb68f.
//
// Solidity: function mintVod(string tokenURI) payable returns(uint256)
func (_Lillybox *LillyboxSession) MintVod(tokenURI string) (*types.Transaction, error) {
	return _Lillybox.Contract.MintVod(&_Lillybox.TransactOpts, tokenURI)
}

// MintVod is a paid mutator transaction binding the contract method 0xb04fb68f.
//
// Solidity: function mintVod(string tokenURI) payable returns(uint256)
func (_Lillybox *LillyboxTransactorSession) MintVod(tokenURI string) (*types.Transaction, error) {
	return _Lillybox.Contract.MintVod(&_Lillybox.TransactOpts, tokenURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lillybox *LillyboxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lillybox.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lillybox *LillyboxSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lillybox.Contract.RenounceOwnership(&_Lillybox.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lillybox *LillyboxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lillybox.Contract.RenounceOwnership(&_Lillybox.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Lillybox *LillyboxTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Lillybox.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Lillybox *LillyboxSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Lillybox.Contract.SafeBatchTransferFrom(&_Lillybox.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Lillybox *LillyboxTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Lillybox.Contract.SafeBatchTransferFrom(&_Lillybox.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Lillybox *LillyboxTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Lillybox.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Lillybox *LillyboxSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Lillybox.Contract.SafeTransferFrom(&_Lillybox.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Lillybox *LillyboxTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Lillybox.Contract.SafeTransferFrom(&_Lillybox.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Lillybox *LillyboxTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Lillybox.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Lillybox *LillyboxSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Lillybox.Contract.SetApprovalForAll(&_Lillybox.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Lillybox *LillyboxTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Lillybox.Contract.SetApprovalForAll(&_Lillybox.TransactOpts, operator, approved)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _stakeAmount) returns(uint256)
func (_Lillybox *LillyboxTransactor) Stake(opts *bind.TransactOpts, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _Lillybox.contract.Transact(opts, "stake", _stakeAmount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _stakeAmount) returns(uint256)
func (_Lillybox *LillyboxSession) Stake(_stakeAmount *big.Int) (*types.Transaction, error) {
	return _Lillybox.Contract.Stake(&_Lillybox.TransactOpts, _stakeAmount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _stakeAmount) returns(uint256)
func (_Lillybox *LillyboxTransactorSession) Stake(_stakeAmount *big.Int) (*types.Transaction, error) {
	return _Lillybox.Contract.Stake(&_Lillybox.TransactOpts, _stakeAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lillybox *LillyboxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Lillybox.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lillybox *LillyboxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lillybox.Contract.TransferOwnership(&_Lillybox.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lillybox *LillyboxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lillybox.Contract.TransferOwnership(&_Lillybox.TransactOpts, newOwner)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _amount) returns(bool)
func (_Lillybox *LillyboxTransactor) Unstake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Lillybox.contract.Transact(opts, "unstake", _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _amount) returns(bool)
func (_Lillybox *LillyboxSession) Unstake(_amount *big.Int) (*types.Transaction, error) {
	return _Lillybox.Contract.Unstake(&_Lillybox.TransactOpts, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _amount) returns(bool)
func (_Lillybox *LillyboxTransactorSession) Unstake(_amount *big.Int) (*types.Transaction, error) {
	return _Lillybox.Contract.Unstake(&_Lillybox.TransactOpts, _amount)
}

// VacuumAd is a paid mutator transaction binding the contract method 0xc8052b9f.
//
// Solidity: function vacuumAd(address _owner, uint256 _tokenId) returns(bool)
func (_Lillybox *LillyboxTransactor) VacuumAd(opts *bind.TransactOpts, _owner common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Lillybox.contract.Transact(opts, "vacuumAd", _owner, _tokenId)
}

// VacuumAd is a paid mutator transaction binding the contract method 0xc8052b9f.
//
// Solidity: function vacuumAd(address _owner, uint256 _tokenId) returns(bool)
func (_Lillybox *LillyboxSession) VacuumAd(_owner common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Lillybox.Contract.VacuumAd(&_Lillybox.TransactOpts, _owner, _tokenId)
}

// VacuumAd is a paid mutator transaction binding the contract method 0xc8052b9f.
//
// Solidity: function vacuumAd(address _owner, uint256 _tokenId) returns(bool)
func (_Lillybox *LillyboxTransactorSession) VacuumAd(_owner common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Lillybox.Contract.VacuumAd(&_Lillybox.TransactOpts, _owner, _tokenId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _amount, uint256 _lilAmount) returns(bool)
func (_Lillybox *LillyboxTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int, _lilAmount *big.Int) (*types.Transaction, error) {
	return _Lillybox.contract.Transact(opts, "withdraw", _amount, _lilAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _amount, uint256 _lilAmount) returns(bool)
func (_Lillybox *LillyboxSession) Withdraw(_amount *big.Int, _lilAmount *big.Int) (*types.Transaction, error) {
	return _Lillybox.Contract.Withdraw(&_Lillybox.TransactOpts, _amount, _lilAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _amount, uint256 _lilAmount) returns(bool)
func (_Lillybox *LillyboxTransactorSession) Withdraw(_amount *big.Int, _lilAmount *big.Int) (*types.Transaction, error) {
	return _Lillybox.Contract.Withdraw(&_Lillybox.TransactOpts, _amount, _lilAmount)
}

// WithdrawRevenue is a paid mutator transaction binding the contract method 0x4f573cb2.
//
// Solidity: function withdrawRevenue() returns()
func (_Lillybox *LillyboxTransactor) WithdrawRevenue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lillybox.contract.Transact(opts, "withdrawRevenue")
}

// WithdrawRevenue is a paid mutator transaction binding the contract method 0x4f573cb2.
//
// Solidity: function withdrawRevenue() returns()
func (_Lillybox *LillyboxSession) WithdrawRevenue() (*types.Transaction, error) {
	return _Lillybox.Contract.WithdrawRevenue(&_Lillybox.TransactOpts)
}

// WithdrawRevenue is a paid mutator transaction binding the contract method 0x4f573cb2.
//
// Solidity: function withdrawRevenue() returns()
func (_Lillybox *LillyboxTransactorSession) WithdrawRevenue() (*types.Transaction, error) {
	return _Lillybox.Contract.WithdrawRevenue(&_Lillybox.TransactOpts)
}

// LillyboxApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Lillybox contract.
type LillyboxApprovalForAllIterator struct {
	Event *LillyboxApprovalForAll // Event containing the contract specifics and raw log

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
func (it *LillyboxApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LillyboxApprovalForAll)
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
		it.Event = new(LillyboxApprovalForAll)
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
func (it *LillyboxApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LillyboxApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LillyboxApprovalForAll represents a ApprovalForAll event raised by the Lillybox contract.
type LillyboxApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Lillybox *LillyboxFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*LillyboxApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Lillybox.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LillyboxApprovalForAllIterator{contract: _Lillybox.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Lillybox *LillyboxFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LillyboxApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Lillybox.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LillyboxApprovalForAll)
				if err := _Lillybox.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Lillybox *LillyboxFilterer) ParseApprovalForAll(log types.Log) (*LillyboxApprovalForAll, error) {
	event := new(LillyboxApprovalForAll)
	if err := _Lillybox.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LillyboxDepositKlayIterator is returned from FilterDepositKlay and is used to iterate over the raw logs and unpacked data for DepositKlay events raised by the Lillybox contract.
type LillyboxDepositKlayIterator struct {
	Event *LillyboxDepositKlay // Event containing the contract specifics and raw log

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
func (it *LillyboxDepositKlayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LillyboxDepositKlay)
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
		it.Event = new(LillyboxDepositKlay)
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
func (it *LillyboxDepositKlayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LillyboxDepositKlayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LillyboxDepositKlay represents a DepositKlay event raised by the Lillybox contract.
type LillyboxDepositKlay struct {
	User        common.Address
	Amount      *big.Int
	Blocknumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDepositKlay is a free log retrieval operation binding the contract event 0xe7cb3a3e3091c69af1533470533fe83a8b870d553dea09cfa02fceb84223d123.
//
// Solidity: event DepositKlay(address indexed user, uint256 amount, uint256 blocknumber)
func (_Lillybox *LillyboxFilterer) FilterDepositKlay(opts *bind.FilterOpts, user []common.Address) (*LillyboxDepositKlayIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lillybox.contract.FilterLogs(opts, "DepositKlay", userRule)
	if err != nil {
		return nil, err
	}
	return &LillyboxDepositKlayIterator{contract: _Lillybox.contract, event: "DepositKlay", logs: logs, sub: sub}, nil
}

// WatchDepositKlay is a free log subscription operation binding the contract event 0xe7cb3a3e3091c69af1533470533fe83a8b870d553dea09cfa02fceb84223d123.
//
// Solidity: event DepositKlay(address indexed user, uint256 amount, uint256 blocknumber)
func (_Lillybox *LillyboxFilterer) WatchDepositKlay(opts *bind.WatchOpts, sink chan<- *LillyboxDepositKlay, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lillybox.contract.WatchLogs(opts, "DepositKlay", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LillyboxDepositKlay)
				if err := _Lillybox.contract.UnpackLog(event, "DepositKlay", log); err != nil {
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

// ParseDepositKlay is a log parse operation binding the contract event 0xe7cb3a3e3091c69af1533470533fe83a8b870d553dea09cfa02fceb84223d123.
//
// Solidity: event DepositKlay(address indexed user, uint256 amount, uint256 blocknumber)
func (_Lillybox *LillyboxFilterer) ParseDepositKlay(log types.Log) (*LillyboxDepositKlay, error) {
	event := new(LillyboxDepositKlay)
	if err := _Lillybox.contract.UnpackLog(event, "DepositKlay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LillyboxFlushRewardIterator is returned from FilterFlushReward and is used to iterate over the raw logs and unpacked data for FlushReward events raised by the Lillybox contract.
type LillyboxFlushRewardIterator struct {
	Event *LillyboxFlushReward // Event containing the contract specifics and raw log

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
func (it *LillyboxFlushRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LillyboxFlushReward)
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
		it.Event = new(LillyboxFlushReward)
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
func (it *LillyboxFlushRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LillyboxFlushRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LillyboxFlushReward represents a FlushReward event raised by the Lillybox contract.
type LillyboxFlushReward struct {
	User        common.Address
	Amount      *big.Int
	Blocknumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFlushReward is a free log retrieval operation binding the contract event 0x7500cb8a15ad852dbd1b3a6cc70ada5d619b23224d22efbd599f765f77fa45c4.
//
// Solidity: event FlushReward(address indexed user, uint256 amount, uint256 blocknumber)
func (_Lillybox *LillyboxFilterer) FilterFlushReward(opts *bind.FilterOpts, user []common.Address) (*LillyboxFlushRewardIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lillybox.contract.FilterLogs(opts, "FlushReward", userRule)
	if err != nil {
		return nil, err
	}
	return &LillyboxFlushRewardIterator{contract: _Lillybox.contract, event: "FlushReward", logs: logs, sub: sub}, nil
}

// WatchFlushReward is a free log subscription operation binding the contract event 0x7500cb8a15ad852dbd1b3a6cc70ada5d619b23224d22efbd599f765f77fa45c4.
//
// Solidity: event FlushReward(address indexed user, uint256 amount, uint256 blocknumber)
func (_Lillybox *LillyboxFilterer) WatchFlushReward(opts *bind.WatchOpts, sink chan<- *LillyboxFlushReward, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lillybox.contract.WatchLogs(opts, "FlushReward", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LillyboxFlushReward)
				if err := _Lillybox.contract.UnpackLog(event, "FlushReward", log); err != nil {
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

// ParseFlushReward is a log parse operation binding the contract event 0x7500cb8a15ad852dbd1b3a6cc70ada5d619b23224d22efbd599f765f77fa45c4.
//
// Solidity: event FlushReward(address indexed user, uint256 amount, uint256 blocknumber)
func (_Lillybox *LillyboxFilterer) ParseFlushReward(log types.Log) (*LillyboxFlushReward, error) {
	event := new(LillyboxFlushReward)
	if err := _Lillybox.contract.UnpackLog(event, "FlushReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LillyboxNewDonationIterator is returned from FilterNewDonation and is used to iterate over the raw logs and unpacked data for NewDonation events raised by the Lillybox contract.
type LillyboxNewDonationIterator struct {
	Event *LillyboxNewDonation // Event containing the contract specifics and raw log

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
func (it *LillyboxNewDonationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LillyboxNewDonation)
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
		it.Event = new(LillyboxNewDonation)
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
func (it *LillyboxNewDonationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LillyboxNewDonationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LillyboxNewDonation represents a NewDonation event raised by the Lillybox contract.
type LillyboxNewDonation struct {
	From        common.Address
	To          common.Address
	Amount      *big.Int
	UsedLIL     *big.Int
	Blocknumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewDonation is a free log retrieval operation binding the contract event 0x2fd09b49f4e4fa0e29915f1b8c0a6ebe8466b03fdc2205800ab541fa8c09678f.
//
// Solidity: event NewDonation(address indexed from, address indexed to, uint256 amount, uint256 usedLIL, uint256 blocknumber)
func (_Lillybox *LillyboxFilterer) FilterNewDonation(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LillyboxNewDonationIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lillybox.contract.FilterLogs(opts, "NewDonation", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LillyboxNewDonationIterator{contract: _Lillybox.contract, event: "NewDonation", logs: logs, sub: sub}, nil
}

// WatchNewDonation is a free log subscription operation binding the contract event 0x2fd09b49f4e4fa0e29915f1b8c0a6ebe8466b03fdc2205800ab541fa8c09678f.
//
// Solidity: event NewDonation(address indexed from, address indexed to, uint256 amount, uint256 usedLIL, uint256 blocknumber)
func (_Lillybox *LillyboxFilterer) WatchNewDonation(opts *bind.WatchOpts, sink chan<- *LillyboxNewDonation, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lillybox.contract.WatchLogs(opts, "NewDonation", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LillyboxNewDonation)
				if err := _Lillybox.contract.UnpackLog(event, "NewDonation", log); err != nil {
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

// ParseNewDonation is a log parse operation binding the contract event 0x2fd09b49f4e4fa0e29915f1b8c0a6ebe8466b03fdc2205800ab541fa8c09678f.
//
// Solidity: event NewDonation(address indexed from, address indexed to, uint256 amount, uint256 usedLIL, uint256 blocknumber)
func (_Lillybox *LillyboxFilterer) ParseNewDonation(log types.Log) (*LillyboxNewDonation, error) {
	event := new(LillyboxNewDonation)
	if err := _Lillybox.contract.UnpackLog(event, "NewDonation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LillyboxNewStakingIterator is returned from FilterNewStaking and is used to iterate over the raw logs and unpacked data for NewStaking events raised by the Lillybox contract.
type LillyboxNewStakingIterator struct {
	Event *LillyboxNewStaking // Event containing the contract specifics and raw log

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
func (it *LillyboxNewStakingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LillyboxNewStaking)
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
		it.Event = new(LillyboxNewStaking)
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
func (it *LillyboxNewStakingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LillyboxNewStakingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LillyboxNewStaking represents a NewStaking event raised by the Lillybox contract.
type LillyboxNewStaking struct {
	User        common.Address
	Amount      *big.Int
	Blocknumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewStaking is a free log retrieval operation binding the contract event 0x86ef88f9e8fad1e7c0a766303b62ffb522daeff6ebeca675e3b5c973d8d0132b.
//
// Solidity: event NewStaking(address indexed user, uint256 amount, uint256 blocknumber)
func (_Lillybox *LillyboxFilterer) FilterNewStaking(opts *bind.FilterOpts, user []common.Address) (*LillyboxNewStakingIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lillybox.contract.FilterLogs(opts, "NewStaking", userRule)
	if err != nil {
		return nil, err
	}
	return &LillyboxNewStakingIterator{contract: _Lillybox.contract, event: "NewStaking", logs: logs, sub: sub}, nil
}

// WatchNewStaking is a free log subscription operation binding the contract event 0x86ef88f9e8fad1e7c0a766303b62ffb522daeff6ebeca675e3b5c973d8d0132b.
//
// Solidity: event NewStaking(address indexed user, uint256 amount, uint256 blocknumber)
func (_Lillybox *LillyboxFilterer) WatchNewStaking(opts *bind.WatchOpts, sink chan<- *LillyboxNewStaking, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lillybox.contract.WatchLogs(opts, "NewStaking", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LillyboxNewStaking)
				if err := _Lillybox.contract.UnpackLog(event, "NewStaking", log); err != nil {
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

// ParseNewStaking is a log parse operation binding the contract event 0x86ef88f9e8fad1e7c0a766303b62ffb522daeff6ebeca675e3b5c973d8d0132b.
//
// Solidity: event NewStaking(address indexed user, uint256 amount, uint256 blocknumber)
func (_Lillybox *LillyboxFilterer) ParseNewStaking(log types.Log) (*LillyboxNewStaking, error) {
	event := new(LillyboxNewStaking)
	if err := _Lillybox.contract.UnpackLog(event, "NewStaking", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LillyboxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Lillybox contract.
type LillyboxOwnershipTransferredIterator struct {
	Event *LillyboxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LillyboxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LillyboxOwnershipTransferred)
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
		it.Event = new(LillyboxOwnershipTransferred)
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
func (it *LillyboxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LillyboxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LillyboxOwnershipTransferred represents a OwnershipTransferred event raised by the Lillybox contract.
type LillyboxOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lillybox *LillyboxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LillyboxOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lillybox.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LillyboxOwnershipTransferredIterator{contract: _Lillybox.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lillybox *LillyboxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LillyboxOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lillybox.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LillyboxOwnershipTransferred)
				if err := _Lillybox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Lillybox *LillyboxFilterer) ParseOwnershipTransferred(log types.Log) (*LillyboxOwnershipTransferred, error) {
	event := new(LillyboxOwnershipTransferred)
	if err := _Lillybox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LillyboxTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Lillybox contract.
type LillyboxTransferBatchIterator struct {
	Event *LillyboxTransferBatch // Event containing the contract specifics and raw log

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
func (it *LillyboxTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LillyboxTransferBatch)
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
		it.Event = new(LillyboxTransferBatch)
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
func (it *LillyboxTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LillyboxTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LillyboxTransferBatch represents a TransferBatch event raised by the Lillybox contract.
type LillyboxTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Amounts  []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] amounts)
func (_Lillybox *LillyboxFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*LillyboxTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lillybox.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LillyboxTransferBatchIterator{contract: _Lillybox.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] amounts)
func (_Lillybox *LillyboxFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *LillyboxTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lillybox.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LillyboxTransferBatch)
				if err := _Lillybox.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] amounts)
func (_Lillybox *LillyboxFilterer) ParseTransferBatch(log types.Log) (*LillyboxTransferBatch, error) {
	event := new(LillyboxTransferBatch)
	if err := _Lillybox.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LillyboxTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Lillybox contract.
type LillyboxTransferSingleIterator struct {
	Event *LillyboxTransferSingle // Event containing the contract specifics and raw log

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
func (it *LillyboxTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LillyboxTransferSingle)
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
		it.Event = new(LillyboxTransferSingle)
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
func (it *LillyboxTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LillyboxTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LillyboxTransferSingle represents a TransferSingle event raised by the Lillybox contract.
type LillyboxTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 amount)
func (_Lillybox *LillyboxFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*LillyboxTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lillybox.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LillyboxTransferSingleIterator{contract: _Lillybox.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 amount)
func (_Lillybox *LillyboxFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *LillyboxTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lillybox.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LillyboxTransferSingle)
				if err := _Lillybox.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 amount)
func (_Lillybox *LillyboxFilterer) ParseTransferSingle(log types.Log) (*LillyboxTransferSingle, error) {
	event := new(LillyboxTransferSingle)
	if err := _Lillybox.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LillyboxURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Lillybox contract.
type LillyboxURIIterator struct {
	Event *LillyboxURI // Event containing the contract specifics and raw log

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
func (it *LillyboxURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LillyboxURI)
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
		it.Event = new(LillyboxURI)
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
func (it *LillyboxURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LillyboxURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LillyboxURI represents a URI event raised by the Lillybox contract.
type LillyboxURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Lillybox *LillyboxFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*LillyboxURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Lillybox.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &LillyboxURIIterator{contract: _Lillybox.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Lillybox *LillyboxFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *LillyboxURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Lillybox.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LillyboxURI)
				if err := _Lillybox.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Lillybox *LillyboxFilterer) ParseURI(log types.Log) (*LillyboxURI, error) {
	event := new(LillyboxURI)
	if err := _Lillybox.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LillyboxUnStakingIterator is returned from FilterUnStaking and is used to iterate over the raw logs and unpacked data for UnStaking events raised by the Lillybox contract.
type LillyboxUnStakingIterator struct {
	Event *LillyboxUnStaking // Event containing the contract specifics and raw log

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
func (it *LillyboxUnStakingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LillyboxUnStaking)
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
		it.Event = new(LillyboxUnStaking)
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
func (it *LillyboxUnStakingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LillyboxUnStakingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LillyboxUnStaking represents a UnStaking event raised by the Lillybox contract.
type LillyboxUnStaking struct {
	User        common.Address
	Amount      *big.Int
	Blocknumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnStaking is a free log retrieval operation binding the contract event 0xc7c828c532c03d3105bb4d86d724a0b9aae548a0fb733a5eaee9fea00959c5e5.
//
// Solidity: event UnStaking(address indexed user, uint256 amount, uint256 blocknumber)
func (_Lillybox *LillyboxFilterer) FilterUnStaking(opts *bind.FilterOpts, user []common.Address) (*LillyboxUnStakingIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lillybox.contract.FilterLogs(opts, "UnStaking", userRule)
	if err != nil {
		return nil, err
	}
	return &LillyboxUnStakingIterator{contract: _Lillybox.contract, event: "UnStaking", logs: logs, sub: sub}, nil
}

// WatchUnStaking is a free log subscription operation binding the contract event 0xc7c828c532c03d3105bb4d86d724a0b9aae548a0fb733a5eaee9fea00959c5e5.
//
// Solidity: event UnStaking(address indexed user, uint256 amount, uint256 blocknumber)
func (_Lillybox *LillyboxFilterer) WatchUnStaking(opts *bind.WatchOpts, sink chan<- *LillyboxUnStaking, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lillybox.contract.WatchLogs(opts, "UnStaking", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LillyboxUnStaking)
				if err := _Lillybox.contract.UnpackLog(event, "UnStaking", log); err != nil {
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

// ParseUnStaking is a log parse operation binding the contract event 0xc7c828c532c03d3105bb4d86d724a0b9aae548a0fb733a5eaee9fea00959c5e5.
//
// Solidity: event UnStaking(address indexed user, uint256 amount, uint256 blocknumber)
func (_Lillybox *LillyboxFilterer) ParseUnStaking(log types.Log) (*LillyboxUnStaking, error) {
	event := new(LillyboxUnStaking)
	if err := _Lillybox.contract.UnpackLog(event, "UnStaking", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LillyboxWithdrawKlayIterator is returned from FilterWithdrawKlay and is used to iterate over the raw logs and unpacked data for WithdrawKlay events raised by the Lillybox contract.
type LillyboxWithdrawKlayIterator struct {
	Event *LillyboxWithdrawKlay // Event containing the contract specifics and raw log

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
func (it *LillyboxWithdrawKlayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LillyboxWithdrawKlay)
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
		it.Event = new(LillyboxWithdrawKlay)
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
func (it *LillyboxWithdrawKlayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LillyboxWithdrawKlayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LillyboxWithdrawKlay represents a WithdrawKlay event raised by the Lillybox contract.
type LillyboxWithdrawKlay struct {
	User        common.Address
	Amount      *big.Int
	UsedLIL     *big.Int
	Blocknumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawKlay is a free log retrieval operation binding the contract event 0xbf5c1db369f8ae0ee068ab9de68bb43360fd3fcc985d0d2cf98d4b1f6ce48306.
//
// Solidity: event WithdrawKlay(address indexed user, uint256 amount, uint256 usedLIL, uint256 blocknumber)
func (_Lillybox *LillyboxFilterer) FilterWithdrawKlay(opts *bind.FilterOpts, user []common.Address) (*LillyboxWithdrawKlayIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lillybox.contract.FilterLogs(opts, "WithdrawKlay", userRule)
	if err != nil {
		return nil, err
	}
	return &LillyboxWithdrawKlayIterator{contract: _Lillybox.contract, event: "WithdrawKlay", logs: logs, sub: sub}, nil
}

// WatchWithdrawKlay is a free log subscription operation binding the contract event 0xbf5c1db369f8ae0ee068ab9de68bb43360fd3fcc985d0d2cf98d4b1f6ce48306.
//
// Solidity: event WithdrawKlay(address indexed user, uint256 amount, uint256 usedLIL, uint256 blocknumber)
func (_Lillybox *LillyboxFilterer) WatchWithdrawKlay(opts *bind.WatchOpts, sink chan<- *LillyboxWithdrawKlay, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lillybox.contract.WatchLogs(opts, "WithdrawKlay", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LillyboxWithdrawKlay)
				if err := _Lillybox.contract.UnpackLog(event, "WithdrawKlay", log); err != nil {
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

// ParseWithdrawKlay is a log parse operation binding the contract event 0xbf5c1db369f8ae0ee068ab9de68bb43360fd3fcc985d0d2cf98d4b1f6ce48306.
//
// Solidity: event WithdrawKlay(address indexed user, uint256 amount, uint256 usedLIL, uint256 blocknumber)
func (_Lillybox *LillyboxFilterer) ParseWithdrawKlay(log types.Log) (*LillyboxWithdrawKlay, error) {
	event := new(LillyboxWithdrawKlay)
	if err := _Lillybox.contract.UnpackLog(event, "WithdrawKlay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
