// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/shared"
)

type ListSuppliersRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Retrieve the next page of results using the continuation token from the previous response.
	ContinuationToken *string `queryParam:"style=form,explode=true,name=continuationToken"`
	// Codat query string allows you to filter by `sourceModifiedDate` or if a supplier is `Active` or `Archived` in the accounting software. Learn more about Codat's query string [here](https://docs.codat.io/using-the-api/querying).
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

func (o *ListSuppliersRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *ListSuppliersRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListSuppliersRequest) GetContinuationToken() *string {
	if o == nil {
		return nil
	}
	return o.ContinuationToken
}

func (o *ListSuppliersRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

type ListSuppliersResponse struct {
	HTTPMeta shared.HTTPMetadata `json:"-"`
	// Success
	Suppliers *shared.Suppliers
}

func (o *ListSuppliersResponse) GetHTTPMeta() shared.HTTPMetadata {
	if o == nil {
		return shared.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListSuppliersResponse) GetSuppliers() *shared.Suppliers {
	if o == nil {
		return nil
	}
	return o.Suppliers
}
