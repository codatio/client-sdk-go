package shared

type Items struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}
