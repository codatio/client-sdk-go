package shared

// Profile
// Describes your Codat client instance
type Profile struct {
	AlertAuthHeader    *string  `json:"alertAuthHeader,omitempty"`
	APIKey             *string  `json:"apiKey,omitempty"`
	ConfirmCompanyName *bool    `json:"confirmCompanyName,omitempty"`
	IconURL            *string  `json:"iconUrl,omitempty"`
	LogoURL            *string  `json:"logoUrl,omitempty"`
	Name               string   `json:"name"`
	RedirectURL        string   `json:"redirectUrl"`
	WhiteListUrls      []string `json:"whiteListUrls,omitempty"`
}
