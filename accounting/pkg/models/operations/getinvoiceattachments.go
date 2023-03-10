package operations

import (
	"net/http"
	"time"
)

type GetInvoiceAttachmentsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	InvoiceID    string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

type GetInvoiceAttachmentsRequest struct {
	PathParams GetInvoiceAttachmentsPathParams
}

type GetInvoiceAttachmentsAttachmentsAttachment struct {
	ContentType        *string    `json:"contentType,omitempty"`
	DateCreated        *time.Time `json:"dateCreated,omitempty"`
	FileSize           *int       `json:"fileSize,omitempty"`
	ID                 *string    `json:"id,omitempty"`
	IncludeWhenSent    *bool      `json:"includeWhenSent,omitempty"`
	ModifiedDate       *time.Time `json:"modifiedDate,omitempty"`
	Name               *string    `json:"name,omitempty"`
	SourceModifiedDate *time.Time `json:"sourceModifiedDate,omitempty"`
}

type GetInvoiceAttachmentsAttachments struct {
	Attachments []GetInvoiceAttachmentsAttachmentsAttachment `json:"attachments,omitempty"`
}

type GetInvoiceAttachmentsResponse struct {
	Attachments *GetInvoiceAttachmentsAttachments
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
