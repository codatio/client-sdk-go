package operations

import (
	"net/http"
	"time"
)

type GetDataconnectionsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetDataconnectionsQueryParams struct {
	OrderBy  *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page     int     `queryParam:"style=form,explode=true,name=page"`
	PageSize *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string `queryParam:"style=form,explode=true,name=query"`
}

type GetDataconnectionsRequest struct {
	PathParams  GetDataconnectionsPathParams
	QueryParams GetDataconnectionsQueryParams
}

type GetDataconnections200ApplicationJSONLinksCurrent struct {
	Href *string `json:"href,omitempty"`
}

type GetDataconnections200ApplicationJSONLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type GetDataconnections200ApplicationJSONLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type GetDataconnections200ApplicationJSONLinksSelf struct {
	Href *string `json:"href,omitempty"`
}

type GetDataconnections200ApplicationJSONLinks struct {
	Current  *GetDataconnections200ApplicationJSONLinksCurrent  `json:"current,omitempty"`
	Next     *GetDataconnections200ApplicationJSONLinksNext     `json:"next,omitempty"`
	Previous *GetDataconnections200ApplicationJSONLinksPrevious `json:"previous,omitempty"`
	Self     *GetDataconnections200ApplicationJSONLinksSelf     `json:"self,omitempty"`
}

type GetDataconnections200ApplicationJSONResultsDataConnectionErrors struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

type GetDataconnections200ApplicationJSONResults struct {
	Created              *time.Time                                                        `json:"created,omitempty"`
	DataConnectionErrors []GetDataconnections200ApplicationJSONResultsDataConnectionErrors `json:"dataConnectionErrors,omitempty"`
	ID                   string                                                            `json:"id"`
	IntegrationID        string                                                            `json:"integrationId"`
	LastSync             *time.Time                                                        `json:"lastSync,omitempty"`
	LinkURL              string                                                            `json:"linkUrl"`
	PlatformName         string                                                            `json:"platformName"`
	SourceID             string                                                            `json:"sourceId"`
	SourceType           *string                                                           `json:"sourceType,omitempty"`
	Status               *string                                                           `json:"status,omitempty"`
}

// GetDataconnections200ApplicationJSON
// Used to represent what can be returned by an endpoint that supports paging.
// Usable with the [ProducesResponseType] attribute on a controller action.
type GetDataconnections200ApplicationJSON struct {
	Links        *GetDataconnections200ApplicationJSONLinks    `json:"_links,omitempty"`
	PageNumber   *int                                          `json:"pageNumber,omitempty"`
	PageSize     *int                                          `json:"pageSize,omitempty"`
	Results      []GetDataconnections200ApplicationJSONResults `json:"results,omitempty"`
	TotalResults *int                                          `json:"totalResults,omitempty"`
}

type GetDataconnectionsResponse struct {
	ContentType                                string
	StatusCode                                 int
	RawResponse                                *http.Response
	GetDataconnections200ApplicationJSONObject *GetDataconnections200ApplicationJSON
}
