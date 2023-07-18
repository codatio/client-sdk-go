// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type GetCreateDirectIncomesModelRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *GetCreateDirectIncomesModelRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetCreateDirectIncomesModelRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type GetCreateDirectIncomesModelResponse struct {
	ContentType string
	// OK
	PushOption  *shared.PushOption
	StatusCode  int
	RawResponse *http.Response
	// Your API request was not properly authorized.
	Schema *shared.Schema
}

func (o *GetCreateDirectIncomesModelResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCreateDirectIncomesModelResponse) GetPushOption() *shared.PushOption {
	if o == nil {
		return nil
	}
	return o.PushOption
}

func (o *GetCreateDirectIncomesModelResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCreateDirectIncomesModelResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetCreateDirectIncomesModelResponse) GetSchema() *shared.Schema {
	if o == nil {
		return nil
	}
	return o.Schema
}
