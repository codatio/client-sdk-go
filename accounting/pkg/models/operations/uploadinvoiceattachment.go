// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type UploadInvoiceAttachmentRequestBody struct {
	Content     []byte `multipartForm:"content"`
	RequestBody string `multipartForm:"name=requestBody"`
}

func (o *UploadInvoiceAttachmentRequestBody) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

func (o *UploadInvoiceAttachmentRequestBody) GetRequestBody() string {
	if o == nil {
		return ""
	}
	return o.RequestBody
}

type UploadInvoiceAttachmentRequest struct {
	RequestBody  *UploadInvoiceAttachmentRequestBody `multipartForm:"file" request:"mediaType=multipart/form-data"`
	CompanyID    string                              `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string                              `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for an invoice
	InvoiceID string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

func (o *UploadInvoiceAttachmentRequest) GetRequestBody() *UploadInvoiceAttachmentRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UploadInvoiceAttachmentRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UploadInvoiceAttachmentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UploadInvoiceAttachmentRequest) GetInvoiceID() string {
	if o == nil {
		return ""
	}
	return o.InvoiceID
}

type UploadInvoiceAttachmentResponse struct {
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *UploadInvoiceAttachmentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UploadInvoiceAttachmentResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *UploadInvoiceAttachmentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UploadInvoiceAttachmentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
