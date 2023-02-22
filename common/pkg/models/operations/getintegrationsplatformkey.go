package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

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

type GetIntegrationsPlatformKeyResponse struct {
	ContentType                                        string
	Integration                                        *shared.Integration
	StatusCode                                         int
	GetIntegrationsPlatformKey401ApplicationJSONObject *GetIntegrationsPlatformKey401ApplicationJSON
	GetIntegrationsPlatformKey404ApplicationJSONObject *GetIntegrationsPlatformKey404ApplicationJSON
}
