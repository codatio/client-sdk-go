// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
	"net/http"
)

type GetCreateUpdateModelOptionsByDataTypeRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// The key of a Codat data type
	DataType shared.DataType `pathParam:"style=simple,explode=false,name=dataType"`
}

func (o *GetCreateUpdateModelOptionsByDataTypeRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetCreateUpdateModelOptionsByDataTypeRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetCreateUpdateModelOptionsByDataTypeRequest) GetDataType() shared.DataType {
	if o == nil {
		return shared.DataType("")
	}
	return o.DataType
}

type GetCreateUpdateModelOptionsByDataTypeResponse struct {
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	// OK
	PushOption  *shared.PushOption
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetCreateUpdateModelOptionsByDataTypeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCreateUpdateModelOptionsByDataTypeResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetCreateUpdateModelOptionsByDataTypeResponse) GetPushOption() *shared.PushOption {
	if o == nil {
		return nil
	}
	return o.PushOption
}

func (o *GetCreateUpdateModelOptionsByDataTypeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCreateUpdateModelOptionsByDataTypeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
