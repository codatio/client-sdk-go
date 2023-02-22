package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsPathParams struct {
	CompanyID string                           `pathParam:"style=simple,explode=false,name=companyId"`
	DataType  shared.DataIntegritydataTypeEnum `pathParam:"style=simple,explode=false,name=dataType"`
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

// GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinks
// Codat's Paging Model
type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinks struct {
	Links        GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinksLinks `json:"_links"`
	PageNumber   int64                                                                          `json:"pageNumber"`
	PageSize     int64                                                                          `json:"pageSize"`
	Results      []shared.DataIntegrityDetails                                                  `json:"results,omitempty"`
	TotalResults int64                                                                          `json:"totalResults"`
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsResponse struct {
	ContentType string
	StatusCode  int
	Links       *GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinks
}
