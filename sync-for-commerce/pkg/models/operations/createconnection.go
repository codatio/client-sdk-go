// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"net/http"
)

type CreateConnectionRequestBody struct {
	PlatformKey *string `json:"platformKey,omitempty"`
}

func (o *CreateConnectionRequestBody) GetPlatformKey() *string {
	if o == nil {
		return nil
	}
	return o.PlatformKey
}

type CreateConnectionRequest struct {
	RequestBody *CreateConnectionRequestBody `request:"mediaType=application/json"`
	CompanyID   string                       `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *CreateConnectionRequest) GetRequestBody() *CreateConnectionRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *CreateConnectionRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type CreateConnectionResponse struct {
	// OK
	Connection  *shared.Connection
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *CreateConnectionResponse) GetConnection() *shared.Connection {
	if o == nil {
		return nil
	}
	return o.Connection
}

func (o *CreateConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateConnectionResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *CreateConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
