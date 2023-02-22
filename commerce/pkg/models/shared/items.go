package shared

type Items struct {
	City       *string      `json:"city,omitempty"`
	Country    *string      `json:"country,omitempty"`
	Line1      *string      `json:"line1,omitempty"`
	Line2      *string      `json:"line2,omitempty"`
	PostalCode *string      `json:"postalCode,omitempty"`
	Region     *string      `json:"region,omitempty"`
	Type       *interface{} `json:"type,omitempty"`
}
