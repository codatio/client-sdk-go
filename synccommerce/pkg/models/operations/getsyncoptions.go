package operations

import (
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
)

type GetSyncOptionsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetSyncOptionsRequest struct {
	PathParams GetSyncOptionsPathParams
}

type GetSyncOptionsResponse struct {
	CodatCommerceDataContractsConfigCompanyConfiguration *shared.CodatCommerceDataContractsConfigCompanyConfiguration
	ContentType                                          string
	StatusCode                                           int
}
