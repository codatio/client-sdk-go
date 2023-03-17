package operations

import (
	"net/http"
	"time"
)

type GetSyncByIDRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	SyncID    string `pathParam:"style=simple,explode=false,name=syncId"`
}

type GetSyncByID200ApplicationJSON struct {
	CompanyID            *string    `json:"companyId,omitempty"`
	DataPushed           *bool      `json:"dataPushed,omitempty"`
	ErrorMessage         *string    `json:"errorMessage,omitempty"`
	SyncExceptionMessage *string    `json:"syncExceptionMessage,omitempty"`
	SyncID               *string    `json:"syncId,omitempty"`
	SyncStatus           *string    `json:"syncStatus,omitempty"`
	SyncStatusCode       *int       `json:"syncStatusCode,omitempty"`
	SyncUtc              *time.Time `json:"syncUtc,omitempty"`
}

type GetSyncByIDResponse struct {
	ContentType                         string
	StatusCode                          int
	RawResponse                         *http.Response
	GetSyncByID200ApplicationJSONObject *GetSyncByID200ApplicationJSON
}
