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
	_ = abi.ConvertType
)

// IBioRegistryBiometricData is an auto generated low-level Go binding around an user-defined struct.
type IBioRegistryBiometricData struct {
	Uuid          string
	UserAddress   common.Address
	BiometricInfo []byte
	UserMetadata  string
}

// BioRegistryMetaData contains all meta data concerning the BioRegistry contract.
var BioRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractState\",\"name\":\"stateContract_\",\"type\":\"address\"},{\"internalType\":\"contractAccountFactory\",\"name\":\"accountFactory_\",\"type\":\"address\"}],\"name\":\"__BioRegistry_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountFactory\",\"outputs\":[{\"internalType\":\"contractAccountFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"issuerId_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"biometricInfo_\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"userAddress_\",\"type\":\"address\"}],\"name\":\"getMetadataByFAndUser\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"issuerId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"userAddress_\",\"type\":\"address\"}],\"name\":\"getMetadataByUser\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"getUserAccountByUUID\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"issuerId_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"getUserByUUID\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"biometricInfo\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"userMetadata\",\"type\":\"string\"}],\"internalType\":\"structIBioRegistry.BiometricData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"isUUIDRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"registerAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateContract\",\"outputs\":[{\"internalType\":\"contractState\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldState_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newState_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOldStateGenesis_\",\"type\":\"bool\"},{\"internalType\":\"uint256[2]\",\"name\":\"a_\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b_\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c_\",\"type\":\"uint256[2]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"biometricInfo\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"userMetadata\",\"type\":\"string\"}],\"internalType\":\"structIBioRegistry.BiometricData[]\",\"name\":\"biometricData_\",\"type\":\"tuple[]\"}],\"name\":\"transitStateAndBiometricData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// BioRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use BioRegistryMetaData.ABI instead.
var BioRegistryABI = BioRegistryMetaData.ABI

// BioRegistry is an auto generated Go binding around an Ethereum contract.
type BioRegistry struct {
	BioRegistryCaller     // Read-only binding to the contract
	BioRegistryTransactor // Write-only binding to the contract
	BioRegistryFilterer   // Log filterer for contract events
}

// BioRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BioRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BioRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BioRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BioRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BioRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BioRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BioRegistrySession struct {
	Contract     *BioRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BioRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BioRegistryCallerSession struct {
	Contract *BioRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BioRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BioRegistryTransactorSession struct {
	Contract     *BioRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BioRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BioRegistryRaw struct {
	Contract *BioRegistry // Generic contract binding to access the raw methods on
}

// BioRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BioRegistryCallerRaw struct {
	Contract *BioRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// BioRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BioRegistryTransactorRaw struct {
	Contract *BioRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBioRegistry creates a new instance of BioRegistry, bound to a specific deployed contract.
func NewBioRegistry(address common.Address, backend bind.ContractBackend) (*BioRegistry, error) {
	contract, err := bindBioRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BioRegistry{BioRegistryCaller: BioRegistryCaller{contract: contract}, BioRegistryTransactor: BioRegistryTransactor{contract: contract}, BioRegistryFilterer: BioRegistryFilterer{contract: contract}}, nil
}

// NewBioRegistryCaller creates a new read-only instance of BioRegistry, bound to a specific deployed contract.
func NewBioRegistryCaller(address common.Address, caller bind.ContractCaller) (*BioRegistryCaller, error) {
	contract, err := bindBioRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BioRegistryCaller{contract: contract}, nil
}

// NewBioRegistryTransactor creates a new write-only instance of BioRegistry, bound to a specific deployed contract.
func NewBioRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*BioRegistryTransactor, error) {
	contract, err := bindBioRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BioRegistryTransactor{contract: contract}, nil
}

// NewBioRegistryFilterer creates a new log filterer instance of BioRegistry, bound to a specific deployed contract.
func NewBioRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*BioRegistryFilterer, error) {
	contract, err := bindBioRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BioRegistryFilterer{contract: contract}, nil
}

// bindBioRegistry binds a generic wrapper to an already deployed contract.
func bindBioRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BioRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BioRegistry *BioRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BioRegistry.Contract.BioRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BioRegistry *BioRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BioRegistry.Contract.BioRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BioRegistry *BioRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BioRegistry.Contract.BioRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BioRegistry *BioRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BioRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BioRegistry *BioRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BioRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BioRegistry *BioRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BioRegistry.Contract.contract.Transact(opts, method, params...)
}

