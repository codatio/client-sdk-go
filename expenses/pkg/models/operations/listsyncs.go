package operations

import (
	"github.com/codatio/client-sdk-go/expenses/pkg/models/shared"
)

type ListSyncsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListSyncsRequest struct {
	PathParams ListSyncsPathParams
}

type ListSyncsResponse struct {
	CompanySyncStatuses []shared.CompanySyncStatus
	ContentType         string
	StatusCode          int
}
