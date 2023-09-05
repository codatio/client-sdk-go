// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
	"net/http"
)

type GetTaxRateRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
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
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
	// Success
	TaxRate *shared.TaxRate
}

func (o *GetTaxRateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTaxRateResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
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