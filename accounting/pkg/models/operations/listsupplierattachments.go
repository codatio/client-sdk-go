// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type ListSupplierAttachmentsRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a supplier
	SupplierID string `pathParam:"style=simple,explode=false,name=supplierId"`
}

func (o *ListSupplierAttachmentsRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *ListSupplierAttachmentsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListSupplierAttachmentsRequest) GetSupplierID() string {
	if o == nil {
		return ""
	}
	return o.SupplierID
}

type ListSupplierAttachmentsResponse struct {
	// Success
	AttachmentsDataset *shared.AttachmentsDataset
	ContentType        string
	StatusCode         int
	RawResponse        *http.Response
	// Your API request was not properly authorized.
	Schema *shared.Schema
}

func (o *ListSupplierAttachmentsResponse) GetAttachmentsDataset() *shared.AttachmentsDataset {
	if o == nil {
		return nil
	}
	return o.AttachmentsDataset
}

func (o *ListSupplierAttachmentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListSupplierAttachmentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListSupplierAttachmentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListSupplierAttachmentsResponse) GetSchema() *shared.Schema {
	if o == nil {
		return nil
	}
	return o.Schema
}
