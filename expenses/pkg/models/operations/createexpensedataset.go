package operations

import (
	"github.com/codatio/client-sdk-go/expenses/pkg/models/shared"
)

type CreateExpenseDatasetPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type CreateExpenseDatasetRequestBody struct {
	Items []shared.ExpenseReconciliation `json:"items,omitempty"`
}

type CreateExpenseDatasetRequest struct {
	PathParams CreateExpenseDatasetPathParams
	Request    *CreateExpenseDatasetRequestBody `request:"mediaType=application/json"`
}

type CreateExpenseDataset200ApplicationJSON struct {
	DatasetID *string `json:"datasetId,omitempty"`
}

type CreateExpenseDatasetResponse struct {
	ContentType                                  string
	StatusCode                                   int
	CreateExpenseDataset200ApplicationJSONObject *CreateExpenseDataset200ApplicationJSON
}
