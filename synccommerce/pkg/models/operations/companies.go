package operations

import (
	"net/http"
	"time"
)

type CompaniesQueryParams struct {
	OrderBy  *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page     int     `queryParam:"style=form,explode=true,name=page"`
	PageSize *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string `queryParam:"style=form,explode=true,name=query"`
}

type CompaniesRequest struct {
	QueryParams CompaniesQueryParams
}

type Companies200ApplicationJSONLinksCurrent struct {
	Href *string `json:"href,omitempty"`
}

type Companies200ApplicationJSONLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type Companies200ApplicationJSONLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type Companies200ApplicationJSONLinksSelf struct {
	Href *string `json:"href,omitempty"`
}

type Companies200ApplicationJSONLinks struct {
	Current  *Companies200ApplicationJSONLinksCurrent  `json:"current,omitempty"`
	Next     *Companies200ApplicationJSONLinksNext     `json:"next,omitempty"`
	Previous *Companies200ApplicationJSONLinksPrevious `json:"previous,omitempty"`
	Self     *Companies200ApplicationJSONLinksSelf     `json:"self,omitempty"`
}

type Companies200ApplicationJSONResultsDataConnectionsDataConnectionErrors struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

type Companies200ApplicationJSONResultsDataConnections struct {
	Created              *time.Time                                                              `json:"created,omitempty"`
	DataConnectionErrors []Companies200ApplicationJSONResultsDataConnectionsDataConnectionErrors `json:"dataConnectionErrors,omitempty"`
	ID                   string                                                                  `json:"id"`
	IntegrationID        string                                                                  `json:"integrationId"`
	LastSync             *time.Time                                                              `json:"lastSync,omitempty"`
	LinkURL              string                                                                  `json:"linkUrl"`
	PlatformName         string                                                                  `json:"platformName"`
	SourceID             string                                                                  `json:"sourceId"`
	SourceType           *string                                                                 `json:"sourceType,omitempty"`
	Status               *string                                                                 `json:"status,omitempty"`
}

type Companies200ApplicationJSONResults struct {
	Created           *time.Time                                          `json:"created,omitempty"`
	CreatedByUserName *string                                             `json:"createdByUserName,omitempty"`
	DataConnections   []Companies200ApplicationJSONResultsDataConnections `json:"dataConnections"`
	ID                string                                              `json:"id"`
	LastSync          *time.Time                                          `json:"lastSync,omitempty"`
	Name              string                                              `json:"name"`
	Platform          string                                              `json:"platform"`
	Redirect          string                                              `json:"redirect"`
}

// Companies200ApplicationJSON
// Used to represent what can be returned by an endpoint that supports paging.
// Usable with the [ProducesResponseType] attribute on a controller action.
type Companies200ApplicationJSON struct {
	Links        *Companies200ApplicationJSONLinks    `json:"_links,omitempty"`
	PageNumber   *int                                 `json:"pageNumber,omitempty"`
	PageSize     *int                                 `json:"pageSize,omitempty"`
	Results      []Companies200ApplicationJSONResults `json:"results,omitempty"`
	TotalResults *int                                 `json:"totalResults,omitempty"`
}

type CompaniesResponse struct {
	ContentType                       string
	StatusCode                        int
	RawResponse                       *http.Response
	Companies200ApplicationJSONObject *Companies200ApplicationJSON
}
