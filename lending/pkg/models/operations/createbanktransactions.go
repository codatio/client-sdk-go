// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v4/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v4/pkg/utils"
	"net/http"
)

type CreateBankTransactionsRequest struct {
	AccountingCreateBankTransactions *shared.AccountingCreateBankTransactions `request:"mediaType=application/json"`
	// Unique identifier for an account.
	AccountID string `pathParam:"style=simple,explode=false,name=accountId"`
	// Allow a sync upon push completion.
	AllowSyncOnPushComplete *bool `default:"true" queryParam:"style=form,explode=true,name=allowSyncOnPushComplete"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Time limit for the push operation to complete before it is timed out.
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (c CreateBankTransactionsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateBankTransactionsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateBankTransactionsRequest) GetAccountingCreateBankTransactions() *shared.AccountingCreateBankTransactions {
	if o == nil {
		return nil
	}
	return o.AccountingCreateBankTransactions
}

func (o *CreateBankTransactionsRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *CreateBankTransactionsRequest) GetAllowSyncOnPushComplete() *bool {
	if o == nil {
		return nil
	}
	return o.AllowSyncOnPushComplete
}

func (o *CreateBankTransactionsRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateBankTransactionsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateBankTransactionsRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type CreateBankTransactionsResponse struct {
	// Success
	AccountingCreateBankTransactionsResponse *shared.AccountingCreateBankTransactionsResponse
	// HTTP response content type for this operation
	ContentType string
	// The request made is not valid.
	ErrorMessage *shared.ErrorMessage
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateBankTransactionsResponse) GetAccountingCreateBankTransactionsResponse() *shared.AccountingCreateBankTransactionsResponse {
	if o == nil {
		return nil
	}
	return o.AccountingCreateBankTransactionsResponse
}

func (o *CreateBankTransactionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateBankTransactionsResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *CreateBankTransactionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateBankTransactionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
