// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/codatio/client-sdk-go/lending/v7/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// SourceAccountStatus - Status of the source account.
type SourceAccountStatus string

const (
	SourceAccountStatusPending      SourceAccountStatus = "pending"
	SourceAccountStatusConnected    SourceAccountStatus = "connected"
	SourceAccountStatusConnecting   SourceAccountStatus = "connecting"
	SourceAccountStatusDisconnected SourceAccountStatus = "disconnected"
	SourceAccountStatusUnknown      SourceAccountStatus = "unknown"
)

func (e SourceAccountStatus) ToPointer() *SourceAccountStatus {
	return &e
}
func (e *SourceAccountStatus) UnmarshalJSON(data []byte) error {
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
		*e = SourceAccountStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAccountStatus: %v", v)
	}
}

// SourceAccount - The target bank account in a supported accounting software for ingestion into a bank feed.
type SourceAccount struct {
	// The bank account name.
	AccountName *string `json:"accountName,omitempty"`
	// The account number.
	AccountNumber *string `json:"accountNumber,omitempty"`
	// The type of bank account e.g. Credit.
	AccountType *string `json:"accountType,omitempty"`
	// The latest balance for the bank account.
	Balance *decimal.Big `decimal:"number" json:"balance,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
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
	// The sort code.
	SortCode *string `json:"sortCode,omitempty"`
	// Status of the source account.
	Status *SourceAccountStatus `json:"status,omitempty"`
}

func (s SourceAccount) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAccount) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceAccount) GetAccountName() *string {
	if o == nil {
		return nil
	}
	return o.AccountName
}

func (o *SourceAccount) GetAccountNumber() *string {
	if o == nil {
		return nil
	}
	return o.AccountNumber
}

func (o *SourceAccount) GetAccountType() *string {
	if o == nil {
		return nil
	}
	return o.AccountType
}

func (o *SourceAccount) GetBalance() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Balance
}

func (o *SourceAccount) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *SourceAccount) GetFeedStartDate() *string {
	if o == nil {
		return nil
	}
	return o.FeedStartDate
}

func (o *SourceAccount) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *SourceAccount) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *SourceAccount) GetSortCode() *string {
	if o == nil {
		return nil
	}
	return o.SortCode
}

func (o *SourceAccount) GetStatus() *SourceAccountStatus {
	if o == nil {
		return nil
	}
	return o.Status
}
