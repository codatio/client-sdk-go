package operations

import (
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
)

type SendTransactionsDataPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type SendTransactionsDataRequest struct {
	PathParams SendTransactionsDataPathParams
	Request    *shared.CodatDataContractsDatasetsCommerceTransactionsDataset `request:"mediaType=application/json"`
}

type SendTransactionsDataResponse struct {
	CodatDataContractsDatasetsCommerceTransactionsDatasetSyncOperation *shared.CodatDataContractsDatasetsCommerceTransactionsDatasetSyncOperation
	ContentType                                                        string
	StatusCode                                                         int
}
