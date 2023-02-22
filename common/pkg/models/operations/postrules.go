package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

type PostRulesRequest struct {
	Request *shared.Rule `request:"mediaType=application/json"`
}

type PostRules401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type PostRulesResponse struct {
	ContentType                       string
	Rule                              *shared.Rule
	StatusCode                        int
	PostRules401ApplicationJSONObject *PostRules401ApplicationJSON
}
