package operations

import (
	"github.com/codatio/client-sdk-go/expenses/pkg/models/shared"
)

type GetSyncTransactionsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	SyncID    string `pathParam:"style=simple,explode=false,name=syncId"`
}

type GetSyncTransactionsQueryParams struct {
	Page     int  `queryParam:"style=form,explode=true,name=page"`
	PageSize *int `queryParam:"style=form,explode=true,name=pageSize"`
}

type GetSyncTransactionsRequest struct {
	PathParams  GetSyncTransactionsPathParams
	QueryParams GetSyncTransactionsQueryParams
}

type GetSyncTransactionsResponse struct {
	ContentType                      string
	StatusCode                       int
	TransactionMetadataPagedResponse *shared.TransactionMetadataPagedResponse
}
