package operations

import (
	"github.com/codatio/client-sdk-go/expenses/pkg/models/shared"
)

type GetSyncTransactionPathParams struct {
	TransactionID string `pathParam:"style=simple,explode=false,name=TransactionId"`
	CompanyID     string `pathParam:"style=simple,explode=false,name=companyId"`
	SyncID        string `pathParam:"style=simple,explode=false,name=syncId"`
}

type GetSyncTransactionRequest struct {
	PathParams GetSyncTransactionPathParams
}

type GetSyncTransactionResponse struct {
	ContentType         string
	StatusCode          int
	TransactionMetadata []shared.TransactionMetadata
}
