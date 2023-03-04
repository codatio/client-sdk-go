package operations

import (
	"net/http"
	"time"
)

// PostProfileSyncSettingsRequestBodySyncSetting
// Describes how often, and how much history, should be fetched for the given data type when a pull operation is queued.
type PostProfileSyncSettingsRequestBodySyncSetting struct {
	DataType         string     `json:"dataType"`
	FetchOnFirstLink bool       `json:"fetchOnFirstLink"`
	IsLocked         *bool      `json:"isLocked,omitempty"`
	MonthsToSync     *int64     `json:"monthsToSync,omitempty"`
	SyncFromUtc      *time.Time `json:"syncFromUtc,omitempty"`
	SyncFromWindow   *int64     `json:"syncFromWindow,omitempty"`
	SyncOrder        int64      `json:"syncOrder"`
	SyncSchedule     int64      `json:"syncSchedule"`
}

type PostProfileSyncSettingsRequestBody struct {
	ClientID          string                                          `json:"clientId"`
	OverridesDefaults bool                                            `json:"overridesDefaults"`
	Settings          []PostProfileSyncSettingsRequestBodySyncSetting `json:"settings"`
}

type PostProfileSyncSettingsRequest struct {
	Request *PostProfileSyncSettingsRequestBody `request:"mediaType=application/json"`
}

type PostProfileSyncSettings401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type PostProfileSyncSettingsResponse struct {
	ContentType                                     string
	StatusCode                                      int
	RawResponse                                     *http.Response
	PostProfileSyncSettings401ApplicationJSONObject *PostProfileSyncSettings401ApplicationJSON
}
