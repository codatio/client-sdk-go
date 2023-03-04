package operations

import (
	"net/http"
)

type IsAgedDebtorReportAvailablePathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type IsAgedDebtorReportAvailableRequest struct {
	PathParams IsAgedDebtorReportAvailablePathParams
}

type IsAgedDebtorReportAvailableResponse struct {
	ContentType                                          string
	StatusCode                                           int
	RawResponse                                          *http.Response
	IsAgedDebtorReportAvailable200ApplicationJSONBoolean *bool
}
