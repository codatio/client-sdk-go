package operations

import (
	"net/http"
	"time"
)

type CreateDataConnectionRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type CreateDataConnectionConnectionConnectionInfo struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty"`
}

type CreateDataConnectionConnectionDataConnectionErrors struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

type CreateDataConnectionConnectionSourceTypeEnum string

const (
	CreateDataConnectionConnectionSourceTypeEnumAccounting CreateDataConnectionConnectionSourceTypeEnum = "Accounting"
	CreateDataConnectionConnectionSourceTypeEnumBanking    CreateDataConnectionConnectionSourceTypeEnum = "Banking"
	CreateDataConnectionConnectionSourceTypeEnumCommerce   CreateDataConnectionConnectionSourceTypeEnum = "Commerce"
	CreateDataConnectionConnectionSourceTypeEnumOther      CreateDataConnectionConnectionSourceTypeEnum = "Other"
	CreateDataConnectionConnectionSourceTypeEnumUnknown    CreateDataConnectionConnectionSourceTypeEnum = "Unknown"
)

type CreateDataConnectionConnectionDataConnectionStatusEnum string

const (
	CreateDataConnectionConnectionDataConnectionStatusEnumPendingAuth  CreateDataConnectionConnectionDataConnectionStatusEnum = "PendingAuth"
	CreateDataConnectionConnectionDataConnectionStatusEnumLinked       CreateDataConnectionConnectionDataConnectionStatusEnum = "Linked"
	CreateDataConnectionConnectionDataConnectionStatusEnumUnlinked     CreateDataConnectionConnectionDataConnectionStatusEnum = "Unlinked"
	CreateDataConnectionConnectionDataConnectionStatusEnumDeauthorized CreateDataConnectionConnectionDataConnectionStatusEnum = "Deauthorized"
)

// CreateDataConnectionConnection
// A connection represents the link between a `company` and a source of data.
type CreateDataConnectionConnection struct {
	ConnectionInfo       *CreateDataConnectionConnectionConnectionInfo          `json:"connectionInfo,omitempty"`
	Created              time.Time                                              `json:"created"`
	DataConnectionErrors []CreateDataConnectionConnectionDataConnectionErrors   `json:"dataConnectionErrors,omitempty"`
	ID                   string                                                 `json:"id"`
	IntegrationID        string                                                 `json:"integrationId"`
	IntegrationKey       string                                                 `json:"integrationKey"`
	LastSync             *time.Time                                             `json:"lastSync,omitempty"`
	LinkURL              string                                                 `json:"linkUrl"`
	PlatformName         string                                                 `json:"platformName"`
	SourceID             string                                                 `json:"sourceId"`
	SourceType           CreateDataConnectionConnectionSourceTypeEnum           `json:"sourceType"`
	Status               CreateDataConnectionConnectionDataConnectionStatusEnum `json:"status"`
}

type CreateDataConnectionResponse struct {
	Connection  *CreateDataConnectionConnection
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
