package shared

type SchemeAPIKey struct {
	APIKey string `security:"name=Your Base 64 encoded API Key"`
}

type Security struct {
	APIKey SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}
