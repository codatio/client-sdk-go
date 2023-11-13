// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// SourceType - The type of platform of the connection.
type SourceType string

const (
	SourceTypeAccounting SourceType = "Accounting"
	SourceTypeBanking    SourceType = "Banking"
	SourceTypeCommerce   SourceType = "Commerce"
	SourceTypeOther      SourceType = "Other"
	SourceTypeUnknown    SourceType = "Unknown"
)

func (e SourceType) ToPointer() *SourceType {
	return &e
}

func (e *SourceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Accounting":
		fallthrough
	case "Banking":
		fallthrough
	case "Commerce":
		fallthrough
	case "Other":
		fallthrough
	case "Unknown":
		*e = SourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceType: %v", v)
	}
}

// Connection - A connection represents a [company's](https://docs.codat.io/bank-feeds-api#/schemas/Company) connection to a data source and allows you to synchronize data (pull and/or push) with that source.
//
// A company can have multiple data connections depending on the type of data source it is connecting to. For example, a single company can link to:
//
// - [Accounting data](https://docs.codat.io/accounting-api/overview) - 1 active connection.
// - [Banking data](https://docs.codat.io/banking-api/overview) - Multiple active connections.
// - [Commerce data](https://docs.codat.io/commerce-api/overview) - Multiple active connections.
// Any combination of accounting, banking, and commerce data connections is allowed.
//
// Before you can use a data connection to pull or push data, the company must grant you access to their business data by [linking the connection](https://docs.codat.io/auth-flow/overview).
type Connection struct {
	AdditionalProperties interface{}       `json:"additionalProperties,omitempty"`
	ConnectionInfo       map[string]string `json:"connectionInfo,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	Created              string                `json:"created"`
	DataConnectionErrors []DataConnectionError `json:"dataConnectionErrors,omitempty"`
	// Unique identifier for a company's data connection.
	ID string `json:"id"`
	// A Codat ID representing the integration.
	IntegrationID string `json:"integrationId"`
	// A unique four-character ID that identifies the platform of the company's data connection. This ensures continuity if the platform changes its name in the future.
	IntegrationKey string `json:"integrationKey"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	LastSync *string `json:"lastSync,omitempty"`
	// The link URL your customers can use to authorize access to their business application.
	LinkURL string `json:"linkUrl"`
	// Name of integration connected to company.
	PlatformName string `json:"platformName"`
	// A source-specific ID used to distinguish between different sources originating from the same data connection. In general, a data connection is a single data source. However, for TrueLayer, `sourceId` is associated with a specific bank and has a many-to-one relationship with the `integrationId`.
	SourceID string `json:"sourceId"`
	// The type of platform of the connection.
	SourceType SourceType `json:"sourceType"`
	// The current authorization status of the data connection.
	Status DataConnectionStatus `json:"status"`
}

func (o *Connection) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *Connection) GetConnectionInfo() map[string]string {
	if o == nil {
		return nil
	}
	return o.ConnectionInfo
}

func (o *Connection) GetCreated() string {
	if o == nil {
		return ""
	}
	return o.Created
}

func (o *Connection) GetDataConnectionErrors() []DataConnectionError {
	if o == nil {
		return nil
	}
	return o.DataConnectionErrors
}

func (o *Connection) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Connection) GetIntegrationID() string {
	if o == nil {
		return ""
	}
	return o.IntegrationID
}

func (o *Connection) GetIntegrationKey() string {
	if o == nil {
		return ""
	}
	return o.IntegrationKey
}

func (o *Connection) GetLastSync() *string {
	if o == nil {
		return nil
	}
	return o.LastSync
}

func (o *Connection) GetLinkURL() string {
	if o == nil {
		return ""
	}
	return o.LinkURL
}

func (o *Connection) GetPlatformName() string {
	if o == nil {
		return ""
	}
	return o.PlatformName
}

func (o *Connection) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *Connection) GetSourceType() SourceType {
	if o == nil {
		return SourceType("")
	}
	return o.SourceType
}

func (o *Connection) GetStatus() DataConnectionStatus {
	if o == nil {
		return DataConnectionStatus("")
	}
	return o.Status
}
