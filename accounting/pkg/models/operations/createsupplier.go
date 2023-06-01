// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type CreateSupplierRequest struct {
	Supplier         *shared.Supplier `request:"mediaType=application/json"`
	CompanyID        string           `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string           `pathParam:"style=simple,explode=false,name=connectionId"`
	TimeoutInMinutes *int             `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type CreateSupplierResponse struct {
	ContentType string
	// Success
	CreateSupplierResponse *shared.CreateSupplierResponse
	StatusCode             int
	RawResponse            *http.Response
	// The request made is not valid.
	Schema *shared.Schema
}
