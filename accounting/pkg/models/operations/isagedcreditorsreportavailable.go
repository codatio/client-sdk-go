package operations

import (
	"net/http"
)

type IsAgedCreditorsReportAvailablePathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type IsAgedCreditorsReportAvailableRequest struct {
	PathParams IsAgedCreditorsReportAvailablePathParams
}

type IsAgedCreditorsReportAvailableResponse struct {
	ContentType                                             string
	StatusCode                                              int
	RawResponse                                             *http.Response
	IsAgedCreditorsReportAvailable200ApplicationJSONBoolean *bool
}
