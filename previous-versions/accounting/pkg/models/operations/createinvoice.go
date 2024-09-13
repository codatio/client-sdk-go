// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"net/http"
)

type CreateInvoiceRequest struct {
	Invoice *shared.Invoice `request:"mediaType=application/json"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Time limit for the push operation to complete before it is timed out.
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (o *CreateInvoiceRequest) GetInvoice() *shared.Invoice {
	if o == nil {
		return nil
	}
	return o.Invoice
}

func (o *CreateInvoiceRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateInvoiceRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateInvoiceRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type CreateInvoiceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	CreateInvoiceResponse *shared.CreateInvoiceResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateInvoiceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateInvoiceResponse) GetCreateInvoiceResponse() *shared.CreateInvoiceResponse {
	if o == nil {
		return nil
	}
	return o.CreateInvoiceResponse
}

func (o *CreateInvoiceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateInvoiceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
