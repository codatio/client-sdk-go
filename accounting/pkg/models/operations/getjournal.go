// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type GetJournalRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
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

// GetJournal409ApplicationJSON - The data type's dataset has not been requested or is still syncing.
type GetJournal409ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

func (o *GetJournal409ApplicationJSON) GetCanBeRetried() *string {
	if o == nil {
		return nil
	}
	return o.CanBeRetried
}

func (o *GetJournal409ApplicationJSON) GetCorrelationID() *string {
	if o == nil {
		return nil
	}
	return o.CorrelationID
}

func (o *GetJournal409ApplicationJSON) GetDetailedErrorCode() *int64 {
	if o == nil {
		return nil
	}
	return o.DetailedErrorCode
}

func (o *GetJournal409ApplicationJSON) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetJournal409ApplicationJSON) GetService() *string {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *GetJournal409ApplicationJSON) GetStatusCode() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

type GetJournalResponse struct {
	ContentType string
	// Success
	Journal     *shared.Journal
	StatusCode  int
	RawResponse *http.Response
	// The data type's dataset has not been requested or is still syncing.
	GetJournal409ApplicationJSONObject *GetJournal409ApplicationJSON
	// Your API request was not properly authorized.
	Schema *shared.Schema
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

func (o *GetJournalResponse) GetGetJournal409ApplicationJSONObject() *GetJournal409ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetJournal409ApplicationJSONObject
}

func (o *GetJournalResponse) GetSchema() *shared.Schema {
	if o == nil {
		return nil
	}
	return o.Schema
}
