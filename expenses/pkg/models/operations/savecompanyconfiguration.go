// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/expenses/pkg/models/shared"
	"net/http"
)

type SaveCompanyConfigurationRequest struct {
	CompanyConfiguration *shared.CompanyConfiguration `request:"mediaType=application/json"`
	CompanyID            string                       `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *SaveCompanyConfigurationRequest) GetCompanyConfiguration() *shared.CompanyConfiguration {
	if o == nil {
		return nil
	}
	return o.CompanyConfiguration
}

func (o *SaveCompanyConfigurationRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type SaveCompanyConfigurationResponse struct {
	// Success
	CompanyConfiguration *shared.CompanyConfiguration
	ContentType          string
	StatusCode           int
	RawResponse          *http.Response
	// The request made is not valid.
	Schema *shared.Schema
}

func (o *SaveCompanyConfigurationResponse) GetCompanyConfiguration() *shared.CompanyConfiguration {
	if o == nil {
		return nil
	}
	return o.CompanyConfiguration
}

func (o *SaveCompanyConfigurationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SaveCompanyConfigurationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SaveCompanyConfigurationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *SaveCompanyConfigurationResponse) GetSchema() *shared.Schema {
	if o == nil {
		return nil
	}
	return o.Schema
}
