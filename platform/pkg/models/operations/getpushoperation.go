// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/platform/v5/pkg/models/shared"
	"net/http"
)

type GetPushOperationRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Push operation key.
	PushOperationKey string `pathParam:"style=simple,explode=false,name=pushOperationKey"`
}

func (o *GetPushOperationRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetPushOperationRequest) GetPushOperationKey() string {
	if o == nil {
		return ""
	}
	return o.PushOperationKey
}

type GetPushOperationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	PushOperation *shared.PushOperation
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetPushOperationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPushOperationResponse) GetPushOperation() *shared.PushOperation {
	if o == nil {
		return nil
	}
	return o.PushOperation
}

func (o *GetPushOperationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPushOperationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
