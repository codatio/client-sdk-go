package shared

type Security struct {
	AuthHeader string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}
