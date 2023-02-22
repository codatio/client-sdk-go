package operations

import (
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
)

type CompaniesQueryParams struct {
	OrderBy  *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page     int     `queryParam:"style=form,explode=true,name=page"`
	PageSize *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string `queryParam:"style=form,explode=true,name=query"`
}

type CompaniesRequest struct {
	QueryParams CompaniesQueryParams
}

type CompaniesResponse struct {
	CodatPublicAPIModelsCompanyCompanyPagedResponseModel *shared.CodatPublicAPIModelsCompanyCompanyPagedResponseModel
	ContentType                                          string
	StatusCode                                           int
}
