package shared

type IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum string

const (
	IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnumRelease        IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum = "Release"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnumBeta           IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum = "Beta"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnumDeprecated     IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum = "Deprecated"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnumNotSupported   IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum = "NotSupported"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnumNotImplemented IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum = "NotImplemented"
)

type IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum string

const (
	IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumGet                IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "Get"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumPost               IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "Post"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumCategorization     IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "Categorization"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumDelete             IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "Delete"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumPut                IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "Put"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumGetAsPdf           IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "GetAsPdf"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumDownloadAttachment IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "DownloadAttachment"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumGetAttachment      IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "GetAttachment"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumGetAttachments     IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "GetAttachments"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumUploadAttachment   IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "UploadAttachment"
)

type IntegrationDatatypeFeatureSupportedFeatures struct {
	FeatureState IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum `json:"featureState"`
	FeatureType  IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum  `json:"featureType"`
}

// IntegrationDatatypeFeature
// Describes support for a given datatype and associated operations
type IntegrationDatatypeFeature struct {
	Datatype          string                                        `json:"datatype"`
	SupportedFeatures []IntegrationDatatypeFeatureSupportedFeatures `json:"supportedFeatures"`
}

// Integration
// An integration that Codat supports
type Integration struct {
	DataProvidedBy     *string                      `json:"dataProvidedBy,omitempty"`
	DatatypeFeatures   []IntegrationDatatypeFeature `json:"datatypeFeatures,omitempty"`
	Enabled            bool                         `json:"enabled"`
	IntegrationID      *string                      `json:"integrationId,omitempty"`
	IsBeta             *bool                        `json:"isBeta,omitempty"`
	IsOfflineConnector *bool                        `json:"isOfflineConnector,omitempty"`
	Key                string                       `json:"key"`
	LogoURL            string                       `json:"logoUrl"`
	Name               string                       `json:"name"`
	SourceID           *string                      `json:"sourceId,omitempty"`
	SourceType         *SourceTypeEnum              `json:"sourceType,omitempty"`
}
