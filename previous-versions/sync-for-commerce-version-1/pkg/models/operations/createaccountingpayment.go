// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"net/http"
)

type CreateAccountingPaymentRequest struct {
	AccountingPayment *shared.AccountingPayment `request:"mediaType=application/json"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Time limit for the push operation to complete before it is timed out.
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (o *CreateAccountingPaymentRequest) GetAccountingPayment() *shared.AccountingPayment {
	if o == nil {
		return nil
	}
	return o.AccountingPayment
}

func (o *CreateAccountingPaymentRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateAccountingPaymentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateAccountingPaymentRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type CreateAccountingPaymentResponse struct {
	// Success
	AccountingCreatePaymentResponse *shared.AccountingCreatePaymentResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateAccountingPaymentResponse) GetAccountingCreatePaymentResponse() *shared.AccountingCreatePaymentResponse {
	if o == nil {
		return nil
	}
	return o.AccountingCreatePaymentResponse
}

func (o *CreateAccountingPaymentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAccountingPaymentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAccountingPaymentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
