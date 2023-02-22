package operations

import (
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
)

type SendOrdersDataPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type SendOrdersDataRequest struct {
	PathParams SendOrdersDataPathParams
	Request    *shared.CodatDataContractsDatasetsCommerceOrdersDataset `request:"mediaType=application/json"`
}

type SendOrdersDataResponse struct {
	CodatDataContractsDatasetsCommerceOrdersDatasetSyncOperation *shared.CodatDataContractsDatasetsCommerceOrdersDatasetSyncOperation
	ContentType                                                  string
	StatusCode                                                   int
}
