// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v7/pkg/models/shared"
	"net/http"
)

type GetAccountingCustomerAttachmentRequest struct {
	// Unique identifier for an attachment.
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a customer.
	CustomerID string `pathParam:"style=simple,explode=false,name=customerId"`
}

func (o *GetAccountingCustomerAttachmentRequest) GetAttachmentID() string {
	if o == nil {
		return ""
	}
	return o.AttachmentID
}

func (o *GetAccountingCustomerAttachmentRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountingCustomerAttachmentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetAccountingCustomerAttachmentRequest) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

type GetAccountingCustomerAttachmentResponse struct {
	// Success
	AccountingAttachment *shared.AccountingAttachment
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAccountingCustomerAttachmentResponse) GetAccountingAttachment() *shared.AccountingAttachment {
	if o == nil {
		return nil
	}
	return o.AccountingAttachment
}

func (o *GetAccountingCustomerAttachmentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingCustomerAttachmentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingCustomerAttachmentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
