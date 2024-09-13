// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/shared"
	"net/http"
)

type RefreshDataTypeRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Optionally, provide a data connection id to only queue pull operations on that connection.
	ConnectionID *string `queryParam:"style=form,explode=true,name=connectionId"`
	// A key for a Codat data type.
	DataType shared.SchemaDataType `pathParam:"style=simple,explode=false,name=dataType"`
}

func (o *RefreshDataTypeRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *RefreshDataTypeRequest) GetConnectionID() *string {
	if o == nil {
		return nil
	}
	return o.ConnectionID
}

func (o *RefreshDataTypeRequest) GetDataType() shared.SchemaDataType {
	if o == nil {
		return shared.SchemaDataType("")
	}
	return o.DataType
}

type RefreshDataTypeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	PullOperation *shared.PullOperation
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RefreshDataTypeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RefreshDataTypeResponse) GetPullOperation() *shared.PullOperation {
	if o == nil {
		return nil
	}
	return o.PullOperation
}

func (o *RefreshDataTypeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RefreshDataTypeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
