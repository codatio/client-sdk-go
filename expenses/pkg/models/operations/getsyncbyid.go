package operations

import (
	"time"
)

type GetSyncByIDPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	SyncID    string `pathParam:"style=simple,explode=false,name=syncId"`
}

type GetSyncByIDRequest struct {
	PathParams GetSyncByIDPathParams
}

type GetSyncByID200TextJSON struct {
	CompanyID            *string    `json:"companyId,omitempty"`
	DataPushed           *bool      `json:"dataPushed,omitempty"`
	ErrorMessage         *string    `json:"errorMessage,omitempty"`
	SyncExceptionMessage *string    `json:"syncExceptionMessage,omitempty"`
	SyncID               *string    `json:"syncId,omitempty"`
	SyncStatus           *string    `json:"syncStatus,omitempty"`
	SyncStatusCode       *int       `json:"syncStatusCode,omitempty"`
	SyncUtc              *time.Time `json:"syncUtc,omitempty"`
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
	GetSyncByID200ApplicationJSONObject *GetSyncByID200ApplicationJSON
	GetSyncByID200TextJSONObject        *GetSyncByID200TextJSON
	GetSyncByID200TextPlainObject       *string
}
