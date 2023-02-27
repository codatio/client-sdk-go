package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type GetBillAttachmentsPathParams struct {
	BillID       string `pathParam:"style=simple,explode=false,name=billId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetBillAttachmentsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetBillAttachmentsRequest struct {
	PathParams GetBillAttachmentsPathParams
	Security   GetBillAttachmentsSecurity
}

type GetBillAttachmentsAttachmentsAttachment struct {
	ContentType        *string    `json:"contentType,omitempty"`
	DateCreated        *time.Time `json:"dateCreated,omitempty"`
	FileSize           *int       `json:"fileSize,omitempty"`
	ID                 *string    `json:"id,omitempty"`
	IncludeWhenSent    *bool      `json:"includeWhenSent,omitempty"`
	ModifiedDate       *time.Time `json:"modifiedDate,omitempty"`
	Name               *string    `json:"name,omitempty"`
	SourceModifiedDate *time.Time `json:"sourceModifiedDate,omitempty"`
}

type GetBillAttachmentsAttachments struct {
	Attachments []GetBillAttachmentsAttachmentsAttachment `json:"attachments,omitempty"`
}

type GetBillAttachmentsResponse struct {
	Attachments *GetBillAttachmentsAttachments
	ContentType string
	StatusCode  int
}
