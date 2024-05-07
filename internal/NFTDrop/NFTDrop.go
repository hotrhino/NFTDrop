// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package NFTDrop

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

// IClaimConditionClaimCondition is an auto generated low-level Go binding around an user-defined struct.
type IClaimConditionClaimCondition struct {
	StartTimestamp         *big.Int
	MaxClaimableSupply     *big.Int
	SupplyClaimed          *big.Int
	QuantityLimitPerWallet *big.Int
	MerkleRoot             [32]byte
	PricePerToken          *big.Int
	Currency               common.Address
	Metadata               string
}

// IDropAllowlistProof is an auto generated low-level Go binding around an user-defined struct.
type IDropAllowlistProof struct {
	Proof                  [][32]byte
	QuantityLimitPerWallet *big.Int
	PricePerToken          *big.Int
	Currency               common.Address
}

// NFTDropMetaData contains all meta data concerning the NFTDrop contract.
var NFTDropMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"name\":\"\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"inputs\":[],\"outputs\":[]},{\"type\":\"error\",\"name\":\"ApprovalQueryForNonexistentToken\",\"inputs\":[],\"outputs\":[]},{\"type\":\"error\",\"name\":\"ApprovalToCurrentOwner\",\"inputs\":[],\"outputs\":[]},{\"type\":\"error\",\"name\":\"ApproveToCaller\",\"inputs\":[],\"outputs\":[]},{\"type\":\"error\",\"name\":\"BalanceQueryForZeroAddress\",\"inputs\":[],\"outputs\":[]},{\"type\":\"error\",\"name\":\"MintToZeroAddress\",\"inputs\":[],\"outputs\":[]},{\"type\":\"error\",\"name\":\"MintZeroQuantity\",\"inputs\":[],\"outputs\":[]},{\"type\":\"error\",\"name\":\"OwnerQueryForNonexistentToken\",\"inputs\":[],\"outputs\":[]},{\"type\":\"error\",\"name\":\"TransferCallerNotOwnerNorApproved\",\"inputs\":[],\"outputs\":[]},{\"type\":\"error\",\"name\":\"TransferFromIncorrectOwner\",\"inputs\":[],\"outputs\":[]},{\"type\":\"error\",\"name\":\"TransferToNonERC721ReceiverImplementer\",\"inputs\":[],\"outputs\":[]},{\"type\":\"error\",\"name\":\"TransferToZeroAddress\",\"inputs\":[],\"outputs\":[]},{\"type\":\"error\",\"name\":\"URIQueryForNonexistentToken\",\"inputs\":[],\"outputs\":[]},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"approved\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"indexed\":true,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"operator\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"bool\",\"name\":\"approved\",\"indexed\":false,\"internalType\":\"bool\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BatchMetadataUpdate\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_fromTokenId\",\"indexed\":false,\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_toTokenId\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClaimConditionsUpdated\",\"inputs\":[{\"type\":\"tuple[]\",\"name\":\"claimConditions\",\"components\":[{\"type\":\"uint256\",\"name\":\"startTimestamp\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"maxClaimableSupply\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"supplyClaimed\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"quantityLimitPerWallet\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"merkleRoot\",\"internalType\":\"bytes32\"},{\"type\":\"uint256\",\"name\":\"pricePerToken\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"currency\",\"internalType\":\"address\"},{\"type\":\"string\",\"name\":\"metadata\",\"internalType\":\"string\"}],\"indexed\":false,\"internalType\":\"structIClaimCondition.ClaimCondition[]\"},{\"type\":\"bool\",\"name\":\"resetEligibility\",\"indexed\":false,\"internalType\":\"bool\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractURIUpdated\",\"inputs\":[{\"type\":\"string\",\"name\":\"prevURI\",\"indexed\":false,\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"newURI\",\"indexed\":false,\"internalType\":\"string\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultRoyalty\",\"inputs\":[{\"type\":\"address\",\"name\":\"newRoyaltyRecipient\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"newRoyaltyBps\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FlatPlatformFeeUpdated\",\"inputs\":[{\"type\":\"address\",\"name\":\"platformFeeRecipient\",\"indexed\":false,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"flatFee\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"version\",\"indexed\":false,\"internalType\":\"uint8\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxTotalSupplyUpdated\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"maxTotalSupply\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataFrozen\",\"inputs\":[],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnerUpdated\",\"inputs\":[{\"type\":\"address\",\"name\":\"prevOwner\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"newOwner\",\"indexed\":true,\"internalType\":\"address\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PlatformFeeInfoUpdated\",\"inputs\":[{\"type\":\"address\",\"name\":\"platformFeeRecipient\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"platformFeeBps\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PlatformFeeTypeUpdated\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"feeType\",\"indexed\":false,\"internalType\":\"enumIPlatformFee.PlatformFeeType\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PrimarySaleRecipientUpdated\",\"inputs\":[{\"type\":\"address\",\"name\":\"recipient\",\"indexed\":true,\"internalType\":\"address\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"previousAdminRole\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"newAdminRole\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"sender\",\"indexed\":true,\"internalType\":\"address\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"sender\",\"indexed\":true,\"internalType\":\"address\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoyaltyForToken\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"indexed\":true,\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"royaltyRecipient\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"royaltyBps\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenURIRevealed\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"index\",\"indexed\":true,\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"revealedURI\",\"indexed\":false,\"internalType\":\"string\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensClaimed\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"claimConditionIndex\",\"indexed\":true,\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"claimer\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"receiver\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"startTokenId\",\"indexed\":false,\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"quantityClaimed\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensLazyMinted\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"startTokenId\",\"indexed\":true,\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"endTokenId\",\"indexed\":false,\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"baseURI\",\"indexed\":false,\"internalType\":\"string\"},{\"type\":\"bytes\",\"name\":\"encryptedBaseURI\",\"indexed\":false,\"internalType\":\"bytes\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"indexed\":true,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchFrozen\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"type\":\"address\",\"name\":\"_receiver\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_quantity\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"_currency\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_pricePerToken\",\"internalType\":\"uint256\"},{\"type\":\"tuple\",\"name\":\"_allowlistProof\",\"components\":[{\"type\":\"bytes32[]\",\"name\":\"proof\",\"internalType\":\"bytes32[]\"},{\"type\":\"uint256\",\"name\":\"quantityLimitPerWallet\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"pricePerToken\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"currency\",\"internalType\":\"address\"}],\"internalType\":\"structIDrop.AllowlistProof\"},{\"type\":\"bytes\",\"name\":\"_data\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claimCondition\",\"inputs\":[],\"outputs\":[{\"type\":\"uint256\",\"name\":\"currentStartId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"count\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractType\",\"inputs\":[],\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"contractURI\",\"inputs\":[],\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractVersion\",\"inputs\":[],\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encryptDecrypt\",\"inputs\":[{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"},{\"type\":\"bytes\",\"name\":\"key\",\"internalType\":\"bytes\"}],\"outputs\":[{\"type\":\"bytes\",\"name\":\"result\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encryptedData\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"bytes\",\"name\":\"\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"freezeBatchBaseURI\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_index\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getActiveClaimConditionId\",\"inputs\":[],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBaseURICount\",\"inputs\":[],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchIdAtIndex\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_index\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getClaimConditionById\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_conditionId\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"tuple\",\"name\":\"condition\",\"components\":[{\"type\":\"uint256\",\"name\":\"startTimestamp\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"maxClaimableSupply\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"supplyClaimed\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"quantityLimitPerWallet\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"merkleRoot\",\"internalType\":\"bytes32\"},{\"type\":\"uint256\",\"name\":\"pricePerToken\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"currency\",\"internalType\":\"address\"},{\"type\":\"string\",\"name\":\"metadata\",\"internalType\":\"string\"}],\"internalType\":\"structIClaimCondition.ClaimCondition\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDefaultRoyaltyInfo\",\"inputs\":[],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"uint16\",\"name\":\"\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFlatPlatformFeeInfo\",\"inputs\":[],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlatformFeeInfo\",\"inputs\":[],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"uint16\",\"name\":\"\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlatformFeeType\",\"inputs\":[],\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"enumIPlatformFee.PlatformFeeType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRevealURI\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_batchId\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"_key\",\"internalType\":\"bytes\"}],\"outputs\":[{\"type\":\"string\",\"name\":\"revealedURI\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"uint256\",\"name\":\"index\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"member\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"count\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoyaltyInfoForToken\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_tokenId\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"uint16\",\"name\":\"\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupplyClaimedByWallet\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_conditionId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"_claimer\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"supplyClaimedByWallet\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRoleWithSwitch\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"type\":\"address\",\"name\":\"_defaultAdmin\",\"internalType\":\"address\"},{\"type\":\"string\",\"name\":\"_name\",\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"_symbol\",\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"_contractURI\",\"internalType\":\"string\"},{\"type\":\"address[]\",\"name\":\"_trustedForwarders\",\"internalType\":\"address[]\"},{\"type\":\"address\",\"name\":\"_saleRecipient\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_royaltyRecipient\",\"internalType\":\"address\"},{\"type\":\"uint128\",\"name\":\"_royaltyBps\",\"internalType\":\"uint128\"},{\"type\":\"uint128\",\"name\":\"_platformFeeBps\",\"internalType\":\"uint128\"},{\"type\":\"address\",\"name\":\"_platformFeeRecipient\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isEncryptedBatch\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_batchId\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isTrustedForwarder\",\"inputs\":[{\"type\":\"address\",\"name\":\"forwarder\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lazyMint\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"_baseURIForTokens\",\"internalType\":\"string\"},{\"type\":\"bytes\",\"name\":\"_data\",\"internalType\":\"bytes\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"batchId\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxTotalSupply\",\"inputs\":[],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"multicall\",\"inputs\":[{\"type\":\"bytes[]\",\"name\":\"data\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"type\":\"bytes[]\",\"name\":\"results\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextTokenIdToClaim\",\"inputs\":[],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextTokenIdToMint\",\"inputs\":[],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"primarySaleRecipient\",\"inputs\":[],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reveal\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_index\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"_key\",\"internalType\":\"bytes\"}],\"outputs\":[{\"type\":\"string\",\"name\":\"revealedURI\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"royaltyInfo\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"salePrice\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"receiver\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"royaltyAmount\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"_data\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\"},{\"type\":\"bool\",\"name\":\"approved\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setClaimConditions\",\"inputs\":[{\"type\":\"tuple[]\",\"name\":\"_conditions\",\"components\":[{\"type\":\"uint256\",\"name\":\"startTimestamp\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"maxClaimableSupply\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"supplyClaimed\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"quantityLimitPerWallet\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"merkleRoot\",\"internalType\":\"bytes32\"},{\"type\":\"uint256\",\"name\":\"pricePerToken\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"currency\",\"internalType\":\"address\"},{\"type\":\"string\",\"name\":\"metadata\",\"internalType\":\"string\"}],\"internalType\":\"structIClaimCondition.ClaimCondition[]\"},{\"type\":\"bool\",\"name\":\"_resetClaimEligibility\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setContractURI\",\"inputs\":[{\"type\":\"string\",\"name\":\"_uri\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDefaultRoyaltyInfo\",\"inputs\":[{\"type\":\"address\",\"name\":\"_royaltyRecipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_royaltyBps\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFlatPlatformFeeInfo\",\"inputs\":[{\"type\":\"address\",\"name\":\"_platformFeeRecipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_flatFee\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxTotalSupply\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_maxTotalSupply\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOwner\",\"inputs\":[{\"type\":\"address\",\"name\":\"_newOwner\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPlatformFeeInfo\",\"inputs\":[{\"type\":\"address\",\"name\":\"_platformFeeRecipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_platformFeeBps\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPlatformFeeType\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"_feeType\",\"internalType\":\"enumIPlatformFee.PlatformFeeType\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPrimarySaleRecipient\",\"inputs\":[{\"type\":\"address\",\"name\":\"_saleRecipient\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRoyaltyInfoForToken\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_tokenId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"_recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_bps\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"type\":\"bytes4\",\"name\":\"interfaceId\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_tokenId\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalMinted\",\"inputs\":[],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateBatchBaseURI\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_index\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"_uri\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyClaim\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_conditionId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"_claimer\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_quantity\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"_currency\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_pricePerToken\",\"internalType\":\"uint256\"},{\"type\":\"tuple\",\"name\":\"_allowlistProof\",\"components\":[{\"type\":\"bytes32[]\",\"name\":\"proof\",\"internalType\":\"bytes32[]\"},{\"type\":\"uint256\",\"name\":\"quantityLimitPerWallet\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"pricePerToken\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"currency\",\"internalType\":\"address\"}],\"internalType\":\"structIDrop.AllowlistProof\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"isOverride\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"}]",
}

// NFTDropABI is the input ABI used to generate the binding from.
// Deprecated: Use NFTDropMetaData.ABI instead.
var NFTDropABI = NFTDropMetaData.ABI

// NFTDrop is an auto generated Go binding around an Ethereum contract.
type NFTDrop struct {
	NFTDropCaller     // Read-only binding to the contract
	NFTDropTransactor // Write-only binding to the contract
	NFTDropFilterer   // Log filterer for contract events
}

