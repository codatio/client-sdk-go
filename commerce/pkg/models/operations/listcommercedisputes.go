package operations

import (
	"time"
)

type ListCommerceDisputesPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListCommerceDisputesQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListCommerceDisputesRequest struct {
	PathParams  ListCommerceDisputesPathParams
	QueryParams ListCommerceDisputesQueryParams
}

type ListCommerceDisputesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCommerceDisputesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceDisputesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceDisputesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCommerceDisputesLinksLinks struct {
	Current  ListCommerceDisputesLinksLinksCurrent   `json:"current"`
	Next     *ListCommerceDisputesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCommerceDisputesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCommerceDisputesLinksLinksSelf      `json:"self"`
}

type ListCommerceDisputesLinksSourceModifiedDateRecordRef struct {
	ID   string      `json:"id"`
	Type interface{} `json:"type"`
}

type ListCommerceDisputesLinksSourceModifiedDateStatusEnum string

const (
	ListCommerceDisputesLinksSourceModifiedDateStatusEnumWon                     ListCommerceDisputesLinksSourceModifiedDateStatusEnum = "Won"
	ListCommerceDisputesLinksSourceModifiedDateStatusEnumLost                    ListCommerceDisputesLinksSourceModifiedDateStatusEnum = "Lost"
	ListCommerceDisputesLinksSourceModifiedDateStatusEnumAccepted                ListCommerceDisputesLinksSourceModifiedDateStatusEnum = "Accepted"
	ListCommerceDisputesLinksSourceModifiedDateStatusEnumProcessing              ListCommerceDisputesLinksSourceModifiedDateStatusEnum = "Processing"
	ListCommerceDisputesLinksSourceModifiedDateStatusEnumChargeRefunded          ListCommerceDisputesLinksSourceModifiedDateStatusEnum = "ChargeRefunded"
	ListCommerceDisputesLinksSourceModifiedDateStatusEnumEvidenceRequired        ListCommerceDisputesLinksSourceModifiedDateStatusEnum = "EvidenceRequired"
	ListCommerceDisputesLinksSourceModifiedDateStatusEnumInquiryEvidenceRequired ListCommerceDisputesLinksSourceModifiedDateStatusEnum = "InquiryEvidenceRequired"
	ListCommerceDisputesLinksSourceModifiedDateStatusEnumInquiryProcessing       ListCommerceDisputesLinksSourceModifiedDateStatusEnum = "InquiryProcessing"
	ListCommerceDisputesLinksSourceModifiedDateStatusEnumInquiryClosed           ListCommerceDisputesLinksSourceModifiedDateStatusEnum = "InquiryClosed"
	ListCommerceDisputesLinksSourceModifiedDateStatusEnumWaitingThirdParty       ListCommerceDisputesLinksSourceModifiedDateStatusEnum = "WaitingThirdParty"
	ListCommerceDisputesLinksSourceModifiedDateStatusEnumUnknown                 ListCommerceDisputesLinksSourceModifiedDateStatusEnum = "Unknown"
)

type ListCommerceDisputesLinksSourceModifiedDate struct {
	CreatedDate          *time.Time                                             `json:"createdDate,omitempty"`
	Currency             string                                                 `json:"currency"`
	DisputedTransactions *ListCommerceDisputesLinksSourceModifiedDateRecordRef  `json:"disputedTransactions,omitempty"`
	DueDate              *time.Time                                             `json:"dueDate,omitempty"`
	ID                   string                                                 `json:"id"`
	ModifiedDate         *time.Time                                             `json:"modifiedDate,omitempty"`
	Reason               *string                                                `json:"reason,omitempty"`
	SourceModifiedDate   *time.Time                                             `json:"sourceModifiedDate,omitempty"`
	Status               *ListCommerceDisputesLinksSourceModifiedDateStatusEnum `json:"status,omitempty"`
	TotalAmount          *float64                                               `json:"totalAmount,omitempty"`
}

// ListCommerceDisputesLinks
// Codat's Paging Model
type ListCommerceDisputesLinks struct {
	Links        ListCommerceDisputesLinksLinks                `json:"_links"`
	PageNumber   int64                                         `json:"pageNumber"`
	PageSize     int64                                         `json:"pageSize"`
	Results      []ListCommerceDisputesLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                         `json:"totalResults"`
}

type ListCommerceDisputesResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListCommerceDisputesLinks
}
