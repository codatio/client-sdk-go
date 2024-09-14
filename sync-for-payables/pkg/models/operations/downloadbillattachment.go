// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
	"io"
)

type DownloadBillAttachmentRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a bill.
	BillID string `pathParam:"style=simple,explode=false,name=billId"`
	// Unique identifier for an attachment.
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
}

func (o *DownloadBillAttachmentRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *DownloadBillAttachmentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *DownloadBillAttachmentRequest) GetBillID() string {
	if o == nil {
		return ""
	}
	return o.BillID
}

func (o *DownloadBillAttachmentRequest) GetAttachmentID() string {
	if o == nil {
		return ""
	}
	return o.AttachmentID
}

type DownloadBillAttachmentResponse struct {
	HTTPMeta shared.HTTPMetadata `json:"-"`
	// Success
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	Data io.ReadCloser
}

func (o *DownloadBillAttachmentResponse) GetHTTPMeta() shared.HTTPMetadata {
	if o == nil {
		return shared.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DownloadBillAttachmentResponse) GetData() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.Data
}