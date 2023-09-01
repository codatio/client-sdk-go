// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TransactionCategory struct {
	// Returns the confidence of the suggested category for the transaction. The value is between 0 and 100.
	Confidence *float64 `json:"confidence,omitempty"`
	// The suggested category is an ordered array of category levels where each element (or level) is a subcategory of the previous element (or level).
	Levels []string `json:"levels,omitempty"`
}

func (o *TransactionCategory) GetConfidence() *float64 {
	if o == nil {
		return nil
	}
	return o.Confidence
}

func (o *TransactionCategory) GetLevels() []string {
	if o == nil {
		return nil
	}
	return o.Levels
}
