// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/commerce/pkg/models/shared"
	"net/http"
)

type GetPaymentMethodRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a payment method.
	PaymentMethodID string `pathParam:"style=simple,explode=false,name=paymentMethodId"`
}

func (o *GetPaymentMethodRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetPaymentMethodRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetPaymentMethodRequest) GetPaymentMethodID() string {
	if o == nil {
		return ""
	}
	return o.PaymentMethodID
}

type GetPaymentMethodResponse struct {
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	// OK
	PaymentMethod *shared.PaymentMethod
	StatusCode    int
	RawResponse   *http.Response
}

func (o *GetPaymentMethodResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPaymentMethodResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
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
