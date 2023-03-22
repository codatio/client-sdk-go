// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"time"
)

// GetDataIntegrityDetailsDataTypeEnum - A key for a Codat data type.
type GetDataIntegrityDetailsDataTypeEnum string

const (
	GetDataIntegrityDetailsDataTypeEnumBankingAccounts     GetDataIntegrityDetailsDataTypeEnum = "banking-accounts"
	GetDataIntegrityDetailsDataTypeEnumBankingTransactions GetDataIntegrityDetailsDataTypeEnum = "banking-transactions"
	GetDataIntegrityDetailsDataTypeEnumBankAccounts        GetDataIntegrityDetailsDataTypeEnum = "bankAccounts"
	GetDataIntegrityDetailsDataTypeEnumAccountTransactions GetDataIntegrityDetailsDataTypeEnum = "accountTransactions"
)

type GetDataIntegrityDetailsRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// A key for a Codat data type.
	DataType GetDataIntegrityDetailsDataTypeEnum `pathParam:"style=simple,explode=false,name=dataType"`
	// Field to order results by. [Read more](https://docs.codat.io/using-the-api/ordering-results).
	OrderBy *string `queryParam:"style=form,explode=true,name=orderBy"`
	// Page number. [Read more](https://docs.codat.io/using-the-api/paging).
	Page int `queryParam:"style=form,explode=true,name=page"`
	// Number of records to return in a page. [Read more](https://docs.codat.io/using-the-api/paging).
	PageSize *int `queryParam:"style=form,explode=true,name=pageSize"`
	// Codat query string. [Read more](https://docs.codat.io/using-the-api/querying).
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

type GetDataIntegrityDetailsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type GetDataIntegrityDetailsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type GetDataIntegrityDetailsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type GetDataIntegrityDetailsLinksLinksSelf struct {
	Href string `json:"href"`
}

type GetDataIntegrityDetailsLinksLinks struct {
	Current  GetDataIntegrityDetailsLinksLinksCurrent   `json:"current"`
	Next     *GetDataIntegrityDetailsLinksLinksNext     `json:"next,omitempty"`
	Previous *GetDataIntegrityDetailsLinksLinksPrevious `json:"previous,omitempty"`
	Self     GetDataIntegrityDetailsLinksLinksSelf      `json:"self"`
}

type GetDataIntegrityDetailsLinksDataIntegrityDetailsMatches struct {
	// The transaction value.
	Amount *string `json:"amount,omitempty"`
	// ID GUID representing the connection of the accounting or banking platform.
	ConnectionID *string `json:"connectionId,omitempty"`
	// The currency of the transaction.
	Currency *string `json:"currency,omitempty"`
	// The date of the transaction.
	Date *string `json:"date,omitempty"`
	// The transaction description.
	Description *string `json:"description,omitempty"`
	// ID GUID of the transaction.
	ID *string `json:"id,omitempty"`
	// The data type which the data type in the URL has been matched against. For example, if you've matched accountTransactions and banking-transactions, and you call this endpoint with accountTransactions in the URL, this property would be banking-transactions.
	Type *string `json:"type,omitempty"`
}

type GetDataIntegrityDetailsLinksDataIntegrityDetails struct {
	// The transaction value.
	Amount *float64 `json:"amount,omitempty"`
	// ID GUID representing the connection of the accounting or banking platform.
	ConnectionID *string `json:"connectionId,omitempty"`
	// The currency of the transaction.
	Currency *string `json:"currency,omitempty"`
	// The date of the transaction.
	Date *time.Time `json:"date,omitempty"`
	// The transaction description.
	Description *string `json:"description,omitempty"`
	// ID GUID of the transaction.
	ID      *string                                                   `json:"id,omitempty"`
	Matches []GetDataIntegrityDetailsLinksDataIntegrityDetailsMatches `json:"matches,omitempty"`
	// The data type of the record.
	Type *string `json:"type,omitempty"`
}

// GetDataIntegrityDetailsLinks - Codat's Paging Model
type GetDataIntegrityDetailsLinks struct {
	Links        GetDataIntegrityDetailsLinksLinks                  `json:"_links"`
	PageNumber   int64                                              `json:"pageNumber"`
	PageSize     int64                                              `json:"pageSize"`
	Results      []GetDataIntegrityDetailsLinksDataIntegrityDetails `json:"results,omitempty"`
	TotalResults int64                                              `json:"totalResults"`
}

type GetDataIntegrityDetailsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	Links *GetDataIntegrityDetailsLinks
}
