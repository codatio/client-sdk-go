package operations

type GetIntegrationsPlatformKeyPathParams struct {
	PlatformKey string `pathParam:"style=simple,explode=false,name=platformKey"`
}

type GetIntegrationsPlatformKeyRequest struct {
	PathParams GetIntegrationsPlatformKeyPathParams
}

type GetIntegrationsPlatformKey404ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetIntegrationsPlatformKey401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum string

const (
	GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnumRelease        GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum = "Release"
	GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnumBeta           GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum = "Beta"
	GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnumDeprecated     GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum = "Deprecated"
	GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnumNotSupported   GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum = "NotSupported"
	GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnumNotImplemented GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum = "NotImplemented"
)

type GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum string

const (
	GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumGet                GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "Get"
	GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumPost               GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "Post"
	GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumCategorization     GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "Categorization"
	GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumDelete             GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "Delete"
	GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumPut                GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "Put"
	GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumGetAsPdf           GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "GetAsPdf"
	GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumDownloadAttachment GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "DownloadAttachment"
	GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumGetAttachment      GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "GetAttachment"
	GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumGetAttachments     GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "GetAttachments"
	GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumUploadAttachment   GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "UploadAttachment"
)

type GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeatures struct {
	FeatureState GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum `json:"featureState"`
	FeatureType  GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum  `json:"featureType"`
}

// GetIntegrationsPlatformKeyIntegrationDatatypeFeature
// Describes support for a given datatype and associated operations
type GetIntegrationsPlatformKeyIntegrationDatatypeFeature struct {
	Datatype          string                                                                  `json:"datatype"`
	SupportedFeatures []GetIntegrationsPlatformKeyIntegrationDatatypeFeatureSupportedFeatures `json:"supportedFeatures"`
}

type GetIntegrationsPlatformKeyIntegrationSourceTypeEnum string

const (
	GetIntegrationsPlatformKeyIntegrationSourceTypeEnumAccounting GetIntegrationsPlatformKeyIntegrationSourceTypeEnum = "Accounting"
	GetIntegrationsPlatformKeyIntegrationSourceTypeEnumBanking    GetIntegrationsPlatformKeyIntegrationSourceTypeEnum = "Banking"
	GetIntegrationsPlatformKeyIntegrationSourceTypeEnumCommerce   GetIntegrationsPlatformKeyIntegrationSourceTypeEnum = "Commerce"
	GetIntegrationsPlatformKeyIntegrationSourceTypeEnumOther      GetIntegrationsPlatformKeyIntegrationSourceTypeEnum = "Other"
	GetIntegrationsPlatformKeyIntegrationSourceTypeEnumUnknown    GetIntegrationsPlatformKeyIntegrationSourceTypeEnum = "Unknown"
)

// GetIntegrationsPlatformKeyIntegration
// An integration that Codat supports
type GetIntegrationsPlatformKeyIntegration struct {
	DataProvidedBy     *string                                                `json:"dataProvidedBy,omitempty"`
	DatatypeFeatures   []GetIntegrationsPlatformKeyIntegrationDatatypeFeature `json:"datatypeFeatures,omitempty"`
	Enabled            bool                                                   `json:"enabled"`
	IntegrationID      *string                                                `json:"integrationId,omitempty"`
	IsBeta             *bool                                                  `json:"isBeta,omitempty"`
	IsOfflineConnector *bool                                                  `json:"isOfflineConnector,omitempty"`
	Key                string                                                 `json:"key"`
	LogoURL            string                                                 `json:"logoUrl"`
	Name               string                                                 `json:"name"`
	SourceID           *string                                                `json:"sourceId,omitempty"`
	SourceType         *GetIntegrationsPlatformKeyIntegrationSourceTypeEnum   `json:"sourceType,omitempty"`
}

type GetIntegrationsPlatformKeyResponse struct {
	ContentType                                        string
	Integration                                        *GetIntegrationsPlatformKeyIntegration
	StatusCode                                         int
	GetIntegrationsPlatformKey401ApplicationJSONObject *GetIntegrationsPlatformKey401ApplicationJSON
	GetIntegrationsPlatformKey404ApplicationJSONObject *GetIntegrationsPlatformKey404ApplicationJSON
}
