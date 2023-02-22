package operations

import (
	"github.com/codatio/client-sdk-go/expenses/pkg/models/shared"
)

type GetLatestSyncPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetLatestSyncRequest struct {
	PathParams GetLatestSyncPathParams
}

type GetLatestSyncResponse struct {
	CompanySyncStatus *shared.CompanySyncStatus
	ContentType       string
	StatusCode        int
}
