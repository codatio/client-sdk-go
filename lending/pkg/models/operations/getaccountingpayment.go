// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"net/http"
)

type GetAccountingPaymentRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a payment.
	PaymentID string `pathParam:"style=simple,explode=false,name=paymentId"`
}

func (o *GetAccountingPaymentRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountingPaymentRequest) GetPaymentID() string {
	if o == nil {
		return ""
	}
	return o.PaymentID
}

type GetAccountingPaymentResponse struct {
	// Success
	AccountingPayment *shared.AccountingPayment
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAccountingPaymentResponse) GetAccountingPayment() *shared.AccountingPayment {
	if o == nil {
		return nil
	}
	return o.AccountingPayment
}

func (o *GetAccountingPaymentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingPaymentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingPaymentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
