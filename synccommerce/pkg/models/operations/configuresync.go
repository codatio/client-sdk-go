package operations

import (
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
)

type ConfigureSyncPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ConfigureSyncRequest struct {
	PathParams ConfigureSyncPathParams
	Request    *shared.CodatCommerceDataContractsConfigCompanyConfiguration `request:"mediaType=application/json"`
}

type ConfigureSyncResponse struct {
	CodatCommerceDataContractsConfigCompanyConfiguration *shared.CodatCommerceDataContractsConfigCompanyConfiguration
	ContentType                                          string
	StatusCode                                           int
}
