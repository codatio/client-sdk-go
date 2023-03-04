package operations

import (
	"net/http"
)

type GetCompanyCommerceSyncStatusPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCompanyCommerceSyncStatusRequest struct {
	PathParams GetCompanyCommerceSyncStatusPathParams
}

type GetCompanyCommerceSyncStatusResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
