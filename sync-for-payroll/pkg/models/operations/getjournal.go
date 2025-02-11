// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/shared"
	"net/http"
)

type GetJournalRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a journal.
	JournalID string `pathParam:"style=simple,explode=false,name=journalId"`
}

func (o *GetJournalRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetJournalRequest) GetJournalID() string {
	if o == nil {
		return ""
	}
	return o.JournalID
}

type GetJournalResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	Journal *shared.Journal
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetJournalResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetJournalResponse) GetJournal() *shared.Journal {
	if o == nil {
		return nil
	}
	return o.Journal
}

func (o *GetJournalResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetJournalResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
