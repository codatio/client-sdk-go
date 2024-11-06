// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v6/pkg/models/shared"
	"net/http"
)

type StartBankStatementUploadSessionRequest struct {
	StartUploadSessionRequest *shared.StartUploadSessionRequest `request:"mediaType=application/json"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *StartBankStatementUploadSessionRequest) GetStartUploadSessionRequest() *shared.StartUploadSessionRequest {
	if o == nil {
		return nil
	}
	return o.StartUploadSessionRequest
}

func (o *StartBankStatementUploadSessionRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *StartBankStatementUploadSessionRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type StartBankStatementUploadSessionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	PullOperation *shared.PullOperation
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *StartBankStatementUploadSessionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *StartBankStatementUploadSessionResponse) GetPullOperation() *shared.PullOperation {
	if o == nil {
		return nil
	}
	return o.PullOperation
}

func (o *StartBankStatementUploadSessionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *StartBankStatementUploadSessionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
