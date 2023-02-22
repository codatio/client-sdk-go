package shared

type ExpenseReconciliationLines struct {
	AccountRef   RecordRef   `json:"accountRef"`
	NetAmount    float64     `json:"netAmount"`
	TaxAmount    float64     `json:"taxAmount"`
	TaxRateRef   RecordRef   `json:"taxRateRef"`
	TrackingRefs []RecordRef `json:"trackingRefs,omitempty"`
}
