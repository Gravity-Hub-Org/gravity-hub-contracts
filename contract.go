// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ISubscriptionABI is the input ABI used to generate the binding from.
const ISubscriptionABI = "[{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"data\",\"type\":\"uint256[]\"}],\"name\":\"attachData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ISubscriptionFuncSigs maps the 4-byte function signature to its string representation.
var ISubscriptionFuncSigs = map[string]string{
	"297d5f3f": "attachData(uint256[])",
}

// ISubscription is an auto generated Go binding around an Ethereum contract.
type ISubscription struct {
	ISubscriptionCaller     // Read-only binding to the contract
	ISubscriptionTransactor // Write-only binding to the contract
	ISubscriptionFilterer   // Log filterer for contract events
}

// ISubscriptionCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISubscriptionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISubscriptionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISubscriptionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISubscriptionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISubscriptionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISubscriptionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISubscriptionSession struct {
	Contract     *ISubscription    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISubscriptionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISubscriptionCallerSession struct {
	Contract *ISubscriptionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ISubscriptionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISubscriptionTransactorSession struct {
	Contract     *ISubscriptionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ISubscriptionRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISubscriptionRaw struct {
	Contract *ISubscription // Generic contract binding to access the raw methods on
}

// ISubscriptionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISubscriptionCallerRaw struct {
	Contract *ISubscriptionCaller // Generic read-only contract binding to access the raw methods on
}

// ISubscriptionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISubscriptionTransactorRaw struct {
	Contract *ISubscriptionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISubscription creates a new instance of ISubscription, bound to a specific deployed contract.
func NewISubscription(address common.Address, backend bind.ContractBackend) (*ISubscription, error) {
	contract, err := bindISubscription(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISubscription{ISubscriptionCaller: ISubscriptionCaller{contract: contract}, ISubscriptionTransactor: ISubscriptionTransactor{contract: contract}, ISubscriptionFilterer: ISubscriptionFilterer{contract: contract}}, nil
}

// NewISubscriptionCaller creates a new read-only instance of ISubscription, bound to a specific deployed contract.
func NewISubscriptionCaller(address common.Address, caller bind.ContractCaller) (*ISubscriptionCaller, error) {
	contract, err := bindISubscription(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISubscriptionCaller{contract: contract}, nil
}

// NewISubscriptionTransactor creates a new write-only instance of ISubscription, bound to a specific deployed contract.
func NewISubscriptionTransactor(address common.Address, transactor bind.ContractTransactor) (*ISubscriptionTransactor, error) {
	contract, err := bindISubscription(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISubscriptionTransactor{contract: contract}, nil
}

// NewISubscriptionFilterer creates a new log filterer instance of ISubscription, bound to a specific deployed contract.
func NewISubscriptionFilterer(address common.Address, filterer bind.ContractFilterer) (*ISubscriptionFilterer, error) {
	contract, err := bindISubscription(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISubscriptionFilterer{contract: contract}, nil
}

// bindISubscription binds a generic wrapper to an already deployed contract.
func bindISubscription(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISubscriptionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISubscription *ISubscriptionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ISubscription.Contract.ISubscriptionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISubscription *ISubscriptionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISubscription.Contract.ISubscriptionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISubscription *ISubscriptionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISubscription.Contract.ISubscriptionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISubscription *ISubscriptionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ISubscription.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISubscription *ISubscriptionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISubscription.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISubscription *ISubscriptionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISubscription.Contract.contract.Transact(opts, method, params...)
}

// AttachData is a paid mutator transaction binding the contract method 0x297d5f3f.
//
// Solidity: function attachData(uint256[] data) returns()
func (_ISubscription *ISubscriptionTransactor) AttachData(opts *bind.TransactOpts, data []*big.Int) (*types.Transaction, error) {
	return _ISubscription.contract.Transact(opts, "attachData", data)
}

// AttachData is a paid mutator transaction binding the contract method 0x297d5f3f.
//
// Solidity: function attachData(uint256[] data) returns()
func (_ISubscription *ISubscriptionSession) AttachData(data []*big.Int) (*types.Transaction, error) {
	return _ISubscription.Contract.AttachData(&_ISubscription.TransactOpts, data)
}

// AttachData is a paid mutator transaction binding the contract method 0x297d5f3f.
//
// Solidity: function attachData(uint256[] data) returns()
func (_ISubscription *ISubscriptionTransactorSession) AttachData(data []*big.Int) (*types.Transaction, error) {
	return _ISubscription.Contract.AttachData(&_ISubscription.TransactOpts, data)
}

// ModelsABI is the input ABI used to generate the binding from.
const ModelsABI = "[]"

// ModelsBin is the compiled bytecode used for deploying new contracts.
var ModelsBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206110ce79fac950e740e909bcecfe970e320b833b04a38939aa6a5ed9a40f11f164736f6c63430006060033"

// DeployModels deploys a new Ethereum contract, binding an instance of Models to it.
func DeployModels(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Models, error) {
	parsed, err := abi.JSON(strings.NewReader(ModelsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ModelsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Models{ModelsCaller: ModelsCaller{contract: contract}, ModelsTransactor: ModelsTransactor{contract: contract}, ModelsFilterer: ModelsFilterer{contract: contract}}, nil
}

// Models is an auto generated Go binding around an Ethereum contract.
type Models struct {
	ModelsCaller     // Read-only binding to the contract
	ModelsTransactor // Write-only binding to the contract
	ModelsFilterer   // Log filterer for contract events
}

// ModelsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModelsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModelsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModelsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModelsSession struct {
	Contract     *Models           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModelsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModelsCallerSession struct {
	Contract *ModelsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ModelsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModelsTransactorSession struct {
	Contract     *ModelsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModelsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModelsRaw struct {
	Contract *Models // Generic contract binding to access the raw methods on
}

// ModelsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModelsCallerRaw struct {
	Contract *ModelsCaller // Generic read-only contract binding to access the raw methods on
}

// ModelsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModelsTransactorRaw struct {
	Contract *ModelsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewModels creates a new instance of Models, bound to a specific deployed contract.
func NewModels(address common.Address, backend bind.ContractBackend) (*Models, error) {
	contract, err := bindModels(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Models{ModelsCaller: ModelsCaller{contract: contract}, ModelsTransactor: ModelsTransactor{contract: contract}, ModelsFilterer: ModelsFilterer{contract: contract}}, nil
}

// NewModelsCaller creates a new read-only instance of Models, bound to a specific deployed contract.
func NewModelsCaller(address common.Address, caller bind.ContractCaller) (*ModelsCaller, error) {
	contract, err := bindModels(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModelsCaller{contract: contract}, nil
}

// NewModelsTransactor creates a new write-only instance of Models, bound to a specific deployed contract.
func NewModelsTransactor(address common.Address, transactor bind.ContractTransactor) (*ModelsTransactor, error) {
	contract, err := bindModels(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModelsTransactor{contract: contract}, nil
}

// NewModelsFilterer creates a new log filterer instance of Models, bound to a specific deployed contract.
func NewModelsFilterer(address common.Address, filterer bind.ContractFilterer) (*ModelsFilterer, error) {
	contract, err := bindModels(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModelsFilterer{contract: contract}, nil
}

// bindModels binds a generic wrapper to an already deployed contract.
func bindModels(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ModelsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Models *ModelsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Models.Contract.ModelsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Models *ModelsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Models.Contract.ModelsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Models *ModelsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Models.Contract.ModelsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Models *ModelsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Models.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Models *ModelsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Models.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Models *ModelsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Models.Contract.contract.Transact(opts, method, params...)
}

// NebulaABI is the input ABI used to generate the binding from.
const NebulaABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"newOracle\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"data\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"confirmData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentOracleId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endEpochHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isPublseSubSent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"oracleInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOnline\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"idInQueue\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleQueue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"first\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"last\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"oracles\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oraclesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pulseQueue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"first\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"last\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pulses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"confirmationCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"subscriptionId\",\"type\":\"bytes32\"}],\"name\":\"sendData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"minConfirmations\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"subscribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"subscriptions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"minConfirmations\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subscriptionsQueue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"first\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"last\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trunOffOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trunOnOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"unsubscribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// NebulaFuncSigs maps the 4-byte function signature to its string representation.
var NebulaFuncSigs = map[string]string{
	"02cb72d3": "confirmData(uint256[],uint8[],bytes32[],bytes32[])",
	"76671808": "currentEpoch()",
	"524ce31f": "currentOracleId()",
	"996211a5": "endEpochHeight()",
	"6148d3f3": "isPublseSubSent(uint256,bytes32)",
	"06fdde03": "name()",
	"c5b1d9aa": "newRound()",
	"bc623feb": "oracleInfo(address)",
	"69a4246d": "oracleQueue()",
	"a81a2677": "oracles(bytes32)",
	"76f70900": "oraclesCount()",
	"1d11f944": "pulseQueue()",
	"0694fbb3": "pulses(uint256)",
	"342e9741": "sendData(uint256,bytes32)",
	"3527715d": "subscribe(address,uint8,uint256)",
	"94259c6c": "subscriptions(bytes32)",
	"b48a9c9b": "subscriptionsQueue()",
	"0a0d00d1": "trunOffOracle()",
	"51d81e09": "trunOnOracle()",
	"f8c60146": "unsubscribe(bytes32)",
}

// NebulaBin is the compiled bytecode used for deploying new contracts.
var NebulaBin = "0x60806040523480156200001157600080fd5b50604051620018b9380380620018b9833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010a57600080fd5b9083019060208201858111156200012057600080fd5b82518660208202830111640100000000821117156200013e57600080fd5b82525081516020918201928201910280838360005b838110156200016d57818101518382015260200162000153565b505050509190910160405250508351620001919250600c9150602085019062000377565b508051600d5560005b81518110156200036e576000828281518110620001b357fe5b602002602001015160405160200180826001600160a01b03166001600160a01b0316815260200191505060405160208183030381529060405280519060200120905073__$289c8261718617bebf80053c733e0efd61$__63a506d9546000836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b1580156200024e57600080fd5b505af415801562000263573d6000803e3d6000fd5b505050508282815181106200027457fe5b60200260200101516011600083815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b031602179055506040518060600160405280848481518110620002cb57fe5b60200260200101516001600160a01b031681526020016001151581526020018281525060126000858581518110620002ff57fe5b6020908102919091018101516001600160a01b03908116835282820193909352604091820160002084518154928601516001600160a01b031990931694169390931760ff60a01b1916600160a01b9115159190910217825591909101516001918201559190910190506200019a565b5050506200041c565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003ba57805160ff1916838001178555620003ea565b82800160010185558215620003ea579182015b82811115620003ea578251825591602001919060010190620003cd565b50620003f8929150620003fc565b5090565b6200041991905b80821115620003f8576000815560010162000403565b90565b61148d806200042c6000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c806369a4246d116100ad578063a81a267711610071578063a81a26771461053f578063b48a9c9b14610578578063bc623feb14610580578063c5b1d9aa146105ce578063f8c60146146105d65761012c565b806369a4246d146104ce57806376671808146104d657806376f70900146104de57806394259c6c146104e6578063996211a5146105375761012c565b8063342e9741116100f4578063342e97411461042f5780633527715d1461045257806351d81e0914610487578063524ce31f1461048f5780636148d3f3146104975761012c565b806302cb72d3146101315780630694fbb31461035a57806306fdde03146103895780630a0d00d1146104065780631d11f9441461040e575b600080fd5b6103586004803603608081101561014757600080fd5b810190602081018135600160201b81111561016157600080fd5b82018360208201111561017357600080fd5b803590602001918460208302840111600160201b8311171561019457600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156101e357600080fd5b8201836020820111156101f557600080fd5b803590602001918460208302840111600160201b8311171561021657600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561026557600080fd5b82018360208201111561027757600080fd5b803590602001918460208302840111600160201b8311171561029857600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156102e757600080fd5b8201836020820111156102f957600080fd5b803590602001918460208302840111600160201b8311171561031a57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506105f3945050505050565b005b6103776004803603602081101561037057600080fd5b50356109b6565b60408051918252519081900360200190f35b6103916109cb565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103cb5781810151838201526020016103b3565b50505050905090810190601f1680156103f85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610358610a59565b610416610b55565b6040805192835260208301919091528051918290030190f35b6103586004803603604081101561044557600080fd5b5080359060200135610b5e565b6103586004803603606081101561046857600080fd5b506001600160a01b038135169060ff6020820135169060400135610d42565b610358610f67565b6103776110c9565b6104ba600480360360408110156104ad57600080fd5b50803590602001356110cf565b604080519115158252519081900360200190f35b6104166110ef565b6103776110f8565b6103776110fe565b610503600480360360208110156104fc57600080fd5b5035611104565b604080516001600160a01b03958616815293909416602084015260ff90911682840152606082015290519081900360800190f35b61037761113d565b61055c6004803603602081101561055557600080fd5b5035611143565b604080516001600160a01b039092168252519081900360200190f35b61041661115e565b6105a66004803603602081101561059657600080fd5b50356001600160a01b0316611167565b604080516001600160a01b039094168452911515602084015282820152519081900360600190f35b610358611196565b610358600480360360208110156105ec57600080fd5b5035611295565b60006105fd611359565b600e5460005b60058110156106d0576000828152601160205260409020546001600160a01b031683826005811061063057fe5b6001600160a01b03909216602092830291909101526040805163870a61ff60e01b81526000600482015260248101859052905173__$289c8261718617bebf80053c733e0efd61$__9263870a61ff9260448082019391829003018186803b15801561069a57600080fd5b505af41580156106ae573d6000803e3d6000fd5b505050506040513d60208110156106c457600080fd5b50519150600101610603565b5060006005430690508281600581106106e557fe5b60200201516001600160a01b0316336001600160a01b03161461073f576040805162461bcd60e51b815260206004820152600d60248201526c34b73b30b634b21037bbb732b960991b604482015290519081900360640190fd5b60005b60058110156109495783816005811061075757fe5b60200201516001600160a01b03166001600c8b848151811061077557fe5b602002602001015160405160200180838054600181600116156101000203166002900480156107db5780601f106107b95761010080835404028352918201916107db565b820191906000526020600020905b8154815290600101906020018083116107c7575b505091825250604080518083038152602080840183527f19457468657265756d205369676e6564204d6573736167653a0a33330000000092840192835281519194509192605c01918401908083835b602083106108495780518252601f19909201916020918201910161082a565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001208a848151811061089057fe5b60200260200101518a85815181106108a457fe5b60200260200101518a86815181106108b857fe5b602002602001015160405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610917573d6000803e3d6000fd5b505050602060405103516001600160a01b031614610936576000610939565b60015b60ff169490940193600101610742565b50604080516060810182528981526020808201869052818301879052436000908152601482529290922081518051929391926109889284920190611377565b50602082015161099e90600183019060056113c2565b50604082015181600601559050505050505050505050565b60146020526000908152604090206006015481565b600c805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610a515780601f10610a2657610100808354040283529160200191610a51565b820191906000526020600020905b815481529060010190602001808311610a3457829003601f168201915b505050505081565b336000908152601260205260409020546001600160a01b031615610ab5576040805162461bcd60e51b815260206004820152600e60248201526d34b73b30b634b21039b2b73232b960911b604482015290519081900360640190fd5b33600090815260126020526040808220805460ff60a01b19168155600101548151639d6ad84b60e01b8152600481018490526024810191909152905173__$289c8261718617bebf80053c733e0efd61$__92639d6ad84b9260448082019391829003018186803b158015610b2857600080fd5b505af4158015610b3c573d6000803e3d6000fd5b5050336000908152601260205260408120600101555050565b60085460095482565b43600101821115610bad576040805162461bcd60e51b815260206004820152601460248201527334b73b30b634b210313637b1b590373ab6b132b960611b604482015290519081900360640190fd5b600082815260156020908152604080832084845290915290205460ff1615610c07576040805162461bcd60e51b81526020600482015260086024820152671cdd58881cd95b9d60c21b604482015290519081900360640190fd5b60008281526015602090815260408083208484528252808320805460ff191660019081179091556013835281842001548584526014835292819020905163297d5f3f60e01b81526004810192835281546024820181905247946001600160a01b03169363297d5f3f9392909182916044019084908015610ca657602002820191906000526020600020905b815481526020019060010190808311610c92575b505092505050600060405180830381600087803b158015610cc657600080fd5b505af1158015610cda573d6000803e3d6000fd5b50505060008381526013602052604090206002015447915082820390811015610d3b576040805162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a59081c995dd85c9960921b604482015290519081900360640190fd5b5050505050565b60408051600080356001600160e01b03191660208084019190915233606090811b602485015287901b6bffffffffffffffffffffffff1916603884015260f886901b6001600160f81b031916604c8401528351808403602d018152604d84019094528351919392606d019182918401908083835b60208310610dd55780518252601f199092019160209182019101610db6565b51815160209384036101000a60001901801990921691161790526040805192909401828103601f1901835284528151918101919091206000818152601390925292902054919450506001600160a01b0316159150610e6a9050576040805162461bcd60e51b815260206004820152600b60248201526a1c9c481a5cc8195e1a5cdd60aa1b604482015290519081900360640190fd5b604080516080810182523381526001600160a01b03868116602080840191825260ff8881168587019081526060860189815260008981526013909452878420965187546001600160a01b0319908116918816919091178855945160018801805493519390961696169590951760ff60a01b1916600160a01b91909216021790915590516002909201919091558151632941b65560e21b815260048082015260248101849052915173__$289c8261718617bebf80053c733e0efd61$__9263a506d954926044808301939192829003018186803b158015610f4957600080fd5b505af4158015610f5d573d6000803e3d6000fd5b5050505050505050565b336000908152601260205260409020546001600160a01b031615610fc3576040805162461bcd60e51b815260206004820152600e60248201526d34b73b30b634b21039b2b73232b960911b604482015290519081900360640190fd5b3360009081526012602090815260408083205481516001600160a01b039091168184015281518082038401815281830180845281519190940120632941b65560e21b90935260448101849052606481018390529051919273__$289c8261718617bebf80053c733e0efd61$__9263a506d954926084808201939291829003018186803b15801561105257600080fd5b505af4158015611066573d6000803e3d6000fd5b505033600081815260126020818152604080842080549885526011835290842080546001600160a01b0319166001600160a01b039990991698909817909755928252909152835460ff60a01b1916600160a01b1784556001909301929092555050565b600e5481565b601560209081526000928352604080842090915290825290205460ff1681565b60005460015482565b600f5481565b600d5481565b6013602052600090815260409020805460018201546002909201546001600160a01b039182169291821691600160a01b900460ff169084565b60105481565b6011602052600090815260409020546001600160a01b031681565b60045460055482565b601260205260009081526040902080546001909101546001600160a01b03821691600160a01b900460ff169083565b60105443116111e1576040805162461bcd60e51b81526020600482015260126024820152711c9bdd5b99081a5cc81b9bdd08195b99195960721b604482015290519081900360640190fd5b600f8054600190810190915554600e54141561120257600054600e5561128c565b73__$289c8261718617bebf80053c733e0efd61$__63870a61ff6000600e546040518363ffffffff1660e01b8152600401808381526020018281526020019250505060206040518083038186803b15801561125c57600080fd5b505af4158015611270573d6000803e3d6000fd5b505050506040513d602081101561128657600080fd5b5051600e555b43600a01601055565b6000818152601360205260409020546001600160a01b031633146112ed576040805162461bcd60e51b815260206004820152600a602482015269696e76616c696420727160b01b604482015290519081900360640190fd5b73__$289c8261718617bebf80053c733e0efd61$__639d6ad84b6004836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b15801561134557600080fd5b505af4158015610d3b573d6000803e3d6000fd5b6040518060a001604052806005906020820280368337509192915050565b8280548282559060005260206000209081019282156113b2579160200282015b828111156113b2578251825591602001919060010190611397565b506113be929150611416565b5090565b826005810192821561140a579160200282015b8281111561140a57825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906113d5565b506113be929150611433565b61143091905b808211156113be576000815560010161141c565b90565b61143091905b808211156113be5780546001600160a01b031916815560010161143956fea2646970667358221220fd74799930e20834e8cb7266595cca68ad8e5c58421c547cb99ce8d3d85fe2d964736f6c63430006060033"

// DeployNebula deploys a new Ethereum contract, binding an instance of Nebula to it.
func DeployNebula(auth *bind.TransactOpts, backend bind.ContractBackend, newName string, newOracle []common.Address) (common.Address, *types.Transaction, *Nebula, error) {
	parsed, err := abi.JSON(strings.NewReader(NebulaABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	queueLibAddr, _, _, _ := DeployQueueLib(auth, backend)
	NebulaBin = strings.Replace(NebulaBin, "__$289c8261718617bebf80053c733e0efd61$__", queueLibAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NebulaBin), backend, newName, newOracle)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Nebula{NebulaCaller: NebulaCaller{contract: contract}, NebulaTransactor: NebulaTransactor{contract: contract}, NebulaFilterer: NebulaFilterer{contract: contract}}, nil
}

// Nebula is an auto generated Go binding around an Ethereum contract.
type Nebula struct {
	NebulaCaller     // Read-only binding to the contract
	NebulaTransactor // Write-only binding to the contract
	NebulaFilterer   // Log filterer for contract events
}

// NebulaCaller is an auto generated read-only Go binding around an Ethereum contract.
type NebulaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NebulaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NebulaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NebulaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NebulaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NebulaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NebulaSession struct {
	Contract     *Nebula           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NebulaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NebulaCallerSession struct {
	Contract *NebulaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NebulaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NebulaTransactorSession struct {
	Contract     *NebulaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NebulaRaw is an auto generated low-level Go binding around an Ethereum contract.
type NebulaRaw struct {
	Contract *Nebula // Generic contract binding to access the raw methods on
}

// NebulaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NebulaCallerRaw struct {
	Contract *NebulaCaller // Generic read-only contract binding to access the raw methods on
}

// NebulaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NebulaTransactorRaw struct {
	Contract *NebulaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNebula creates a new instance of Nebula, bound to a specific deployed contract.
func NewNebula(address common.Address, backend bind.ContractBackend) (*Nebula, error) {
	contract, err := bindNebula(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nebula{NebulaCaller: NebulaCaller{contract: contract}, NebulaTransactor: NebulaTransactor{contract: contract}, NebulaFilterer: NebulaFilterer{contract: contract}}, nil
}

// NewNebulaCaller creates a new read-only instance of Nebula, bound to a specific deployed contract.
func NewNebulaCaller(address common.Address, caller bind.ContractCaller) (*NebulaCaller, error) {
	contract, err := bindNebula(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NebulaCaller{contract: contract}, nil
}

// NewNebulaTransactor creates a new write-only instance of Nebula, bound to a specific deployed contract.
func NewNebulaTransactor(address common.Address, transactor bind.ContractTransactor) (*NebulaTransactor, error) {
	contract, err := bindNebula(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NebulaTransactor{contract: contract}, nil
}

// NewNebulaFilterer creates a new log filterer instance of Nebula, bound to a specific deployed contract.
func NewNebulaFilterer(address common.Address, filterer bind.ContractFilterer) (*NebulaFilterer, error) {
	contract, err := bindNebula(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NebulaFilterer{contract: contract}, nil
}

// bindNebula binds a generic wrapper to an already deployed contract.
func bindNebula(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NebulaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nebula *NebulaRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Nebula.Contract.NebulaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nebula *NebulaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nebula.Contract.NebulaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nebula *NebulaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nebula.Contract.NebulaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nebula *NebulaCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Nebula.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nebula *NebulaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nebula.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nebula *NebulaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nebula.Contract.contract.Transact(opts, method, params...)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Nebula *NebulaCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Nebula.contract.Call(opts, out, "currentEpoch")
	return *ret0, err
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Nebula *NebulaSession) CurrentEpoch() (*big.Int, error) {
	return _Nebula.Contract.CurrentEpoch(&_Nebula.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Nebula *NebulaCallerSession) CurrentEpoch() (*big.Int, error) {
	return _Nebula.Contract.CurrentEpoch(&_Nebula.CallOpts)
}

// CurrentOracleId is a free data retrieval call binding the contract method 0x524ce31f.
//
// Solidity: function currentOracleId() view returns(bytes32)
func (_Nebula *NebulaCaller) CurrentOracleId(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Nebula.contract.Call(opts, out, "currentOracleId")
	return *ret0, err
}

// CurrentOracleId is a free data retrieval call binding the contract method 0x524ce31f.
//
// Solidity: function currentOracleId() view returns(bytes32)
func (_Nebula *NebulaSession) CurrentOracleId() ([32]byte, error) {
	return _Nebula.Contract.CurrentOracleId(&_Nebula.CallOpts)
}

// CurrentOracleId is a free data retrieval call binding the contract method 0x524ce31f.
//
// Solidity: function currentOracleId() view returns(bytes32)
func (_Nebula *NebulaCallerSession) CurrentOracleId() ([32]byte, error) {
	return _Nebula.Contract.CurrentOracleId(&_Nebula.CallOpts)
}

// EndEpochHeight is a free data retrieval call binding the contract method 0x996211a5.
//
// Solidity: function endEpochHeight() view returns(uint256)
func (_Nebula *NebulaCaller) EndEpochHeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Nebula.contract.Call(opts, out, "endEpochHeight")
	return *ret0, err
}

// EndEpochHeight is a free data retrieval call binding the contract method 0x996211a5.
//
// Solidity: function endEpochHeight() view returns(uint256)
func (_Nebula *NebulaSession) EndEpochHeight() (*big.Int, error) {
	return _Nebula.Contract.EndEpochHeight(&_Nebula.CallOpts)
}

// EndEpochHeight is a free data retrieval call binding the contract method 0x996211a5.
//
// Solidity: function endEpochHeight() view returns(uint256)
func (_Nebula *NebulaCallerSession) EndEpochHeight() (*big.Int, error) {
	return _Nebula.Contract.EndEpochHeight(&_Nebula.CallOpts)
}

// IsPublseSubSent is a free data retrieval call binding the contract method 0x6148d3f3.
//
// Solidity: function isPublseSubSent(uint256 , bytes32 ) view returns(bool)
func (_Nebula *NebulaCaller) IsPublseSubSent(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Nebula.contract.Call(opts, out, "isPublseSubSent", arg0, arg1)
	return *ret0, err
}

// IsPublseSubSent is a free data retrieval call binding the contract method 0x6148d3f3.
//
// Solidity: function isPublseSubSent(uint256 , bytes32 ) view returns(bool)
func (_Nebula *NebulaSession) IsPublseSubSent(arg0 *big.Int, arg1 [32]byte) (bool, error) {
	return _Nebula.Contract.IsPublseSubSent(&_Nebula.CallOpts, arg0, arg1)
}

// IsPublseSubSent is a free data retrieval call binding the contract method 0x6148d3f3.
//
// Solidity: function isPublseSubSent(uint256 , bytes32 ) view returns(bool)
func (_Nebula *NebulaCallerSession) IsPublseSubSent(arg0 *big.Int, arg1 [32]byte) (bool, error) {
	return _Nebula.Contract.IsPublseSubSent(&_Nebula.CallOpts, arg0, arg1)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nebula *NebulaCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Nebula.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nebula *NebulaSession) Name() (string, error) {
	return _Nebula.Contract.Name(&_Nebula.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nebula *NebulaCallerSession) Name() (string, error) {
	return _Nebula.Contract.Name(&_Nebula.CallOpts)
}

// OracleInfo is a free data retrieval call binding the contract method 0xbc623feb.
//
// Solidity: function oracleInfo(address ) view returns(address owner, bool isOnline, bytes32 idInQueue)
func (_Nebula *NebulaCaller) OracleInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Owner     common.Address
	IsOnline  bool
	IdInQueue [32]byte
}, error) {
	ret := new(struct {
		Owner     common.Address
		IsOnline  bool
		IdInQueue [32]byte
	})
	out := ret
	err := _Nebula.contract.Call(opts, out, "oracleInfo", arg0)
	return *ret, err
}

// OracleInfo is a free data retrieval call binding the contract method 0xbc623feb.
//
// Solidity: function oracleInfo(address ) view returns(address owner, bool isOnline, bytes32 idInQueue)
func (_Nebula *NebulaSession) OracleInfo(arg0 common.Address) (struct {
	Owner     common.Address
	IsOnline  bool
	IdInQueue [32]byte
}, error) {
	return _Nebula.Contract.OracleInfo(&_Nebula.CallOpts, arg0)
}

// OracleInfo is a free data retrieval call binding the contract method 0xbc623feb.
//
// Solidity: function oracleInfo(address ) view returns(address owner, bool isOnline, bytes32 idInQueue)
func (_Nebula *NebulaCallerSession) OracleInfo(arg0 common.Address) (struct {
	Owner     common.Address
	IsOnline  bool
	IdInQueue [32]byte
}, error) {
	return _Nebula.Contract.OracleInfo(&_Nebula.CallOpts, arg0)
}

// OracleQueue is a free data retrieval call binding the contract method 0x69a4246d.
//
// Solidity: function oracleQueue() view returns(bytes32 first, bytes32 last)
func (_Nebula *NebulaCaller) OracleQueue(opts *bind.CallOpts) (struct {
	First [32]byte
	Last  [32]byte
}, error) {
	ret := new(struct {
		First [32]byte
		Last  [32]byte
	})
	out := ret
	err := _Nebula.contract.Call(opts, out, "oracleQueue")
	return *ret, err
}

// OracleQueue is a free data retrieval call binding the contract method 0x69a4246d.
//
// Solidity: function oracleQueue() view returns(bytes32 first, bytes32 last)
func (_Nebula *NebulaSession) OracleQueue() (struct {
	First [32]byte
	Last  [32]byte
}, error) {
	return _Nebula.Contract.OracleQueue(&_Nebula.CallOpts)
}

// OracleQueue is a free data retrieval call binding the contract method 0x69a4246d.
//
// Solidity: function oracleQueue() view returns(bytes32 first, bytes32 last)
func (_Nebula *NebulaCallerSession) OracleQueue() (struct {
	First [32]byte
	Last  [32]byte
}, error) {
	return _Nebula.Contract.OracleQueue(&_Nebula.CallOpts)
}

// Oracles is a free data retrieval call binding the contract method 0xa81a2677.
//
// Solidity: function oracles(bytes32 ) view returns(address)
func (_Nebula *NebulaCaller) Oracles(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Nebula.contract.Call(opts, out, "oracles", arg0)
	return *ret0, err
}

// Oracles is a free data retrieval call binding the contract method 0xa81a2677.
//
// Solidity: function oracles(bytes32 ) view returns(address)
func (_Nebula *NebulaSession) Oracles(arg0 [32]byte) (common.Address, error) {
	return _Nebula.Contract.Oracles(&_Nebula.CallOpts, arg0)
}

// Oracles is a free data retrieval call binding the contract method 0xa81a2677.
//
// Solidity: function oracles(bytes32 ) view returns(address)
func (_Nebula *NebulaCallerSession) Oracles(arg0 [32]byte) (common.Address, error) {
	return _Nebula.Contract.Oracles(&_Nebula.CallOpts, arg0)
}

// OraclesCount is a free data retrieval call binding the contract method 0x76f70900.
//
// Solidity: function oraclesCount() view returns(uint256)
func (_Nebula *NebulaCaller) OraclesCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Nebula.contract.Call(opts, out, "oraclesCount")
	return *ret0, err
}

// OraclesCount is a free data retrieval call binding the contract method 0x76f70900.
//
// Solidity: function oraclesCount() view returns(uint256)
func (_Nebula *NebulaSession) OraclesCount() (*big.Int, error) {
	return _Nebula.Contract.OraclesCount(&_Nebula.CallOpts)
}

// OraclesCount is a free data retrieval call binding the contract method 0x76f70900.
//
// Solidity: function oraclesCount() view returns(uint256)
func (_Nebula *NebulaCallerSession) OraclesCount() (*big.Int, error) {
	return _Nebula.Contract.OraclesCount(&_Nebula.CallOpts)
}

// PulseQueue is a free data retrieval call binding the contract method 0x1d11f944.
//
// Solidity: function pulseQueue() view returns(bytes32 first, bytes32 last)
func (_Nebula *NebulaCaller) PulseQueue(opts *bind.CallOpts) (struct {
	First [32]byte
	Last  [32]byte
}, error) {
	ret := new(struct {
		First [32]byte
		Last  [32]byte
	})
	out := ret
	err := _Nebula.contract.Call(opts, out, "pulseQueue")
	return *ret, err
}

// PulseQueue is a free data retrieval call binding the contract method 0x1d11f944.
//
// Solidity: function pulseQueue() view returns(bytes32 first, bytes32 last)
func (_Nebula *NebulaSession) PulseQueue() (struct {
	First [32]byte
	Last  [32]byte
}, error) {
	return _Nebula.Contract.PulseQueue(&_Nebula.CallOpts)
}

// PulseQueue is a free data retrieval call binding the contract method 0x1d11f944.
//
// Solidity: function pulseQueue() view returns(bytes32 first, bytes32 last)
func (_Nebula *NebulaCallerSession) PulseQueue() (struct {
	First [32]byte
	Last  [32]byte
}, error) {
	return _Nebula.Contract.PulseQueue(&_Nebula.CallOpts)
}

// Pulses is a free data retrieval call binding the contract method 0x0694fbb3.
//
// Solidity: function pulses(uint256 ) view returns(uint256 confirmationCount)
func (_Nebula *NebulaCaller) Pulses(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Nebula.contract.Call(opts, out, "pulses", arg0)
	return *ret0, err
}

// Pulses is a free data retrieval call binding the contract method 0x0694fbb3.
//
// Solidity: function pulses(uint256 ) view returns(uint256 confirmationCount)
func (_Nebula *NebulaSession) Pulses(arg0 *big.Int) (*big.Int, error) {
	return _Nebula.Contract.Pulses(&_Nebula.CallOpts, arg0)
}

// Pulses is a free data retrieval call binding the contract method 0x0694fbb3.
//
// Solidity: function pulses(uint256 ) view returns(uint256 confirmationCount)
func (_Nebula *NebulaCallerSession) Pulses(arg0 *big.Int) (*big.Int, error) {
	return _Nebula.Contract.Pulses(&_Nebula.CallOpts, arg0)
}

// Subscriptions is a free data retrieval call binding the contract method 0x94259c6c.
//
// Solidity: function subscriptions(bytes32 ) view returns(address owner, address contractAddress, uint8 minConfirmations, uint256 reward)
func (_Nebula *NebulaCaller) Subscriptions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Owner            common.Address
	ContractAddress  common.Address
	MinConfirmations uint8
	Reward           *big.Int
}, error) {
	ret := new(struct {
		Owner            common.Address
		ContractAddress  common.Address
		MinConfirmations uint8
		Reward           *big.Int
	})
	out := ret
	err := _Nebula.contract.Call(opts, out, "subscriptions", arg0)
	return *ret, err
}

// Subscriptions is a free data retrieval call binding the contract method 0x94259c6c.
//
// Solidity: function subscriptions(bytes32 ) view returns(address owner, address contractAddress, uint8 minConfirmations, uint256 reward)
func (_Nebula *NebulaSession) Subscriptions(arg0 [32]byte) (struct {
	Owner            common.Address
	ContractAddress  common.Address
	MinConfirmations uint8
	Reward           *big.Int
}, error) {
	return _Nebula.Contract.Subscriptions(&_Nebula.CallOpts, arg0)
}

// Subscriptions is a free data retrieval call binding the contract method 0x94259c6c.
//
// Solidity: function subscriptions(bytes32 ) view returns(address owner, address contractAddress, uint8 minConfirmations, uint256 reward)
func (_Nebula *NebulaCallerSession) Subscriptions(arg0 [32]byte) (struct {
	Owner            common.Address
	ContractAddress  common.Address
	MinConfirmations uint8
	Reward           *big.Int
}, error) {
	return _Nebula.Contract.Subscriptions(&_Nebula.CallOpts, arg0)
}

// SubscriptionsQueue is a free data retrieval call binding the contract method 0xb48a9c9b.
//
// Solidity: function subscriptionsQueue() view returns(bytes32 first, bytes32 last)
func (_Nebula *NebulaCaller) SubscriptionsQueue(opts *bind.CallOpts) (struct {
	First [32]byte
	Last  [32]byte
}, error) {
	ret := new(struct {
		First [32]byte
		Last  [32]byte
	})
	out := ret
	err := _Nebula.contract.Call(opts, out, "subscriptionsQueue")
	return *ret, err
}

// SubscriptionsQueue is a free data retrieval call binding the contract method 0xb48a9c9b.
//
// Solidity: function subscriptionsQueue() view returns(bytes32 first, bytes32 last)
func (_Nebula *NebulaSession) SubscriptionsQueue() (struct {
	First [32]byte
	Last  [32]byte
}, error) {
	return _Nebula.Contract.SubscriptionsQueue(&_Nebula.CallOpts)
}

// SubscriptionsQueue is a free data retrieval call binding the contract method 0xb48a9c9b.
//
// Solidity: function subscriptionsQueue() view returns(bytes32 first, bytes32 last)
func (_Nebula *NebulaCallerSession) SubscriptionsQueue() (struct {
	First [32]byte
	Last  [32]byte
}, error) {
	return _Nebula.Contract.SubscriptionsQueue(&_Nebula.CallOpts)
}

// ConfirmData is a paid mutator transaction binding the contract method 0x02cb72d3.
//
// Solidity: function confirmData(uint256[] data, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_Nebula *NebulaTransactor) ConfirmData(opts *bind.TransactOpts, data []*big.Int, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Nebula.contract.Transact(opts, "confirmData", data, v, r, s)
}

// ConfirmData is a paid mutator transaction binding the contract method 0x02cb72d3.
//
// Solidity: function confirmData(uint256[] data, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_Nebula *NebulaSession) ConfirmData(data []*big.Int, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Nebula.Contract.ConfirmData(&_Nebula.TransactOpts, data, v, r, s)
}

// ConfirmData is a paid mutator transaction binding the contract method 0x02cb72d3.
//
// Solidity: function confirmData(uint256[] data, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_Nebula *NebulaTransactorSession) ConfirmData(data []*big.Int, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Nebula.Contract.ConfirmData(&_Nebula.TransactOpts, data, v, r, s)
}

// NewRound is a paid mutator transaction binding the contract method 0xc5b1d9aa.
//
// Solidity: function newRound() returns()
func (_Nebula *NebulaTransactor) NewRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nebula.contract.Transact(opts, "newRound")
}

// NewRound is a paid mutator transaction binding the contract method 0xc5b1d9aa.
//
// Solidity: function newRound() returns()
func (_Nebula *NebulaSession) NewRound() (*types.Transaction, error) {
	return _Nebula.Contract.NewRound(&_Nebula.TransactOpts)
}

// NewRound is a paid mutator transaction binding the contract method 0xc5b1d9aa.
//
// Solidity: function newRound() returns()
func (_Nebula *NebulaTransactorSession) NewRound() (*types.Transaction, error) {
	return _Nebula.Contract.NewRound(&_Nebula.TransactOpts)
}

// SendData is a paid mutator transaction binding the contract method 0x342e9741.
//
// Solidity: function sendData(uint256 blockNumber, bytes32 subscriptionId) returns()
func (_Nebula *NebulaTransactor) SendData(opts *bind.TransactOpts, blockNumber *big.Int, subscriptionId [32]byte) (*types.Transaction, error) {
	return _Nebula.contract.Transact(opts, "sendData", blockNumber, subscriptionId)
}

// SendData is a paid mutator transaction binding the contract method 0x342e9741.
//
// Solidity: function sendData(uint256 blockNumber, bytes32 subscriptionId) returns()
func (_Nebula *NebulaSession) SendData(blockNumber *big.Int, subscriptionId [32]byte) (*types.Transaction, error) {
	return _Nebula.Contract.SendData(&_Nebula.TransactOpts, blockNumber, subscriptionId)
}

// SendData is a paid mutator transaction binding the contract method 0x342e9741.
//
// Solidity: function sendData(uint256 blockNumber, bytes32 subscriptionId) returns()
func (_Nebula *NebulaTransactorSession) SendData(blockNumber *big.Int, subscriptionId [32]byte) (*types.Transaction, error) {
	return _Nebula.Contract.SendData(&_Nebula.TransactOpts, blockNumber, subscriptionId)
}

// Subscribe is a paid mutator transaction binding the contract method 0x3527715d.
//
// Solidity: function subscribe(address contractAddress, uint8 minConfirmations, uint256 reward) returns()
func (_Nebula *NebulaTransactor) Subscribe(opts *bind.TransactOpts, contractAddress common.Address, minConfirmations uint8, reward *big.Int) (*types.Transaction, error) {
	return _Nebula.contract.Transact(opts, "subscribe", contractAddress, minConfirmations, reward)
}

// Subscribe is a paid mutator transaction binding the contract method 0x3527715d.
//
// Solidity: function subscribe(address contractAddress, uint8 minConfirmations, uint256 reward) returns()
func (_Nebula *NebulaSession) Subscribe(contractAddress common.Address, minConfirmations uint8, reward *big.Int) (*types.Transaction, error) {
	return _Nebula.Contract.Subscribe(&_Nebula.TransactOpts, contractAddress, minConfirmations, reward)
}

// Subscribe is a paid mutator transaction binding the contract method 0x3527715d.
//
// Solidity: function subscribe(address contractAddress, uint8 minConfirmations, uint256 reward) returns()
func (_Nebula *NebulaTransactorSession) Subscribe(contractAddress common.Address, minConfirmations uint8, reward *big.Int) (*types.Transaction, error) {
	return _Nebula.Contract.Subscribe(&_Nebula.TransactOpts, contractAddress, minConfirmations, reward)
}

// TrunOffOracle is a paid mutator transaction binding the contract method 0x0a0d00d1.
//
// Solidity: function trunOffOracle() returns()
func (_Nebula *NebulaTransactor) TrunOffOracle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nebula.contract.Transact(opts, "trunOffOracle")
}

// TrunOffOracle is a paid mutator transaction binding the contract method 0x0a0d00d1.
//
// Solidity: function trunOffOracle() returns()
func (_Nebula *NebulaSession) TrunOffOracle() (*types.Transaction, error) {
	return _Nebula.Contract.TrunOffOracle(&_Nebula.TransactOpts)
}

// TrunOffOracle is a paid mutator transaction binding the contract method 0x0a0d00d1.
//
// Solidity: function trunOffOracle() returns()
func (_Nebula *NebulaTransactorSession) TrunOffOracle() (*types.Transaction, error) {
	return _Nebula.Contract.TrunOffOracle(&_Nebula.TransactOpts)
}

// TrunOnOracle is a paid mutator transaction binding the contract method 0x51d81e09.
//
// Solidity: function trunOnOracle() returns()
func (_Nebula *NebulaTransactor) TrunOnOracle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nebula.contract.Transact(opts, "trunOnOracle")
}

// TrunOnOracle is a paid mutator transaction binding the contract method 0x51d81e09.
//
// Solidity: function trunOnOracle() returns()
func (_Nebula *NebulaSession) TrunOnOracle() (*types.Transaction, error) {
	return _Nebula.Contract.TrunOnOracle(&_Nebula.TransactOpts)
}

// TrunOnOracle is a paid mutator transaction binding the contract method 0x51d81e09.
//
// Solidity: function trunOnOracle() returns()
func (_Nebula *NebulaTransactorSession) TrunOnOracle() (*types.Transaction, error) {
	return _Nebula.Contract.TrunOnOracle(&_Nebula.TransactOpts)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xf8c60146.
//
// Solidity: function unsubscribe(bytes32 id) returns()
func (_Nebula *NebulaTransactor) Unsubscribe(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _Nebula.contract.Transact(opts, "unsubscribe", id)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xf8c60146.
//
// Solidity: function unsubscribe(bytes32 id) returns()
func (_Nebula *NebulaSession) Unsubscribe(id [32]byte) (*types.Transaction, error) {
	return _Nebula.Contract.Unsubscribe(&_Nebula.TransactOpts, id)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xf8c60146.
//
// Solidity: function unsubscribe(bytes32 id) returns()
func (_Nebula *NebulaTransactorSession) Unsubscribe(id [32]byte) (*types.Transaction, error) {
	return _Nebula.Contract.Unsubscribe(&_Nebula.TransactOpts, id)
}

// QueueLibABI is the input ABI used to generate the binding from.
const QueueLibABI = "[]"

// QueueLibFuncSigs maps the 4-byte function signature to its string representation.
var QueueLibFuncSigs = map[string]string{
	"9d6ad84b": "drop(QueueLib.Queue storage,bytes32)",
	"870a61ff": "next(QueueLib.Queue storage,bytes32)",
	"a506d954": "push(QueueLib.Queue storage,bytes32)",
}

// QueueLibBin is the compiled bytecode used for deploying new contracts.
var QueueLibBin = "0x6101fc610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c8063870a61ff146100505780639d6ad84b14610085578063a506d954146100b7575b600080fd5b6100736004803603604081101561006657600080fd5b50803590602001356100e7565b60408051918252519081900360200190f35b81801561009157600080fd5b506100b5600480360360408110156100a857600080fd5b508035906020013561010f565b005b8180156100c357600080fd5b506100b5600480360360408110156100da57600080fd5b508035906020013561017e565b6000816100f657508154610109565b5060008181526002830160205260409020545b92915050565b8154811480156101225750808260010154145b15610136576000808355600183015561017a565b8154811415610157576000818152600283016020526040902054825561017a565b808260010154141561017a57600081815260038301602052604090205460018301555b5050565b8154610193578082556001820181905561017a565b6001820180546000908152600284016020908152604080832085905583548584526003870190925290912055819055505056fea264697066735822122041550c8f352f4370a09f8d78a896f20b6124b22e3dd289894eecae66aefbf19b64736f6c63430006060033"

// DeployQueueLib deploys a new Ethereum contract, binding an instance of QueueLib to it.
func DeployQueueLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *QueueLib, error) {
	parsed, err := abi.JSON(strings.NewReader(QueueLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(QueueLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &QueueLib{QueueLibCaller: QueueLibCaller{contract: contract}, QueueLibTransactor: QueueLibTransactor{contract: contract}, QueueLibFilterer: QueueLibFilterer{contract: contract}}, nil
}

// QueueLib is an auto generated Go binding around an Ethereum contract.
type QueueLib struct {
	QueueLibCaller     // Read-only binding to the contract
	QueueLibTransactor // Write-only binding to the contract
	QueueLibFilterer   // Log filterer for contract events
}

// QueueLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type QueueLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueueLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QueueLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueueLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QueueLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueueLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QueueLibSession struct {
	Contract     *QueueLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QueueLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QueueLibCallerSession struct {
	Contract *QueueLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// QueueLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QueueLibTransactorSession struct {
	Contract     *QueueLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// QueueLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type QueueLibRaw struct {
	Contract *QueueLib // Generic contract binding to access the raw methods on
}

// QueueLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QueueLibCallerRaw struct {
	Contract *QueueLibCaller // Generic read-only contract binding to access the raw methods on
}

// QueueLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QueueLibTransactorRaw struct {
	Contract *QueueLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQueueLib creates a new instance of QueueLib, bound to a specific deployed contract.
func NewQueueLib(address common.Address, backend bind.ContractBackend) (*QueueLib, error) {
	contract, err := bindQueueLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QueueLib{QueueLibCaller: QueueLibCaller{contract: contract}, QueueLibTransactor: QueueLibTransactor{contract: contract}, QueueLibFilterer: QueueLibFilterer{contract: contract}}, nil
}

// NewQueueLibCaller creates a new read-only instance of QueueLib, bound to a specific deployed contract.
func NewQueueLibCaller(address common.Address, caller bind.ContractCaller) (*QueueLibCaller, error) {
	contract, err := bindQueueLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QueueLibCaller{contract: contract}, nil
}

// NewQueueLibTransactor creates a new write-only instance of QueueLib, bound to a specific deployed contract.
func NewQueueLibTransactor(address common.Address, transactor bind.ContractTransactor) (*QueueLibTransactor, error) {
	contract, err := bindQueueLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QueueLibTransactor{contract: contract}, nil
}

// NewQueueLibFilterer creates a new log filterer instance of QueueLib, bound to a specific deployed contract.
func NewQueueLibFilterer(address common.Address, filterer bind.ContractFilterer) (*QueueLibFilterer, error) {
	contract, err := bindQueueLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QueueLibFilterer{contract: contract}, nil
}

// bindQueueLib binds a generic wrapper to an already deployed contract.
func bindQueueLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QueueLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QueueLib *QueueLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _QueueLib.Contract.QueueLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QueueLib *QueueLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QueueLib.Contract.QueueLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QueueLib *QueueLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QueueLib.Contract.QueueLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QueueLib *QueueLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _QueueLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QueueLib *QueueLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QueueLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QueueLib *QueueLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QueueLib.Contract.contract.Transact(opts, method, params...)
}
