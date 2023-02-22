package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetTrackingCategoryPathParams struct {
	CompanyID          string `pathParam:"style=simple,explode=false,name=companyId"`
	TrackingCategoryID string `pathParam:"style=simple,explode=false,name=trackingCategoryId"`
}

type GetTrackingCategorySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetTrackingCategoryRequest struct {
	PathParams GetTrackingCategoryPathParams
	Security   GetTrackingCategorySecurity
}

type GetTrackingCategoryResponse struct {
	ContentType          string
	StatusCode           int
	TrackingCategoryTree *shared.TrackingCategoryTree
}
