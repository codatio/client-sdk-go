package operations

import (
	"net/http"
)

// PutProfileProfile
// Describes your Codat client instance
type PutProfileProfile struct {
	AlertAuthHeader    *string  `json:"alertAuthHeader,omitempty"`
	APIKey             *string  `json:"apiKey,omitempty"`
	ConfirmCompanyName *bool    `json:"confirmCompanyName,omitempty"`
	IconURL            *string  `json:"iconUrl,omitempty"`
	LogoURL            *string  `json:"logoUrl,omitempty"`
	Name               string   `json:"name"`
	RedirectURL        string   `json:"redirectUrl"`
	WhiteListUrls      []string `json:"whiteListUrls,omitempty"`
}

type PutProfileRequest struct {
	Request *PutProfileProfile `request:"mediaType=application/json"`
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
	Profile                            *PutProfileProfile
	StatusCode                         int
	RawResponse                        *http.Response
	PutProfile401ApplicationJSONObject *PutProfile401ApplicationJSON
}
