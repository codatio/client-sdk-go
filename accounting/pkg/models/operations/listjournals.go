package operations

import (
	"net/http"
	"time"
)

type ListJournalsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListJournalsQueryParams struct {
	OrderBy  *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page     int     `queryParam:"style=form,explode=true,name=page"`
	PageSize *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string `queryParam:"style=form,explode=true,name=query"`
}

type ListJournalsRequest struct {
	PathParams  ListJournalsPathParams
	QueryParams ListJournalsQueryParams
}

type ListJournalsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListJournalsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListJournalsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListJournalsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListJournalsLinksLinks struct {
	Current  ListJournalsLinksLinksCurrent   `json:"current"`
	Next     *ListJournalsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListJournalsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListJournalsLinksLinksSelf      `json:"self"`
}

// ListJournalsLinksSourceModifiedDateMetadataMetadata
// Additional information about the entity
type ListJournalsLinksSourceModifiedDateMetadataMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type ListJournalsLinksSourceModifiedDateMetadata struct {
	Metadata *ListJournalsLinksSourceModifiedDateMetadataMetadata `json:"metadata,omitempty"`
}

type ListJournalsLinksSourceModifiedDateStatusEnum string

const (
	ListJournalsLinksSourceModifiedDateStatusEnumUnknown  ListJournalsLinksSourceModifiedDateStatusEnum = "Unknown"
	ListJournalsLinksSourceModifiedDateStatusEnumActive   ListJournalsLinksSourceModifiedDateStatusEnum = "Active"
	ListJournalsLinksSourceModifiedDateStatusEnumArchived ListJournalsLinksSourceModifiedDateStatusEnum = "Archived"
)

// ListJournalsLinksSourceModifiedDate
// > **Language tip:** For line items, or individual transactions, of a company's financial documents, refer to the [Journal entries](https://docs.codat.io/accounting-api#/schemas/JournalEntry) data type
//
// > View the coverage for journals in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journals" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// In accounting software, journals are used to record all the financial transactions of a company. Each transaction in a journal is represented by a separate [journal entry](https://docs.codat.io/accounting-api#/schemas/JournalEntry). These entries are used to create the general ledger, which is then used to create the financial statements of a business.
//
// When a company records all their transactions in a single journal, it can become large and difficult to maintain and track. This is why large companies often use multiple journals (also known as subjournals) to categorize and manage journal entries.
//
// Such journals can be divided into two categories:
//
// - Special journals: journals used to record specific types of transactions; for example, a purchases journal, a sales journal, or a cash management journal.
// - General journals: journals used to record transactions that fall outside the scope of the special journals.
//
// Multiple journals or subjournals are used in the following Codat integrations:
//
// - [Sage Intacct](https://docs.codat.io/integrations/accounting/sage-intacct/accounting-sage-intacct)  (mandatory)
// - [Exact Online](https://docs.codat.io/integrations/accounting/exact-online/accounting-exact-online)  (mandatory)
// - [Oracle NetSuite](https://docs.codat.io/integrations/accounting/netsuite/accounting-netsuite) (optional)
//
// > When pushing journal entries to an accounting platform that doesn’t support multiple journals (multi-book accounting), the entries will be linked to the platform-generic journal. The Journals data type will only include one object.
type ListJournalsLinksSourceModifiedDate struct {
	CreatedOn          *time.Time                                     `json:"createdOn,omitempty"`
	HasChildren        *bool                                          `json:"hasChildren,omitempty"`
	ID                 *string                                        `json:"id,omitempty"`
	JournalCode        *string                                        `json:"journalCode,omitempty"`
	Metadata           *ListJournalsLinksSourceModifiedDateMetadata   `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                     `json:"modifiedDate,omitempty"`
	Name               *string                                        `json:"name,omitempty"`
	ParentID           *string                                        `json:"parentId,omitempty"`
	SourceModifiedDate *time.Time                                     `json:"sourceModifiedDate,omitempty"`
	Status             *ListJournalsLinksSourceModifiedDateStatusEnum `json:"status,omitempty"`
	Type               *string                                        `json:"type,omitempty"`
}

// ListJournalsLinks
// Codat's Paging Model
type ListJournalsLinks struct {
	Links        ListJournalsLinksLinks                `json:"_links"`
	PageNumber   int64                                 `json:"pageNumber"`
	PageSize     int64                                 `json:"pageSize"`
	Results      []ListJournalsLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                 `json:"totalResults"`
}

type ListJournalsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListJournalsLinks
}
