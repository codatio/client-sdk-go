package operations

import (
	"net/http"
)

type IsAgedCreditorsReportAvailableRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type IsAgedCreditorsReportAvailableResponse struct {
	ContentType                                             string
	StatusCode                                              int
	RawResponse                                             *http.Response
	IsAgedCreditorsReportAvailable200ApplicationJSONBoolean *bool
}
