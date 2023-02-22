package shared

// Details
// Dictionary list outlining the missing properties or allowed values.
type Details struct {
	PropertyDetail1 []string `json:"propertyDetail1,omitempty"`
	PropertyDetail2 []string `json:"propertyDetail2,omitempty"`
	PropertyDetailN []string `json:"propertyDetailN,omitempty"`
}
