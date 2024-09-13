// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	"net/http"
)

type RefreshCustomDataTypeRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a custom data type.
	CustomDataIdentifier string `pathParam:"style=simple,explode=false,name=customDataIdentifier"`
}

func (o *RefreshCustomDataTypeRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *RefreshCustomDataTypeRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *RefreshCustomDataTypeRequest) GetCustomDataIdentifier() string {
	if o == nil {
		return ""
	}
	return o.CustomDataIdentifier
}

type RefreshCustomDataTypeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	PullOperation *shared.PullOperation
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RefreshCustomDataTypeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RefreshCustomDataTypeResponse) GetPullOperation() *shared.PullOperation {
	if o == nil {
		return nil
	}
	return o.PullOperation
}

func (o *RefreshCustomDataTypeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RefreshCustomDataTypeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
