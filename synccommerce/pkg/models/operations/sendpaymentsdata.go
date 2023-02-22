package operations

import (
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
)

type SendPaymentsDataPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type SendPaymentsDataRequest struct {
	PathParams SendPaymentsDataPathParams
	Request    *shared.CodatDataContractsDatasetsCommercePaymentsDataset `request:"mediaType=application/json"`
}

type SendPaymentsDataResponse struct {
	CodatDataContractsDatasetsCommercePaymentsDatasetSyncOperation *shared.CodatDataContractsDatasetsCommercePaymentsDatasetSyncOperation
	ContentType                                                    string
	StatusCode                                                     int
}
