package operations

import (
	"net/http"
)

type GetSettingsProfile401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

// GetSettingsProfileProfile
// Describes your Codat client instance
type GetSettingsProfileProfile struct {
	AlertAuthHeader    *string  `json:"alertAuthHeader,omitempty"`
	APIKey             *string  `json:"apiKey,omitempty"`
	ConfirmCompanyName *bool    `json:"confirmCompanyName,omitempty"`
	IconURL            *string  `json:"iconUrl,omitempty"`
	LogoURL            *string  `json:"logoUrl,omitempty"`
	Name               string   `json:"name"`
	RedirectURL        string   `json:"redirectUrl"`
	WhiteListUrls      []string `json:"whiteListUrls,omitempty"`
}

type GetSettingsProfileResponse struct {
	ContentType                                string
	Profile                                    *GetSettingsProfileProfile
	StatusCode                                 int
	RawResponse                                *http.Response
	GetSettingsProfile401ApplicationJSONObject *GetSettingsProfile401ApplicationJSON
}
