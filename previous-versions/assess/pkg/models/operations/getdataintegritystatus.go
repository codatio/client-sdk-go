// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"net/http"
)

type GetDataIntegrityStatusRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// A key for a Codat data type.
	DataType shared.DataIntegrityDataType `pathParam:"style=simple,explode=false,name=dataType"`
}

func (o *GetDataIntegrityStatusRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetDataIntegrityStatusRequest) GetDataType() shared.DataIntegrityDataType {
	if o == nil {
		return shared.DataIntegrityDataType("")
	}
	return o.DataType
}

type GetDataIntegrityStatusResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	Status *shared.Status
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetDataIntegrityStatusResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDataIntegrityStatusResponse) GetStatus() *shared.Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetDataIntegrityStatusResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDataIntegrityStatusResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
