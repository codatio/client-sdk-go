// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type UpdateInvoiceRequest struct {
	Invoice      *shared.Invoice `request:"mediaType=application/json"`
	CompanyID    string          `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string          `pathParam:"style=simple,explode=false,name=connectionId"`
	// When updating data in the destination platform Codat checks the `sourceModifiedDate` against the `lastupdated` date from the accounting platform, if they're different Codat will return an error suggesting you should initiate another pull of the data. If this is set to `true` then the update will override this check.
	ForceUpdate *bool `queryParam:"style=form,explode=true,name=forceUpdate"`
	// Unique identifier for an invoice
	InvoiceID        string `pathParam:"style=simple,explode=false,name=invoiceId"`
	TimeoutInMinutes *int   `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type UpdateInvoiceResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	UpdateInvoiceResponse *shared.UpdateInvoiceResponse
	// The request made is not valid.
	Schema *shared.Schema
}
