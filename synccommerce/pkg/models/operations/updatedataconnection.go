package operations

import (
	"time"
)

type UpdateDataConnectionPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type UpdateDataConnectionRequestBody struct {
	Status *string `json:"status,omitempty"`
}

type UpdateDataConnectionRequest struct {
	PathParams UpdateDataConnectionPathParams
	Request    *UpdateDataConnectionRequestBody `request:"mediaType=application/json"`
}

type UpdateDataConnection200ApplicationJSONDataConnectionErrors struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

type UpdateDataConnection200ApplicationJSON struct {
	Created              *time.Time                                                   `json:"created,omitempty"`
	DataConnectionErrors []UpdateDataConnection200ApplicationJSONDataConnectionErrors `json:"dataConnectionErrors,omitempty"`
	ID                   string                                                       `json:"id"`
	IntegrationID        string                                                       `json:"integrationId"`
	LastSync             *time.Time                                                   `json:"lastSync,omitempty"`
	LinkURL              string                                                       `json:"linkUrl"`
	PlatformName         string                                                       `json:"platformName"`
	SourceID             string                                                       `json:"sourceId"`
	SourceType           *string                                                      `json:"sourceType,omitempty"`
	Status               *string                                                      `json:"status,omitempty"`
}

type UpdateDataConnectionResponse struct {
	ContentType                                  string
	StatusCode                                   int
	UpdateDataConnection200ApplicationJSONObject *UpdateDataConnection200ApplicationJSON
}
