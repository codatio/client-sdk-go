package operations

import (
	"net/http"
	"time"
)

type GetDataIntegrityDetailsDataTypeEnum string

const (
	GetDataIntegrityDetailsDataTypeEnumBankingAccounts     GetDataIntegrityDetailsDataTypeEnum = "banking-accounts"
	GetDataIntegrityDetailsDataTypeEnumBankingTransactions GetDataIntegrityDetailsDataTypeEnum = "banking-transactions"
	GetDataIntegrityDetailsDataTypeEnumBankAccounts        GetDataIntegrityDetailsDataTypeEnum = "bankAccounts"
	GetDataIntegrityDetailsDataTypeEnumAccountTransactions GetDataIntegrityDetailsDataTypeEnum = "accountTransactions"
)

type GetDataIntegrityDetailsRequest struct {
	CompanyID string                              `pathParam:"style=simple,explode=false,name=companyId"`
	DataType  GetDataIntegrityDetailsDataTypeEnum `pathParam:"style=simple,explode=false,name=dataType"`
	OrderBy   *string                             `queryParam:"style=form,explode=true,name=orderBy"`
	Page      int                                 `queryParam:"style=form,explode=true,name=page"`
	PageSize  *int                                `queryParam:"style=form,explode=true,name=pageSize"`
	Query     *string                             `queryParam:"style=form,explode=true,name=query"`
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
	Amount       *string `json:"amount,omitempty"`
	ConnectionID *string `json:"connectionId,omitempty"`
	Currency     *string `json:"currency,omitempty"`
	Date         *string `json:"date,omitempty"`
	Description  *string `json:"description,omitempty"`
	ID           *string `json:"id,omitempty"`
	Type         *string `json:"type,omitempty"`
}

type GetDataIntegrityDetailsLinksDataIntegrityDetails struct {
	Amount       *float64                                                  `json:"amount,omitempty"`
	ConnectionID *string                                                   `json:"connectionId,omitempty"`
	Currency     *string                                                   `json:"currency,omitempty"`
	Date         *time.Time                                                `json:"date,omitempty"`
	Description  *string                                                   `json:"description,omitempty"`
	ID           *string                                                   `json:"id,omitempty"`
	Matches      []GetDataIntegrityDetailsLinksDataIntegrityDetailsMatches `json:"matches,omitempty"`
	Type         *string                                                   `json:"type,omitempty"`
}

// GetDataIntegrityDetailsLinks
// Codat's Paging Model
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
	Links       *GetDataIntegrityDetailsLinks
}
