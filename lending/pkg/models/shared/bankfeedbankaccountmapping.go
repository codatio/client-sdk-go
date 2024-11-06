// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type BankFeedBankAccountMapping struct {
	// Unique ID for the source account
	SourceAccountID string `json:"sourceAccountId"`
	// Unique ID for the target account
	TargetAccountID *string `json:"targetAccountId,omitempty"`
}

func (o *BankFeedBankAccountMapping) GetSourceAccountID() string {
	if o == nil {
		return ""
	}
	return o.SourceAccountID
}

func (o *BankFeedBankAccountMapping) GetTargetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.TargetAccountID
}
