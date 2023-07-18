// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
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

// GetTaxRate409ApplicationJSON - The data type's dataset has not been requested or is still syncing.
type GetTaxRate409ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

func (o *GetTaxRate409ApplicationJSON) GetCanBeRetried() *string {
	if o == nil {
		return nil
	}
	return o.CanBeRetried
}

func (o *GetTaxRate409ApplicationJSON) GetCorrelationID() *string {
	if o == nil {
		return nil
	}
	return o.CorrelationID
}

func (o *GetTaxRate409ApplicationJSON) GetDetailedErrorCode() *int64 {
	if o == nil {
		return nil
	}
	return o.DetailedErrorCode
}

func (o *GetTaxRate409ApplicationJSON) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetTaxRate409ApplicationJSON) GetService() *string {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *GetTaxRate409ApplicationJSON) GetStatusCode() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

type GetTaxRateResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	TaxRate *shared.TaxRate
	// The data type's dataset has not been requested or is still syncing.
	GetTaxRate409ApplicationJSONObject *GetTaxRate409ApplicationJSON
	// Your API request was not properly authorized.
	Schema *shared.Schema
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

func (o *GetTaxRateResponse) GetGetTaxRate409ApplicationJSONObject() *GetTaxRate409ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetTaxRate409ApplicationJSONObject
}

func (o *GetTaxRateResponse) GetSchema() *shared.Schema {
	if o == nil {
		return nil
	}
	return o.Schema
}
