package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

type PutProfileRequest struct {
	Request *shared.Profile `request:"mediaType=application/json"`
}

type PutProfile401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type PutProfileResponse struct {
	ContentType                        string
	Profile                            *shared.Profile
	StatusCode                         int
	PutProfile401ApplicationJSONObject *PutProfile401ApplicationJSON
}
