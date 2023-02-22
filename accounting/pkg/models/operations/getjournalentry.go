package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetJournalEntryPathParams struct {
	CompanyID      string `pathParam:"style=simple,explode=false,name=companyId"`
	JournalEntryID string `pathParam:"style=simple,explode=false,name=journalEntryId"`
}

type GetJournalEntrySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetJournalEntryRequest struct {
	PathParams GetJournalEntryPathParams
	Security   GetJournalEntrySecurity
}

type GetJournalEntryResponse struct {
	ContentType  string
	JournalEntry *shared.JournalEntry
	StatusCode   int
}
