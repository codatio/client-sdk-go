package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

type PostProfileSyncSettingsRequestBody struct {
	ClientID          string               `json:"clientId"`
	OverridesDefaults bool                 `json:"overridesDefaults"`
	Settings          []shared.SyncSetting `json:"settings"`
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
	PostProfileSyncSettings401ApplicationJSONObject *PostProfileSyncSettings401ApplicationJSON
}
