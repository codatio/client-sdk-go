// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
	"net/http"
)

type UnlinkConnectionRequestBody struct {
	Status *string `json:"status,omitempty"`
}

type UnlinkConnectionRequest struct {
	RequestBody  *UnlinkConnectionRequestBody `request:"mediaType=application/json"`
	CompanyID    string                       `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string                       `pathParam:"style=simple,explode=false,name=connectionId"`
}

type UnlinkConnectionResponse struct {
	// OK
	Connection  *shared.Connection
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}