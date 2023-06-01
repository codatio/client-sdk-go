// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type CreateBillPaymentRequest struct {
	BillPayment      *shared.BillPayment `request:"mediaType=application/json"`
	CompanyID        string              `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string              `pathParam:"style=simple,explode=false,name=connectionId"`
	TimeoutInMinutes *int                `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type CreateBillPaymentResponse struct {
	ContentType string
	// Success
	CreateBillPaymentResponse *shared.CreateBillPaymentResponse
	StatusCode                int
	RawResponse               *http.Response
	// The request made is not valid.
	Schema *shared.Schema
}
