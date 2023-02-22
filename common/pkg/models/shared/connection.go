package shared

import (
	"time"
)

type ConnectionConnectionInfo struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty"`
}

type ConnectionDataConnectionErrors struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

type ConnectionSourceTypeEnum string

const (
	ConnectionSourceTypeEnumAccounting ConnectionSourceTypeEnum = "Accounting"
	ConnectionSourceTypeEnumBanking    ConnectionSourceTypeEnum = "Banking"
	ConnectionSourceTypeEnumCommerce   ConnectionSourceTypeEnum = "Commerce"
	ConnectionSourceTypeEnumOther      ConnectionSourceTypeEnum = "Other"
	ConnectionSourceTypeEnumUnknown    ConnectionSourceTypeEnum = "Unknown"
)

type ConnectionStatusEnum string

const (
	ConnectionStatusEnumPendingAuth  ConnectionStatusEnum = "PendingAuth"
	ConnectionStatusEnumLinked       ConnectionStatusEnum = "Linked"
	ConnectionStatusEnumUnlinked     ConnectionStatusEnum = "Unlinked"
	ConnectionStatusEnumDeauthorized ConnectionStatusEnum = "Deauthorized"
)

// Connection
// A connection represents the link between a `company` and a source of data.
type Connection struct {
	ConnectionInfo       *ConnectionConnectionInfo        `json:"connectionInfo,omitempty"`
	Created              time.Time                        `json:"created"`
	DataConnectionErrors []ConnectionDataConnectionErrors `json:"dataConnectionErrors,omitempty"`
	ID                   string                           `json:"id"`
	IntegrationID        string                           `json:"integrationId"`
	IntegrationKey       string                           `json:"integrationKey"`
	LastSync             *time.Time                       `json:"lastSync,omitempty"`
	LinkURL              string                           `json:"linkUrl"`
	PlatformName         string                           `json:"platformName"`
	SourceID             string                           `json:"sourceId"`
	SourceType           ConnectionSourceTypeEnum         `json:"sourceType"`
	Status               ConnectionStatusEnum             `json:"status"`
}
