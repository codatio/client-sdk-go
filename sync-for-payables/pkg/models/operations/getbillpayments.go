// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"net/http"
)

type GetBillPaymentsRequest struct {
	// Unique identifier for a bill payment.
	BillPaymentID string `pathParam:"style=simple,explode=false,name=billPaymentId"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *GetBillPaymentsRequest) GetBillPaymentID() string {
	if o == nil {
		return ""
	}
	return o.BillPaymentID
}

func (o *GetBillPaymentsRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type GetBillPaymentsResponse struct {
	// Success
	BillPayment *shared.BillPayment
	// HTTP response content type for this operation
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetBillPaymentsResponse) GetBillPayment() *shared.BillPayment {
	if o == nil {
		return nil
	}
	return o.BillPayment
}

func (o *GetBillPaymentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetBillPaymentsResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetBillPaymentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetBillPaymentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
