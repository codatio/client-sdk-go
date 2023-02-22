package shared

type ProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}
