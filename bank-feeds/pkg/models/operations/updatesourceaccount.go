// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/bank-feeds/v6/pkg/models/shared"
	"net/http"
)

type UpdateSourceAccountRequest struct {
	SourceAccount *shared.SourceAccount `request:"mediaType=application/json"`
	// Unique identifier for an account.
	AccountID string `pathParam:"style=simple,explode=false,name=accountId"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *UpdateSourceAccountRequest) GetSourceAccount() *shared.SourceAccount {
	if o == nil {
		return nil
	}
	return o.SourceAccount
}

func (o *UpdateSourceAccountRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *UpdateSourceAccountRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UpdateSourceAccountRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type UpdateSourceAccountResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	SourceAccount *shared.SourceAccount
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateSourceAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSourceAccountResponse) GetSourceAccount() *shared.SourceAccount {
	if o == nil {
		return nil
	}
	return o.SourceAccount
}

func (o *UpdateSourceAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSourceAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
