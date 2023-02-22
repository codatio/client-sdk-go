package shared

import (
	"time"
)

type CodatSyncDirectAPIModelsDataConnectionError struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}
