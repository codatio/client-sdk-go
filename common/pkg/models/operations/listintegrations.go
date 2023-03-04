package operations

import (
	"net/http"
)

type ListIntegrationsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListIntegrationsRequest struct {
	QueryParams ListIntegrationsQueryParams
}

type ListIntegrations401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type ListIntegrations400ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type ListIntegrationsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListIntegrationsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListIntegrationsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListIntegrationsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListIntegrationsLinksLinks struct {
	Current  ListIntegrationsLinksLinksCurrent   `json:"current"`
	Next     *ListIntegrationsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListIntegrationsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListIntegrationsLinksLinksSelf      `json:"self"`
}

type ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum string

const (
	ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnumRelease        ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum = "Release"
	ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnumBeta           ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum = "Beta"
	ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnumDeprecated     ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum = "Deprecated"
	ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnumNotSupported   ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum = "NotSupported"
	ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnumNotImplemented ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum = "NotImplemented"
)

type ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum string

const (
	ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumGet                ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "Get"
	ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumPost               ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "Post"
	ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumCategorization     ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "Categorization"
	ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumDelete             ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "Delete"
	ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumPut                ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "Put"
	ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumGetAsPdf           ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "GetAsPdf"
	ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumDownloadAttachment ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "DownloadAttachment"
	ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumGetAttachment      ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "GetAttachment"
	ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumGetAttachments     ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "GetAttachments"
	ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumUploadAttachment   ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "UploadAttachment"
)

type ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeatures struct {
	FeatureState ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum `json:"featureState"`
	FeatureType  ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum  `json:"featureType"`
}

// ListIntegrationsLinksIntegrationDatatypeFeature
// Describes support for a given datatype and associated operations
type ListIntegrationsLinksIntegrationDatatypeFeature struct {
	Datatype          string                                                             `json:"datatype"`
	SupportedFeatures []ListIntegrationsLinksIntegrationDatatypeFeatureSupportedFeatures `json:"supportedFeatures"`
}

type ListIntegrationsLinksIntegrationSourceTypeEnum string

const (
	ListIntegrationsLinksIntegrationSourceTypeEnumAccounting ListIntegrationsLinksIntegrationSourceTypeEnum = "Accounting"
	ListIntegrationsLinksIntegrationSourceTypeEnumBanking    ListIntegrationsLinksIntegrationSourceTypeEnum = "Banking"
	ListIntegrationsLinksIntegrationSourceTypeEnumCommerce   ListIntegrationsLinksIntegrationSourceTypeEnum = "Commerce"
	ListIntegrationsLinksIntegrationSourceTypeEnumOther      ListIntegrationsLinksIntegrationSourceTypeEnum = "Other"
	ListIntegrationsLinksIntegrationSourceTypeEnumUnknown    ListIntegrationsLinksIntegrationSourceTypeEnum = "Unknown"
)

// ListIntegrationsLinksIntegration
// An integration that Codat supports
type ListIntegrationsLinksIntegration struct {
	DataProvidedBy     *string                                           `json:"dataProvidedBy,omitempty"`
	DatatypeFeatures   []ListIntegrationsLinksIntegrationDatatypeFeature `json:"datatypeFeatures,omitempty"`
	Enabled            bool                                              `json:"enabled"`
	IntegrationID      *string                                           `json:"integrationId,omitempty"`
	IsBeta             *bool                                             `json:"isBeta,omitempty"`
	IsOfflineConnector *bool                                             `json:"isOfflineConnector,omitempty"`
	Key                string                                            `json:"key"`
	LogoURL            string                                            `json:"logoUrl"`
	Name               string                                            `json:"name"`
	SourceID           *string                                           `json:"sourceId,omitempty"`
	SourceType         *ListIntegrationsLinksIntegrationSourceTypeEnum   `json:"sourceType,omitempty"`
}

// ListIntegrationsLinks
// Codat's Paging Model
type ListIntegrationsLinks struct {
	Links        ListIntegrationsLinksLinks         `json:"_links"`
	PageNumber   int64                              `json:"pageNumber"`
	PageSize     int64                              `json:"pageSize"`
	Results      []ListIntegrationsLinksIntegration `json:"results,omitempty"`
	TotalResults int64                              `json:"totalResults"`
}

type ListIntegrationsResponse struct {
	ContentType                              string
	StatusCode                               int
	RawResponse                              *http.Response
	Links                                    *ListIntegrationsLinks
	ListIntegrations400ApplicationJSONObject *ListIntegrations400ApplicationJSON
	ListIntegrations401ApplicationJSONObject *ListIntegrations401ApplicationJSON
}
