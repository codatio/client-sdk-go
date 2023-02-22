package operations

import (
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
)

type PostSyncLatestPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type PostSyncLatestRequests struct {
	CodatCommerceAPIModelsSyncToLatestArgs  *shared.CodatCommerceAPIModelsSyncToLatestArgs `request:"mediaType=application/*+json"`
	CodatCommerceAPIModelsSyncToLatestArgs1 *shared.CodatCommerceAPIModelsSyncToLatestArgs `request:"mediaType=application/json"`
	CodatCommerceAPIModelsSyncToLatestArgs2 *shared.CodatCommerceAPIModelsSyncToLatestArgs `request:"mediaType=application/json-patch+json"`
	CodatCommerceAPIModelsSyncToLatestArgs3 *shared.CodatCommerceAPIModelsSyncToLatestArgs `request:"mediaType=text/json"`
}

type PostSyncLatestRequest struct {
	PathParams PostSyncLatestPathParams
	Request    *PostSyncLatestRequests
}

type PostSyncLatestResponse struct {
	CodatCommerceDataContractsModelsCommerceSyncCreateSyncResponse *shared.CodatCommerceDataContractsModelsCommerceSyncCreateSyncResponse
	ContentType                                                    string
	StatusCode                                                     int
}
