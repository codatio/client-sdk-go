package operations

import (
	"net/http"
)

type UploadAttachmentPathParams struct {
	CompanyID     string `pathParam:"style=simple,explode=false,name=companyId"`
	SyncID        string `pathParam:"style=simple,explode=false,name=syncId"`
	TransactionID string `pathParam:"style=simple,explode=false,name=transactionId"`
}

type UploadAttachmentRequest struct {
	PathParams UploadAttachmentPathParams
}

type UploadAttachment200ApplicationJSON struct {
	CompanyID     *string `json:"companyId,omitempty"`
	ID            *string `json:"id,omitempty"`
	TransactionID *string `json:"transactionId,omitempty"`
}

type UploadAttachmentResponse struct {
	ContentType                              string
	StatusCode                               int
	RawResponse                              *http.Response
	UploadAttachment200ApplicationJSONObject *UploadAttachment200ApplicationJSON
}
