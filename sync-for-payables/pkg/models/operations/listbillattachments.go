// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
)

type ListBillAttachmentsRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a bill.
	BillID string `pathParam:"style=simple,explode=false,name=billId"`
}

func (o *ListBillAttachmentsRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *ListBillAttachmentsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListBillAttachmentsRequest) GetBillID() string {
	if o == nil {
		return ""
	}
	return o.BillID
}

type ListBillAttachmentsResponse struct {
	HTTPMeta shared.HTTPMetadata `json:"-"`
	// Success
	Attachment *shared.Attachment
}

func (o *ListBillAttachmentsResponse) GetHTTPMeta() shared.HTTPMetadata {
	if o == nil {
		return shared.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListBillAttachmentsResponse) GetAttachment() *shared.Attachment {
	if o == nil {
		return nil
	}
	return o.Attachment
}