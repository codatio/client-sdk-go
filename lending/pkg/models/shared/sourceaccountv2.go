// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// SourceAccountV2AccountType - The type of bank account e.g. checking, savings, loan, creditCard, prepaidCard.
type SourceAccountV2AccountType string

const (
	SourceAccountV2AccountTypeChecking    SourceAccountV2AccountType = "checking"
	SourceAccountV2AccountTypeSavings     SourceAccountV2AccountType = "savings"
	SourceAccountV2AccountTypeLoan        SourceAccountV2AccountType = "loan"
	SourceAccountV2AccountTypeCreditCard  SourceAccountV2AccountType = "creditCard"
	SourceAccountV2AccountTypePrepaidCard SourceAccountV2AccountType = "prepaidCard"
)

func (e SourceAccountV2AccountType) ToPointer() *SourceAccountV2AccountType {
	return &e
}
func (e *SourceAccountV2AccountType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "checking":
		fallthrough
	case "savings":
		fallthrough
	case "loan":
		fallthrough
	case "creditCard":
		fallthrough
	case "prepaidCard":
		*e = SourceAccountV2AccountType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAccountV2AccountType: %v", v)
	}
}

// SourceAccountV2Status - Status of the source account.
type SourceAccountV2Status string

const (
	SourceAccountV2StatusPending      SourceAccountV2Status = "pending"
	SourceAccountV2StatusConnected    SourceAccountV2Status = "connected"
	SourceAccountV2StatusConnecting   SourceAccountV2Status = "connecting"
	SourceAccountV2StatusDisconnected SourceAccountV2Status = "disconnected"
	SourceAccountV2StatusUnknown      SourceAccountV2Status = "unknown"
)

func (e SourceAccountV2Status) ToPointer() *SourceAccountV2Status {
	return &e
}
func (e *SourceAccountV2Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pending":
		fallthrough
	case "connected":
		fallthrough
	case "connecting":
		fallthrough
	case "disconnected":
		fallthrough
	case "unknown":
		*e = SourceAccountV2Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAccountV2Status: %v", v)
	}
}

// SourceAccountV2 - The target bank account in a supported accounting software for ingestion into a bank feed.
type SourceAccountV2 struct {
	AccountInfo *AccountInfo `json:"accountInfo,omitempty"`
	// The bank account name.
	AccountName string `json:"accountName"`
	// The account number.
	AccountNumber string `json:"accountNumber"`
	// The type of bank account e.g. checking, savings, loan, creditCard, prepaidCard.
	AccountType SourceAccountV2AccountType `json:"accountType"`
	// The latest balance for the bank account.
	Balance *decimal.Big `decimal:"number" json:"balance"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency string `json:"currency"`
	// In Codat's data model, dates are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date fields are formatted as strings; for example:
	// ```
	// 2020-10-08
	// ```
	FeedStartDate *string `json:"feedStartDate,omitempty"`
	// Unique ID for the bank account.
	ID string `json:"id"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// Routing information for the bank. This does not include account number.
	RoutingInfo *RoutingInfo `json:"routingInfo,omitempty"`
	// The sort code.
	SortCode *string `json:"sortCode,omitempty"`
	// Status of the source account.
	Status *SourceAccountV2Status `json:"status,omitempty"`
}

func (s SourceAccountV2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAccountV2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceAccountV2) GetAccountInfo() *AccountInfo {
	if o == nil {
		return nil
	}
	return o.AccountInfo
}

func (o *SourceAccountV2) GetAccountName() string {
	if o == nil {
		return ""
	}
	return o.AccountName
}

func (o *SourceAccountV2) GetAccountNumber() string {
	if o == nil {
		return ""
	}
	return o.AccountNumber
}

func (o *SourceAccountV2) GetAccountType() SourceAccountV2AccountType {
	if o == nil {
		return SourceAccountV2AccountType("")
	}
	return o.AccountType
}

func (o *SourceAccountV2) GetBalance() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.Balance
}

func (o *SourceAccountV2) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *SourceAccountV2) GetFeedStartDate() *string {
	if o == nil {
		return nil
	}
	return o.FeedStartDate
}

func (o *SourceAccountV2) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *SourceAccountV2) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *SourceAccountV2) GetRoutingInfo() *RoutingInfo {
	if o == nil {
		return nil
	}
	return o.RoutingInfo
}

func (o *SourceAccountV2) GetSortCode() *string {
	if o == nil {
		return nil
	}
	return o.SortCode
}

func (o *SourceAccountV2) GetStatus() *SourceAccountV2Status {
	if o == nil {
		return nil
	}
	return o.Status
}
