package shared

type SchemeAuthHeader struct {
	APIKey string `security:"name=Authorization"`
}

type Security struct {
	AuthHeader SchemeAuthHeader `security:"scheme,type=apiKey,subtype=header"`
}
