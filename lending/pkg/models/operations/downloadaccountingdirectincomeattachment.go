// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"io"
	"net/http"
)

type DownloadAccountingDirectIncomeAttachmentRequest struct {
	// Unique identifier for an attachment.
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a direct income.
	DirectIncomeID string `pathParam:"style=simple,explode=false,name=directIncomeId"`
}

func (o *DownloadAccountingDirectIncomeAttachmentRequest) GetAttachmentID() string {
	if o == nil {
		return ""
	}
	return o.AttachmentID
}

func (o *DownloadAccountingDirectIncomeAttachmentRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *DownloadAccountingDirectIncomeAttachmentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *DownloadAccountingDirectIncomeAttachmentRequest) GetDirectIncomeID() string {
	if o == nil {
		return ""
	}
	return o.DirectIncomeID
}

type DownloadAccountingDirectIncomeAttachmentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	Data io.ReadCloser
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DownloadAccountingDirectIncomeAttachmentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DownloadAccountingDirectIncomeAttachmentResponse) GetData() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *DownloadAccountingDirectIncomeAttachmentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DownloadAccountingDirectIncomeAttachmentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
