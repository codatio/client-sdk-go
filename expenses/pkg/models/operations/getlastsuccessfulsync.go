package operations

import (
	"github.com/codatio/client-sdk-go/expenses/pkg/models/shared"
)

type GetLastSuccessfulSyncPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetLastSuccessfulSyncRequest struct {
	PathParams GetLastSuccessfulSyncPathParams
}

type GetLastSuccessfulSyncResponse struct {
	CompanySyncStatus *shared.CompanySyncStatus
	ContentType       string
	StatusCode        int
}
