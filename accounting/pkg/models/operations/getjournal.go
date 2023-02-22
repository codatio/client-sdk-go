package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetJournalPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	JournalID string `pathParam:"style=simple,explode=false,name=journalId"`
}

type GetJournalSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetJournalRequest struct {
	PathParams GetJournalPathParams
	Security   GetJournalSecurity
}

type GetJournalResponse struct {
	ContentType string
	Journal     *shared.Journal
	StatusCode  int
}
