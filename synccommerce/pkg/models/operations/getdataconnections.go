package operations

import (
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
)

type GetDataconnectionsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetDataconnectionsQueryParams struct {
	OrderBy  *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page     int     `queryParam:"style=form,explode=true,name=page"`
	PageSize *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string `queryParam:"style=form,explode=true,name=query"`
}

type GetDataconnectionsRequest struct {
	PathParams  GetDataconnectionsPathParams
	QueryParams GetDataconnectionsQueryParams
}

type GetDataconnectionsResponse struct {
	CodatPublicAPIModelsCompanyDataConnectionPagedResponseModel *shared.CodatPublicAPIModelsCompanyDataConnectionPagedResponseModel
	ContentType                                                 string
	StatusCode                                                  int
}
