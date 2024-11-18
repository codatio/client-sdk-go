// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v7/pkg/models/shared"
	"net/http"
)

type GetAccountingCustomerRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a customer.
	CustomerID string `pathParam:"style=simple,explode=false,name=customerId"`
}

func (o *GetAccountingCustomerRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountingCustomerRequest) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

type GetAccountingCustomerResponse struct {
	// Success
	AccountingCustomer *shared.AccountingCustomer
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAccountingCustomerResponse) GetAccountingCustomer() *shared.AccountingCustomer {
	if o == nil {
		return nil
	}
	return o.AccountingCustomer
}

func (o *GetAccountingCustomerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingCustomerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingCustomerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
