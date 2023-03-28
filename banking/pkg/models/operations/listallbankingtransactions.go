// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ListAllBankingTransactionsRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Field to order results by. [Read more](https://docs.codat.io/using-the-api/ordering-results).
	OrderBy *string `queryParam:"style=form,explode=true,name=orderBy"`
	// Page number. [Read more](https://docs.codat.io/using-the-api/paging).
	Page int `queryParam:"style=form,explode=true,name=page"`
	// Number of records to return in a page. [Read more](https://docs.codat.io/using-the-api/paging).
	PageSize *int `queryParam:"style=form,explode=true,name=pageSize"`
	// Codat query string. [Read more](https://docs.codat.io/using-the-api/querying).
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

// ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum - Code to identify the underlying transaction.
type ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum string

const (
	ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumUnknown       ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Unknown"
	ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumFee           ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Fee"
	ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumPayment       ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Payment"
	ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumCash          ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Cash"
	ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumTransfer      ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Transfer"
	ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumInterest      ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Interest"
	ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumCashback      ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Cashback"
	ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumCheque        ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Cheque"
	ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumDirectDebit   ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "DirectDebit"
	ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumPurchase      ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Purchase"
	ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumStandingOrder ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "StandingOrder"
	ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumAdjustment    ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Adjustment"
	ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumCredit        ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Credit"
	ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumOther         ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "Other"
	ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnumNotSupported  ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum = "NotSupported"
)

func (e *ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum) UnmarshalJSON(data []byte) error {
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
		*e = ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum: %s", s)
	}
}

// ListAllBankingTransactions200ApplicationJSONSourceModifiedDateTransactionCategoryRef - An object of bank transaction category reference data.
type ListAllBankingTransactions200ApplicationJSONSourceModifiedDateTransactionCategoryRef struct {
	// The unique category reference id for the bank transaction.
	ID string `json:"id"`
	// The category name reference for the bank transaction.
	Name *string `json:"name,omitempty"`
}

// ListAllBankingTransactions200ApplicationJSONSourceModifiedDate - The Banking Transactions data type provides an immutable source of up-to-date information on income and expenditure.
//
// View the coverage for banking transactions in the [Data Coverage Explorer](https://knowledge.codat.io/supported-features/banking?view=tab-by-data-type&dataType=banking-transactions).
type ListAllBankingTransactions200ApplicationJSONSourceModifiedDate struct {
	// The unique identifier of the bank account.
	AccountID string `json:"accountId"`
	// The amount of the bank transaction.
	Amount *float64 `json:"amount,omitempty"`
	// The date the bank transaction was authorized.
	AuthorizedDate *string `json:"authorizedDate,omitempty"`
	// Code to identify the underlying transaction.
	Code *ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum `json:"code,omitempty"`
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
	TransactionCategoryRef *ListAllBankingTransactions200ApplicationJSONSourceModifiedDateTransactionCategoryRef `json:"transactionCategoryRef,omitempty"`
}

// ListAllBankingTransactions200ApplicationJSON - Success
type ListAllBankingTransactions200ApplicationJSON struct {
	// The Banking Transactions data type provides an immutable source of up-to-date information on income and expenditure.
	//
	// View the coverage for banking transactions in the [Data Coverage Explorer](https://knowledge.codat.io/supported-features/banking?view=tab-by-data-type&dataType=banking-transactions).
	Results *ListAllBankingTransactions200ApplicationJSONSourceModifiedDate `json:"results,omitempty"`
}

type ListAllBankingTransactionsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	ListAllBankingTransactions200ApplicationJSONObject *ListAllBankingTransactions200ApplicationJSON
}
