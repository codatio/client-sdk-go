// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"net/http"
)

type CreateAccountingInvoiceRequest struct {
	AccountingInvoice *shared.AccountingInvoice `request:"mediaType=application/json"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Time limit for the push operation to complete before it is timed out.
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (o *CreateAccountingInvoiceRequest) GetAccountingInvoice() *shared.AccountingInvoice {
	if o == nil {
		return nil
	}
	return o.AccountingInvoice
}

func (o *CreateAccountingInvoiceRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateAccountingInvoiceRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateAccountingInvoiceRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type CreateAccountingInvoiceResponse struct {
	// Success
	AccountingCreateInvoiceResponse *shared.AccountingCreateInvoiceResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateAccountingInvoiceResponse) GetAccountingCreateInvoiceResponse() *shared.AccountingCreateInvoiceResponse {
	if o == nil {
		return nil
	}
	return o.AccountingCreateInvoiceResponse
}

func (o *CreateAccountingInvoiceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAccountingInvoiceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAccountingInvoiceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
