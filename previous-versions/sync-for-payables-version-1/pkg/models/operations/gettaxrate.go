// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v3/pkg/models/shared"
	"net/http"
)

type GetTaxRateRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a tax rate.
	TaxRateID string `pathParam:"style=simple,explode=false,name=taxRateId"`
}

func (o *GetTaxRateRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetTaxRateRequest) GetTaxRateID() string {
	if o == nil {
		return ""
	}
	return o.TaxRateID
}

type GetTaxRateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	TaxRate *shared.TaxRate
}

func (o *GetTaxRateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTaxRateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTaxRateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTaxRateResponse) GetTaxRate() *shared.TaxRate {
	if o == nil {
		return nil
	}
	return o.TaxRate
}
