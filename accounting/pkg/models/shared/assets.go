package shared

type Assets struct {
	AccountID *string  `json:"accountId,omitempty"`
	Items     []Assets `json:"items,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Value     float64  `json:"value"`
}
