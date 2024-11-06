// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v6/pkg/models/shared"
	"net/http"
)

type GetAccountingSupplierRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a supplier.
	SupplierID string `pathParam:"style=simple,explode=false,name=supplierId"`
}

func (o *GetAccountingSupplierRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountingSupplierRequest) GetSupplierID() string {
	if o == nil {
		return ""
	}
	return o.SupplierID
}

type GetAccountingSupplierResponse struct {
	// Success
	AccountingSupplier *shared.AccountingSupplier
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAccountingSupplierResponse) GetAccountingSupplier() *shared.AccountingSupplier {
	if o == nil {
		return nil
	}
	return o.AccountingSupplier
}

func (o *GetAccountingSupplierResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingSupplierResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingSupplierResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
