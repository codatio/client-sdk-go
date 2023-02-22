package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

type GetProfileSyncSettings401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetProfileSyncSettings200ApplicationJSON struct {
	ClientID          *string              `json:"clientId,omitempty"`
	OverridesDefaults *bool                `json:"overridesDefaults,omitempty"`
	Settings          []shared.SyncSetting `json:"settings,omitempty"`
}

type GetProfileSyncSettingsResponse struct {
	ContentType                                    string
	StatusCode                                     int
	GetProfileSyncSettings200ApplicationJSONObject *GetProfileSyncSettings200ApplicationJSON
	GetProfileSyncSettings401ApplicationJSONObject *GetProfileSyncSettings401ApplicationJSON
}
