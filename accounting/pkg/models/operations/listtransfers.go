package operations

import (
	"net/http"
	"time"
)

type ListTransfersPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListTransfersQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListTransfersRequest struct {
	PathParams  ListTransfersPathParams
	QueryParams ListTransfersQueryParams
}

type ListTransfersLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListTransfersLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListTransfersLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListTransfersLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListTransfersLinksLinks struct {
	Current  ListTransfersLinksLinksCurrent   `json:"current"`
	Next     *ListTransfersLinksLinksNext     `json:"next,omitempty"`
	Previous *ListTransfersLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListTransfersLinksLinksSelf      `json:"self"`
}

// ListTransfersLinksSourceModifiedDateContactRef
// The customer or supplier for the transfer, if available.
type ListTransfersLinksSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// ListTransfersLinksSourceModifiedDateTransferAccountAccountRef
// The account that the transfer is moving from or to.
type ListTransfersLinksSourceModifiedDateTransferAccountAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// ListTransfersLinksSourceModifiedDateTransferAccount
// The details of the accounts the transfer is moving from.
type ListTransfersLinksSourceModifiedDateTransferAccount struct {
	AccountRef *ListTransfersLinksSourceModifiedDateTransferAccountAccountRef `json:"accountRef,omitempty"`
	Amount     *float64                                                       `json:"amount,omitempty"`
	Currency   *string                                                        `json:"currency,omitempty"`
}

type ListTransfersLinksSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type ListTransfersLinksSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

type ListTransfersLinksSourceModifiedDateTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// ListTransfersLinksSourceModifiedDate
// > View the coverage for transfers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=transfers" target="_blank">Data coverage explorer</a>.
//
// From the **Transfers** endpoints, you can:
//
// - [Retrieve a list of all transfers for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers)
// - [Retrieve a single transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers__transferId_)
// - [Add a new transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/post_companies__companyId__connections__connectionId__push_transfers)
//
// **Transfers** is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
type ListTransfersLinksSourceModifiedDate struct {
	ContactRef           *ListTransfersLinksSourceModifiedDateContactRef            `json:"contactRef,omitempty"`
	Date                 *time.Time                                                 `json:"date,omitempty"`
	DepositedRecordRefs  []string                                                   `json:"depositedRecordRefs,omitempty"`
	Description          *string                                                    `json:"description,omitempty"`
	From                 *ListTransfersLinksSourceModifiedDateTransferAccount       `json:"from,omitempty"`
	ID                   *string                                                    `json:"id,omitempty"`
	Metadata             *ListTransfersLinksSourceModifiedDateMetadata              `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                                 `json:"modifiedDate,omitempty"`
	SourceModifiedDate   *time.Time                                                 `json:"sourceModifiedDate,omitempty"`
	SupplementalData     *ListTransfersLinksSourceModifiedDateSupplementalData      `json:"supplementalData,omitempty"`
	To                   *ListTransfersLinksSourceModifiedDateTransferAccount       `json:"to,omitempty"`
	TrackingCategoryRefs []ListTransfersLinksSourceModifiedDateTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
}

// ListTransfersLinks
// Codat's Paging Model
type ListTransfersLinks struct {
	Links        ListTransfersLinksLinks                `json:"_links"`
	PageNumber   int64                                  `json:"pageNumber"`
	PageSize     int64                                  `json:"pageSize"`
	Results      []ListTransfersLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                  `json:"totalResults"`
}

type ListTransfersResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListTransfersLinks
}
