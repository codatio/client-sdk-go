// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ListBankingTransactionsRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Field to order results by. [Read more](https://docs.codat.io/using-the-api/ordering-results).
	OrderBy *string `queryParam:"style=form,explode=true,name=orderBy"`
	// Page number. [Read more](https://docs.codat.io/using-the-api/paging).
	Page int `queryParam:"style=form,explode=true,name=page"`
	// Number of records to return in a page. [Read more](https://docs.codat.io/using-the-api/paging).
	PageSize *int `queryParam:"style=form,explode=true,name=pageSize"`
	// Codat query string. [Read more](https://docs.codat.io/using-the-api/querying).
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

// ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum - Code to identify the underlying transaction.
type ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum string

const (
	ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumUnknown       ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Unknown"
	ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumFee           ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Fee"
	ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumPayment       ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Payment"
	ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumCash          ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Cash"
	ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumTransfer      ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Transfer"
	ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumInterest      ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Interest"
	ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumCashback      ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Cashback"
	ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumCheque        ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Cheque"
	ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumDirectDebit   ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "DirectDebit"
	ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumPurchase      ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Purchase"
	ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumStandingOrder ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "StandingOrder"
	ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumAdjustment    ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Adjustment"
	ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumCredit        ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Credit"
	ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumOther         ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Other"
	ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumNotSupported  ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "NotSupported"
)

func (e *ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Fee":
		fallthrough
	case "Payment":
		fallthrough
	case "Cash":
		fallthrough
	case "Transfer":
		fallthrough
	case "Interest":
		fallthrough
	case "Cashback":
		fallthrough
	case "Cheque":
		fallthrough
	case "DirectDebit":
		fallthrough
	case "Purchase":
		fallthrough
	case "StandingOrder":
		fallthrough
	case "Adjustment":
		fallthrough
	case "Credit":
		fallthrough
	case "Other":
		fallthrough
	case "NotSupported":
		*e = ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum: %s", s)
	}
}

// ListBankingTransactions200ApplicationJSONSourceModifiedDateTransactionCategoryRef - An object of bank transaction category reference data.
type ListBankingTransactions200ApplicationJSONSourceModifiedDateTransactionCategoryRef struct {
	// The unique category reference id for the bank transaction.
	ID string `json:"id"`
	// The category name reference for the bank transaction.
	Name *string `json:"name,omitempty"`
}

// ListBankingTransactions200ApplicationJSONSourceModifiedDate - The Banking Transactions data type provides an immutable source of up-to-date information on income and expenditure.
//
// View the coverage for banking transactions in the [Data Coverage Explorer](https://knowledge.codat.io/supported-features/banking?view=tab-by-data-type&dataType=banking-transactions).
type ListBankingTransactions200ApplicationJSONSourceModifiedDate struct {
	// The unique identifier of the bank account.
	AccountID string `json:"accountId"`
	// The amount of the bank transaction.
	Amount *float64 `json:"amount,omitempty"`
	// The date the bank transaction was authorized.
	AuthorizedDate *string `json:"authorizedDate,omitempty"`
	// Code to identify the underlying transaction.
	Code *ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum `json:"code,omitempty"`
	// The currency of the bank transaction.
	Currency string `json:"currency"`
	// The description of the bank transaction.
	Description *string `json:"description,omitempty"`
	// The unique identifier of the bank transaction.
	ID string `json:"id"`
	// The name of the merchant.
	MerchantName *string `json:"merchantName,omitempty"`
	// The date on which this record was last modified in Codat.
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// The date the bank transaction was cleared.
	PostedDate *string `json:"postedDate,omitempty"`
	// The date on which this record was last modified in the originating system
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// An object of bank transaction category reference data.
	TransactionCategoryRef *ListBankingTransactions200ApplicationJSONSourceModifiedDateTransactionCategoryRef `json:"transactionCategoryRef,omitempty"`
}

// ListBankingTransactions200ApplicationJSON - Success
type ListBankingTransactions200ApplicationJSON struct {
	// The Banking Transactions data type provides an immutable source of up-to-date information on income and expenditure.
	//
	// View the coverage for banking transactions in the [Data Coverage Explorer](https://knowledge.codat.io/supported-features/banking?view=tab-by-data-type&dataType=banking-transactions).
	Results *ListBankingTransactions200ApplicationJSONSourceModifiedDate `json:"results,omitempty"`
}

type ListBankingTransactionsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	ListBankingTransactions200ApplicationJSONObject *ListBankingTransactions200ApplicationJSON
}
