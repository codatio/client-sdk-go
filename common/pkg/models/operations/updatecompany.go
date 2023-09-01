// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
	"net/http"
)

type UpdateCompanyRequest struct {
	CompanyRequestBody *shared.CompanyRequestBody `request:"mediaType=application/json"`
	CompanyID          string                     `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *UpdateCompanyRequest) GetCompanyRequestBody() *shared.CompanyRequestBody {
	if o == nil {
		return nil
	}
	return o.CompanyRequestBody
}

func (o *UpdateCompanyRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type UpdateCompanyResponse struct {
	// OK
	Company     *shared.Company
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *UpdateCompanyResponse) GetCompany() *shared.Company {
	if o == nil {
		return nil
	}
	return o.Company
}

func (o *UpdateCompanyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateCompanyResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *UpdateCompanyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateCompanyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
