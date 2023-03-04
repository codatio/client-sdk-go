package operations

import (
	"net/http"
	"time"
)

type GetDirectIncomeAttachmentPathParams struct {
	AttachmentID   string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID      string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID   string `pathParam:"style=simple,explode=false,name=connectionId"`
	DirectIncomeID string `pathParam:"style=simple,explode=false,name=directIncomeId"`
}

type GetDirectIncomeAttachmentQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type GetDirectIncomeAttachmentRequest struct {
	PathParams  GetDirectIncomeAttachmentPathParams
	QueryParams GetDirectIncomeAttachmentQueryParams
}

type GetDirectIncomeAttachmentAttachment struct {
	ContentType        *string    `json:"contentType,omitempty"`
	DateCreated        *time.Time `json:"dateCreated,omitempty"`
	FileSize           *int       `json:"fileSize,omitempty"`
	ID                 *string    `json:"id,omitempty"`
	IncludeWhenSent    *bool      `json:"includeWhenSent,omitempty"`
	ModifiedDate       *time.Time `json:"modifiedDate,omitempty"`
	Name               *string    `json:"name,omitempty"`
	SourceModifiedDate *time.Time `json:"sourceModifiedDate,omitempty"`
}

type GetDirectIncomeAttachmentResponse struct {
	Attachment  *GetDirectIncomeAttachmentAttachment
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
