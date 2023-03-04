package operations

import (
	"net/http"
	"time"
)

type ListSyncsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListSyncsRequest struct {
	PathParams ListSyncsPathParams
}

type ListSyncs200ApplicationJSON struct {
	CompanyID            *string    `json:"companyId,omitempty"`
	DataPushed           *bool      `json:"dataPushed,omitempty"`
	ErrorMessage         *string    `json:"errorMessage,omitempty"`
	SyncExceptionMessage *string    `json:"syncExceptionMessage,omitempty"`
	SyncID               *string    `json:"syncId,omitempty"`
	SyncStatus           *string    `json:"syncStatus,omitempty"`
	SyncStatusCode       *int       `json:"syncStatusCode,omitempty"`
	SyncUtc              *time.Time `json:"syncUtc,omitempty"`
}

type ListSyncsResponse struct {
	ContentType                        string
	StatusCode                         int
	RawResponse                        *http.Response
	ListSyncs200ApplicationJSONObjects []ListSyncs200ApplicationJSON
}
