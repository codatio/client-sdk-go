package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type ListDirectCostAttachmentsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	DirectCostID string `pathParam:"style=simple,explode=false,name=directCostId"`
}

type ListDirectCostAttachmentsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListDirectCostAttachmentsRequest struct {
	PathParams ListDirectCostAttachmentsPathParams
	Security   ListDirectCostAttachmentsSecurity
}

type ListDirectCostAttachmentsAttachmentsAttachment struct {
	ContentType        *string    `json:"contentType,omitempty"`
	DateCreated        *time.Time `json:"dateCreated,omitempty"`
	FileSize           *int       `json:"fileSize,omitempty"`
	ID                 *string    `json:"id,omitempty"`
	IncludeWhenSent    *bool      `json:"includeWhenSent,omitempty"`
	ModifiedDate       *time.Time `json:"modifiedDate,omitempty"`
	Name               *string    `json:"name,omitempty"`
	SourceModifiedDate *time.Time `json:"sourceModifiedDate,omitempty"`
}

type ListDirectCostAttachmentsAttachments struct {
	Attachments []ListDirectCostAttachmentsAttachmentsAttachment `json:"attachments,omitempty"`
}

type ListDirectCostAttachmentsResponse struct {
	Attachments *ListDirectCostAttachmentsAttachments
	ContentType string
	StatusCode  int
}