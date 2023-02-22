package shared

import (
	"time"
)

type CodatSyncDirectAPIModelsDataConnection struct {
	Created              *time.Time                                    `json:"created,omitempty"`
	DataConnectionErrors []CodatSyncDirectAPIModelsDataConnectionError `json:"dataConnectionErrors,omitempty"`
	ID                   string                                        `json:"id"`
	IntegrationID        string                                        `json:"integrationId"`
	LastSync             *time.Time                                    `json:"lastSync,omitempty"`
	LinkURL              string                                        `json:"linkUrl"`
	PlatformName         string                                        `json:"platformName"`
	SourceID             string                                        `json:"sourceId"`
	SourceType           *string                                       `json:"sourceType,omitempty"`
	Status               *string                                       `json:"status,omitempty"`
}
