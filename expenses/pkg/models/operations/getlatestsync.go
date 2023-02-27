package operations

import (
	"time"
)

type GetLatestSyncPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetLatestSyncRequest struct {
	PathParams GetLatestSyncPathParams
}

type GetLatestSync200ApplicationJSON struct {
	CompanyID            *string    `json:"companyId,omitempty"`
	DataPushed           *bool      `json:"dataPushed,omitempty"`
	ErrorMessage         *string    `json:"errorMessage,omitempty"`
	SyncExceptionMessage *string    `json:"syncExceptionMessage,omitempty"`
	SyncID               *string    `json:"syncId,omitempty"`
	SyncStatus           *string    `json:"syncStatus,omitempty"`
	SyncStatusCode       *int       `json:"syncStatusCode,omitempty"`
	SyncUtc              *time.Time `json:"syncUtc,omitempty"`
}

type GetLatestSyncResponse struct {
	ContentType                           string
	StatusCode                            int
	GetLatestSync200ApplicationJSONObject *GetLatestSync200ApplicationJSON
}
