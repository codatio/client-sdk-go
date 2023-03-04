package operations

import (
	"net/http"
	"time"
)

type GetCustomerAttachmentsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	CustomerID   string `pathParam:"style=simple,explode=false,name=customerId"`
}

type GetCustomerAttachmentsRequest struct {
	PathParams GetCustomerAttachmentsPathParams
}

type GetCustomerAttachmentsAttachmentsAttachment struct {
	ContentType        *string    `json:"contentType,omitempty"`
	DateCreated        *time.Time `json:"dateCreated,omitempty"`
	FileSize           *int       `json:"fileSize,omitempty"`
	ID                 *string    `json:"id,omitempty"`
	IncludeWhenSent    *bool      `json:"includeWhenSent,omitempty"`
	ModifiedDate       *time.Time `json:"modifiedDate,omitempty"`
	Name               *string    `json:"name,omitempty"`
	SourceModifiedDate *time.Time `json:"sourceModifiedDate,omitempty"`
}

type GetCustomerAttachmentsAttachments struct {
	Attachments []GetCustomerAttachmentsAttachmentsAttachment `json:"attachments,omitempty"`
}

type GetCustomerAttachmentsResponse struct {
	Attachments *GetCustomerAttachmentsAttachments
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