// AccountFactory is a free data retrieval call binding the contract method 0x687cd9c1.
//
// Solidity: function accountFactory() view returns(address)
func (_BioRegistry *BioRegistryCaller) AccountFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BioRegistry.contract.Call(opts, &out, "accountFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountFactory is a free data retrieval call binding the contract method 0x687cd9c1.
//
// Solidity: function accountFactory() view returns(address)
func (_BioRegistry *BioRegistrySession) AccountFactory() (common.Address, error) {
	return _BioRegistry.Contract.AccountFactory(&_BioRegistry.CallOpts)
}

// AccountFactory is a free data retrieval call binding the contract method 0x687cd9c1.
//
// Solidity: function accountFactory() view returns(address)
func (_BioRegistry *BioRegistryCallerSession) AccountFactory() (common.Address, error) {
	return _BioRegistry.Contract.AccountFactory(&_BioRegistry.CallOpts)
}

// GetMetadataByFAndUser is a free data retrieval call binding the contract method 0x1ad66555.
//
// Solidity: function getMetadataByFAndUser(uint256 issuerId_, bytes biometricInfo_, address userAddress_) view returns(string)
func (_BioRegistry *BioRegistryCaller) GetMetadataByFAndUser(opts *bind.CallOpts, issuerId_ *big.Int, biometricInfo_ []byte, userAddress_ common.Address) (string, error) {
	var out []interface{}
	err := _BioRegistry.contract.Call(opts, &out, "getMetadataByFAndUser", issuerId_, biometricInfo_, userAddress_)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMetadataByFAndUser is a free data retrieval call binding the contract method 0x1ad66555.
//
// Solidity: function getMetadataByFAndUser(uint256 issuerId_, bytes biometricInfo_, address userAddress_) view returns(string)
func (_BioRegistry *BioRegistrySession) GetMetadataByFAndUser(issuerId_ *big.Int, biometricInfo_ []byte, userAddress_ common.Address) (string, error) {
	return _BioRegistry.Contract.GetMetadataByFAndUser(&_BioRegistry.CallOpts, issuerId_, biometricInfo_, userAddress_)
}

// GetMetadataByFAndUser is a free data retrieval call binding the contract method 0x1ad66555.
//
// Solidity: function getMetadataByFAndUser(uint256 issuerId_, bytes biometricInfo_, address userAddress_) view returns(string)
func (_BioRegistry *BioRegistryCallerSession) GetMetadataByFAndUser(issuerId_ *big.Int, biometricInfo_ []byte, userAddress_ common.Address) (string, error) {
	return _BioRegistry.Contract.GetMetadataByFAndUser(&_BioRegistry.CallOpts, issuerId_, biometricInfo_, userAddress_)
}

// GetMetadataByUser is a free data retrieval call binding the contract method 0x773b3c36.
//
// Solidity: function getMetadataByUser(uint256 issuerId_, address userAddress_) view returns(string)
func (_BioRegistry *BioRegistryCaller) GetMetadataByUser(opts *bind.CallOpts, issuerId_ *big.Int, userAddress_ common.Address) (string, error) {
	var out []interface{}
	err := _BioRegistry.contract.Call(opts, &out, "getMetadataByUser", issuerId_, userAddress_)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMetadataByUser is a free data retrieval call binding the contract method 0x773b3c36.
//
// Solidity: function getMetadataByUser(uint256 issuerId_, address userAddress_) view returns(string)
func (_BioRegistry *BioRegistrySession) GetMetadataByUser(issuerId_ *big.Int, userAddress_ common.Address) (string, error) {
	return _BioRegistry.Contract.GetMetadataByUser(&_BioRegistry.CallOpts, issuerId_, userAddress_)
}

// GetMetadataByUser is a free data retrieval call binding the contract method 0x773b3c36.
//
// Solidity: function getMetadataByUser(uint256 issuerId_, address userAddress_) view returns(string)
func (_BioRegistry *BioRegistryCallerSession) GetMetadataByUser(issuerId_ *big.Int, userAddress_ common.Address) (string, error) {
	return _BioRegistry.Contract.GetMetadataByUser(&_BioRegistry.CallOpts, issuerId_, userAddress_)
}

// GetUserAccountByUUID is a free data retrieval call binding the contract method 0x0d16a096.
//
// Solidity: function getUserAccountByUUID(string uuid_) view returns(address)
func (_BioRegistry *BioRegistryCaller) GetUserAccountByUUID(opts *bind.CallOpts, uuid_ string) (common.Address, error) {
	var out []interface{}
	err := _BioRegistry.contract.Call(opts, &out, "getUserAccountByUUID", uuid_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUserAccountByUUID is a free data retrieval call binding the contract method 0x0d16a096.
//
// Solidity: function getUserAccountByUUID(string uuid_) view returns(address)
func (_BioRegistry *BioRegistrySession) GetUserAccountByUUID(uuid_ string) (common.Address, error) {
	return _BioRegistry.Contract.GetUserAccountByUUID(&_BioRegistry.CallOpts, uuid_)
}

// GetUserAccountByUUID is a free data retrieval call binding the contract method 0x0d16a096.
//
// Solidity: function getUserAccountByUUID(string uuid_) view returns(address)
func (_BioRegistry *BioRegistryCallerSession) GetUserAccountByUUID(uuid_ string) (common.Address, error) {
	return _BioRegistry.Contract.GetUserAccountByUUID(&_BioRegistry.CallOpts, uuid_)
}

// GetUserByUUID is a free data retrieval call binding the contract method 0x7bcc3698.
//
// Solidity: function getUserByUUID(uint256 issuerId_, string uuid_) view returns((string,address,bytes,string))
func (_BioRegistry *BioRegistryCaller) GetUserByUUID(opts *bind.CallOpts, issuerId_ *big.Int, uuid_ string) (IBioRegistryBiometricData, error) {
	var out []interface{}
	err := _BioRegistry.contract.Call(opts, &out, "getUserByUUID", issuerId_, uuid_)

	if err != nil {
		return *new(IBioRegistryBiometricData), err
	}

	out0 := *abi.ConvertType(out[0], new(IBioRegistryBiometricData)).(*IBioRegistryBiometricData)

	return out0, err

}

// GetUserByUUID is a free data retrieval call binding the contract method 0x7bcc3698.
//
// Solidity: function getUserByUUID(uint256 issuerId_, string uuid_) view returns((string,address,bytes,string))
func (_BioRegistry *BioRegistrySession) GetUserByUUID(issuerId_ *big.Int, uuid_ string) (IBioRegistryBiometricData, error) {
	return _BioRegistry.Contract.GetUserByUUID(&_BioRegistry.CallOpts, issuerId_, uuid_)
}

// GetUserByUUID is a free data retrieval call binding the contract method 0x7bcc3698.
//
// Solidity: function getUserByUUID(uint256 issuerId_, string uuid_) view returns((string,address,bytes,string))
func (_BioRegistry *BioRegistryCallerSession) GetUserByUUID(issuerId_ *big.Int, uuid_ string) (IBioRegistryBiometricData, error) {
	return _BioRegistry.Contract.GetUserByUUID(&_BioRegistry.CallOpts, issuerId_, uuid_)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_BioRegistry *BioRegistryCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BioRegistry.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_BioRegistry *BioRegistrySession) Implementation() (common.Address, error) {
	return _BioRegistry.Contract.Implementation(&_BioRegistry.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_BioRegistry *BioRegistryCallerSession) Implementation() (common.Address, error) {
	return _BioRegistry.Contract.Implementation(&_BioRegistry.CallOpts)
}

// IsUUIDRegistered is a free data retrieval call binding the contract method 0xed8a138f.
//
// Solidity: function isUUIDRegistered(string uuid_) view returns(bool)
func (_BioRegistry *BioRegistryCaller) IsUUIDRegistered(opts *bind.CallOpts, uuid_ string) (bool, error) {
	var out []interface{}
	err := _BioRegistry.contract.Call(opts, &out, "isUUIDRegistered", uuid_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUUIDRegistered is a free data retrieval call binding the contract method 0xed8a138f.
//
// Solidity: function isUUIDRegistered(string uuid_) view returns(bool)
func (_BioRegistry *BioRegistrySession) IsUUIDRegistered(uuid_ string) (bool, error) {
	return _BioRegistry.Contract.IsUUIDRegistered(&_BioRegistry.CallOpts, uuid_)
}

// IsUUIDRegistered is a free data retrieval call binding the contract method 0xed8a138f.
//
// Solidity: function isUUIDRegistered(string uuid_) view returns(bool)
func (_BioRegistry *BioRegistryCallerSession) IsUUIDRegistered(uuid_ string) (bool, error) {
	return _BioRegistry.Contract.IsUUIDRegistered(&_BioRegistry.CallOpts, uuid_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BioRegistry *BioRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BioRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BioRegistry *BioRegistrySession) Owner() (common.Address, error) {
	return _BioRegistry.Contract.Owner(&_BioRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BioRegistry *BioRegistryCallerSession) Owner() (common.Address, error) {
	return _BioRegistry.Contract.Owner(&_BioRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BioRegistry *BioRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BioRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BioRegistry *BioRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _BioRegistry.Contract.ProxiableUUID(&_BioRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BioRegistry *BioRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BioRegistry.Contract.ProxiableUUID(&_BioRegistry.CallOpts)
}

// StateContract is a free data retrieval call binding the contract method 0xdb5f30e1.
//
// Solidity: function stateContract() view returns(address)
func (_BioRegistry *BioRegistryCaller) StateContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BioRegistry.contract.Call(opts, &out, "stateContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StateContract is a free data retrieval call binding the contract method 0xdb5f30e1.
//
// Solidity: function stateContract() view returns(address)
func (_BioRegistry *BioRegistrySession) StateContract() (common.Address, error) {
	return _BioRegistry.Contract.StateContract(&_BioRegistry.CallOpts)
}

// StateContract is a free data retrieval call binding the contract method 0xdb5f30e1.
//
// Solidity: function stateContract() view returns(address)
func (_BioRegistry *BioRegistryCallerSession) StateContract() (common.Address, error) {
	return _BioRegistry.Contract.StateContract(&_BioRegistry.CallOpts)
}

// BioRegistryInit is a paid mutator transaction binding the contract method 0x3404748f.
//
// Solidity: function __BioRegistry_init(address stateContract_, address accountFactory_) returns()
func (_BioRegistry *BioRegistryTransactor) BioRegistryInit(opts *bind.TransactOpts, stateContract_ common.Address, accountFactory_ common.Address) (*types.Transaction, error) {
	return _BioRegistry.contract.Transact(opts, "__BioRegistry_init", stateContract_, accountFactory_)
}

// BioRegistryInit is a paid mutator transaction binding the contract method 0x3404748f.
//
// Solidity: function __BioRegistry_init(address stateContract_, address accountFactory_) returns()
func (_BioRegistry *BioRegistrySession) BioRegistryInit(stateContract_ common.Address, accountFactory_ common.Address) (*types.Transaction, error) {
	return _BioRegistry.Contract.BioRegistryInit(&_BioRegistry.TransactOpts, stateContract_, accountFactory_)
}

// BioRegistryInit is a paid mutator transaction binding the contract method 0x3404748f.
//
// Solidity: function __BioRegistry_init(address stateContract_, address accountFactory_) returns()
func (_BioRegistry *BioRegistryTransactorSession) BioRegistryInit(stateContract_ common.Address, accountFactory_ common.Address) (*types.Transaction, error) {
	return _BioRegistry.Contract.BioRegistryInit(&_BioRegistry.TransactOpts, stateContract_, accountFactory_)
}

// RegisterAccount is a paid mutator transaction binding the contract method 0x95ad1543.
//
// Solidity: function registerAccount(string uuid_, address account_) returns()
func (_BioRegistry *BioRegistryTransactor) RegisterAccount(opts *bind.TransactOpts, uuid_ string, account_ common.Address) (*types.Transaction, error) {
	return _BioRegistry.contract.Transact(opts, "registerAccount", uuid_, account_)
}

// RegisterAccount is a paid mutator transaction binding the contract method 0x95ad1543.
//
// Solidity: function registerAccount(string uuid_, address account_) returns()
func (_BioRegistry *BioRegistrySession) RegisterAccount(uuid_ string, account_ common.Address) (*types.Transaction, error) {
	return _BioRegistry.Contract.RegisterAccount(&_BioRegistry.TransactOpts, uuid_, account_)
}

// RegisterAccount is a paid mutator transaction binding the contract method 0x95ad1543.
//
// Solidity: function registerAccount(string uuid_, address account_) returns()
func (_BioRegistry *BioRegistryTransactorSession) RegisterAccount(uuid_ string, account_ common.Address) (*types.Transaction, error) {
	return _BioRegistry.Contract.RegisterAccount(&_BioRegistry.TransactOpts, uuid_, account_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BioRegistry *BioRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BioRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BioRegistry *BioRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _BioRegistry.Contract.RenounceOwnership(&_BioRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BioRegistry *BioRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BioRegistry.Contract.RenounceOwnership(&_BioRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BioRegistry *BioRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BioRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BioRegistry *BioRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BioRegistry.Contract.TransferOwnership(&_BioRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BioRegistry *BioRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BioRegistry.Contract.TransferOwnership(&_BioRegistry.TransactOpts, newOwner)
}

// TransitStateAndBiometricData is a paid mutator transaction binding the contract method 0x3711d483.
//
// Solidity: function transitStateAndBiometricData(uint256 id_, uint256 oldState_, uint256 newState_, bool isOldStateGenesis_, uint256[2] a_, uint256[2][2] b_, uint256[2] c_, (string,address,bytes,string)[] biometricData_) returns()
func (_BioRegistry *BioRegistryTransactor) TransitStateAndBiometricData(opts *bind.TransactOpts, id_ *big.Int, oldState_ *big.Int, newState_ *big.Int, isOldStateGenesis_ bool, a_ [2]*big.Int, b_ [2][2]*big.Int, c_ [2]*big.Int, biometricData_ []IBioRegistryBiometricData) (*types.Transaction, error) {
	return _BioRegistry.contract.Transact(opts, "transitStateAndBiometricData", id_, oldState_, newState_, isOldStateGenesis_, a_, b_, c_, biometricData_)
}

// TransitStateAndBiometricData is a paid mutator transaction binding the contract method 0x3711d483.
//
// Solidity: function transitStateAndBiometricData(uint256 id_, uint256 oldState_, uint256 newState_, bool isOldStateGenesis_, uint256[2] a_, uint256[2][2] b_, uint256[2] c_, (string,address,bytes,string)[] biometricData_) returns()
func (_BioRegistry *BioRegistrySession) TransitStateAndBiometricData(id_ *big.Int, oldState_ *big.Int, newState_ *big.Int, isOldStateGenesis_ bool, a_ [2]*big.Int, b_ [2][2]*big.Int, c_ [2]*big.Int, biometricData_ []IBioRegistryBiometricData) (*types.Transaction, error) {
	return _BioRegistry.Contract.TransitStateAndBiometricData(&_BioRegistry.TransactOpts, id_, oldState_, newState_, isOldStateGenesis_, a_, b_, c_, biometricData_)
}

// TransitStateAndBiometricData is a paid mutator transaction binding the contract method 0x3711d483.
//
// Solidity: function transitStateAndBiometricData(uint256 id_, uint256 oldState_, uint256 newState_, bool isOldStateGenesis_, uint256[2] a_, uint256[2][2] b_, uint256[2] c_, (string,address,bytes,string)[] biometricData_) returns()
func (_BioRegistry *BioRegistryTransactorSession) TransitStateAndBiometricData(id_ *big.Int, oldState_ *big.Int, newState_ *big.Int, isOldStateGenesis_ bool, a_ [2]*big.Int, b_ [2][2]*big.Int, c_ [2]*big.Int, biometricData_ []IBioRegistryBiometricData) (*types.Transaction, error) {
	return _BioRegistry.Contract.TransitStateAndBiometricData(&_BioRegistry.TransactOpts, id_, oldState_, newState_, isOldStateGenesis_, a_, b_, c_, biometricData_)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BioRegistry *BioRegistryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _BioRegistry.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BioRegistry *BioRegistrySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _BioRegistry.Contract.UpgradeTo(&_BioRegistry.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BioRegistry *BioRegistryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _BioRegistry.Contract.UpgradeTo(&_BioRegistry.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BioRegistry *BioRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BioRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BioRegistry *BioRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BioRegistry.Contract.UpgradeToAndCall(&_BioRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BioRegistry *BioRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BioRegistry.Contract.UpgradeToAndCall(&_BioRegistry.TransactOpts, newImplementation, data)
}

// BioRegistryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the BioRegistry contract.
type BioRegistryAdminChangedIterator struct {
	Event *BioRegistryAdminChanged // Event containing the contract specifics and raw log

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
func (it *BioRegistryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BioRegistryAdminChanged)
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
		it.Event = new(BioRegistryAdminChanged)
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
func (it *BioRegistryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BioRegistryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BioRegistryAdminChanged represents a AdminChanged event raised by the BioRegistry contract.
type BioRegistryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BioRegistry *BioRegistryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*BioRegistryAdminChangedIterator, error) {

	logs, sub, err := _BioRegistry.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &BioRegistryAdminChangedIterator{contract: _BioRegistry.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BioRegistry *BioRegistryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *BioRegistryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _BioRegistry.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BioRegistryAdminChanged)
				if err := _BioRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BioRegistry *BioRegistryFilterer) ParseAdminChanged(log types.Log) (*BioRegistryAdminChanged, error) {
	event := new(BioRegistryAdminChanged)
	if err := _BioRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BioRegistryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the BioRegistry contract.
type BioRegistryBeaconUpgradedIterator struct {
	Event *BioRegistryBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *BioRegistryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BioRegistryBeaconUpgraded)
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
		it.Event = new(BioRegistryBeaconUpgraded)
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
func (it *BioRegistryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BioRegistryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BioRegistryBeaconUpgraded represents a BeaconUpgraded event raised by the BioRegistry contract.
type BioRegistryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BioRegistry *BioRegistryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*BioRegistryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _BioRegistry.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &BioRegistryBeaconUpgradedIterator{contract: _BioRegistry.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BioRegistry *BioRegistryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *BioRegistryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _BioRegistry.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BioRegistryBeaconUpgraded)
				if err := _BioRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BioRegistry *BioRegistryFilterer) ParseBeaconUpgraded(log types.Log) (*BioRegistryBeaconUpgraded, error) {
	event := new(BioRegistryBeaconUpgraded)
	if err := _BioRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BioRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BioRegistry contract.
type BioRegistryInitializedIterator struct {
	Event *BioRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *BioRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BioRegistryInitialized)
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
		it.Event = new(BioRegistryInitialized)
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
func (it *BioRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BioRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BioRegistryInitialized represents a Initialized event raised by the BioRegistry contract.
type BioRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BioRegistry *BioRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*BioRegistryInitializedIterator, error) {

	logs, sub, err := _BioRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BioRegistryInitializedIterator{contract: _BioRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BioRegistry *BioRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BioRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _BioRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BioRegistryInitialized)
				if err := _BioRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BioRegistry *BioRegistryFilterer) ParseInitialized(log types.Log) (*BioRegistryInitialized, error) {
	event := new(BioRegistryInitialized)
	if err := _BioRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BioRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BioRegistry contract.
type BioRegistryOwnershipTransferredIterator struct {
	Event *BioRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BioRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BioRegistryOwnershipTransferred)
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
		it.Event = new(BioRegistryOwnershipTransferred)
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
func (it *BioRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BioRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BioRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the BioRegistry contract.
type BioRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BioRegistry *BioRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BioRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BioRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BioRegistryOwnershipTransferredIterator{contract: _BioRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BioRegistry *BioRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BioRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BioRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BioRegistryOwnershipTransferred)
				if err := _BioRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BioRegistry *BioRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*BioRegistryOwnershipTransferred, error) {
	event := new(BioRegistryOwnershipTransferred)
	if err := _BioRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BioRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BioRegistry contract.
type BioRegistryUpgradedIterator struct {
	Event *BioRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *BioRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BioRegistryUpgraded)
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
		it.Event = new(BioRegistryUpgraded)
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
func (it *BioRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BioRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BioRegistryUpgraded represents a Upgraded event raised by the BioRegistry contract.
type BioRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BioRegistry *BioRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BioRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BioRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BioRegistryUpgradedIterator{contract: _BioRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BioRegistry *BioRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BioRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BioRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BioRegistryUpgraded)
				if err := _BioRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BioRegistry *BioRegistryFilterer) ParseUpgraded(log types.Log) (*BioRegistryUpgraded, error) {
	event := new(BioRegistryUpgraded)
	if err := _BioRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
