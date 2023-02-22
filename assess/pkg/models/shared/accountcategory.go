package shared

type AccountCategory struct {
	DetailType *string `json:"detailType,omitempty"`
	Subtype    *string `json:"subtype,omitempty"`
	Type       *string `json:"type,omitempty"`
}
