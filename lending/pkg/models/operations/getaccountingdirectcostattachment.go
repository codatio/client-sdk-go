// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v2/pkg/models/shared"
	"net/http"
)

type GetAccountingDirectCostAttachmentRequest struct {
	// Unique identifier for an attachment
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a direct cost
	DirectCostID string `pathParam:"style=simple,explode=false,name=directCostId"`
}

func (o *GetAccountingDirectCostAttachmentRequest) GetAttachmentID() string {
	if o == nil {
		return ""
	}
	return o.AttachmentID
}

func (o *GetAccountingDirectCostAttachmentRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountingDirectCostAttachmentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetAccountingDirectCostAttachmentRequest) GetDirectCostID() string {
	if o == nil {
		return ""
	}
	return o.DirectCostID
}

type GetAccountingDirectCostAttachmentResponse struct {
	// Success
	AccountingAttachment *shared.AccountingAttachment
	ContentType          string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetAccountingDirectCostAttachmentResponse) GetAccountingAttachment() *shared.AccountingAttachment {
	if o == nil {
		return nil
	}
	return o.AccountingAttachment
}

func (o *GetAccountingDirectCostAttachmentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingDirectCostAttachmentResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetAccountingDirectCostAttachmentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingDirectCostAttachmentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
