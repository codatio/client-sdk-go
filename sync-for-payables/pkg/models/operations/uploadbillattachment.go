// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/shared"
)

type UploadBillAttachmentRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a bill.
	BillID           string                   `pathParam:"style=simple,explode=false,name=billId"`
	AttachmentUpload *shared.AttachmentUpload `request:"mediaType=multipart/form-data"`
}

func (o *UploadBillAttachmentRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UploadBillAttachmentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UploadBillAttachmentRequest) GetBillID() string {
	if o == nil {
		return ""
	}
	return o.BillID
}

func (o *UploadBillAttachmentRequest) GetAttachmentUpload() *shared.AttachmentUpload {
	if o == nil {
		return nil
	}
	return o.AttachmentUpload
}

type UploadBillAttachmentResponse struct {
	HTTPMeta shared.HTTPMetadata `json:"-"`
	// Created
	Attachment *shared.Attachment
}

func (o *UploadBillAttachmentResponse) GetHTTPMeta() shared.HTTPMetadata {
	if o == nil {
		return shared.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UploadBillAttachmentResponse) GetAttachment() *shared.Attachment {
	if o == nil {
		return nil
	}
	return o.Attachment
}
