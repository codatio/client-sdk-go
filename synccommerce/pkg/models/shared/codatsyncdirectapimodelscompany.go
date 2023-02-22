package shared

import (
	"time"
)

type CodatSyncDirectAPIModelsCompany struct {
	Created         *time.Time                               `json:"created,omitempty"`
	DataConnections []CodatSyncDirectAPIModelsDataConnection `json:"dataConnections,omitempty"`
	ID              string                                   `json:"id"`
	Name            string                                   `json:"name"`
}
