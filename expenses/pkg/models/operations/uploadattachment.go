package operations

import (
	"github.com/codatio/client-sdk-go/expenses/pkg/models/shared"
)

type UploadAttachmentPathParams struct {
	CompanyID     string `pathParam:"style=simple,explode=false,name=companyId"`
	SyncID        string `pathParam:"style=simple,explode=false,name=syncId"`
	TransactionID string `pathParam:"style=simple,explode=false,name=transactionId"`
}

type UploadAttachmentRequest struct {
	PathParams UploadAttachmentPathParams
}

type UploadAttachmentResponse struct {
	ContentType string
	StatusCode  int
	Attachment  *shared.Attachment
}
