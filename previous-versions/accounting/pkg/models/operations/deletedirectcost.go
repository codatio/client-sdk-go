// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"net/http"
)

type DeleteDirectCostRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a direct cost.
	DirectCostID string `pathParam:"style=simple,explode=false,name=directCostId"`
}

func (o *DeleteDirectCostRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *DeleteDirectCostRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *DeleteDirectCostRequest) GetDirectCostID() string {
	if o == nil {
		return ""
	}
	return o.DirectCostID
}

type DeleteDirectCostResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	PushOperationSummary *shared.PushOperationSummary
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteDirectCostResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteDirectCostResponse) GetPushOperationSummary() *shared.PushOperationSummary {
	if o == nil {
		return nil
	}
	return o.PushOperationSummary
}

func (o *DeleteDirectCostResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteDirectCostResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}