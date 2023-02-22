package shared

type DataIntegritySummaryByAmount struct {
	Currency        *string  `json:"currency,omitempty"`
	MatchPercentage *float64 `json:"matchPercentage,omitempty"`
	Matched         *float64 `json:"matched,omitempty"`
	Total           *float64 `json:"total,omitempty"`
	Unmatched       *float64 `json:"unmatched,omitempty"`
}

type DataIntegritySummaryByCount struct {
	MatchPercentage *float64 `json:"matchPercentage,omitempty"`
	Matched         *float64 `json:"matched,omitempty"`
	Total           *float64 `json:"total,omitempty"`
	Unmatched       *float64 `json:"unmatched,omitempty"`
}

type DataIntegritySummary struct {
	ByAmount *DataIntegritySummaryByAmount `json:"byAmount,omitempty"`
	ByCount  *DataIntegritySummaryByCount  `json:"byCount,omitempty"`
	Type     *string                       `json:"type,omitempty"`
}
