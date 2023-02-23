package operations

import (
	"time"
)

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsDataTypeEnum string

const (
	GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsDataTypeEnumBankingAccounts     GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsDataTypeEnum = "banking-accounts"
	GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsDataTypeEnumBankingTransactions GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsDataTypeEnum = "banking-transactions"
	GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsDataTypeEnumBankAccounts        GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsDataTypeEnum = "bankAccounts"
	GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsDataTypeEnumAccountTransactions GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsDataTypeEnum = "accountTransactions"
)

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsPathParams struct {
	CompanyID string                                                                           `pathParam:"style=simple,explode=false,name=companyId"`
	DataType  GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsDataTypeEnum `pathParam:"style=simple,explode=false,name=dataType"`
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsRequest struct {
	PathParams  GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsPathParams
	QueryParams GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsQueryParams
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinksLinksSelf struct {
	Href string `json:"href"`
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinksLinks struct {
	Current  GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinksLinksCurrent   `json:"current"`
	Next     *GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinksLinksNext     `json:"next,omitempty"`
	Previous *GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinksLinksPrevious `json:"previous,omitempty"`
	Self     GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinksLinksSelf      `json:"self"`
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinksDataIntegrityDetailsMatches struct {
	Amount       *string `json:"amount,omitempty"`
	ConnectionID *string `json:"connectionId,omitempty"`
	Currency     *string `json:"currency,omitempty"`
	Date         *string `json:"date,omitempty"`
	Description  *string `json:"description,omitempty"`
	ID           *string `json:"id,omitempty"`
	Type         *string `json:"type,omitempty"`
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinksDataIntegrityDetails struct {
	Amount       *float64                                                                                               `json:"amount,omitempty"`
	ConnectionID *string                                                                                                `json:"connectionId,omitempty"`
	Currency     *string                                                                                                `json:"currency,omitempty"`
	Date         *time.Time                                                                                             `json:"date,omitempty"`
	Description  *string                                                                                                `json:"description,omitempty"`
	ID           *string                                                                                                `json:"id,omitempty"`
	Matches      []GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinksDataIntegrityDetailsMatches `json:"matches,omitempty"`
	Type         *string                                                                                                `json:"type,omitempty"`
}

// GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinks
// Codat's Paging Model
type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinks struct {
	Links        GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinksLinks                  `json:"_links"`
	PageNumber   int64                                                                                           `json:"pageNumber"`
	PageSize     int64                                                                                           `json:"pageSize"`
	Results      []GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinksDataIntegrityDetails `json:"results,omitempty"`
	TotalResults int64                                                                                           `json:"totalResults"`
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsResponse struct {
	ContentType string
	StatusCode  int
	Links       *GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinks
}
