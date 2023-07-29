// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type UpdateCreditNoteRequest struct {
	CreditNote   *shared.CreditNote `request:"mediaType=application/json"`
	CompanyID    string             `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string             `pathParam:"style=simple,explode=false,name=connectionId"`
	CreditNoteID string             `pathParam:"style=simple,explode=false,name=creditNoteId"`
	// When updating data in the destination platform Codat checks the `sourceModifiedDate` against the `lastupdated` date from the accounting platform, if they're different Codat will return an error suggesting you should initiate another pull of the data. If this is set to `true` then the update will override this check.
	ForceUpdate      *bool `queryParam:"style=form,explode=true,name=forceUpdate"`
	TimeoutInMinutes *int  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (o *UpdateCreditNoteRequest) GetCreditNote() *shared.CreditNote {
	if o == nil {
		return nil
	}
	return o.CreditNote
}

func (o *UpdateCreditNoteRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UpdateCreditNoteRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateCreditNoteRequest) GetCreditNoteID() string {
	if o == nil {
		return ""
	}
	return o.CreditNoteID
}

func (o *UpdateCreditNoteRequest) GetForceUpdate() *bool {
	if o == nil {
		return nil
	}
	return o.ForceUpdate
}

func (o *UpdateCreditNoteRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type UpdateCreditNoteResponse struct {
	ContentType string
	// The request made is not valid.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
	// Success
	UpdateCreditNoteResponse *shared.UpdateCreditNoteResponse
}

func (o *UpdateCreditNoteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateCreditNoteResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *UpdateCreditNoteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateCreditNoteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateCreditNoteResponse) GetUpdateCreditNoteResponse() *shared.UpdateCreditNoteResponse {
	if o == nil {
		return nil
	}
	return o.UpdateCreditNoteResponse
}
