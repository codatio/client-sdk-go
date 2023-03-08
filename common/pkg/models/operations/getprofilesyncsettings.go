package operations

import (
	"net/http"
	"time"
)

type GetProfileSyncSettings401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

// GetProfileSyncSettings200ApplicationJSONSyncSetting
// Describes how often, and how much history, should be fetched for the given data type when a pull operation is queued.
type GetProfileSyncSettings200ApplicationJSONSyncSetting struct {
	DataType         string     `json:"dataType"`
	FetchOnFirstLink bool       `json:"fetchOnFirstLink"`
	IsLocked         *bool      `json:"isLocked,omitempty"`
	MonthsToSync     *int64     `json:"monthsToSync,omitempty"`
	SyncFromUtc      *time.Time `json:"syncFromUtc,omitempty"`
	SyncFromWindow   *int64     `json:"syncFromWindow,omitempty"`
	SyncOrder        int64      `json:"syncOrder"`
	SyncSchedule     int64      `json:"syncSchedule"`
}

type GetProfileSyncSettings200ApplicationJSON struct {
	ClientID          *string                                               `json:"clientId,omitempty"`
	OverridesDefaults *bool                                                 `json:"overridesDefaults,omitempty"`
	Settings          []GetProfileSyncSettings200ApplicationJSONSyncSetting `json:"settings,omitempty"`
}

type GetProfileSyncSettingsResponse struct {
	ContentType                                    string
	StatusCode                                     int
	RawResponse                                    *http.Response
	GetProfileSyncSettings200ApplicationJSONObject *GetProfileSyncSettings200ApplicationJSON
	GetProfileSyncSettings401ApplicationJSONObject *GetProfileSyncSettings401ApplicationJSON
}
