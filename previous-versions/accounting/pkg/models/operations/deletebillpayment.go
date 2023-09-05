// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"net/http"
)

type DeleteBillPaymentRequest struct {
	BillPaymentID string `pathParam:"style=simple,explode=false,name=billPaymentId"`
	CompanyID     string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID  string `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *DeleteBillPaymentRequest) GetBillPaymentID() string {
	if o == nil {
		return ""
	}
	return o.BillPaymentID
}

func (o *DeleteBillPaymentRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *DeleteBillPaymentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type DeleteBillPaymentResponse struct {
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	// OK
	PushOperationSummary *shared.PushOperationSummary
	StatusCode           int
	RawResponse          *http.Response
}

func (o *DeleteBillPaymentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteBillPaymentResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *DeleteBillPaymentResponse) GetPushOperationSummary() *shared.PushOperationSummary {
	if o == nil {
		return nil
	}
	return o.PushOperationSummary
}

func (o *DeleteBillPaymentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteBillPaymentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}