package systemcontract

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const validatorsInteractiveABI = `
[
	{
		"inputs": [
		  {
			"internalType": "address[]",
			"name": "vals",
			"type": "address[]"
		  }
		],
		"name": "initialize",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "distributeBlockReward",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "getTopValidators",
		"outputs": [
		  {
			"internalType": "address[]",
			"name": "",
			"type": "address[]"
		  }
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
		  {
			"internalType": "address[]",
			"name": "newSet",
			"type": "address[]"
		  },
		  {
			"internalType": "uint256",
			"name": "epoch",
			"type": "uint256"
		  }
		],
		"name": "updateActiveValidatorSet",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]
`

const punishInteractiveABI = `
[
	{
		"inputs": [],
		"name": "initialize",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
		  {
			"internalType": "address",
			"name": "val",
			"type": "address"
		  }
		],
		"name": "punish",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
		  {
			"internalType": "uint256",
			"name": "epoch",
			"type": "uint256"
		  }
		],
		"name": "decreaseMissedBlocksCounter",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	  }
]
`

const proposalInteractiveABI = `
[
	{
		"inputs": [
		  {
			"internalType": "address[]",
			"name": "vals",
			"type": "address[]"
		  }
		],
		"name": "initialize",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]
`
const addrListInteractiveABI = `
[
    {
        "inputs": [],
        "name": "devVerifyEnabled",
        "outputs": [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "getBlacksFrom",
        "outputs": [
            {
                "internalType": "address[]",
                "name": "",
                "type": "address[]"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "getBlacksTo",
        "outputs": [
            {
                "internalType": "address[]",
                "name": "",
                "type": "address[]"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "initialize",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "addr",
                "type": "address"
            }
        ],
        "name": "isDeveloper",
        "outputs": [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    }
]`

const incentiveInteractiveABI = `
[
    {
        "inputs": [],
        "name": "initialize",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
      }
  ]
`

var (
	ValidatorsContractName  = "validators"
	PunishContractName      = "punish"
	ProposalContractName    = "proposal"
	AddressListContractName = "address_list"
	IncentiveContractName   = "incentive"
)

// DevMappingPosition is the position of the state variable `devs`.
// Since the state variables are as follow:
//    bool public initialized;
//    bool public enabled;
//    address public admin;
//    address public pendingAdmin;
//    mapping(address => bool) private devs;
//
// according to [Layout of State Variables in Storage](https://docs.soliditylang.org/en/v0.8.4/internals/layout_in_storage.html),
// and after optimizer enabled, the `initialized`, `enabled` and `admin` will be packed, and stores at slot 0,
// `pendingAdmin` stores at slot 1, so the position for `devs` is 2.
const DevMappingPosition = 2

var abiMap map[string]abi.ABI

func init() {
	abiMap = make(map[string]abi.ABI)
	tmpABI, _ := abi.JSON(strings.NewReader(proposalInteractiveABI))
	abiMap[ProposalContractName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(validatorsInteractiveABI))
	abiMap[ValidatorsContractName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(punishInteractiveABI))
	abiMap[PunishContractName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(addrListInteractiveABI))
	abiMap[AddressListContractName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(incentiveInteractiveABI))
	abiMap[IncentiveContractName] = tmpABI
}

func GetInteractiveABI() map[string]abi.ABI {
	return abiMap
}
