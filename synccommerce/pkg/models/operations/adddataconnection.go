package operations

import (
	"net/http"
	"time"
)

type AddDataConnectionRequest struct {
	RequestBody *string `request:"mediaType=application/json"`
	CompanyID   string  `pathParam:"style=simple,explode=false,name=companyId"`
}

type AddDataConnection200ApplicationJSONDataConnectionErrors struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

type AddDataConnection200ApplicationJSON struct {
	Created              *time.Time                                                `json:"created,omitempty"`
	DataConnectionErrors []AddDataConnection200ApplicationJSONDataConnectionErrors `json:"dataConnectionErrors,omitempty"`
	ID                   string                                                    `json:"id"`
	IntegrationID        string                                                    `json:"integrationId"`
	LastSync             *time.Time                                                `json:"lastSync,omitempty"`
	LinkURL              string                                                    `json:"linkUrl"`
	PlatformName         string                                                    `json:"platformName"`
	SourceID             string                                                    `json:"sourceId"`
	SourceType           *string                                                   `json:"sourceType,omitempty"`
	Status               *string                                                   `json:"status,omitempty"`
}

type AddDataConnectionResponse struct {
	ContentType                               string
	StatusCode                                int
	RawResponse                               *http.Response
	AddDataConnection200ApplicationJSONObject *AddDataConnection200ApplicationJSON
}
