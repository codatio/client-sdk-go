// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type EnhancedCashFlowItem struct {
	// An array of transaction data.
	Transactions []EnhancedCashFlowTransaction `json:"transactions,omitempty"`
}

func (o *EnhancedCashFlowItem) GetTransactions() []EnhancedCashFlowTransaction {
	if o == nil {
		return nil
	}
	return o.Transactions
}
