package operations

import (
	"net/http"
)

type IsAgedDebtorReportAvailableRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type IsAgedDebtorReportAvailableResponse struct {
	ContentType                                          string
	StatusCode                                           int
	RawResponse                                          *http.Response
	IsAgedDebtorReportAvailable200ApplicationJSONBoolean *bool
}
