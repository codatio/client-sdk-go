// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v3/pkg/models/shared"
)

type UpdateCompanyRequest struct {
	// Unique identifier for a company.
	CompanyID          string                     `pathParam:"style=simple,explode=false,name=companyId"`
	CompanyRequestBody *shared.CompanyRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateCompanyRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UpdateCompanyRequest) GetCompanyRequestBody() *shared.CompanyRequestBody {
	if o == nil {
		return nil
	}
	return o.CompanyRequestBody
}

type UpdateCompanyResponse struct {
	HTTPMeta shared.HTTPMetadata `json:"-"`
	// OK
	Company *shared.Company
}

func (o *UpdateCompanyResponse) GetHTTPMeta() shared.HTTPMetadata {
	if o == nil {
		return shared.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateCompanyResponse) GetCompany() *shared.Company {
	if o == nil {
		return nil
	}
	return o.Company
}
