package operations

import (
	"time"
)

type PostCompaniesRequestBody struct {
	Name string `json:"name"`
}

type PostCompaniesRequest struct {
	Request *PostCompaniesRequestBody `request:"mediaType=application/json"`
}

type PostCompanies200ApplicationJSONDataConnectionsDataConnectionErrors struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

type PostCompanies200ApplicationJSONDataConnections struct {
	Created              *time.Time                                                           `json:"created,omitempty"`
	DataConnectionErrors []PostCompanies200ApplicationJSONDataConnectionsDataConnectionErrors `json:"dataConnectionErrors,omitempty"`
	ID                   string                                                               `json:"id"`
	IntegrationID        string                                                               `json:"integrationId"`
	LastSync             *time.Time                                                           `json:"lastSync,omitempty"`
	LinkURL              string                                                               `json:"linkUrl"`
	PlatformName         string                                                               `json:"platformName"`
	SourceID             string                                                               `json:"sourceId"`
	SourceType           *string                                                              `json:"sourceType,omitempty"`
	Status               *string                                                              `json:"status,omitempty"`
}

type PostCompanies200ApplicationJSON struct {
	Created           *time.Time                                       `json:"created,omitempty"`
	CreatedByUserName *string                                          `json:"createdByUserName,omitempty"`
	DataConnections   []PostCompanies200ApplicationJSONDataConnections `json:"dataConnections"`
	ID                string                                           `json:"id"`
	LastSync          *time.Time                                       `json:"lastSync,omitempty"`
	Name              string                                           `json:"name"`
	Platform          string                                           `json:"platform"`
	Redirect          string                                           `json:"redirect"`
}

type PostCompaniesResponse struct {
	ContentType                           string
	StatusCode                            int
	PostCompanies200ApplicationJSONObject *PostCompanies200ApplicationJSON
}
