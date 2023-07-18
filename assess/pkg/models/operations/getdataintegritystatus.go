// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"net/http"
)

type GetDataIntegrityStatusRequest struct {
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
		return DataIntegrityDataType("")
	}
	return o.DataType
}

type GetDataIntegrityStatusResponse struct {
	ContentType string
	// OK
	Status      *shared.Status
	StatusCode  int
	RawResponse *http.Response
	// Your API request was not properly authorized.
	Schema *shared.Schema
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

func (o *GetDataIntegrityStatusResponse) GetSchema() *shared.Schema {
	if o == nil {
		return nil
	}
	return o.Schema
}
