package operations

import (
	"net/http"
	"time"
)

type GetLastSuccessfulSyncRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
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
	RawResponse                                   *http.Response
	GetLastSuccessfulSync200ApplicationJSONObject *GetLastSuccessfulSync200ApplicationJSON
}
