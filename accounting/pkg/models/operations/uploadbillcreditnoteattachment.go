// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type UploadBillCreditNoteAttachmentRequestBody struct {
	Content     []byte `multipartForm:"content"`
	RequestBody string `multipartForm:"name=requestBody"`
}

func (o *UploadBillCreditNoteAttachmentRequestBody) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

func (o *UploadBillCreditNoteAttachmentRequestBody) GetRequestBody() string {
	if o == nil {
		return ""
	}
	return o.RequestBody
}

type UploadBillCreditNoteAttachmentRequest struct {
	RequestBody      *UploadBillCreditNoteAttachmentRequestBody `multipartForm:"file" request:"mediaType=multipart/form-data"`
	BillCreditNoteID string                                     `pathParam:"style=simple,explode=false,name=billCreditNoteId"`
	CompanyID        string                                     `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string                                     `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *UploadBillCreditNoteAttachmentRequest) GetRequestBody() *UploadBillCreditNoteAttachmentRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UploadBillCreditNoteAttachmentRequest) GetBillCreditNoteID() string {
	if o == nil {
		return ""
	}
	return o.BillCreditNoteID
}

func (o *UploadBillCreditNoteAttachmentRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UploadBillCreditNoteAttachmentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type UploadBillCreditNoteAttachmentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Your API request was not properly authorized.
	Schema *shared.Schema
}

func (o *UploadBillCreditNoteAttachmentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UploadBillCreditNoteAttachmentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UploadBillCreditNoteAttachmentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UploadBillCreditNoteAttachmentResponse) GetSchema() *shared.Schema {
	if o == nil {
		return nil
	}
	return o.Schema
}
