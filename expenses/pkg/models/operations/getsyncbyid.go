package operations

import (
	"github.com/codatio/client-sdk-go/expenses/pkg/models/shared"
)

type GetSyncByIDPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	SyncID    string `pathParam:"style=simple,explode=false,name=syncId"`
}

type GetSyncByIDRequest struct {
	PathParams GetSyncByIDPathParams
}

type GetSyncByIDResponse struct {
	CompanySyncStatus *shared.CompanySyncStatus
	ContentType       string
	StatusCode        int
}
