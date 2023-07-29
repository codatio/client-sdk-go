// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
	"net/http"
)

type CreateBankAccountMappingRequest struct {
	BankFeedAccountMapping *shared.BankFeedAccountMapping `request:"mediaType=application/json"`
	CompanyID              string                         `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID           string                         `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *CreateBankAccountMappingRequest) GetBankFeedAccountMapping() *shared.BankFeedAccountMapping {
	if o == nil {
		return nil
	}
	return o.BankFeedAccountMapping
}

func (o *CreateBankAccountMappingRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateBankAccountMappingRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type CreateBankAccountMappingResponse struct {
	// Success
	AccountMappingResult *shared.AccountMappingResult
	ContentType          string
	// The request made is not valid.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *CreateBankAccountMappingResponse) GetAccountMappingResult() *shared.AccountMappingResult {
	if o == nil {
		return nil
	}
	return o.AccountMappingResult
}

func (o *CreateBankAccountMappingResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateBankAccountMappingResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *CreateBankAccountMappingResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateBankAccountMappingResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
