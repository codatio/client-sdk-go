package operations

import (
	"net/http"
	"time"
)

type ListSupplierAttachmentsRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	SupplierID   string `pathParam:"style=simple,explode=false,name=supplierId"`
}

type ListSupplierAttachmentsAttachmentsAttachment struct {
	ContentType        *string    `json:"contentType,omitempty"`
	DateCreated        *time.Time `json:"dateCreated,omitempty"`
	FileSize           *int       `json:"fileSize,omitempty"`
	ID                 *string    `json:"id,omitempty"`
	IncludeWhenSent    *bool      `json:"includeWhenSent,omitempty"`
	ModifiedDate       *time.Time `json:"modifiedDate,omitempty"`
	Name               *string    `json:"name,omitempty"`
	SourceModifiedDate *time.Time `json:"sourceModifiedDate,omitempty"`
}

type ListSupplierAttachmentsAttachments struct {
	Attachments []ListSupplierAttachmentsAttachmentsAttachment `json:"attachments,omitempty"`
}

type ListSupplierAttachmentsResponse struct {
	Attachments *ListSupplierAttachmentsAttachments
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
