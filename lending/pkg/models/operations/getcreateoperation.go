// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"net/http"
)

type GetCreateOperationRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for the push operation.
	PushOperationKey string `pathParam:"style=simple,explode=false,name=pushOperationKey"`
}

func (o *GetCreateOperationRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetCreateOperationRequest) GetPushOperationKey() string {
	if o == nil {
		return ""
	}
	return o.PushOperationKey
}

type GetCreateOperationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	PushOperation *shared.PushOperation
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetCreateOperationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCreateOperationResponse) GetPushOperation() *shared.PushOperation {
	if o == nil {
		return nil
	}
	return o.PushOperation
}

func (o *GetCreateOperationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCreateOperationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
