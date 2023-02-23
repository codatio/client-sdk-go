package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type ListDirectIncomeAttachmentsPathParams struct {
	CompanyID      string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID   string `pathParam:"style=simple,explode=false,name=connectionId"`
	DirectIncomeID string `pathParam:"style=simple,explode=false,name=directIncomeId"`
}

type ListDirectIncomeAttachmentsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListDirectIncomeAttachmentsRequest struct {
	PathParams ListDirectIncomeAttachmentsPathParams
	Security   ListDirectIncomeAttachmentsSecurity
}

type ListDirectIncomeAttachmentsAttachmentsAttachment struct {
	ContentType        *string    `json:"contentType,omitempty"`
	DateCreated        *time.Time `json:"dateCreated,omitempty"`
	FileSize           *int       `json:"fileSize,omitempty"`
	ID                 *string    `json:"id,omitempty"`
	IncludeWhenSent    *bool      `json:"includeWhenSent,omitempty"`
	ModifiedDate       *time.Time `json:"modifiedDate,omitempty"`
	Name               *string    `json:"name,omitempty"`
	SourceModifiedDate *time.Time `json:"sourceModifiedDate,omitempty"`
}

type ListDirectIncomeAttachmentsAttachments struct {
	Attachments []ListDirectIncomeAttachmentsAttachmentsAttachment `json:"attachments,omitempty"`
}

type ListDirectIncomeAttachmentsResponse struct {
	Attachments *ListDirectIncomeAttachmentsAttachments
	ContentType string
	StatusCode  int
}
