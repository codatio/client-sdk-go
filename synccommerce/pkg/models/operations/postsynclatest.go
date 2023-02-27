package operations

import (
	"time"
)

type PostSyncLatestPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type PostSyncLatestApplicationWildcardPlusJSON struct {
	SyncTo *time.Time `json:"syncTo,omitempty"`
}

type PostSyncLatestApplicationJSON struct {
	SyncTo *time.Time `json:"syncTo,omitempty"`
}

type PostSyncLatestApplicationJSONPatchPlusJSON struct {
	SyncTo *time.Time `json:"syncTo,omitempty"`
}

type PostSyncLatestTextJSON struct {
	SyncTo *time.Time `json:"syncTo,omitempty"`
}

type PostSyncLatestRequests struct {
	Object  *PostSyncLatestApplicationWildcardPlusJSON  `request:"mediaType=application/*+json"`
	Object1 *PostSyncLatestApplicationJSON              `request:"mediaType=application/json"`
	Object2 *PostSyncLatestApplicationJSONPatchPlusJSON `request:"mediaType=application/json-patch+json"`
	Object3 *PostSyncLatestTextJSON                     `request:"mediaType=text/json"`
}

type PostSyncLatestRequest struct {
	PathParams PostSyncLatestPathParams
	Request    *PostSyncLatestRequests
}

type PostSyncLatest200TextJSONDataConnectionsDataConnectionErrors struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

type PostSyncLatest200TextJSONDataConnections struct {
	Created              *time.Time                                                     `json:"created,omitempty"`
	DataConnectionErrors []PostSyncLatest200TextJSONDataConnectionsDataConnectionErrors `json:"dataConnectionErrors,omitempty"`
	ID                   string                                                         `json:"id"`
	IntegrationID        string                                                         `json:"integrationId"`
	LastSync             *time.Time                                                     `json:"lastSync,omitempty"`
	LinkURL              string                                                         `json:"linkUrl"`
	PlatformName         string                                                         `json:"platformName"`
	SourceID             string                                                         `json:"sourceId"`
	SourceType           *string                                                        `json:"sourceType,omitempty"`
	Status               *string                                                        `json:"status,omitempty"`
}

type PostSyncLatest200TextJSONSyncDateRangeUtc struct {
	Finish *time.Time `json:"finish,omitempty"`
	Start  *time.Time `json:"start,omitempty"`
}

type PostSyncLatest200TextJSON struct {
	CommerceSyncID       *string                                    `json:"commerceSyncId,omitempty"`
	CompanyID            *string                                    `json:"companyId,omitempty"`
	DataConnections      []PostSyncLatest200TextJSONDataConnections `json:"dataConnections,omitempty"`
	DataPushed           *bool                                      `json:"dataPushed,omitempty"`
	ErrorMessage         *string                                    `json:"errorMessage,omitempty"`
	SyncDateRangeUtc     *PostSyncLatest200TextJSONSyncDateRangeUtc `json:"syncDateRangeUtc,omitempty"`
	SyncExceptionMessage *string                                    `json:"syncExceptionMessage,omitempty"`
	SyncStatus           *string                                    `json:"syncStatus,omitempty"`
	SyncStatusCode       *int                                       `json:"syncStatusCode,omitempty"`
	SyncUtc              *time.Time                                 `json:"syncUtc,omitempty"`
}

type PostSyncLatest200ApplicationJSONDataConnectionsDataConnectionErrors struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

type PostSyncLatest200ApplicationJSONDataConnections struct {
	Created              *time.Time                                                            `json:"created,omitempty"`
	DataConnectionErrors []PostSyncLatest200ApplicationJSONDataConnectionsDataConnectionErrors `json:"dataConnectionErrors,omitempty"`
	ID                   string                                                                `json:"id"`
	IntegrationID        string                                                                `json:"integrationId"`
	LastSync             *time.Time                                                            `json:"lastSync,omitempty"`
	LinkURL              string                                                                `json:"linkUrl"`
	PlatformName         string                                                                `json:"platformName"`
	SourceID             string                                                                `json:"sourceId"`
	SourceType           *string                                                               `json:"sourceType,omitempty"`
	Status               *string                                                               `json:"status,omitempty"`
}

type PostSyncLatest200ApplicationJSONSyncDateRangeUtc struct {
	Finish *time.Time `json:"finish,omitempty"`
	Start  *time.Time `json:"start,omitempty"`
}

type PostSyncLatest200ApplicationJSON struct {
	CommerceSyncID       *string                                           `json:"commerceSyncId,omitempty"`
	CompanyID            *string                                           `json:"companyId,omitempty"`
	DataConnections      []PostSyncLatest200ApplicationJSONDataConnections `json:"dataConnections,omitempty"`
	DataPushed           *bool                                             `json:"dataPushed,omitempty"`
	ErrorMessage         *string                                           `json:"errorMessage,omitempty"`
	SyncDateRangeUtc     *PostSyncLatest200ApplicationJSONSyncDateRangeUtc `json:"syncDateRangeUtc,omitempty"`
	SyncExceptionMessage *string                                           `json:"syncExceptionMessage,omitempty"`
	SyncStatus           *string                                           `json:"syncStatus,omitempty"`
	SyncStatusCode       *int                                              `json:"syncStatusCode,omitempty"`
	SyncUtc              *time.Time                                        `json:"syncUtc,omitempty"`
}

type PostSyncLatestResponse struct {
	ContentType                            string
	StatusCode                             int
	PostSyncLatest200ApplicationJSONObject *PostSyncLatest200ApplicationJSON
	PostSyncLatest200TextJSONObject        *PostSyncLatest200TextJSON
	PostSyncLatest200TextPlainObject       *string
}
