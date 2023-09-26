// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type IsAgedCreditorsReportAvailableRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *IsAgedCreditorsReportAvailableRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type IsAgedCreditorsReportAvailableResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	IsAgedCreditorsReportAvailable200ApplicationJSONBoolean *bool
}

func (o *IsAgedCreditorsReportAvailableResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *IsAgedCreditorsReportAvailableResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *IsAgedCreditorsReportAvailableResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *IsAgedCreditorsReportAvailableResponse) GetIsAgedCreditorsReportAvailable200ApplicationJSONBoolean() *bool {
	if o == nil {
		return nil
	}
	return o.IsAgedCreditorsReportAvailable200ApplicationJSONBoolean
}
