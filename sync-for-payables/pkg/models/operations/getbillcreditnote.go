// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"net/http"
)

type GetBillCreditNoteRequest struct {
	BillCreditNoteID string `pathParam:"style=simple,explode=false,name=billCreditNoteId"`
	CompanyID        string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *GetBillCreditNoteRequest) GetBillCreditNoteID() string {
	if o == nil {
		return ""
	}
	return o.BillCreditNoteID
}

func (o *GetBillCreditNoteRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type GetBillCreditNoteResponse struct {
	// Success
	BillCreditNote *shared.BillCreditNote
	ContentType    string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetBillCreditNoteResponse) GetBillCreditNote() *shared.BillCreditNote {
	if o == nil {
		return nil
	}
	return o.BillCreditNote
}

func (o *GetBillCreditNoteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetBillCreditNoteResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetBillCreditNoteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetBillCreditNoteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
