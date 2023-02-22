package operations

import (
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
)

type GetIntegrationsQueryParams struct {
	OrderBy  *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page     int     `queryParam:"style=form,explode=true,name=page"`
	PageSize *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string `queryParam:"style=form,explode=true,name=query"`
}

type GetIntegrationsRequest struct {
	QueryParams GetIntegrationsQueryParams
}

type GetIntegrationsResponse struct {
	CodatPublicAPIModelsPlatformCredentialsPlatformSourceModelPagedResponseModel *shared.CodatPublicAPIModelsPlatformCredentialsPlatformSourceModelPagedResponseModel
	ContentType                                                                  string
	StatusCode                                                                   int
}
