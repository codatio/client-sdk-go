// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/platform/v2/pkg/models/shared"
	"net/http"
)

type UnlinkConnectionRequest struct {
	UpdateConnectionStatus *shared.UpdateConnectionStatus `request:"mediaType=application/json"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *UnlinkConnectionRequest) GetUpdateConnectionStatus() *shared.UpdateConnectionStatus {
	if o == nil {
		return nil
	}
	return o.UpdateConnectionStatus
}

func (o *UnlinkConnectionRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UnlinkConnectionRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type UnlinkConnectionResponse struct {
	// OK
	Connection *shared.Connection
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UnlinkConnectionResponse) GetConnection() *shared.Connection {
	if o == nil {
		return nil
	}
	return o.Connection
}

func (o *UnlinkConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UnlinkConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UnlinkConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
