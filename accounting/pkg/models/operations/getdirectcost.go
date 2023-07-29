// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type GetDirectCostRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a direct cost
	DirectCostID string `pathParam:"style=simple,explode=false,name=directCostId"`
}

func (o *GetDirectCostRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetDirectCostRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetDirectCostRequest) GetDirectCostID() string {
	if o == nil {
		return ""
	}
	return o.DirectCostID
}

type GetDirectCostResponse struct {
	ContentType string
	// Success
	DirectCost *shared.DirectCost
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetDirectCostResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDirectCostResponse) GetDirectCost() *shared.DirectCost {
	if o == nil {
		return nil
	}
	return o.DirectCost
}

func (o *GetDirectCostResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetDirectCostResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDirectCostResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
