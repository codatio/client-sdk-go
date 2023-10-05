// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/platform/pkg/models/shared"
	"net/http"
)

type UpdateConnectionAuthorizationRequest struct {
	RequestBody  map[string]string `request:"mediaType=application/json"`
	CompanyID    string            `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string            `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *UpdateConnectionAuthorizationRequest) GetRequestBody() map[string]string {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UpdateConnectionAuthorizationRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UpdateConnectionAuthorizationRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type UpdateConnectionAuthorizationResponse struct {
	// OK
	Connection *shared.Connection
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateConnectionAuthorizationResponse) GetConnection() *shared.Connection {
	if o == nil {
		return nil
	}
	return o.Connection
}

func (o *UpdateConnectionAuthorizationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateConnectionAuthorizationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateConnectionAuthorizationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
