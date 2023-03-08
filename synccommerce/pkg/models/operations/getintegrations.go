package operations

import (
	"net/http"
)

type GetIntegrationsQueryParams struct {
	OrderBy  *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page     int     `queryParam:"style=form,explode=true,name=page"`
	PageSize *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string `queryParam:"style=form,explode=true,name=query"`
}

type GetIntegrationsRequest struct {
	QueryParams GetIntegrationsQueryParams
}

type GetIntegrations200ApplicationJSONLinksCurrent struct {
	Href *string `json:"href,omitempty"`
}

type GetIntegrations200ApplicationJSONLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type GetIntegrations200ApplicationJSONLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type GetIntegrations200ApplicationJSONLinksSelf struct {
	Href *string `json:"href,omitempty"`
}

type GetIntegrations200ApplicationJSONLinks struct {
	Current  *GetIntegrations200ApplicationJSONLinksCurrent  `json:"current,omitempty"`
	Next     *GetIntegrations200ApplicationJSONLinksNext     `json:"next,omitempty"`
	Previous *GetIntegrations200ApplicationJSONLinksPrevious `json:"previous,omitempty"`
	Self     *GetIntegrations200ApplicationJSONLinksSelf     `json:"self,omitempty"`
}

type GetIntegrations200ApplicationJSONResultsDatatypeFeaturesSupportedFeatures struct {
	FeatureState *string `json:"featureState,omitempty"`
	FeatureType  *string `json:"featureType,omitempty"`
}

type GetIntegrations200ApplicationJSONResultsDatatypeFeatures struct {
	Datatype          *string                                                                     `json:"datatype,omitempty"`
	SupportedFeatures []GetIntegrations200ApplicationJSONResultsDatatypeFeaturesSupportedFeatures `json:"supportedFeatures,omitempty"`
}

type GetIntegrations200ApplicationJSONResultsSourceTypeEnum string

const (
	GetIntegrations200ApplicationJSONResultsSourceTypeEnumUnknown    GetIntegrations200ApplicationJSONResultsSourceTypeEnum = "Unknown"
	GetIntegrations200ApplicationJSONResultsSourceTypeEnumAccounting GetIntegrations200ApplicationJSONResultsSourceTypeEnum = "Accounting"
	GetIntegrations200ApplicationJSONResultsSourceTypeEnumBanking    GetIntegrations200ApplicationJSONResultsSourceTypeEnum = "Banking"
	GetIntegrations200ApplicationJSONResultsSourceTypeEnumBankFeed   GetIntegrations200ApplicationJSONResultsSourceTypeEnum = "BankFeed"
	GetIntegrations200ApplicationJSONResultsSourceTypeEnumCommerce   GetIntegrations200ApplicationJSONResultsSourceTypeEnum = "Commerce"
	GetIntegrations200ApplicationJSONResultsSourceTypeEnumExpense    GetIntegrations200ApplicationJSONResultsSourceTypeEnum = "Expense"
	GetIntegrations200ApplicationJSONResultsSourceTypeEnumOther      GetIntegrations200ApplicationJSONResultsSourceTypeEnum = "Other"
)

type GetIntegrations200ApplicationJSONResultsSupportedEnvironmentsEnum string

const (
	GetIntegrations200ApplicationJSONResultsSupportedEnvironmentsEnumUnknown        GetIntegrations200ApplicationJSONResultsSupportedEnvironmentsEnum = "Unknown"
	GetIntegrations200ApplicationJSONResultsSupportedEnvironmentsEnumSandboxOnly    GetIntegrations200ApplicationJSONResultsSupportedEnvironmentsEnum = "SandboxOnly"
	GetIntegrations200ApplicationJSONResultsSupportedEnvironmentsEnumLiveOnly       GetIntegrations200ApplicationJSONResultsSupportedEnvironmentsEnum = "LiveOnly"
	GetIntegrations200ApplicationJSONResultsSupportedEnvironmentsEnumLiveAndSandbox GetIntegrations200ApplicationJSONResultsSupportedEnvironmentsEnum = "LiveAndSandbox"
)

type GetIntegrations200ApplicationJSONResults struct {
	DatatypeFeatures       []GetIntegrations200ApplicationJSONResultsDatatypeFeatures         `json:"datatypeFeatures,omitempty"`
	Enabled                *bool                                                              `json:"enabled,omitempty"`
	IntegrationID          *string                                                            `json:"integrationId,omitempty"`
	IsBeta                 *bool                                                              `json:"isBeta,omitempty"`
	IsOfflineConnector     *bool                                                              `json:"isOfflineConnector,omitempty"`
	Key                    *string                                                            `json:"key,omitempty"`
	LinkedConnectionsCount *int                                                               `json:"linkedConnectionsCount,omitempty"`
	LogoURL                *string                                                            `json:"logoUrl,omitempty"`
	Name                   *string                                                            `json:"name,omitempty"`
	SourceID               *string                                                            `json:"sourceId,omitempty"`
	SourceType             *GetIntegrations200ApplicationJSONResultsSourceTypeEnum            `json:"sourceType,omitempty"`
	SupportedEnvironments  *GetIntegrations200ApplicationJSONResultsSupportedEnvironmentsEnum `json:"supportedEnvironments,omitempty"`
	TotalConnectionsCount  *int                                                               `json:"totalConnectionsCount,omitempty"`
}

// GetIntegrations200ApplicationJSON
// Used to represent what can be returned by an endpoint that supports paging.
// Usable with the [ProducesResponseType] attribute on a controller action.
type GetIntegrations200ApplicationJSON struct {
	Links        *GetIntegrations200ApplicationJSONLinks    `json:"_links,omitempty"`
	PageNumber   *int                                       `json:"pageNumber,omitempty"`
	PageSize     *int                                       `json:"pageSize,omitempty"`
	Results      []GetIntegrations200ApplicationJSONResults `json:"results,omitempty"`
	TotalResults *int                                       `json:"totalResults,omitempty"`
}

type GetIntegrationsResponse struct {
	ContentType                             string
	StatusCode                              int
	RawResponse                             *http.Response
	GetIntegrations200ApplicationJSONObject *GetIntegrations200ApplicationJSON
}
