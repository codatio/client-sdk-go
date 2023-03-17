package operations

import (
	"net/http"
	"time"
)

type CreatePartnerexpenseConnectionRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type CreatePartnerexpenseConnectionConnectionConnectionInfo struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty"`
}

type CreatePartnerexpenseConnectionConnectionDataConnectionErrors struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

type CreatePartnerexpenseConnectionConnectionSourceTypeEnum string

const (
	CreatePartnerexpenseConnectionConnectionSourceTypeEnumAccounting CreatePartnerexpenseConnectionConnectionSourceTypeEnum = "Accounting"
	CreatePartnerexpenseConnectionConnectionSourceTypeEnumBanking    CreatePartnerexpenseConnectionConnectionSourceTypeEnum = "Banking"
	CreatePartnerexpenseConnectionConnectionSourceTypeEnumCommerce   CreatePartnerexpenseConnectionConnectionSourceTypeEnum = "Commerce"
	CreatePartnerexpenseConnectionConnectionSourceTypeEnumOther      CreatePartnerexpenseConnectionConnectionSourceTypeEnum = "Other"
	CreatePartnerexpenseConnectionConnectionSourceTypeEnumUnknown    CreatePartnerexpenseConnectionConnectionSourceTypeEnum = "Unknown"
)

type CreatePartnerexpenseConnectionConnectionDataConnectionStatusEnum string

const (
	CreatePartnerexpenseConnectionConnectionDataConnectionStatusEnumPendingAuth  CreatePartnerexpenseConnectionConnectionDataConnectionStatusEnum = "PendingAuth"
	CreatePartnerexpenseConnectionConnectionDataConnectionStatusEnumLinked       CreatePartnerexpenseConnectionConnectionDataConnectionStatusEnum = "Linked"
	CreatePartnerexpenseConnectionConnectionDataConnectionStatusEnumUnlinked     CreatePartnerexpenseConnectionConnectionDataConnectionStatusEnum = "Unlinked"
	CreatePartnerexpenseConnectionConnectionDataConnectionStatusEnumDeauthorized CreatePartnerexpenseConnectionConnectionDataConnectionStatusEnum = "Deauthorized"
)

// CreatePartnerexpenseConnectionConnection
// A connection represents the link between a `company` and a source of data.
type CreatePartnerexpenseConnectionConnection struct {
	ConnectionInfo       *CreatePartnerexpenseConnectionConnectionConnectionInfo          `json:"connectionInfo,omitempty"`
	Created              time.Time                                                        `json:"created"`
	DataConnectionErrors []CreatePartnerexpenseConnectionConnectionDataConnectionErrors   `json:"dataConnectionErrors,omitempty"`
	ID                   string                                                           `json:"id"`
	IntegrationID        string                                                           `json:"integrationId"`
	IntegrationKey       string                                                           `json:"integrationKey"`
	LastSync             *time.Time                                                       `json:"lastSync,omitempty"`
	LinkURL              string                                                           `json:"linkUrl"`
	PlatformName         string                                                           `json:"platformName"`
	SourceID             string                                                           `json:"sourceId"`
	SourceType           CreatePartnerexpenseConnectionConnectionSourceTypeEnum           `json:"sourceType"`
	Status               CreatePartnerexpenseConnectionConnectionDataConnectionStatusEnum `json:"status"`
}

type CreatePartnerexpenseConnectionResponse struct {
	Connection  *CreatePartnerexpenseConnectionConnection
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