// NFTDropCaller is an auto generated read-only Go binding around an Ethereum contract.
type NFTDropCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTDropTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NFTDropTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTDropFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFTDropFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTDropSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFTDropSession struct {
	Contract     *NFTDrop          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTDropCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NFTDropCallerSession struct {
	Contract *NFTDropCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// NFTDropTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NFTDropTransactorSession struct {
	Contract     *NFTDropTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NFTDropRaw is an auto generated low-level Go binding around an Ethereum contract.
type NFTDropRaw struct {
	Contract *NFTDrop // Generic contract binding to access the raw methods on
}

// NFTDropCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NFTDropCallerRaw struct {
	Contract *NFTDropCaller // Generic read-only contract binding to access the raw methods on
}

// NFTDropTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NFTDropTransactorRaw struct {
	Contract *NFTDropTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNFTDrop creates a new instance of NFTDrop, bound to a specific deployed contract.
func NewNFTDrop(address common.Address, backend bind.ContractBackend) (*NFTDrop, error) {
	contract, err := bindNFTDrop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFTDrop{NFTDropCaller: NFTDropCaller{contract: contract}, NFTDropTransactor: NFTDropTransactor{contract: contract}, NFTDropFilterer: NFTDropFilterer{contract: contract}}, nil
}

// NewNFTDropCaller creates a new read-only instance of NFTDrop, bound to a specific deployed contract.
func NewNFTDropCaller(address common.Address, caller bind.ContractCaller) (*NFTDropCaller, error) {
	contract, err := bindNFTDrop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NFTDropCaller{contract: contract}, nil
}

// NewNFTDropTransactor creates a new write-only instance of NFTDrop, bound to a specific deployed contract.
func NewNFTDropTransactor(address common.Address, transactor bind.ContractTransactor) (*NFTDropTransactor, error) {
	contract, err := bindNFTDrop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NFTDropTransactor{contract: contract}, nil
}

// NewNFTDropFilterer creates a new log filterer instance of NFTDrop, bound to a specific deployed contract.
func NewNFTDropFilterer(address common.Address, filterer bind.ContractFilterer) (*NFTDropFilterer, error) {
	contract, err := bindNFTDrop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFTDropFilterer{contract: contract}, nil
}

// bindNFTDrop binds a generic wrapper to an already deployed contract.
func bindNFTDrop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NFTDropABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTDrop *NFTDropRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTDrop.Contract.NFTDropCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTDrop *NFTDropRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTDrop.Contract.NFTDropTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTDrop *NFTDropRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTDrop.Contract.NFTDropTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTDrop *NFTDropCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTDrop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTDrop *NFTDropTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTDrop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTDrop *NFTDropTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTDrop.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NFTDrop *NFTDropCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NFTDrop *NFTDropSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NFTDrop.Contract.DEFAULTADMINROLE(&_NFTDrop.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NFTDrop *NFTDropCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NFTDrop.Contract.DEFAULTADMINROLE(&_NFTDrop.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NFTDrop *NFTDropCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NFTDrop *NFTDropSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NFTDrop.Contract.BalanceOf(&_NFTDrop.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NFTDrop *NFTDropCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NFTDrop.Contract.BalanceOf(&_NFTDrop.CallOpts, owner)
}

// BatchFrozen is a free data retrieval call binding the contract method 0x83040532.
//
// Solidity: function batchFrozen(uint256 ) view returns(bool)
func (_NFTDrop *NFTDropCaller) BatchFrozen(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "batchFrozen", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BatchFrozen is a free data retrieval call binding the contract method 0x83040532.
//
// Solidity: function batchFrozen(uint256 ) view returns(bool)
func (_NFTDrop *NFTDropSession) BatchFrozen(arg0 *big.Int) (bool, error) {
	return _NFTDrop.Contract.BatchFrozen(&_NFTDrop.CallOpts, arg0)
}

// BatchFrozen is a free data retrieval call binding the contract method 0x83040532.
//
// Solidity: function batchFrozen(uint256 ) view returns(bool)
func (_NFTDrop *NFTDropCallerSession) BatchFrozen(arg0 *big.Int) (bool, error) {
	return _NFTDrop.Contract.BatchFrozen(&_NFTDrop.CallOpts, arg0)
}

// ClaimCondition is a free data retrieval call binding the contract method 0xd637ed59.
//
// Solidity: function claimCondition() view returns(uint256 currentStartId, uint256 count)
func (_NFTDrop *NFTDropCaller) ClaimCondition(opts *bind.CallOpts) (struct {
	CurrentStartId *big.Int
	Count          *big.Int
}, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "claimCondition")

	outstruct := new(struct {
		CurrentStartId *big.Int
		Count          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentStartId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Count = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ClaimCondition is a free data retrieval call binding the contract method 0xd637ed59.
//
// Solidity: function claimCondition() view returns(uint256 currentStartId, uint256 count)
func (_NFTDrop *NFTDropSession) ClaimCondition() (struct {
	CurrentStartId *big.Int
	Count          *big.Int
}, error) {
	return _NFTDrop.Contract.ClaimCondition(&_NFTDrop.CallOpts)
}

// ClaimCondition is a free data retrieval call binding the contract method 0xd637ed59.
//
// Solidity: function claimCondition() view returns(uint256 currentStartId, uint256 count)
func (_NFTDrop *NFTDropCallerSession) ClaimCondition() (struct {
	CurrentStartId *big.Int
	Count          *big.Int
}, error) {
	return _NFTDrop.Contract.ClaimCondition(&_NFTDrop.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() pure returns(bytes32)
func (_NFTDrop *NFTDropCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() pure returns(bytes32)
func (_NFTDrop *NFTDropSession) ContractType() ([32]byte, error) {
	return _NFTDrop.Contract.ContractType(&_NFTDrop.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() pure returns(bytes32)
func (_NFTDrop *NFTDropCallerSession) ContractType() ([32]byte, error) {
	return _NFTDrop.Contract.ContractType(&_NFTDrop.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_NFTDrop *NFTDropCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_NFTDrop *NFTDropSession) ContractURI() (string, error) {
	return _NFTDrop.Contract.ContractURI(&_NFTDrop.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_NFTDrop *NFTDropCallerSession) ContractURI() (string, error) {
	return _NFTDrop.Contract.ContractURI(&_NFTDrop.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint8)
func (_NFTDrop *NFTDropCaller) ContractVersion(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "contractVersion")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint8)
func (_NFTDrop *NFTDropSession) ContractVersion() (uint8, error) {
	return _NFTDrop.Contract.ContractVersion(&_NFTDrop.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint8)
func (_NFTDrop *NFTDropCallerSession) ContractVersion() (uint8, error) {
	return _NFTDrop.Contract.ContractVersion(&_NFTDrop.CallOpts)
}

// EncryptDecrypt is a free data retrieval call binding the contract method 0xe7150322.
//
// Solidity: function encryptDecrypt(bytes data, bytes key) pure returns(bytes result)
func (_NFTDrop *NFTDropCaller) EncryptDecrypt(opts *bind.CallOpts, data []byte, key []byte) ([]byte, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "encryptDecrypt", data, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncryptDecrypt is a free data retrieval call binding the contract method 0xe7150322.
//
// Solidity: function encryptDecrypt(bytes data, bytes key) pure returns(bytes result)
func (_NFTDrop *NFTDropSession) EncryptDecrypt(data []byte, key []byte) ([]byte, error) {
	return _NFTDrop.Contract.EncryptDecrypt(&_NFTDrop.CallOpts, data, key)
}

// EncryptDecrypt is a free data retrieval call binding the contract method 0xe7150322.
//
// Solidity: function encryptDecrypt(bytes data, bytes key) pure returns(bytes result)
func (_NFTDrop *NFTDropCallerSession) EncryptDecrypt(data []byte, key []byte) ([]byte, error) {
	return _NFTDrop.Contract.EncryptDecrypt(&_NFTDrop.CallOpts, data, key)
}

// EncryptedData is a free data retrieval call binding the contract method 0xa05112fc.
//
// Solidity: function encryptedData(uint256 ) view returns(bytes)
func (_NFTDrop *NFTDropCaller) EncryptedData(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "encryptedData", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncryptedData is a free data retrieval call binding the contract method 0xa05112fc.
//
// Solidity: function encryptedData(uint256 ) view returns(bytes)
func (_NFTDrop *NFTDropSession) EncryptedData(arg0 *big.Int) ([]byte, error) {
	return _NFTDrop.Contract.EncryptedData(&_NFTDrop.CallOpts, arg0)
}

// EncryptedData is a free data retrieval call binding the contract method 0xa05112fc.
//
// Solidity: function encryptedData(uint256 ) view returns(bytes)
func (_NFTDrop *NFTDropCallerSession) EncryptedData(arg0 *big.Int) ([]byte, error) {
	return _NFTDrop.Contract.EncryptedData(&_NFTDrop.CallOpts, arg0)
}

// GetActiveClaimConditionId is a free data retrieval call binding the contract method 0xc68907de.
//
// Solidity: function getActiveClaimConditionId() view returns(uint256)
func (_NFTDrop *NFTDropCaller) GetActiveClaimConditionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "getActiveClaimConditionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveClaimConditionId is a free data retrieval call binding the contract method 0xc68907de.
//
// Solidity: function getActiveClaimConditionId() view returns(uint256)
func (_NFTDrop *NFTDropSession) GetActiveClaimConditionId() (*big.Int, error) {
	return _NFTDrop.Contract.GetActiveClaimConditionId(&_NFTDrop.CallOpts)
}

// GetActiveClaimConditionId is a free data retrieval call binding the contract method 0xc68907de.
//
// Solidity: function getActiveClaimConditionId() view returns(uint256)
func (_NFTDrop *NFTDropCallerSession) GetActiveClaimConditionId() (*big.Int, error) {
	return _NFTDrop.Contract.GetActiveClaimConditionId(&_NFTDrop.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NFTDrop *NFTDropCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NFTDrop *NFTDropSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NFTDrop.Contract.GetApproved(&_NFTDrop.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NFTDrop *NFTDropCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NFTDrop.Contract.GetApproved(&_NFTDrop.CallOpts, tokenId)
}

// GetBaseURICount is a free data retrieval call binding the contract method 0x63b45e2d.
//
// Solidity: function getBaseURICount() view returns(uint256)
func (_NFTDrop *NFTDropCaller) GetBaseURICount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "getBaseURICount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBaseURICount is a free data retrieval call binding the contract method 0x63b45e2d.
//
// Solidity: function getBaseURICount() view returns(uint256)
func (_NFTDrop *NFTDropSession) GetBaseURICount() (*big.Int, error) {
	return _NFTDrop.Contract.GetBaseURICount(&_NFTDrop.CallOpts)
}

// GetBaseURICount is a free data retrieval call binding the contract method 0x63b45e2d.
//
// Solidity: function getBaseURICount() view returns(uint256)
func (_NFTDrop *NFTDropCallerSession) GetBaseURICount() (*big.Int, error) {
	return _NFTDrop.Contract.GetBaseURICount(&_NFTDrop.CallOpts)
}

// GetBatchIdAtIndex is a free data retrieval call binding the contract method 0x2419f51b.
//
// Solidity: function getBatchIdAtIndex(uint256 _index) view returns(uint256)
func (_NFTDrop *NFTDropCaller) GetBatchIdAtIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "getBatchIdAtIndex", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBatchIdAtIndex is a free data retrieval call binding the contract method 0x2419f51b.
//
// Solidity: function getBatchIdAtIndex(uint256 _index) view returns(uint256)
func (_NFTDrop *NFTDropSession) GetBatchIdAtIndex(_index *big.Int) (*big.Int, error) {
	return _NFTDrop.Contract.GetBatchIdAtIndex(&_NFTDrop.CallOpts, _index)
}

// GetBatchIdAtIndex is a free data retrieval call binding the contract method 0x2419f51b.
//
// Solidity: function getBatchIdAtIndex(uint256 _index) view returns(uint256)
func (_NFTDrop *NFTDropCallerSession) GetBatchIdAtIndex(_index *big.Int) (*big.Int, error) {
	return _NFTDrop.Contract.GetBatchIdAtIndex(&_NFTDrop.CallOpts, _index)
}

// GetClaimConditionById is a free data retrieval call binding the contract method 0x6f8934f4.
//
// Solidity: function getClaimConditionById(uint256 _conditionId) view returns((uint256,uint256,uint256,uint256,bytes32,uint256,address,string) condition)
func (_NFTDrop *NFTDropCaller) GetClaimConditionById(opts *bind.CallOpts, _conditionId *big.Int) (IClaimConditionClaimCondition, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "getClaimConditionById", _conditionId)

	if err != nil {
		return *new(IClaimConditionClaimCondition), err
	}

	out0 := *abi.ConvertType(out[0], new(IClaimConditionClaimCondition)).(*IClaimConditionClaimCondition)

	return out0, err

}

// GetClaimConditionById is a free data retrieval call binding the contract method 0x6f8934f4.
//
// Solidity: function getClaimConditionById(uint256 _conditionId) view returns((uint256,uint256,uint256,uint256,bytes32,uint256,address,string) condition)
func (_NFTDrop *NFTDropSession) GetClaimConditionById(_conditionId *big.Int) (IClaimConditionClaimCondition, error) {
	return _NFTDrop.Contract.GetClaimConditionById(&_NFTDrop.CallOpts, _conditionId)
}

// GetClaimConditionById is a free data retrieval call binding the contract method 0x6f8934f4.
//
// Solidity: function getClaimConditionById(uint256 _conditionId) view returns((uint256,uint256,uint256,uint256,bytes32,uint256,address,string) condition)
func (_NFTDrop *NFTDropCallerSession) GetClaimConditionById(_conditionId *big.Int) (IClaimConditionClaimCondition, error) {
	return _NFTDrop.Contract.GetClaimConditionById(&_NFTDrop.CallOpts, _conditionId)
}

// GetDefaultRoyaltyInfo is a free data retrieval call binding the contract method 0xb24f2d39.
//
// Solidity: function getDefaultRoyaltyInfo() view returns(address, uint16)
func (_NFTDrop *NFTDropCaller) GetDefaultRoyaltyInfo(opts *bind.CallOpts) (common.Address, uint16, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "getDefaultRoyaltyInfo")

	if err != nil {
		return *new(common.Address), *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return out0, out1, err

}

// GetDefaultRoyaltyInfo is a free data retrieval call binding the contract method 0xb24f2d39.
//
// Solidity: function getDefaultRoyaltyInfo() view returns(address, uint16)
func (_NFTDrop *NFTDropSession) GetDefaultRoyaltyInfo() (common.Address, uint16, error) {
	return _NFTDrop.Contract.GetDefaultRoyaltyInfo(&_NFTDrop.CallOpts)
}

// GetDefaultRoyaltyInfo is a free data retrieval call binding the contract method 0xb24f2d39.
//
// Solidity: function getDefaultRoyaltyInfo() view returns(address, uint16)
func (_NFTDrop *NFTDropCallerSession) GetDefaultRoyaltyInfo() (common.Address, uint16, error) {
	return _NFTDrop.Contract.GetDefaultRoyaltyInfo(&_NFTDrop.CallOpts)
}

// GetFlatPlatformFeeInfo is a free data retrieval call binding the contract method 0xe57553da.
//
// Solidity: function getFlatPlatformFeeInfo() view returns(address, uint256)
func (_NFTDrop *NFTDropCaller) GetFlatPlatformFeeInfo(opts *bind.CallOpts) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "getFlatPlatformFeeInfo")

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetFlatPlatformFeeInfo is a free data retrieval call binding the contract method 0xe57553da.
//
// Solidity: function getFlatPlatformFeeInfo() view returns(address, uint256)
func (_NFTDrop *NFTDropSession) GetFlatPlatformFeeInfo() (common.Address, *big.Int, error) {
	return _NFTDrop.Contract.GetFlatPlatformFeeInfo(&_NFTDrop.CallOpts)
}

// GetFlatPlatformFeeInfo is a free data retrieval call binding the contract method 0xe57553da.
//
// Solidity: function getFlatPlatformFeeInfo() view returns(address, uint256)
func (_NFTDrop *NFTDropCallerSession) GetFlatPlatformFeeInfo() (common.Address, *big.Int, error) {
	return _NFTDrop.Contract.GetFlatPlatformFeeInfo(&_NFTDrop.CallOpts)
}

// GetPlatformFeeInfo is a free data retrieval call binding the contract method 0xd45573f6.
//
// Solidity: function getPlatformFeeInfo() view returns(address, uint16)
func (_NFTDrop *NFTDropCaller) GetPlatformFeeInfo(opts *bind.CallOpts) (common.Address, uint16, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "getPlatformFeeInfo")

	if err != nil {
		return *new(common.Address), *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return out0, out1, err

}

// GetPlatformFeeInfo is a free data retrieval call binding the contract method 0xd45573f6.
//
// Solidity: function getPlatformFeeInfo() view returns(address, uint16)
func (_NFTDrop *NFTDropSession) GetPlatformFeeInfo() (common.Address, uint16, error) {
	return _NFTDrop.Contract.GetPlatformFeeInfo(&_NFTDrop.CallOpts)
}

// GetPlatformFeeInfo is a free data retrieval call binding the contract method 0xd45573f6.
//
// Solidity: function getPlatformFeeInfo() view returns(address, uint16)
func (_NFTDrop *NFTDropCallerSession) GetPlatformFeeInfo() (common.Address, uint16, error) {
	return _NFTDrop.Contract.GetPlatformFeeInfo(&_NFTDrop.CallOpts)
}

// GetPlatformFeeType is a free data retrieval call binding the contract method 0xf28083c3.
//
// Solidity: function getPlatformFeeType() view returns(uint8)
func (_NFTDrop *NFTDropCaller) GetPlatformFeeType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "getPlatformFeeType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPlatformFeeType is a free data retrieval call binding the contract method 0xf28083c3.
//
// Solidity: function getPlatformFeeType() view returns(uint8)
func (_NFTDrop *NFTDropSession) GetPlatformFeeType() (uint8, error) {
	return _NFTDrop.Contract.GetPlatformFeeType(&_NFTDrop.CallOpts)
}

// GetPlatformFeeType is a free data retrieval call binding the contract method 0xf28083c3.
//
// Solidity: function getPlatformFeeType() view returns(uint8)
func (_NFTDrop *NFTDropCallerSession) GetPlatformFeeType() (uint8, error) {
	return _NFTDrop.Contract.GetPlatformFeeType(&_NFTDrop.CallOpts)
}

// GetRevealURI is a free data retrieval call binding the contract method 0x9fc4d68f.
//
// Solidity: function getRevealURI(uint256 _batchId, bytes _key) view returns(string revealedURI)
func (_NFTDrop *NFTDropCaller) GetRevealURI(opts *bind.CallOpts, _batchId *big.Int, _key []byte) (string, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "getRevealURI", _batchId, _key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetRevealURI is a free data retrieval call binding the contract method 0x9fc4d68f.
//
// Solidity: function getRevealURI(uint256 _batchId, bytes _key) view returns(string revealedURI)
func (_NFTDrop *NFTDropSession) GetRevealURI(_batchId *big.Int, _key []byte) (string, error) {
	return _NFTDrop.Contract.GetRevealURI(&_NFTDrop.CallOpts, _batchId, _key)
}

// GetRevealURI is a free data retrieval call binding the contract method 0x9fc4d68f.
//
// Solidity: function getRevealURI(uint256 _batchId, bytes _key) view returns(string revealedURI)
func (_NFTDrop *NFTDropCallerSession) GetRevealURI(_batchId *big.Int, _key []byte) (string, error) {
	return _NFTDrop.Contract.GetRevealURI(&_NFTDrop.CallOpts, _batchId, _key)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NFTDrop *NFTDropCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NFTDrop *NFTDropSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NFTDrop.Contract.GetRoleAdmin(&_NFTDrop.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NFTDrop *NFTDropCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NFTDrop.Contract.GetRoleAdmin(&_NFTDrop.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address member)
func (_NFTDrop *NFTDropCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address member)
func (_NFTDrop *NFTDropSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _NFTDrop.Contract.GetRoleMember(&_NFTDrop.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address member)
func (_NFTDrop *NFTDropCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _NFTDrop.Contract.GetRoleMember(&_NFTDrop.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256 count)
func (_NFTDrop *NFTDropCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256 count)
func (_NFTDrop *NFTDropSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _NFTDrop.Contract.GetRoleMemberCount(&_NFTDrop.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256 count)
func (_NFTDrop *NFTDropCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _NFTDrop.Contract.GetRoleMemberCount(&_NFTDrop.CallOpts, role)
}

// GetRoyaltyInfoForToken is a free data retrieval call binding the contract method 0x4cc157df.
//
// Solidity: function getRoyaltyInfoForToken(uint256 _tokenId) view returns(address, uint16)
func (_NFTDrop *NFTDropCaller) GetRoyaltyInfoForToken(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, uint16, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "getRoyaltyInfoForToken", _tokenId)

	if err != nil {
		return *new(common.Address), *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return out0, out1, err

}

// GetRoyaltyInfoForToken is a free data retrieval call binding the contract method 0x4cc157df.
//
// Solidity: function getRoyaltyInfoForToken(uint256 _tokenId) view returns(address, uint16)
func (_NFTDrop *NFTDropSession) GetRoyaltyInfoForToken(_tokenId *big.Int) (common.Address, uint16, error) {
	return _NFTDrop.Contract.GetRoyaltyInfoForToken(&_NFTDrop.CallOpts, _tokenId)
}

// GetRoyaltyInfoForToken is a free data retrieval call binding the contract method 0x4cc157df.
//
// Solidity: function getRoyaltyInfoForToken(uint256 _tokenId) view returns(address, uint16)
func (_NFTDrop *NFTDropCallerSession) GetRoyaltyInfoForToken(_tokenId *big.Int) (common.Address, uint16, error) {
	return _NFTDrop.Contract.GetRoyaltyInfoForToken(&_NFTDrop.CallOpts, _tokenId)
}

// GetSupplyClaimedByWallet is a free data retrieval call binding the contract method 0xad1eefc5.
//
// Solidity: function getSupplyClaimedByWallet(uint256 _conditionId, address _claimer) view returns(uint256 supplyClaimedByWallet)
func (_NFTDrop *NFTDropCaller) GetSupplyClaimedByWallet(opts *bind.CallOpts, _conditionId *big.Int, _claimer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "getSupplyClaimedByWallet", _conditionId, _claimer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSupplyClaimedByWallet is a free data retrieval call binding the contract method 0xad1eefc5.
//
// Solidity: function getSupplyClaimedByWallet(uint256 _conditionId, address _claimer) view returns(uint256 supplyClaimedByWallet)
func (_NFTDrop *NFTDropSession) GetSupplyClaimedByWallet(_conditionId *big.Int, _claimer common.Address) (*big.Int, error) {
	return _NFTDrop.Contract.GetSupplyClaimedByWallet(&_NFTDrop.CallOpts, _conditionId, _claimer)
}

// GetSupplyClaimedByWallet is a free data retrieval call binding the contract method 0xad1eefc5.
//
// Solidity: function getSupplyClaimedByWallet(uint256 _conditionId, address _claimer) view returns(uint256 supplyClaimedByWallet)
func (_NFTDrop *NFTDropCallerSession) GetSupplyClaimedByWallet(_conditionId *big.Int, _claimer common.Address) (*big.Int, error) {
	return _NFTDrop.Contract.GetSupplyClaimedByWallet(&_NFTDrop.CallOpts, _conditionId, _claimer)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NFTDrop *NFTDropCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NFTDrop *NFTDropSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NFTDrop.Contract.HasRole(&_NFTDrop.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NFTDrop *NFTDropCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NFTDrop.Contract.HasRole(&_NFTDrop.CallOpts, role, account)
}

// HasRoleWithSwitch is a free data retrieval call binding the contract method 0xa32fa5b3.
//
// Solidity: function hasRoleWithSwitch(bytes32 role, address account) view returns(bool)
func (_NFTDrop *NFTDropCaller) HasRoleWithSwitch(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "hasRoleWithSwitch", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRoleWithSwitch is a free data retrieval call binding the contract method 0xa32fa5b3.
//
// Solidity: function hasRoleWithSwitch(bytes32 role, address account) view returns(bool)
func (_NFTDrop *NFTDropSession) HasRoleWithSwitch(role [32]byte, account common.Address) (bool, error) {
	return _NFTDrop.Contract.HasRoleWithSwitch(&_NFTDrop.CallOpts, role, account)
}

// HasRoleWithSwitch is a free data retrieval call binding the contract method 0xa32fa5b3.
//
// Solidity: function hasRoleWithSwitch(bytes32 role, address account) view returns(bool)
func (_NFTDrop *NFTDropCallerSession) HasRoleWithSwitch(role [32]byte, account common.Address) (bool, error) {
	return _NFTDrop.Contract.HasRoleWithSwitch(&_NFTDrop.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NFTDrop *NFTDropCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NFTDrop *NFTDropSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NFTDrop.Contract.IsApprovedForAll(&_NFTDrop.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NFTDrop *NFTDropCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NFTDrop.Contract.IsApprovedForAll(&_NFTDrop.CallOpts, owner, operator)
}

// IsEncryptedBatch is a free data retrieval call binding the contract method 0x492e224b.
//
// Solidity: function isEncryptedBatch(uint256 _batchId) view returns(bool)
func (_NFTDrop *NFTDropCaller) IsEncryptedBatch(opts *bind.CallOpts, _batchId *big.Int) (bool, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "isEncryptedBatch", _batchId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEncryptedBatch is a free data retrieval call binding the contract method 0x492e224b.
//
// Solidity: function isEncryptedBatch(uint256 _batchId) view returns(bool)
func (_NFTDrop *NFTDropSession) IsEncryptedBatch(_batchId *big.Int) (bool, error) {
	return _NFTDrop.Contract.IsEncryptedBatch(&_NFTDrop.CallOpts, _batchId)
}

// IsEncryptedBatch is a free data retrieval call binding the contract method 0x492e224b.
//
// Solidity: function isEncryptedBatch(uint256 _batchId) view returns(bool)
func (_NFTDrop *NFTDropCallerSession) IsEncryptedBatch(_batchId *big.Int) (bool, error) {
	return _NFTDrop.Contract.IsEncryptedBatch(&_NFTDrop.CallOpts, _batchId)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_NFTDrop *NFTDropCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_NFTDrop *NFTDropSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _NFTDrop.Contract.IsTrustedForwarder(&_NFTDrop.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_NFTDrop *NFTDropCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _NFTDrop.Contract.IsTrustedForwarder(&_NFTDrop.CallOpts, forwarder)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_NFTDrop *NFTDropCaller) MaxTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "maxTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_NFTDrop *NFTDropSession) MaxTotalSupply() (*big.Int, error) {
	return _NFTDrop.Contract.MaxTotalSupply(&_NFTDrop.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_NFTDrop *NFTDropCallerSession) MaxTotalSupply() (*big.Int, error) {
	return _NFTDrop.Contract.MaxTotalSupply(&_NFTDrop.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NFTDrop *NFTDropCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NFTDrop *NFTDropSession) Name() (string, error) {
	return _NFTDrop.Contract.Name(&_NFTDrop.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NFTDrop *NFTDropCallerSession) Name() (string, error) {
	return _NFTDrop.Contract.Name(&_NFTDrop.CallOpts)
}

// NextTokenIdToClaim is a free data retrieval call binding the contract method 0xacd083f8.
//
// Solidity: function nextTokenIdToClaim() view returns(uint256)
func (_NFTDrop *NFTDropCaller) NextTokenIdToClaim(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "nextTokenIdToClaim")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenIdToClaim is a free data retrieval call binding the contract method 0xacd083f8.
//
// Solidity: function nextTokenIdToClaim() view returns(uint256)
func (_NFTDrop *NFTDropSession) NextTokenIdToClaim() (*big.Int, error) {
	return _NFTDrop.Contract.NextTokenIdToClaim(&_NFTDrop.CallOpts)
}

// NextTokenIdToClaim is a free data retrieval call binding the contract method 0xacd083f8.
//
// Solidity: function nextTokenIdToClaim() view returns(uint256)
func (_NFTDrop *NFTDropCallerSession) NextTokenIdToClaim() (*big.Int, error) {
	return _NFTDrop.Contract.NextTokenIdToClaim(&_NFTDrop.CallOpts)
}

// NextTokenIdToMint is a free data retrieval call binding the contract method 0x3b1475a7.
//
// Solidity: function nextTokenIdToMint() view returns(uint256)
func (_NFTDrop *NFTDropCaller) NextTokenIdToMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "nextTokenIdToMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenIdToMint is a free data retrieval call binding the contract method 0x3b1475a7.
//
// Solidity: function nextTokenIdToMint() view returns(uint256)
func (_NFTDrop *NFTDropSession) NextTokenIdToMint() (*big.Int, error) {
	return _NFTDrop.Contract.NextTokenIdToMint(&_NFTDrop.CallOpts)
}

// NextTokenIdToMint is a free data retrieval call binding the contract method 0x3b1475a7.
//
// Solidity: function nextTokenIdToMint() view returns(uint256)
func (_NFTDrop *NFTDropCallerSession) NextTokenIdToMint() (*big.Int, error) {
	return _NFTDrop.Contract.NextTokenIdToMint(&_NFTDrop.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTDrop *NFTDropCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTDrop *NFTDropSession) Owner() (common.Address, error) {
	return _NFTDrop.Contract.Owner(&_NFTDrop.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTDrop *NFTDropCallerSession) Owner() (common.Address, error) {
	return _NFTDrop.Contract.Owner(&_NFTDrop.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NFTDrop *NFTDropCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NFTDrop *NFTDropSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NFTDrop.Contract.OwnerOf(&_NFTDrop.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NFTDrop *NFTDropCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NFTDrop.Contract.OwnerOf(&_NFTDrop.CallOpts, tokenId)
}

// PrimarySaleRecipient is a free data retrieval call binding the contract method 0x079fe40e.
//
// Solidity: function primarySaleRecipient() view returns(address)
func (_NFTDrop *NFTDropCaller) PrimarySaleRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "primarySaleRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrimarySaleRecipient is a free data retrieval call binding the contract method 0x079fe40e.
//
// Solidity: function primarySaleRecipient() view returns(address)
func (_NFTDrop *NFTDropSession) PrimarySaleRecipient() (common.Address, error) {
	return _NFTDrop.Contract.PrimarySaleRecipient(&_NFTDrop.CallOpts)
}

// PrimarySaleRecipient is a free data retrieval call binding the contract method 0x079fe40e.
//
// Solidity: function primarySaleRecipient() view returns(address)
func (_NFTDrop *NFTDropCallerSession) PrimarySaleRecipient() (common.Address, error) {
	return _NFTDrop.Contract.PrimarySaleRecipient(&_NFTDrop.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_NFTDrop *NFTDropCaller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_NFTDrop *NFTDropSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _NFTDrop.Contract.RoyaltyInfo(&_NFTDrop.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_NFTDrop *NFTDropCallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _NFTDrop.Contract.RoyaltyInfo(&_NFTDrop.CallOpts, tokenId, salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NFTDrop *NFTDropCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NFTDrop *NFTDropSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NFTDrop.Contract.SupportsInterface(&_NFTDrop.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NFTDrop *NFTDropCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NFTDrop.Contract.SupportsInterface(&_NFTDrop.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NFTDrop *NFTDropCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NFTDrop *NFTDropSession) Symbol() (string, error) {
	return _NFTDrop.Contract.Symbol(&_NFTDrop.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NFTDrop *NFTDropCallerSession) Symbol() (string, error) {
	return _NFTDrop.Contract.Symbol(&_NFTDrop.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_NFTDrop *NFTDropCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_NFTDrop *NFTDropSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _NFTDrop.Contract.TokenURI(&_NFTDrop.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_NFTDrop *NFTDropCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _NFTDrop.Contract.TokenURI(&_NFTDrop.CallOpts, _tokenId)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NFTDrop *NFTDropCaller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "totalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NFTDrop *NFTDropSession) TotalMinted() (*big.Int, error) {
	return _NFTDrop.Contract.TotalMinted(&_NFTDrop.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NFTDrop *NFTDropCallerSession) TotalMinted() (*big.Int, error) {
	return _NFTDrop.Contract.TotalMinted(&_NFTDrop.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NFTDrop *NFTDropCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NFTDrop *NFTDropSession) TotalSupply() (*big.Int, error) {
	return _NFTDrop.Contract.TotalSupply(&_NFTDrop.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NFTDrop *NFTDropCallerSession) TotalSupply() (*big.Int, error) {
	return _NFTDrop.Contract.TotalSupply(&_NFTDrop.CallOpts)
}

// VerifyClaim is a free data retrieval call binding the contract method 0x23a2902b.
//
// Solidity: function verifyClaim(uint256 _conditionId, address _claimer, uint256 _quantity, address _currency, uint256 _pricePerToken, (bytes32[],uint256,uint256,address) _allowlistProof) view returns(bool isOverride)
func (_NFTDrop *NFTDropCaller) VerifyClaim(opts *bind.CallOpts, _conditionId *big.Int, _claimer common.Address, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, _allowlistProof IDropAllowlistProof) (bool, error) {
	var out []interface{}
	err := _NFTDrop.contract.Call(opts, &out, "verifyClaim", _conditionId, _claimer, _quantity, _currency, _pricePerToken, _allowlistProof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyClaim is a free data retrieval call binding the contract method 0x23a2902b.
//
// Solidity: function verifyClaim(uint256 _conditionId, address _claimer, uint256 _quantity, address _currency, uint256 _pricePerToken, (bytes32[],uint256,uint256,address) _allowlistProof) view returns(bool isOverride)
func (_NFTDrop *NFTDropSession) VerifyClaim(_conditionId *big.Int, _claimer common.Address, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, _allowlistProof IDropAllowlistProof) (bool, error) {
	return _NFTDrop.Contract.VerifyClaim(&_NFTDrop.CallOpts, _conditionId, _claimer, _quantity, _currency, _pricePerToken, _allowlistProof)
}

// VerifyClaim is a free data retrieval call binding the contract method 0x23a2902b.
//
// Solidity: function verifyClaim(uint256 _conditionId, address _claimer, uint256 _quantity, address _currency, uint256 _pricePerToken, (bytes32[],uint256,uint256,address) _allowlistProof) view returns(bool isOverride)
func (_NFTDrop *NFTDropCallerSession) VerifyClaim(_conditionId *big.Int, _claimer common.Address, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, _allowlistProof IDropAllowlistProof) (bool, error) {
	return _NFTDrop.Contract.VerifyClaim(&_NFTDrop.CallOpts, _conditionId, _claimer, _quantity, _currency, _pricePerToken, _allowlistProof)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NFTDrop *NFTDropTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NFTDrop *NFTDropSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTDrop.Contract.Approve(&_NFTDrop.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NFTDrop *NFTDropTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTDrop.Contract.Approve(&_NFTDrop.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_NFTDrop *NFTDropTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_NFTDrop *NFTDropSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _NFTDrop.Contract.Burn(&_NFTDrop.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_NFTDrop *NFTDropTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _NFTDrop.Contract.Burn(&_NFTDrop.TransactOpts, tokenId)
}

// Claim is a paid mutator transaction binding the contract method 0x84bb1e42.
//
// Solidity: function claim(address _receiver, uint256 _quantity, address _currency, uint256 _pricePerToken, (bytes32[],uint256,uint256,address) _allowlistProof, bytes _data) payable returns()
func (_NFTDrop *NFTDropTransactor) Claim(opts *bind.TransactOpts, _receiver common.Address, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, _allowlistProof IDropAllowlistProof, _data []byte) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "claim", _receiver, _quantity, _currency, _pricePerToken, _allowlistProof, _data)
}

// Claim is a paid mutator transaction binding the contract method 0x84bb1e42.
//
// Solidity: function claim(address _receiver, uint256 _quantity, address _currency, uint256 _pricePerToken, (bytes32[],uint256,uint256,address) _allowlistProof, bytes _data) payable returns()
func (_NFTDrop *NFTDropSession) Claim(_receiver common.Address, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, _allowlistProof IDropAllowlistProof, _data []byte) (*types.Transaction, error) {
	return _NFTDrop.Contract.Claim(&_NFTDrop.TransactOpts, _receiver, _quantity, _currency, _pricePerToken, _allowlistProof, _data)
}

// Claim is a paid mutator transaction binding the contract method 0x84bb1e42.
//
// Solidity: function claim(address _receiver, uint256 _quantity, address _currency, uint256 _pricePerToken, (bytes32[],uint256,uint256,address) _allowlistProof, bytes _data) payable returns()
func (_NFTDrop *NFTDropTransactorSession) Claim(_receiver common.Address, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, _allowlistProof IDropAllowlistProof, _data []byte) (*types.Transaction, error) {
	return _NFTDrop.Contract.Claim(&_NFTDrop.TransactOpts, _receiver, _quantity, _currency, _pricePerToken, _allowlistProof, _data)
}

// FreezeBatchBaseURI is a paid mutator transaction binding the contract method 0xa07ced9e.
//
// Solidity: function freezeBatchBaseURI(uint256 _index) returns()
func (_NFTDrop *NFTDropTransactor) FreezeBatchBaseURI(opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "freezeBatchBaseURI", _index)
}

// FreezeBatchBaseURI is a paid mutator transaction binding the contract method 0xa07ced9e.
//
// Solidity: function freezeBatchBaseURI(uint256 _index) returns()
func (_NFTDrop *NFTDropSession) FreezeBatchBaseURI(_index *big.Int) (*types.Transaction, error) {
	return _NFTDrop.Contract.FreezeBatchBaseURI(&_NFTDrop.TransactOpts, _index)
}

// FreezeBatchBaseURI is a paid mutator transaction binding the contract method 0xa07ced9e.
//
// Solidity: function freezeBatchBaseURI(uint256 _index) returns()
func (_NFTDrop *NFTDropTransactorSession) FreezeBatchBaseURI(_index *big.Int) (*types.Transaction, error) {
	return _NFTDrop.Contract.FreezeBatchBaseURI(&_NFTDrop.TransactOpts, _index)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NFTDrop *NFTDropTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NFTDrop *NFTDropSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFTDrop.Contract.GrantRole(&_NFTDrop.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NFTDrop *NFTDropTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFTDrop.Contract.GrantRole(&_NFTDrop.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xe1591634.
//
// Solidity: function initialize(address _defaultAdmin, string _name, string _symbol, string _contractURI, address[] _trustedForwarders, address _saleRecipient, address _royaltyRecipient, uint128 _royaltyBps, uint128 _platformFeeBps, address _platformFeeRecipient) returns()
func (_NFTDrop *NFTDropTransactor) Initialize(opts *bind.TransactOpts, _defaultAdmin common.Address, _name string, _symbol string, _contractURI string, _trustedForwarders []common.Address, _saleRecipient common.Address, _royaltyRecipient common.Address, _royaltyBps *big.Int, _platformFeeBps *big.Int, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "initialize", _defaultAdmin, _name, _symbol, _contractURI, _trustedForwarders, _saleRecipient, _royaltyRecipient, _royaltyBps, _platformFeeBps, _platformFeeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xe1591634.
//
// Solidity: function initialize(address _defaultAdmin, string _name, string _symbol, string _contractURI, address[] _trustedForwarders, address _saleRecipient, address _royaltyRecipient, uint128 _royaltyBps, uint128 _platformFeeBps, address _platformFeeRecipient) returns()
func (_NFTDrop *NFTDropSession) Initialize(_defaultAdmin common.Address, _name string, _symbol string, _contractURI string, _trustedForwarders []common.Address, _saleRecipient common.Address, _royaltyRecipient common.Address, _royaltyBps *big.Int, _platformFeeBps *big.Int, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _NFTDrop.Contract.Initialize(&_NFTDrop.TransactOpts, _defaultAdmin, _name, _symbol, _contractURI, _trustedForwarders, _saleRecipient, _royaltyRecipient, _royaltyBps, _platformFeeBps, _platformFeeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xe1591634.
//
// Solidity: function initialize(address _defaultAdmin, string _name, string _symbol, string _contractURI, address[] _trustedForwarders, address _saleRecipient, address _royaltyRecipient, uint128 _royaltyBps, uint128 _platformFeeBps, address _platformFeeRecipient) returns()
func (_NFTDrop *NFTDropTransactorSession) Initialize(_defaultAdmin common.Address, _name string, _symbol string, _contractURI string, _trustedForwarders []common.Address, _saleRecipient common.Address, _royaltyRecipient common.Address, _royaltyBps *big.Int, _platformFeeBps *big.Int, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _NFTDrop.Contract.Initialize(&_NFTDrop.TransactOpts, _defaultAdmin, _name, _symbol, _contractURI, _trustedForwarders, _saleRecipient, _royaltyRecipient, _royaltyBps, _platformFeeBps, _platformFeeRecipient)
}

// LazyMint is a paid mutator transaction binding the contract method 0xd37c353b.
//
// Solidity: function lazyMint(uint256 _amount, string _baseURIForTokens, bytes _data) returns(uint256 batchId)
func (_NFTDrop *NFTDropTransactor) LazyMint(opts *bind.TransactOpts, _amount *big.Int, _baseURIForTokens string, _data []byte) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "lazyMint", _amount, _baseURIForTokens, _data)
}

// LazyMint is a paid mutator transaction binding the contract method 0xd37c353b.
//
// Solidity: function lazyMint(uint256 _amount, string _baseURIForTokens, bytes _data) returns(uint256 batchId)
func (_NFTDrop *NFTDropSession) LazyMint(_amount *big.Int, _baseURIForTokens string, _data []byte) (*types.Transaction, error) {
	return _NFTDrop.Contract.LazyMint(&_NFTDrop.TransactOpts, _amount, _baseURIForTokens, _data)
}

// LazyMint is a paid mutator transaction binding the contract method 0xd37c353b.
//
// Solidity: function lazyMint(uint256 _amount, string _baseURIForTokens, bytes _data) returns(uint256 batchId)
func (_NFTDrop *NFTDropTransactorSession) LazyMint(_amount *big.Int, _baseURIForTokens string, _data []byte) (*types.Transaction, error) {
	return _NFTDrop.Contract.LazyMint(&_NFTDrop.TransactOpts, _amount, _baseURIForTokens, _data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_NFTDrop *NFTDropTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_NFTDrop *NFTDropSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _NFTDrop.Contract.Multicall(&_NFTDrop.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_NFTDrop *NFTDropTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _NFTDrop.Contract.Multicall(&_NFTDrop.TransactOpts, data)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NFTDrop *NFTDropTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NFTDrop *NFTDropSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFTDrop.Contract.RenounceRole(&_NFTDrop.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NFTDrop *NFTDropTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFTDrop.Contract.RenounceRole(&_NFTDrop.TransactOpts, role, account)
}

// Reveal is a paid mutator transaction binding the contract method 0xce805642.
//
// Solidity: function reveal(uint256 _index, bytes _key) returns(string revealedURI)
func (_NFTDrop *NFTDropTransactor) Reveal(opts *bind.TransactOpts, _index *big.Int, _key []byte) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "reveal", _index, _key)
}

// Reveal is a paid mutator transaction binding the contract method 0xce805642.
//
// Solidity: function reveal(uint256 _index, bytes _key) returns(string revealedURI)
func (_NFTDrop *NFTDropSession) Reveal(_index *big.Int, _key []byte) (*types.Transaction, error) {
	return _NFTDrop.Contract.Reveal(&_NFTDrop.TransactOpts, _index, _key)
}

// Reveal is a paid mutator transaction binding the contract method 0xce805642.
//
// Solidity: function reveal(uint256 _index, bytes _key) returns(string revealedURI)
func (_NFTDrop *NFTDropTransactorSession) Reveal(_index *big.Int, _key []byte) (*types.Transaction, error) {
	return _NFTDrop.Contract.Reveal(&_NFTDrop.TransactOpts, _index, _key)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NFTDrop *NFTDropTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NFTDrop *NFTDropSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFTDrop.Contract.RevokeRole(&_NFTDrop.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NFTDrop *NFTDropTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFTDrop.Contract.RevokeRole(&_NFTDrop.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NFTDrop *NFTDropTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NFTDrop *NFTDropSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTDrop.Contract.SafeTransferFrom(&_NFTDrop.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NFTDrop *NFTDropTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTDrop.Contract.SafeTransferFrom(&_NFTDrop.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_NFTDrop *NFTDropTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_NFTDrop *NFTDropSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _NFTDrop.Contract.SafeTransferFrom0(&_NFTDrop.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_NFTDrop *NFTDropTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _NFTDrop.Contract.SafeTransferFrom0(&_NFTDrop.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NFTDrop *NFTDropTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NFTDrop *NFTDropSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetApprovalForAll(&_NFTDrop.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NFTDrop *NFTDropTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetApprovalForAll(&_NFTDrop.TransactOpts, operator, approved)
}

// SetClaimConditions is a paid mutator transaction binding the contract method 0x74bc7db7.
//
// Solidity: function setClaimConditions((uint256,uint256,uint256,uint256,bytes32,uint256,address,string)[] _conditions, bool _resetClaimEligibility) returns()
func (_NFTDrop *NFTDropTransactor) SetClaimConditions(opts *bind.TransactOpts, _conditions []IClaimConditionClaimCondition, _resetClaimEligibility bool) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "setClaimConditions", _conditions, _resetClaimEligibility)
}

// SetClaimConditions is a paid mutator transaction binding the contract method 0x74bc7db7.
//
// Solidity: function setClaimConditions((uint256,uint256,uint256,uint256,bytes32,uint256,address,string)[] _conditions, bool _resetClaimEligibility) returns()
func (_NFTDrop *NFTDropSession) SetClaimConditions(_conditions []IClaimConditionClaimCondition, _resetClaimEligibility bool) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetClaimConditions(&_NFTDrop.TransactOpts, _conditions, _resetClaimEligibility)
}

// SetClaimConditions is a paid mutator transaction binding the contract method 0x74bc7db7.
//
// Solidity: function setClaimConditions((uint256,uint256,uint256,uint256,bytes32,uint256,address,string)[] _conditions, bool _resetClaimEligibility) returns()
func (_NFTDrop *NFTDropTransactorSession) SetClaimConditions(_conditions []IClaimConditionClaimCondition, _resetClaimEligibility bool) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetClaimConditions(&_NFTDrop.TransactOpts, _conditions, _resetClaimEligibility)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_NFTDrop *NFTDropTransactor) SetContractURI(opts *bind.TransactOpts, _uri string) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "setContractURI", _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_NFTDrop *NFTDropSession) SetContractURI(_uri string) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetContractURI(&_NFTDrop.TransactOpts, _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_NFTDrop *NFTDropTransactorSession) SetContractURI(_uri string) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetContractURI(&_NFTDrop.TransactOpts, _uri)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_NFTDrop *NFTDropTransactor) SetDefaultRoyaltyInfo(opts *bind.TransactOpts, _royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "setDefaultRoyaltyInfo", _royaltyRecipient, _royaltyBps)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_NFTDrop *NFTDropSession) SetDefaultRoyaltyInfo(_royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetDefaultRoyaltyInfo(&_NFTDrop.TransactOpts, _royaltyRecipient, _royaltyBps)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_NFTDrop *NFTDropTransactorSession) SetDefaultRoyaltyInfo(_royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetDefaultRoyaltyInfo(&_NFTDrop.TransactOpts, _royaltyRecipient, _royaltyBps)
}

// SetFlatPlatformFeeInfo is a paid mutator transaction binding the contract method 0x7e54523c.
//
// Solidity: function setFlatPlatformFeeInfo(address _platformFeeRecipient, uint256 _flatFee) returns()
func (_NFTDrop *NFTDropTransactor) SetFlatPlatformFeeInfo(opts *bind.TransactOpts, _platformFeeRecipient common.Address, _flatFee *big.Int) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "setFlatPlatformFeeInfo", _platformFeeRecipient, _flatFee)
}

// SetFlatPlatformFeeInfo is a paid mutator transaction binding the contract method 0x7e54523c.
//
// Solidity: function setFlatPlatformFeeInfo(address _platformFeeRecipient, uint256 _flatFee) returns()
func (_NFTDrop *NFTDropSession) SetFlatPlatformFeeInfo(_platformFeeRecipient common.Address, _flatFee *big.Int) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetFlatPlatformFeeInfo(&_NFTDrop.TransactOpts, _platformFeeRecipient, _flatFee)
}

// SetFlatPlatformFeeInfo is a paid mutator transaction binding the contract method 0x7e54523c.
//
// Solidity: function setFlatPlatformFeeInfo(address _platformFeeRecipient, uint256 _flatFee) returns()
func (_NFTDrop *NFTDropTransactorSession) SetFlatPlatformFeeInfo(_platformFeeRecipient common.Address, _flatFee *big.Int) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetFlatPlatformFeeInfo(&_NFTDrop.TransactOpts, _platformFeeRecipient, _flatFee)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x3f3e4c11.
//
// Solidity: function setMaxTotalSupply(uint256 _maxTotalSupply) returns()
func (_NFTDrop *NFTDropTransactor) SetMaxTotalSupply(opts *bind.TransactOpts, _maxTotalSupply *big.Int) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "setMaxTotalSupply", _maxTotalSupply)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x3f3e4c11.
//
// Solidity: function setMaxTotalSupply(uint256 _maxTotalSupply) returns()
func (_NFTDrop *NFTDropSession) SetMaxTotalSupply(_maxTotalSupply *big.Int) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetMaxTotalSupply(&_NFTDrop.TransactOpts, _maxTotalSupply)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x3f3e4c11.
//
// Solidity: function setMaxTotalSupply(uint256 _maxTotalSupply) returns()
func (_NFTDrop *NFTDropTransactorSession) SetMaxTotalSupply(_maxTotalSupply *big.Int) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetMaxTotalSupply(&_NFTDrop.TransactOpts, _maxTotalSupply)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_NFTDrop *NFTDropTransactor) SetOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "setOwner", _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_NFTDrop *NFTDropSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetOwner(&_NFTDrop.TransactOpts, _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_NFTDrop *NFTDropTransactorSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetOwner(&_NFTDrop.TransactOpts, _newOwner)
}

// SetPlatformFeeInfo is a paid mutator transaction binding the contract method 0x1e7ac488.
//
// Solidity: function setPlatformFeeInfo(address _platformFeeRecipient, uint256 _platformFeeBps) returns()
func (_NFTDrop *NFTDropTransactor) SetPlatformFeeInfo(opts *bind.TransactOpts, _platformFeeRecipient common.Address, _platformFeeBps *big.Int) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "setPlatformFeeInfo", _platformFeeRecipient, _platformFeeBps)
}

// SetPlatformFeeInfo is a paid mutator transaction binding the contract method 0x1e7ac488.
//
// Solidity: function setPlatformFeeInfo(address _platformFeeRecipient, uint256 _platformFeeBps) returns()
func (_NFTDrop *NFTDropSession) SetPlatformFeeInfo(_platformFeeRecipient common.Address, _platformFeeBps *big.Int) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetPlatformFeeInfo(&_NFTDrop.TransactOpts, _platformFeeRecipient, _platformFeeBps)
}

// SetPlatformFeeInfo is a paid mutator transaction binding the contract method 0x1e7ac488.
//
// Solidity: function setPlatformFeeInfo(address _platformFeeRecipient, uint256 _platformFeeBps) returns()
func (_NFTDrop *NFTDropTransactorSession) SetPlatformFeeInfo(_platformFeeRecipient common.Address, _platformFeeBps *big.Int) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetPlatformFeeInfo(&_NFTDrop.TransactOpts, _platformFeeRecipient, _platformFeeBps)
}

// SetPlatformFeeType is a paid mutator transaction binding the contract method 0xb6f10c79.
//
// Solidity: function setPlatformFeeType(uint8 _feeType) returns()
func (_NFTDrop *NFTDropTransactor) SetPlatformFeeType(opts *bind.TransactOpts, _feeType uint8) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "setPlatformFeeType", _feeType)
}

// SetPlatformFeeType is a paid mutator transaction binding the contract method 0xb6f10c79.
//
// Solidity: function setPlatformFeeType(uint8 _feeType) returns()
func (_NFTDrop *NFTDropSession) SetPlatformFeeType(_feeType uint8) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetPlatformFeeType(&_NFTDrop.TransactOpts, _feeType)
}

// SetPlatformFeeType is a paid mutator transaction binding the contract method 0xb6f10c79.
//
// Solidity: function setPlatformFeeType(uint8 _feeType) returns()
func (_NFTDrop *NFTDropTransactorSession) SetPlatformFeeType(_feeType uint8) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetPlatformFeeType(&_NFTDrop.TransactOpts, _feeType)
}

// SetPrimarySaleRecipient is a paid mutator transaction binding the contract method 0x6f4f2837.
//
// Solidity: function setPrimarySaleRecipient(address _saleRecipient) returns()
func (_NFTDrop *NFTDropTransactor) SetPrimarySaleRecipient(opts *bind.TransactOpts, _saleRecipient common.Address) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "setPrimarySaleRecipient", _saleRecipient)
}

// SetPrimarySaleRecipient is a paid mutator transaction binding the contract method 0x6f4f2837.
//
// Solidity: function setPrimarySaleRecipient(address _saleRecipient) returns()
func (_NFTDrop *NFTDropSession) SetPrimarySaleRecipient(_saleRecipient common.Address) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetPrimarySaleRecipient(&_NFTDrop.TransactOpts, _saleRecipient)
}

// SetPrimarySaleRecipient is a paid mutator transaction binding the contract method 0x6f4f2837.
//
// Solidity: function setPrimarySaleRecipient(address _saleRecipient) returns()
func (_NFTDrop *NFTDropTransactorSession) SetPrimarySaleRecipient(_saleRecipient common.Address) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetPrimarySaleRecipient(&_NFTDrop.TransactOpts, _saleRecipient)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 _tokenId, address _recipient, uint256 _bps) returns()
func (_NFTDrop *NFTDropTransactor) SetRoyaltyInfoForToken(opts *bind.TransactOpts, _tokenId *big.Int, _recipient common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "setRoyaltyInfoForToken", _tokenId, _recipient, _bps)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 _tokenId, address _recipient, uint256 _bps) returns()
func (_NFTDrop *NFTDropSession) SetRoyaltyInfoForToken(_tokenId *big.Int, _recipient common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetRoyaltyInfoForToken(&_NFTDrop.TransactOpts, _tokenId, _recipient, _bps)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 _tokenId, address _recipient, uint256 _bps) returns()
func (_NFTDrop *NFTDropTransactorSession) SetRoyaltyInfoForToken(_tokenId *big.Int, _recipient common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _NFTDrop.Contract.SetRoyaltyInfoForToken(&_NFTDrop.TransactOpts, _tokenId, _recipient, _bps)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NFTDrop *NFTDropTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NFTDrop *NFTDropSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTDrop.Contract.TransferFrom(&_NFTDrop.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NFTDrop *NFTDropTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTDrop.Contract.TransferFrom(&_NFTDrop.TransactOpts, from, to, tokenId)
}

// UpdateBatchBaseURI is a paid mutator transaction binding the contract method 0xde903ddd.
//
// Solidity: function updateBatchBaseURI(uint256 _index, string _uri) returns()
func (_NFTDrop *NFTDropTransactor) UpdateBatchBaseURI(opts *bind.TransactOpts, _index *big.Int, _uri string) (*types.Transaction, error) {
	return _NFTDrop.contract.Transact(opts, "updateBatchBaseURI", _index, _uri)
}

// UpdateBatchBaseURI is a paid mutator transaction binding the contract method 0xde903ddd.
//
// Solidity: function updateBatchBaseURI(uint256 _index, string _uri) returns()
func (_NFTDrop *NFTDropSession) UpdateBatchBaseURI(_index *big.Int, _uri string) (*types.Transaction, error) {
	return _NFTDrop.Contract.UpdateBatchBaseURI(&_NFTDrop.TransactOpts, _index, _uri)
}

// UpdateBatchBaseURI is a paid mutator transaction binding the contract method 0xde903ddd.
//
// Solidity: function updateBatchBaseURI(uint256 _index, string _uri) returns()
func (_NFTDrop *NFTDropTransactorSession) UpdateBatchBaseURI(_index *big.Int, _uri string) (*types.Transaction, error) {
	return _NFTDrop.Contract.UpdateBatchBaseURI(&_NFTDrop.TransactOpts, _index, _uri)
}

// NFTDropApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NFTDrop contract.
type NFTDropApprovalIterator struct {
	Event *NFTDropApproval // Event containing the contract specifics and raw log

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
func (it *NFTDropApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropApproval)
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
		it.Event = new(NFTDropApproval)
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
func (it *NFTDropApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropApproval represents a Approval event raised by the NFTDrop contract.
type NFTDropApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NFTDrop *NFTDropFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*NFTDropApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTDropApprovalIterator{contract: _NFTDrop.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NFTDrop *NFTDropFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NFTDropApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropApproval)
				if err := _NFTDrop.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NFTDrop *NFTDropFilterer) ParseApproval(log types.Log) (*NFTDropApproval, error) {
	event := new(NFTDropApproval)
	if err := _NFTDrop.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDropApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the NFTDrop contract.
type NFTDropApprovalForAllIterator struct {
	Event *NFTDropApprovalForAll // Event containing the contract specifics and raw log

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
func (it *NFTDropApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropApprovalForAll)
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
		it.Event = new(NFTDropApprovalForAll)
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
func (it *NFTDropApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropApprovalForAll represents a ApprovalForAll event raised by the NFTDrop contract.
type NFTDropApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NFTDrop *NFTDropFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*NFTDropApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NFTDropApprovalForAllIterator{contract: _NFTDrop.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NFTDrop *NFTDropFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NFTDropApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropApprovalForAll)
				if err := _NFTDrop.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_NFTDrop *NFTDropFilterer) ParseApprovalForAll(log types.Log) (*NFTDropApprovalForAll, error) {
	event := new(NFTDropApprovalForAll)
	if err := _NFTDrop.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDropBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the NFTDrop contract.
type NFTDropBatchMetadataUpdateIterator struct {
	Event *NFTDropBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *NFTDropBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropBatchMetadataUpdate)
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
		it.Event = new(NFTDropBatchMetadataUpdate)
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
func (it *NFTDropBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the NFTDrop contract.
type NFTDropBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_NFTDrop *NFTDropFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*NFTDropBatchMetadataUpdateIterator, error) {

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &NFTDropBatchMetadataUpdateIterator{contract: _NFTDrop.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_NFTDrop *NFTDropFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *NFTDropBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropBatchMetadataUpdate)
				if err := _NFTDrop.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_NFTDrop *NFTDropFilterer) ParseBatchMetadataUpdate(log types.Log) (*NFTDropBatchMetadataUpdate, error) {
	event := new(NFTDropBatchMetadataUpdate)
	if err := _NFTDrop.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDropClaimConditionsUpdatedIterator is returned from FilterClaimConditionsUpdated and is used to iterate over the raw logs and unpacked data for ClaimConditionsUpdated events raised by the NFTDrop contract.
type NFTDropClaimConditionsUpdatedIterator struct {
	Event *NFTDropClaimConditionsUpdated // Event containing the contract specifics and raw log

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
func (it *NFTDropClaimConditionsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropClaimConditionsUpdated)
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
		it.Event = new(NFTDropClaimConditionsUpdated)
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
func (it *NFTDropClaimConditionsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropClaimConditionsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropClaimConditionsUpdated represents a ClaimConditionsUpdated event raised by the NFTDrop contract.
type NFTDropClaimConditionsUpdated struct {
	ClaimConditions  []IClaimConditionClaimCondition
	ResetEligibility bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterClaimConditionsUpdated is a free log retrieval operation binding the contract event 0xbf4016fceeaaa4ac5cf4be865b559ff85825ab4ca7aa7b661d16e2f544c03098.
//
// Solidity: event ClaimConditionsUpdated((uint256,uint256,uint256,uint256,bytes32,uint256,address,string)[] claimConditions, bool resetEligibility)
func (_NFTDrop *NFTDropFilterer) FilterClaimConditionsUpdated(opts *bind.FilterOpts) (*NFTDropClaimConditionsUpdatedIterator, error) {

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "ClaimConditionsUpdated")
	if err != nil {
		return nil, err
	}
	return &NFTDropClaimConditionsUpdatedIterator{contract: _NFTDrop.contract, event: "ClaimConditionsUpdated", logs: logs, sub: sub}, nil
}

// WatchClaimConditionsUpdated is a free log subscription operation binding the contract event 0xbf4016fceeaaa4ac5cf4be865b559ff85825ab4ca7aa7b661d16e2f544c03098.
//
// Solidity: event ClaimConditionsUpdated((uint256,uint256,uint256,uint256,bytes32,uint256,address,string)[] claimConditions, bool resetEligibility)
func (_NFTDrop *NFTDropFilterer) WatchClaimConditionsUpdated(opts *bind.WatchOpts, sink chan<- *NFTDropClaimConditionsUpdated) (event.Subscription, error) {

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "ClaimConditionsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropClaimConditionsUpdated)
				if err := _NFTDrop.contract.UnpackLog(event, "ClaimConditionsUpdated", log); err != nil {
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

// ParseClaimConditionsUpdated is a log parse operation binding the contract event 0xbf4016fceeaaa4ac5cf4be865b559ff85825ab4ca7aa7b661d16e2f544c03098.
//
// Solidity: event ClaimConditionsUpdated((uint256,uint256,uint256,uint256,bytes32,uint256,address,string)[] claimConditions, bool resetEligibility)
func (_NFTDrop *NFTDropFilterer) ParseClaimConditionsUpdated(log types.Log) (*NFTDropClaimConditionsUpdated, error) {
	event := new(NFTDropClaimConditionsUpdated)
	if err := _NFTDrop.contract.UnpackLog(event, "ClaimConditionsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDropContractURIUpdatedIterator is returned from FilterContractURIUpdated and is used to iterate over the raw logs and unpacked data for ContractURIUpdated events raised by the NFTDrop contract.
type NFTDropContractURIUpdatedIterator struct {
	Event *NFTDropContractURIUpdated // Event containing the contract specifics and raw log

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
func (it *NFTDropContractURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropContractURIUpdated)
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
		it.Event = new(NFTDropContractURIUpdated)
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
func (it *NFTDropContractURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropContractURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropContractURIUpdated represents a ContractURIUpdated event raised by the NFTDrop contract.
type NFTDropContractURIUpdated struct {
	PrevURI string
	NewURI  string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractURIUpdated is a free log retrieval operation binding the contract event 0xc9c7c3fe08b88b4df9d4d47ef47d2c43d55c025a0ba88ca442580ed9e7348a16.
//
// Solidity: event ContractURIUpdated(string prevURI, string newURI)
func (_NFTDrop *NFTDropFilterer) FilterContractURIUpdated(opts *bind.FilterOpts) (*NFTDropContractURIUpdatedIterator, error) {

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return &NFTDropContractURIUpdatedIterator{contract: _NFTDrop.contract, event: "ContractURIUpdated", logs: logs, sub: sub}, nil
}

// WatchContractURIUpdated is a free log subscription operation binding the contract event 0xc9c7c3fe08b88b4df9d4d47ef47d2c43d55c025a0ba88ca442580ed9e7348a16.
//
// Solidity: event ContractURIUpdated(string prevURI, string newURI)
func (_NFTDrop *NFTDropFilterer) WatchContractURIUpdated(opts *bind.WatchOpts, sink chan<- *NFTDropContractURIUpdated) (event.Subscription, error) {

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropContractURIUpdated)
				if err := _NFTDrop.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
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

// ParseContractURIUpdated is a log parse operation binding the contract event 0xc9c7c3fe08b88b4df9d4d47ef47d2c43d55c025a0ba88ca442580ed9e7348a16.
//
// Solidity: event ContractURIUpdated(string prevURI, string newURI)
func (_NFTDrop *NFTDropFilterer) ParseContractURIUpdated(log types.Log) (*NFTDropContractURIUpdated, error) {
	event := new(NFTDropContractURIUpdated)
	if err := _NFTDrop.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDropDefaultRoyaltyIterator is returned from FilterDefaultRoyalty and is used to iterate over the raw logs and unpacked data for DefaultRoyalty events raised by the NFTDrop contract.
type NFTDropDefaultRoyaltyIterator struct {
	Event *NFTDropDefaultRoyalty // Event containing the contract specifics and raw log

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
func (it *NFTDropDefaultRoyaltyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropDefaultRoyalty)
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
		it.Event = new(NFTDropDefaultRoyalty)
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
func (it *NFTDropDefaultRoyaltyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropDefaultRoyaltyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropDefaultRoyalty represents a DefaultRoyalty event raised by the NFTDrop contract.
type NFTDropDefaultRoyalty struct {
	NewRoyaltyRecipient common.Address
	NewRoyaltyBps       *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDefaultRoyalty is a free log retrieval operation binding the contract event 0x90d7ec04bcb8978719414f82e52e4cb651db41d0e6f8cea6118c2191e6183adb.
//
// Solidity: event DefaultRoyalty(address indexed newRoyaltyRecipient, uint256 newRoyaltyBps)
func (_NFTDrop *NFTDropFilterer) FilterDefaultRoyalty(opts *bind.FilterOpts, newRoyaltyRecipient []common.Address) (*NFTDropDefaultRoyaltyIterator, error) {

	var newRoyaltyRecipientRule []interface{}
	for _, newRoyaltyRecipientItem := range newRoyaltyRecipient {
		newRoyaltyRecipientRule = append(newRoyaltyRecipientRule, newRoyaltyRecipientItem)
	}

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "DefaultRoyalty", newRoyaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return &NFTDropDefaultRoyaltyIterator{contract: _NFTDrop.contract, event: "DefaultRoyalty", logs: logs, sub: sub}, nil
}

// WatchDefaultRoyalty is a free log subscription operation binding the contract event 0x90d7ec04bcb8978719414f82e52e4cb651db41d0e6f8cea6118c2191e6183adb.
//
// Solidity: event DefaultRoyalty(address indexed newRoyaltyRecipient, uint256 newRoyaltyBps)
func (_NFTDrop *NFTDropFilterer) WatchDefaultRoyalty(opts *bind.WatchOpts, sink chan<- *NFTDropDefaultRoyalty, newRoyaltyRecipient []common.Address) (event.Subscription, error) {

	var newRoyaltyRecipientRule []interface{}
	for _, newRoyaltyRecipientItem := range newRoyaltyRecipient {
		newRoyaltyRecipientRule = append(newRoyaltyRecipientRule, newRoyaltyRecipientItem)
	}

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "DefaultRoyalty", newRoyaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropDefaultRoyalty)
				if err := _NFTDrop.contract.UnpackLog(event, "DefaultRoyalty", log); err != nil {
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

// ParseDefaultRoyalty is a log parse operation binding the contract event 0x90d7ec04bcb8978719414f82e52e4cb651db41d0e6f8cea6118c2191e6183adb.
//
// Solidity: event DefaultRoyalty(address indexed newRoyaltyRecipient, uint256 newRoyaltyBps)
func (_NFTDrop *NFTDropFilterer) ParseDefaultRoyalty(log types.Log) (*NFTDropDefaultRoyalty, error) {
	event := new(NFTDropDefaultRoyalty)
	if err := _NFTDrop.contract.UnpackLog(event, "DefaultRoyalty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDropFlatPlatformFeeUpdatedIterator is returned from FilterFlatPlatformFeeUpdated and is used to iterate over the raw logs and unpacked data for FlatPlatformFeeUpdated events raised by the NFTDrop contract.
type NFTDropFlatPlatformFeeUpdatedIterator struct {
	Event *NFTDropFlatPlatformFeeUpdated // Event containing the contract specifics and raw log

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
func (it *NFTDropFlatPlatformFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropFlatPlatformFeeUpdated)
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
		it.Event = new(NFTDropFlatPlatformFeeUpdated)
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
func (it *NFTDropFlatPlatformFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropFlatPlatformFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropFlatPlatformFeeUpdated represents a FlatPlatformFeeUpdated event raised by the NFTDrop contract.
type NFTDropFlatPlatformFeeUpdated struct {
	PlatformFeeRecipient common.Address
	FlatFee              *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterFlatPlatformFeeUpdated is a free log retrieval operation binding the contract event 0xf8086cee80709bd44c82f89dbca54115ebd05e840a88ab81df9cf5be9754eb63.
//
// Solidity: event FlatPlatformFeeUpdated(address platformFeeRecipient, uint256 flatFee)
func (_NFTDrop *NFTDropFilterer) FilterFlatPlatformFeeUpdated(opts *bind.FilterOpts) (*NFTDropFlatPlatformFeeUpdatedIterator, error) {

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "FlatPlatformFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &NFTDropFlatPlatformFeeUpdatedIterator{contract: _NFTDrop.contract, event: "FlatPlatformFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchFlatPlatformFeeUpdated is a free log subscription operation binding the contract event 0xf8086cee80709bd44c82f89dbca54115ebd05e840a88ab81df9cf5be9754eb63.
//
// Solidity: event FlatPlatformFeeUpdated(address platformFeeRecipient, uint256 flatFee)
func (_NFTDrop *NFTDropFilterer) WatchFlatPlatformFeeUpdated(opts *bind.WatchOpts, sink chan<- *NFTDropFlatPlatformFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "FlatPlatformFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropFlatPlatformFeeUpdated)
				if err := _NFTDrop.contract.UnpackLog(event, "FlatPlatformFeeUpdated", log); err != nil {
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

// ParseFlatPlatformFeeUpdated is a log parse operation binding the contract event 0xf8086cee80709bd44c82f89dbca54115ebd05e840a88ab81df9cf5be9754eb63.
//
// Solidity: event FlatPlatformFeeUpdated(address platformFeeRecipient, uint256 flatFee)
func (_NFTDrop *NFTDropFilterer) ParseFlatPlatformFeeUpdated(log types.Log) (*NFTDropFlatPlatformFeeUpdated, error) {
	event := new(NFTDropFlatPlatformFeeUpdated)
	if err := _NFTDrop.contract.UnpackLog(event, "FlatPlatformFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDropInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NFTDrop contract.
type NFTDropInitializedIterator struct {
	Event *NFTDropInitialized // Event containing the contract specifics and raw log

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
func (it *NFTDropInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropInitialized)
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
		it.Event = new(NFTDropInitialized)
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
func (it *NFTDropInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropInitialized represents a Initialized event raised by the NFTDrop contract.
type NFTDropInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NFTDrop *NFTDropFilterer) FilterInitialized(opts *bind.FilterOpts) (*NFTDropInitializedIterator, error) {

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NFTDropInitializedIterator{contract: _NFTDrop.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NFTDrop *NFTDropFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NFTDropInitialized) (event.Subscription, error) {

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropInitialized)
				if err := _NFTDrop.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NFTDrop *NFTDropFilterer) ParseInitialized(log types.Log) (*NFTDropInitialized, error) {
	event := new(NFTDropInitialized)
	if err := _NFTDrop.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDropMaxTotalSupplyUpdatedIterator is returned from FilterMaxTotalSupplyUpdated and is used to iterate over the raw logs and unpacked data for MaxTotalSupplyUpdated events raised by the NFTDrop contract.
type NFTDropMaxTotalSupplyUpdatedIterator struct {
	Event *NFTDropMaxTotalSupplyUpdated // Event containing the contract specifics and raw log

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
func (it *NFTDropMaxTotalSupplyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropMaxTotalSupplyUpdated)
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
		it.Event = new(NFTDropMaxTotalSupplyUpdated)
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
func (it *NFTDropMaxTotalSupplyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropMaxTotalSupplyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropMaxTotalSupplyUpdated represents a MaxTotalSupplyUpdated event raised by the NFTDrop contract.
type NFTDropMaxTotalSupplyUpdated struct {
	MaxTotalSupply *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMaxTotalSupplyUpdated is a free log retrieval operation binding the contract event 0xf2672935fc79f5237559e2e2999dbe743bf65430894ac2b37666890e7c69e1af.
//
// Solidity: event MaxTotalSupplyUpdated(uint256 maxTotalSupply)
func (_NFTDrop *NFTDropFilterer) FilterMaxTotalSupplyUpdated(opts *bind.FilterOpts) (*NFTDropMaxTotalSupplyUpdatedIterator, error) {

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "MaxTotalSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return &NFTDropMaxTotalSupplyUpdatedIterator{contract: _NFTDrop.contract, event: "MaxTotalSupplyUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxTotalSupplyUpdated is a free log subscription operation binding the contract event 0xf2672935fc79f5237559e2e2999dbe743bf65430894ac2b37666890e7c69e1af.
//
// Solidity: event MaxTotalSupplyUpdated(uint256 maxTotalSupply)
func (_NFTDrop *NFTDropFilterer) WatchMaxTotalSupplyUpdated(opts *bind.WatchOpts, sink chan<- *NFTDropMaxTotalSupplyUpdated) (event.Subscription, error) {

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "MaxTotalSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropMaxTotalSupplyUpdated)
				if err := _NFTDrop.contract.UnpackLog(event, "MaxTotalSupplyUpdated", log); err != nil {
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

// ParseMaxTotalSupplyUpdated is a log parse operation binding the contract event 0xf2672935fc79f5237559e2e2999dbe743bf65430894ac2b37666890e7c69e1af.
//
// Solidity: event MaxTotalSupplyUpdated(uint256 maxTotalSupply)
func (_NFTDrop *NFTDropFilterer) ParseMaxTotalSupplyUpdated(log types.Log) (*NFTDropMaxTotalSupplyUpdated, error) {
	event := new(NFTDropMaxTotalSupplyUpdated)
	if err := _NFTDrop.contract.UnpackLog(event, "MaxTotalSupplyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDropMetadataFrozenIterator is returned from FilterMetadataFrozen and is used to iterate over the raw logs and unpacked data for MetadataFrozen events raised by the NFTDrop contract.
type NFTDropMetadataFrozenIterator struct {
	Event *NFTDropMetadataFrozen // Event containing the contract specifics and raw log

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
func (it *NFTDropMetadataFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropMetadataFrozen)
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
		it.Event = new(NFTDropMetadataFrozen)
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
func (it *NFTDropMetadataFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropMetadataFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropMetadataFrozen represents a MetadataFrozen event raised by the NFTDrop contract.
type NFTDropMetadataFrozen struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMetadataFrozen is a free log retrieval operation binding the contract event 0xeef043febddf4e1d1cf1f72ff1407b84e036e805aa0934418cb82095da8d7164.
//
// Solidity: event MetadataFrozen()
func (_NFTDrop *NFTDropFilterer) FilterMetadataFrozen(opts *bind.FilterOpts) (*NFTDropMetadataFrozenIterator, error) {

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "MetadataFrozen")
	if err != nil {
		return nil, err
	}
	return &NFTDropMetadataFrozenIterator{contract: _NFTDrop.contract, event: "MetadataFrozen", logs: logs, sub: sub}, nil
}

// WatchMetadataFrozen is a free log subscription operation binding the contract event 0xeef043febddf4e1d1cf1f72ff1407b84e036e805aa0934418cb82095da8d7164.
//
// Solidity: event MetadataFrozen()
func (_NFTDrop *NFTDropFilterer) WatchMetadataFrozen(opts *bind.WatchOpts, sink chan<- *NFTDropMetadataFrozen) (event.Subscription, error) {

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "MetadataFrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropMetadataFrozen)
				if err := _NFTDrop.contract.UnpackLog(event, "MetadataFrozen", log); err != nil {
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

// ParseMetadataFrozen is a log parse operation binding the contract event 0xeef043febddf4e1d1cf1f72ff1407b84e036e805aa0934418cb82095da8d7164.
//
// Solidity: event MetadataFrozen()
func (_NFTDrop *NFTDropFilterer) ParseMetadataFrozen(log types.Log) (*NFTDropMetadataFrozen, error) {
	event := new(NFTDropMetadataFrozen)
	if err := _NFTDrop.contract.UnpackLog(event, "MetadataFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDropOwnerUpdatedIterator is returned from FilterOwnerUpdated and is used to iterate over the raw logs and unpacked data for OwnerUpdated events raised by the NFTDrop contract.
type NFTDropOwnerUpdatedIterator struct {
	Event *NFTDropOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *NFTDropOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropOwnerUpdated)
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
		it.Event = new(NFTDropOwnerUpdated)
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
func (it *NFTDropOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropOwnerUpdated represents a OwnerUpdated event raised by the NFTDrop contract.
type NFTDropOwnerUpdated struct {
	PrevOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdated is a free log retrieval operation binding the contract event 0x8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d76.
//
// Solidity: event OwnerUpdated(address indexed prevOwner, address indexed newOwner)
func (_NFTDrop *NFTDropFilterer) FilterOwnerUpdated(opts *bind.FilterOpts, prevOwner []common.Address, newOwner []common.Address) (*NFTDropOwnerUpdatedIterator, error) {

	var prevOwnerRule []interface{}
	for _, prevOwnerItem := range prevOwner {
		prevOwnerRule = append(prevOwnerRule, prevOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "OwnerUpdated", prevOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NFTDropOwnerUpdatedIterator{contract: _NFTDrop.contract, event: "OwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdated is a free log subscription operation binding the contract event 0x8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d76.
//
// Solidity: event OwnerUpdated(address indexed prevOwner, address indexed newOwner)
func (_NFTDrop *NFTDropFilterer) WatchOwnerUpdated(opts *bind.WatchOpts, sink chan<- *NFTDropOwnerUpdated, prevOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var prevOwnerRule []interface{}
	for _, prevOwnerItem := range prevOwner {
		prevOwnerRule = append(prevOwnerRule, prevOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "OwnerUpdated", prevOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropOwnerUpdated)
				if err := _NFTDrop.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
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

// ParseOwnerUpdated is a log parse operation binding the contract event 0x8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d76.
//
// Solidity: event OwnerUpdated(address indexed prevOwner, address indexed newOwner)
func (_NFTDrop *NFTDropFilterer) ParseOwnerUpdated(log types.Log) (*NFTDropOwnerUpdated, error) {
	event := new(NFTDropOwnerUpdated)
	if err := _NFTDrop.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDropPlatformFeeInfoUpdatedIterator is returned from FilterPlatformFeeInfoUpdated and is used to iterate over the raw logs and unpacked data for PlatformFeeInfoUpdated events raised by the NFTDrop contract.
type NFTDropPlatformFeeInfoUpdatedIterator struct {
	Event *NFTDropPlatformFeeInfoUpdated // Event containing the contract specifics and raw log

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
func (it *NFTDropPlatformFeeInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropPlatformFeeInfoUpdated)
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
		it.Event = new(NFTDropPlatformFeeInfoUpdated)
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
func (it *NFTDropPlatformFeeInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropPlatformFeeInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropPlatformFeeInfoUpdated represents a PlatformFeeInfoUpdated event raised by the NFTDrop contract.
type NFTDropPlatformFeeInfoUpdated struct {
	PlatformFeeRecipient common.Address
	PlatformFeeBps       *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterPlatformFeeInfoUpdated is a free log retrieval operation binding the contract event 0xe2497bd806ec41a6e0dd992c29a72efc0ef8fec9092d1978fd4a1e00b2f18304.
//
// Solidity: event PlatformFeeInfoUpdated(address indexed platformFeeRecipient, uint256 platformFeeBps)
func (_NFTDrop *NFTDropFilterer) FilterPlatformFeeInfoUpdated(opts *bind.FilterOpts, platformFeeRecipient []common.Address) (*NFTDropPlatformFeeInfoUpdatedIterator, error) {

	var platformFeeRecipientRule []interface{}
	for _, platformFeeRecipientItem := range platformFeeRecipient {
		platformFeeRecipientRule = append(platformFeeRecipientRule, platformFeeRecipientItem)
	}

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "PlatformFeeInfoUpdated", platformFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &NFTDropPlatformFeeInfoUpdatedIterator{contract: _NFTDrop.contract, event: "PlatformFeeInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchPlatformFeeInfoUpdated is a free log subscription operation binding the contract event 0xe2497bd806ec41a6e0dd992c29a72efc0ef8fec9092d1978fd4a1e00b2f18304.
//
// Solidity: event PlatformFeeInfoUpdated(address indexed platformFeeRecipient, uint256 platformFeeBps)
func (_NFTDrop *NFTDropFilterer) WatchPlatformFeeInfoUpdated(opts *bind.WatchOpts, sink chan<- *NFTDropPlatformFeeInfoUpdated, platformFeeRecipient []common.Address) (event.Subscription, error) {

	var platformFeeRecipientRule []interface{}
	for _, platformFeeRecipientItem := range platformFeeRecipient {
		platformFeeRecipientRule = append(platformFeeRecipientRule, platformFeeRecipientItem)
	}

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "PlatformFeeInfoUpdated", platformFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropPlatformFeeInfoUpdated)
				if err := _NFTDrop.contract.UnpackLog(event, "PlatformFeeInfoUpdated", log); err != nil {
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

// ParsePlatformFeeInfoUpdated is a log parse operation binding the contract event 0xe2497bd806ec41a6e0dd992c29a72efc0ef8fec9092d1978fd4a1e00b2f18304.
//
// Solidity: event PlatformFeeInfoUpdated(address indexed platformFeeRecipient, uint256 platformFeeBps)
func (_NFTDrop *NFTDropFilterer) ParsePlatformFeeInfoUpdated(log types.Log) (*NFTDropPlatformFeeInfoUpdated, error) {
	event := new(NFTDropPlatformFeeInfoUpdated)
	if err := _NFTDrop.contract.UnpackLog(event, "PlatformFeeInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDropPlatformFeeTypeUpdatedIterator is returned from FilterPlatformFeeTypeUpdated and is used to iterate over the raw logs and unpacked data for PlatformFeeTypeUpdated events raised by the NFTDrop contract.
type NFTDropPlatformFeeTypeUpdatedIterator struct {
	Event *NFTDropPlatformFeeTypeUpdated // Event containing the contract specifics and raw log

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
func (it *NFTDropPlatformFeeTypeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropPlatformFeeTypeUpdated)
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
		it.Event = new(NFTDropPlatformFeeTypeUpdated)
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
func (it *NFTDropPlatformFeeTypeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropPlatformFeeTypeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropPlatformFeeTypeUpdated represents a PlatformFeeTypeUpdated event raised by the NFTDrop contract.
type NFTDropPlatformFeeTypeUpdated struct {
	FeeType uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPlatformFeeTypeUpdated is a free log retrieval operation binding the contract event 0xd246da9440709ce0dd3f4fd669abc85ada012ab9774b8ecdcc5059ba1486b9c1.
//
// Solidity: event PlatformFeeTypeUpdated(uint8 feeType)
func (_NFTDrop *NFTDropFilterer) FilterPlatformFeeTypeUpdated(opts *bind.FilterOpts) (*NFTDropPlatformFeeTypeUpdatedIterator, error) {

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "PlatformFeeTypeUpdated")
	if err != nil {
		return nil, err
	}
	return &NFTDropPlatformFeeTypeUpdatedIterator{contract: _NFTDrop.contract, event: "PlatformFeeTypeUpdated", logs: logs, sub: sub}, nil
}

// WatchPlatformFeeTypeUpdated is a free log subscription operation binding the contract event 0xd246da9440709ce0dd3f4fd669abc85ada012ab9774b8ecdcc5059ba1486b9c1.
//
// Solidity: event PlatformFeeTypeUpdated(uint8 feeType)
func (_NFTDrop *NFTDropFilterer) WatchPlatformFeeTypeUpdated(opts *bind.WatchOpts, sink chan<- *NFTDropPlatformFeeTypeUpdated) (event.Subscription, error) {

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "PlatformFeeTypeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropPlatformFeeTypeUpdated)
				if err := _NFTDrop.contract.UnpackLog(event, "PlatformFeeTypeUpdated", log); err != nil {
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

// ParsePlatformFeeTypeUpdated is a log parse operation binding the contract event 0xd246da9440709ce0dd3f4fd669abc85ada012ab9774b8ecdcc5059ba1486b9c1.
//
// Solidity: event PlatformFeeTypeUpdated(uint8 feeType)
func (_NFTDrop *NFTDropFilterer) ParsePlatformFeeTypeUpdated(log types.Log) (*NFTDropPlatformFeeTypeUpdated, error) {
	event := new(NFTDropPlatformFeeTypeUpdated)
	if err := _NFTDrop.contract.UnpackLog(event, "PlatformFeeTypeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDropPrimarySaleRecipientUpdatedIterator is returned from FilterPrimarySaleRecipientUpdated and is used to iterate over the raw logs and unpacked data for PrimarySaleRecipientUpdated events raised by the NFTDrop contract.
type NFTDropPrimarySaleRecipientUpdatedIterator struct {
	Event *NFTDropPrimarySaleRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *NFTDropPrimarySaleRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropPrimarySaleRecipientUpdated)
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
		it.Event = new(NFTDropPrimarySaleRecipientUpdated)
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
func (it *NFTDropPrimarySaleRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropPrimarySaleRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropPrimarySaleRecipientUpdated represents a PrimarySaleRecipientUpdated event raised by the NFTDrop contract.
type NFTDropPrimarySaleRecipientUpdated struct {
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPrimarySaleRecipientUpdated is a free log retrieval operation binding the contract event 0x299d17e95023f496e0ffc4909cff1a61f74bb5eb18de6f900f4155bfa1b3b333.
//
// Solidity: event PrimarySaleRecipientUpdated(address indexed recipient)
func (_NFTDrop *NFTDropFilterer) FilterPrimarySaleRecipientUpdated(opts *bind.FilterOpts, recipient []common.Address) (*NFTDropPrimarySaleRecipientUpdatedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "PrimarySaleRecipientUpdated", recipientRule)
	if err != nil {
		return nil, err
	}
	return &NFTDropPrimarySaleRecipientUpdatedIterator{contract: _NFTDrop.contract, event: "PrimarySaleRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchPrimarySaleRecipientUpdated is a free log subscription operation binding the contract event 0x299d17e95023f496e0ffc4909cff1a61f74bb5eb18de6f900f4155bfa1b3b333.
//
// Solidity: event PrimarySaleRecipientUpdated(address indexed recipient)
func (_NFTDrop *NFTDropFilterer) WatchPrimarySaleRecipientUpdated(opts *bind.WatchOpts, sink chan<- *NFTDropPrimarySaleRecipientUpdated, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "PrimarySaleRecipientUpdated", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropPrimarySaleRecipientUpdated)
				if err := _NFTDrop.contract.UnpackLog(event, "PrimarySaleRecipientUpdated", log); err != nil {
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

// ParsePrimarySaleRecipientUpdated is a log parse operation binding the contract event 0x299d17e95023f496e0ffc4909cff1a61f74bb5eb18de6f900f4155bfa1b3b333.
//
// Solidity: event PrimarySaleRecipientUpdated(address indexed recipient)
func (_NFTDrop *NFTDropFilterer) ParsePrimarySaleRecipientUpdated(log types.Log) (*NFTDropPrimarySaleRecipientUpdated, error) {
	event := new(NFTDropPrimarySaleRecipientUpdated)
	if err := _NFTDrop.contract.UnpackLog(event, "PrimarySaleRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDropRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the NFTDrop contract.
type NFTDropRoleAdminChangedIterator struct {
	Event *NFTDropRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *NFTDropRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropRoleAdminChanged)
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
		it.Event = new(NFTDropRoleAdminChanged)
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
func (it *NFTDropRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropRoleAdminChanged represents a RoleAdminChanged event raised by the NFTDrop contract.
type NFTDropRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NFTDrop *NFTDropFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*NFTDropRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &NFTDropRoleAdminChangedIterator{contract: _NFTDrop.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NFTDrop *NFTDropFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *NFTDropRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropRoleAdminChanged)
				if err := _NFTDrop.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NFTDrop *NFTDropFilterer) ParseRoleAdminChanged(log types.Log) (*NFTDropRoleAdminChanged, error) {
	event := new(NFTDropRoleAdminChanged)
	if err := _NFTDrop.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDropRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the NFTDrop contract.
type NFTDropRoleGrantedIterator struct {
	Event *NFTDropRoleGranted // Event containing the contract specifics and raw log

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
func (it *NFTDropRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropRoleGranted)
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
		it.Event = new(NFTDropRoleGranted)
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
func (it *NFTDropRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropRoleGranted represents a RoleGranted event raised by the NFTDrop contract.
type NFTDropRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NFTDrop *NFTDropFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NFTDropRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NFTDropRoleGrantedIterator{contract: _NFTDrop.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NFTDrop *NFTDropFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *NFTDropRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropRoleGranted)
				if err := _NFTDrop.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NFTDrop *NFTDropFilterer) ParseRoleGranted(log types.Log) (*NFTDropRoleGranted, error) {
	event := new(NFTDropRoleGranted)
	if err := _NFTDrop.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDropRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the NFTDrop contract.
type NFTDropRoleRevokedIterator struct {
	Event *NFTDropRoleRevoked // Event containing the contract specifics and raw log

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
func (it *NFTDropRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropRoleRevoked)
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
		it.Event = new(NFTDropRoleRevoked)
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
func (it *NFTDropRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropRoleRevoked represents a RoleRevoked event raised by the NFTDrop contract.
type NFTDropRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NFTDrop *NFTDropFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NFTDropRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NFTDropRoleRevokedIterator{contract: _NFTDrop.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NFTDrop *NFTDropFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *NFTDropRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropRoleRevoked)
				if err := _NFTDrop.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NFTDrop *NFTDropFilterer) ParseRoleRevoked(log types.Log) (*NFTDropRoleRevoked, error) {
	event := new(NFTDropRoleRevoked)
	if err := _NFTDrop.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDropRoyaltyForTokenIterator is returned from FilterRoyaltyForToken and is used to iterate over the raw logs and unpacked data for RoyaltyForToken events raised by the NFTDrop contract.
type NFTDropRoyaltyForTokenIterator struct {
	Event *NFTDropRoyaltyForToken // Event containing the contract specifics and raw log

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
func (it *NFTDropRoyaltyForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropRoyaltyForToken)
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
		it.Event = new(NFTDropRoyaltyForToken)
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
func (it *NFTDropRoyaltyForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropRoyaltyForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropRoyaltyForToken represents a RoyaltyForToken event raised by the NFTDrop contract.
type NFTDropRoyaltyForToken struct {
	TokenId          *big.Int
	RoyaltyRecipient common.Address
	RoyaltyBps       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyForToken is a free log retrieval operation binding the contract event 0x7365cf4122f072a3365c20d54eff9b38d73c096c28e1892ec8f5b0e403a0f12d.
//
// Solidity: event RoyaltyForToken(uint256 indexed tokenId, address indexed royaltyRecipient, uint256 royaltyBps)
func (_NFTDrop *NFTDropFilterer) FilterRoyaltyForToken(opts *bind.FilterOpts, tokenId []*big.Int, royaltyRecipient []common.Address) (*NFTDropRoyaltyForTokenIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var royaltyRecipientRule []interface{}
	for _, royaltyRecipientItem := range royaltyRecipient {
		royaltyRecipientRule = append(royaltyRecipientRule, royaltyRecipientItem)
	}

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "RoyaltyForToken", tokenIdRule, royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return &NFTDropRoyaltyForTokenIterator{contract: _NFTDrop.contract, event: "RoyaltyForToken", logs: logs, sub: sub}, nil
}

// WatchRoyaltyForToken is a free log subscription operation binding the contract event 0x7365cf4122f072a3365c20d54eff9b38d73c096c28e1892ec8f5b0e403a0f12d.
//
// Solidity: event RoyaltyForToken(uint256 indexed tokenId, address indexed royaltyRecipient, uint256 royaltyBps)
func (_NFTDrop *NFTDropFilterer) WatchRoyaltyForToken(opts *bind.WatchOpts, sink chan<- *NFTDropRoyaltyForToken, tokenId []*big.Int, royaltyRecipient []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var royaltyRecipientRule []interface{}
	for _, royaltyRecipientItem := range royaltyRecipient {
		royaltyRecipientRule = append(royaltyRecipientRule, royaltyRecipientItem)
	}

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "RoyaltyForToken", tokenIdRule, royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropRoyaltyForToken)
				if err := _NFTDrop.contract.UnpackLog(event, "RoyaltyForToken", log); err != nil {
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

// ParseRoyaltyForToken is a log parse operation binding the contract event 0x7365cf4122f072a3365c20d54eff9b38d73c096c28e1892ec8f5b0e403a0f12d.
//
// Solidity: event RoyaltyForToken(uint256 indexed tokenId, address indexed royaltyRecipient, uint256 royaltyBps)
func (_NFTDrop *NFTDropFilterer) ParseRoyaltyForToken(log types.Log) (*NFTDropRoyaltyForToken, error) {
	event := new(NFTDropRoyaltyForToken)
	if err := _NFTDrop.contract.UnpackLog(event, "RoyaltyForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDropTokenURIRevealedIterator is returned from FilterTokenURIRevealed and is used to iterate over the raw logs and unpacked data for TokenURIRevealed events raised by the NFTDrop contract.
type NFTDropTokenURIRevealedIterator struct {
	Event *NFTDropTokenURIRevealed // Event containing the contract specifics and raw log

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
func (it *NFTDropTokenURIRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropTokenURIRevealed)
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
		it.Event = new(NFTDropTokenURIRevealed)
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
func (it *NFTDropTokenURIRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropTokenURIRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropTokenURIRevealed represents a TokenURIRevealed event raised by the NFTDrop contract.
type NFTDropTokenURIRevealed struct {
	Index       *big.Int
	RevealedURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenURIRevealed is a free log retrieval operation binding the contract event 0x6df1d8db2a036436ffe0b2d1833f2c5f1e624818dfce2578c0faa4b83ef9998d.
//
// Solidity: event TokenURIRevealed(uint256 indexed index, string revealedURI)
func (_NFTDrop *NFTDropFilterer) FilterTokenURIRevealed(opts *bind.FilterOpts, index []*big.Int) (*NFTDropTokenURIRevealedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "TokenURIRevealed", indexRule)
	if err != nil {
		return nil, err
	}
	return &NFTDropTokenURIRevealedIterator{contract: _NFTDrop.contract, event: "TokenURIRevealed", logs: logs, sub: sub}, nil
}

// WatchTokenURIRevealed is a free log subscription operation binding the contract event 0x6df1d8db2a036436ffe0b2d1833f2c5f1e624818dfce2578c0faa4b83ef9998d.
//
// Solidity: event TokenURIRevealed(uint256 indexed index, string revealedURI)
func (_NFTDrop *NFTDropFilterer) WatchTokenURIRevealed(opts *bind.WatchOpts, sink chan<- *NFTDropTokenURIRevealed, index []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "TokenURIRevealed", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropTokenURIRevealed)
				if err := _NFTDrop.contract.UnpackLog(event, "TokenURIRevealed", log); err != nil {
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

// ParseTokenURIRevealed is a log parse operation binding the contract event 0x6df1d8db2a036436ffe0b2d1833f2c5f1e624818dfce2578c0faa4b83ef9998d.
//
// Solidity: event TokenURIRevealed(uint256 indexed index, string revealedURI)
func (_NFTDrop *NFTDropFilterer) ParseTokenURIRevealed(log types.Log) (*NFTDropTokenURIRevealed, error) {
	event := new(NFTDropTokenURIRevealed)
	if err := _NFTDrop.contract.UnpackLog(event, "TokenURIRevealed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDropTokensClaimedIterator is returned from FilterTokensClaimed and is used to iterate over the raw logs and unpacked data for TokensClaimed events raised by the NFTDrop contract.
type NFTDropTokensClaimedIterator struct {
	Event *NFTDropTokensClaimed // Event containing the contract specifics and raw log

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
func (it *NFTDropTokensClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropTokensClaimed)
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
		it.Event = new(NFTDropTokensClaimed)
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
func (it *NFTDropTokensClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropTokensClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropTokensClaimed represents a TokensClaimed event raised by the NFTDrop contract.
type NFTDropTokensClaimed struct {
	ClaimConditionIndex *big.Int
	Claimer             common.Address
	Receiver            common.Address
	StartTokenId        *big.Int
	QuantityClaimed     *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensClaimed is a free log retrieval operation binding the contract event 0xfa76a4010d9533e3e964f2930a65fb6042a12fa6ff5b08281837a10b0be7321e.
//
// Solidity: event TokensClaimed(uint256 indexed claimConditionIndex, address indexed claimer, address indexed receiver, uint256 startTokenId, uint256 quantityClaimed)
func (_NFTDrop *NFTDropFilterer) FilterTokensClaimed(opts *bind.FilterOpts, claimConditionIndex []*big.Int, claimer []common.Address, receiver []common.Address) (*NFTDropTokensClaimedIterator, error) {

	var claimConditionIndexRule []interface{}
	for _, claimConditionIndexItem := range claimConditionIndex {
		claimConditionIndexRule = append(claimConditionIndexRule, claimConditionIndexItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "TokensClaimed", claimConditionIndexRule, claimerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &NFTDropTokensClaimedIterator{contract: _NFTDrop.contract, event: "TokensClaimed", logs: logs, sub: sub}, nil
}

// WatchTokensClaimed is a free log subscription operation binding the contract event 0xfa76a4010d9533e3e964f2930a65fb6042a12fa6ff5b08281837a10b0be7321e.
//
// Solidity: event TokensClaimed(uint256 indexed claimConditionIndex, address indexed claimer, address indexed receiver, uint256 startTokenId, uint256 quantityClaimed)
func (_NFTDrop *NFTDropFilterer) WatchTokensClaimed(opts *bind.WatchOpts, sink chan<- *NFTDropTokensClaimed, claimConditionIndex []*big.Int, claimer []common.Address, receiver []common.Address) (event.Subscription, error) {

	var claimConditionIndexRule []interface{}
	for _, claimConditionIndexItem := range claimConditionIndex {
		claimConditionIndexRule = append(claimConditionIndexRule, claimConditionIndexItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "TokensClaimed", claimConditionIndexRule, claimerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropTokensClaimed)
				if err := _NFTDrop.contract.UnpackLog(event, "TokensClaimed", log); err != nil {
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

// ParseTokensClaimed is a log parse operation binding the contract event 0xfa76a4010d9533e3e964f2930a65fb6042a12fa6ff5b08281837a10b0be7321e.
//
// Solidity: event TokensClaimed(uint256 indexed claimConditionIndex, address indexed claimer, address indexed receiver, uint256 startTokenId, uint256 quantityClaimed)
func (_NFTDrop *NFTDropFilterer) ParseTokensClaimed(log types.Log) (*NFTDropTokensClaimed, error) {
	event := new(NFTDropTokensClaimed)
	if err := _NFTDrop.contract.UnpackLog(event, "TokensClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDropTokensLazyMintedIterator is returned from FilterTokensLazyMinted and is used to iterate over the raw logs and unpacked data for TokensLazyMinted events raised by the NFTDrop contract.
type NFTDropTokensLazyMintedIterator struct {
	Event *NFTDropTokensLazyMinted // Event containing the contract specifics and raw log

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
func (it *NFTDropTokensLazyMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropTokensLazyMinted)
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
		it.Event = new(NFTDropTokensLazyMinted)
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
func (it *NFTDropTokensLazyMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropTokensLazyMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropTokensLazyMinted represents a TokensLazyMinted event raised by the NFTDrop contract.
type NFTDropTokensLazyMinted struct {
	StartTokenId     *big.Int
	EndTokenId       *big.Int
	BaseURI          string
	EncryptedBaseURI []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokensLazyMinted is a free log retrieval operation binding the contract event 0x2a0365091ef1a40953c670dce28177e37520648a6fdc91506bffac0ab045570d.
//
// Solidity: event TokensLazyMinted(uint256 indexed startTokenId, uint256 endTokenId, string baseURI, bytes encryptedBaseURI)
func (_NFTDrop *NFTDropFilterer) FilterTokensLazyMinted(opts *bind.FilterOpts, startTokenId []*big.Int) (*NFTDropTokensLazyMintedIterator, error) {

	var startTokenIdRule []interface{}
	for _, startTokenIdItem := range startTokenId {
		startTokenIdRule = append(startTokenIdRule, startTokenIdItem)
	}

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "TokensLazyMinted", startTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTDropTokensLazyMintedIterator{contract: _NFTDrop.contract, event: "TokensLazyMinted", logs: logs, sub: sub}, nil
}

// WatchTokensLazyMinted is a free log subscription operation binding the contract event 0x2a0365091ef1a40953c670dce28177e37520648a6fdc91506bffac0ab045570d.
//
// Solidity: event TokensLazyMinted(uint256 indexed startTokenId, uint256 endTokenId, string baseURI, bytes encryptedBaseURI)
func (_NFTDrop *NFTDropFilterer) WatchTokensLazyMinted(opts *bind.WatchOpts, sink chan<- *NFTDropTokensLazyMinted, startTokenId []*big.Int) (event.Subscription, error) {

	var startTokenIdRule []interface{}
	for _, startTokenIdItem := range startTokenId {
		startTokenIdRule = append(startTokenIdRule, startTokenIdItem)
	}

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "TokensLazyMinted", startTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropTokensLazyMinted)
				if err := _NFTDrop.contract.UnpackLog(event, "TokensLazyMinted", log); err != nil {
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

// ParseTokensLazyMinted is a log parse operation binding the contract event 0x2a0365091ef1a40953c670dce28177e37520648a6fdc91506bffac0ab045570d.
//
// Solidity: event TokensLazyMinted(uint256 indexed startTokenId, uint256 endTokenId, string baseURI, bytes encryptedBaseURI)
func (_NFTDrop *NFTDropFilterer) ParseTokensLazyMinted(log types.Log) (*NFTDropTokensLazyMinted, error) {
	event := new(NFTDropTokensLazyMinted)
	if err := _NFTDrop.contract.UnpackLog(event, "TokensLazyMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDropTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NFTDrop contract.
type NFTDropTransferIterator struct {
	Event *NFTDropTransfer // Event containing the contract specifics and raw log

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
func (it *NFTDropTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDropTransfer)
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
		it.Event = new(NFTDropTransfer)
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
func (it *NFTDropTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDropTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDropTransfer represents a Transfer event raised by the NFTDrop contract.
type NFTDropTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NFTDrop *NFTDropFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*NFTDropTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTDrop.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTDropTransferIterator{contract: _NFTDrop.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NFTDrop *NFTDropFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NFTDropTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTDrop.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDropTransfer)
				if err := _NFTDrop.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NFTDrop *NFTDropFilterer) ParseTransfer(log types.Log) (*NFTDropTransfer, error) {
	event := new(NFTDropTransfer)
	if err := _NFTDrop.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
