// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"net/http"
)

type UploadDirectIncomeAttachmentRequest struct {
	AttachmentUpload *shared.AttachmentUpload `request:"mediaType=multipart/form-data"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a direct income.
	DirectIncomeID string `pathParam:"style=simple,explode=false,name=directIncomeId"`
}

func (o *UploadDirectIncomeAttachmentRequest) GetAttachmentUpload() *shared.AttachmentUpload {
	if o == nil {
		return nil
	}
	return o.AttachmentUpload
}

func (o *UploadDirectIncomeAttachmentRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UploadDirectIncomeAttachmentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UploadDirectIncomeAttachmentRequest) GetDirectIncomeID() string {
	if o == nil {
		return ""
	}
	return o.DirectIncomeID
}

type UploadDirectIncomeAttachmentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UploadDirectIncomeAttachmentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UploadDirectIncomeAttachmentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UploadDirectIncomeAttachmentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
