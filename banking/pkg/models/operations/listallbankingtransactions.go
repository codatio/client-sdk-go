package operations

import (
	"net/http"
	"time"
)

type ListAllBankingTransactionsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListAllBankingTransactionsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListAllBankingTransactionsRequest struct {
	PathParams  ListAllBankingTransactionsPathParams
	QueryParams ListAllBankingTransactionsQueryParams
}

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

// ListAllBankingTransactions200ApplicationJSONSourceModifiedDateTransactionCategoryRef
// An object of bank transaction category reference data.
type ListAllBankingTransactions200ApplicationJSONSourceModifiedDateTransactionCategoryRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// ListAllBankingTransactions200ApplicationJSONSourceModifiedDate
// The Banking Transactions data type provides an immutable source of up-to-date information on income and expenditure.
//
// View the coverage for banking transactions in the [Data Coverage Explorer](https://knowledge.codat.io/supported-features/banking?view=tab-by-data-type&dataType=banking-transactions).
type ListAllBankingTransactions200ApplicationJSONSourceModifiedDate struct {
	AccountID              string                                                                                `json:"accountId"`
	Amount                 *float64                                                                              `json:"amount,omitempty"`
	AuthorizedDate         *time.Time                                                                            `json:"authorizedDate,omitempty"`
	Code                   *ListAllBankingTransactions200ApplicationJSONSourceModifiedDateCodeEnum               `json:"code,omitempty"`
	Currency               string                                                                                `json:"currency"`
	Description            *string                                                                               `json:"description,omitempty"`
	ID                     string                                                                                `json:"id"`
	MerchantName           *string                                                                               `json:"merchantName,omitempty"`
	ModifiedDate           *time.Time                                                                            `json:"modifiedDate,omitempty"`
	PostedDate             *time.Time                                                                            `json:"postedDate,omitempty"`
	SourceModifiedDate     *time.Time                                                                            `json:"sourceModifiedDate,omitempty"`
	TransactionCategoryRef *ListAllBankingTransactions200ApplicationJSONSourceModifiedDateTransactionCategoryRef `json:"transactionCategoryRef,omitempty"`
}

type ListAllBankingTransactions200ApplicationJSON struct {
	Results *ListAllBankingTransactions200ApplicationJSONSourceModifiedDate `json:"results,omitempty"`
}

type ListAllBankingTransactionsResponse struct {
	ContentType                                        string
	StatusCode                                         int
	RawResponse                                        *http.Response
	ListAllBankingTransactions200ApplicationJSONObject *ListAllBankingTransactions200ApplicationJSON
}
