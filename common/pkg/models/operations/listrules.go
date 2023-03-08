package operations

import (
	"net/http"
)

type ListRulesQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListRulesRequest struct {
	QueryParams ListRulesQueryParams
}

type ListRules401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type ListRules400ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type ListRulesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListRulesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListRulesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListRulesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListRulesLinksLinks struct {
	Current  ListRulesLinksLinksCurrent   `json:"current"`
	Next     *ListRulesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListRulesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListRulesLinksLinksSelf      `json:"self"`
}

type ListRulesLinksWebhookNotifiers struct {
	Emails  []string `json:"emails,omitempty"`
	Webhook *string  `json:"webhook,omitempty"`
}

// ListRulesLinksWebhook
// Configuration to alert to a url or list of email addresses based on the given type / condition.
type ListRulesLinksWebhook struct {
	CompanyID *string                        `json:"companyId,omitempty"`
	ID        string                         `json:"id"`
	Notifiers ListRulesLinksWebhookNotifiers `json:"notifiers"`
	Type      string                         `json:"type"`
}

// ListRulesLinks
// Codat's Paging Model
type ListRulesLinks struct {
	Links        ListRulesLinksLinks     `json:"_links"`
	PageNumber   int64                   `json:"pageNumber"`
	PageSize     int64                   `json:"pageSize"`
	Results      []ListRulesLinksWebhook `json:"results,omitempty"`
	TotalResults int64                   `json:"totalResults"`
}

type ListRulesResponse struct {
	ContentType                       string
	StatusCode                        int
	RawResponse                       *http.Response
	Links                             *ListRulesLinks
	ListRules400ApplicationJSONObject *ListRules400ApplicationJSON
	ListRules401ApplicationJSONObject *ListRules401ApplicationJSON
}
