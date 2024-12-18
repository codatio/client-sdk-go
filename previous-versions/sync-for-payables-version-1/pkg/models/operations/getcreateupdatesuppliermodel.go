// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v3/pkg/models/shared"
	"net/http"
)

type GetCreateUpdateSupplierModelRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *GetCreateUpdateSupplierModelRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetCreateUpdateSupplierModelRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type GetCreateUpdateSupplierModelResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	PushOption *shared.PushOption
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetCreateUpdateSupplierModelResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCreateUpdateSupplierModelResponse) GetPushOption() *shared.PushOption {
	if o == nil {
		return nil
	}
	return o.PushOption
}

func (o *GetCreateUpdateSupplierModelResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCreateUpdateSupplierModelResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
