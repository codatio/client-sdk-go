// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
)

type ListBillsRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Retrieve the next page of results using the continuation token from the previous response.
	ContinuationToken *string `queryParam:"style=form,explode=true,name=continuationToken"`
	// Codat query string allows you to filter by `status` and `sourceModifiedDate`. Learn more about Codat's query string [here](https://docs.codat.io/using-the-api/querying). Platfrom specfic statuses: Xero supports  Open | PartiallyPaid | Paid | Void | Draft. Qbo supports Open | PartiallyPaid | Paid. FreeAgent supports Open | PartiallyPaid | Paid.
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

func (o *ListBillsRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *ListBillsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListBillsRequest) GetContinuationToken() *string {
	if o == nil {
		return nil
	}
	return o.ContinuationToken
}

func (o *ListBillsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

type ListBillsResponse struct {
	HTTPMeta shared.HTTPMetadata `json:"-"`
	// Success
	Bills *shared.Bills
}

func (o *ListBillsResponse) GetHTTPMeta() shared.HTTPMetadata {
	if o == nil {
		return shared.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListBillsResponse) GetBills() *shared.Bills {
	if o == nil {
		return nil
	}
	return o.Bills
}
