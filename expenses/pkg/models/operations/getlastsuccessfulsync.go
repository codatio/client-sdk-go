package operations

import (
	"time"
)

type GetLastSuccessfulSyncPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetLastSuccessfulSyncRequest struct {
	PathParams GetLastSuccessfulSyncPathParams
}

type GetLastSuccessfulSync200ApplicationJSON struct {
	CompanyID            *string    `json:"companyId,omitempty"`
	DataPushed           *bool      `json:"dataPushed,omitempty"`
	ErrorMessage         *string    `json:"errorMessage,omitempty"`
	SyncExceptionMessage *string    `json:"syncExceptionMessage,omitempty"`
	SyncID               *string    `json:"syncId,omitempty"`
	SyncStatus           *string    `json:"syncStatus,omitempty"`
	SyncStatusCode       *int       `json:"syncStatusCode,omitempty"`
	SyncUtc              *time.Time `json:"syncUtc,omitempty"`
}

type GetLastSuccessfulSyncResponse struct {
	ContentType                                   string
	StatusCode                                    int
	GetLastSuccessfulSync200ApplicationJSONObject *GetLastSuccessfulSync200ApplicationJSON
}
