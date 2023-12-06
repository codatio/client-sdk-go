// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v3/pkg/models/shared"
	"net/http"
)

type GetPaymentMethodRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a payment method.
	PaymentMethodID string `pathParam:"style=simple,explode=false,name=paymentMethodId"`
}

func (o *GetPaymentMethodRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetPaymentMethodRequest) GetPaymentMethodID() string {
	if o == nil {
		return ""
	}
	return o.PaymentMethodID
}

type GetPaymentMethodResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	PaymentMethod *shared.PaymentMethod
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetPaymentMethodResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPaymentMethodResponse) GetPaymentMethod() *shared.PaymentMethod {
	if o == nil {
		return nil
	}
	return o.PaymentMethod
}

func (o *GetPaymentMethodResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPaymentMethodResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
