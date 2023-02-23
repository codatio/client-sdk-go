package operations

import (
	"github.com/codatio/client-sdk-go/banking/pkg/models/shared"
	"time"
)

type ListBankingTransactionsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListBankingTransactionsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListBankingTransactionsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListBankingTransactionsRequest struct {
	PathParams  ListBankingTransactionsPathParams
	QueryParams ListBankingTransactionsQueryParams
	Security    ListBankingTransactionsSecurity
}

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

// ListBankingTransactions200ApplicationJSONSourceModifiedDateTransactionCategoryRef
// An object of bank transaction category reference data.
type ListBankingTransactions200ApplicationJSONSourceModifiedDateTransactionCategoryRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// ListBankingTransactions200ApplicationJSONSourceModifiedDate
// The Banking Transactions data type provides an immutable source of up-to-date information on income and expenditure.
//
// View the coverage for banking transactions in the [Data Coverage Explorer](https://knowledge.codat.io/supported-features/banking?view=tab-by-data-type&dataType=banking-transactions).
type ListBankingTransactions200ApplicationJSONSourceModifiedDate struct {
	AccountID              string                                                                             `json:"accountId"`
	Amount                 *float64                                                                           `json:"amount,omitempty"`
	AuthorizedDate         *time.Time                                                                         `json:"authorizedDate,omitempty"`
	Code                   *ListBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum               `json:"code,omitempty"`
	Currency               string                                                                             `json:"currency"`
	Description            *string                                                                            `json:"description,omitempty"`
	ID                     string                                                                             `json:"id"`
	MerchantName           *string                                                                            `json:"merchantName,omitempty"`
	ModifiedDate           *time.Time                                                                         `json:"modifiedDate,omitempty"`
	PostedDate             *time.Time                                                                         `json:"postedDate,omitempty"`
	SourceModifiedDate     *time.Time                                                                         `json:"sourceModifiedDate,omitempty"`
	TransactionCategoryRef *ListBankingTransactions200ApplicationJSONSourceModifiedDateTransactionCategoryRef `json:"transactionCategoryRef,omitempty"`
}

type ListBankingTransactions200ApplicationJSON struct {
	Results *ListBankingTransactions200ApplicationJSONSourceModifiedDate `json:"results,omitempty"`
}

type ListBankingTransactionsResponse struct {
	ContentType                                     string
	StatusCode                                      int
	ListBankingTransactions200ApplicationJSONObject *ListBankingTransactions200ApplicationJSON
}
