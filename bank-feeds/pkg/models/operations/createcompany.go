// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/bank-feeds/v2/pkg/models/shared"
	"net/http"
)

type CreateCompanyResponse struct {
	// OK
	Company     *shared.Company
	ContentType string
	// The request made is not valid.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *CreateCompanyResponse) GetCompany() *shared.Company {
	if o == nil {
		return nil
	}
	return o.Company
}

func (o *CreateCompanyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateCompanyResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *CreateCompanyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateCompanyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
