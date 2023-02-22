package shared

type SupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}
