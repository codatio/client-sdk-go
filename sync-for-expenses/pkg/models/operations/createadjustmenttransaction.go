// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/models/shared"
	"net/http"
)

type CreateAdjustmentTransactionRequest struct {
	RequestBody []shared.AdjustmentTransactionRequest `request:"mediaType=application/json"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *CreateAdjustmentTransactionRequest) GetRequestBody() []shared.AdjustmentTransactionRequest {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *CreateAdjustmentTransactionRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type CreateAdjustmentTransactionResponse struct {
	// OK
	AdjustmentTransactionResponse *shared.AdjustmentTransactionResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateAdjustmentTransactionResponse) GetAdjustmentTransactionResponse() *shared.AdjustmentTransactionResponse {
	if o == nil {
		return nil
	}
	return o.AdjustmentTransactionResponse
}

func (o *CreateAdjustmentTransactionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAdjustmentTransactionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAdjustmentTransactionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}