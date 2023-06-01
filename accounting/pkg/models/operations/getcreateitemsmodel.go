// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type GetCreateItemsModelRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetCreateItemsModelResponse struct {
	ContentType string
	// OK
	PushOption  *shared.PushOption
	StatusCode  int
	RawResponse *http.Response
	// Your API request was not properly authorized.
	Schema *shared.Schema
}
