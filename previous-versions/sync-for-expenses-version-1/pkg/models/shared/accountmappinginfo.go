// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AccountType - Type of the account.
type AccountType string

const (
	AccountTypeAsset     AccountType = "Asset"
	AccountTypeLiability AccountType = "Liability"
	AccountTypeIncome    AccountType = "Income"
	AccountTypeExpense   AccountType = "Expense"
	AccountTypeEquity    AccountType = "Equity"
)

func (e AccountType) ToPointer() *AccountType {
	return &e
}

func (e *AccountType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Asset":
		fallthrough
	case "Liability":
		fallthrough
	case "Income":
		fallthrough
	case "Expense":
		fallthrough
	case "Equity":
		*e = AccountType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountType: %v", v)
	}
}

type ValidTransactionTypes string

const (
	ValidTransactionTypesPayment       ValidTransactionTypes = "Payment"
	ValidTransactionTypesRefund        ValidTransactionTypes = "Refund"
	ValidTransactionTypesReward        ValidTransactionTypes = "Reward"
	ValidTransactionTypesChargeback    ValidTransactionTypes = "Chargeback"
	ValidTransactionTypesTransferIn    ValidTransactionTypes = "TransferIn"
	ValidTransactionTypesTransferOut   ValidTransactionTypes = "TransferOut"
	ValidTransactionTypesAdjustmentIn  ValidTransactionTypes = "AdjustmentIn"
	ValidTransactionTypesAdjustmentOut ValidTransactionTypes = "AdjustmentOut"
)

func (e ValidTransactionTypes) ToPointer() *ValidTransactionTypes {
	return &e
}

func (e *ValidTransactionTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Payment":
		fallthrough
	case "Refund":
		fallthrough
	case "Reward":
		fallthrough
	case "Chargeback":
		fallthrough
	case "TransferIn":
		fallthrough
	case "TransferOut":
		fallthrough
	case "AdjustmentIn":
		fallthrough
	case "AdjustmentOut":
		*e = ValidTransactionTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ValidTransactionTypes: %v", v)
	}
}

type AccountMappingInfo struct {
	// Type of the account.
	AccountType *AccountType `json:"accountType,omitempty"`
	// Currency of the account.
	Currency *string `json:"currency,omitempty"`
	// Unique identifier of account.
	ID *string `json:"id,omitempty"`
	// Name of the account as it appears in the companies accounting software.
	Name *string `json:"name,omitempty"`
	// Supported transaction types for the account.
	ValidTransactionTypes []ValidTransactionTypes `json:"validTransactionTypes,omitempty"`
}

func (o *AccountMappingInfo) GetAccountType() *AccountType {
	if o == nil {
		return nil
	}
	return o.AccountType
}

func (o *AccountMappingInfo) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AccountMappingInfo) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountMappingInfo) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AccountMappingInfo) GetValidTransactionTypes() []ValidTransactionTypes {
	if o == nil {
		return nil
	}
	return o.ValidTransactionTypes
}
