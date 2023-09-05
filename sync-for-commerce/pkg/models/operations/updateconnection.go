// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-commerce/pkg/models/shared"
	"net/http"
)

type UpdateConnectionRequest struct {
	UpdateConnection *shared.UpdateConnection `request:"mediaType=application/json"`
	CompanyID        string                   `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string                   `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *UpdateConnectionRequest) GetUpdateConnection() *shared.UpdateConnection {
	if o == nil {
		return nil
	}
	return o.UpdateConnection
}

func (o *UpdateConnectionRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UpdateConnectionRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type UpdateConnectionResponse struct {
	// Success
	Connection  *shared.Connection
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *UpdateConnectionResponse) GetConnection() *shared.Connection {
	if o == nil {
		return nil
	}
	return o.Connection
}

func (o *UpdateConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}