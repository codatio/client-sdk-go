// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type BankFeedAccountMappingResponse struct {
	// Error returned during the post request
	Error *string `json:"error,omitempty"`
	// Unique ID for the source account.
	SourceAccountID *string `json:"sourceAccountId,omitempty"`
	// Status of the POST request.
	Status *string `json:"status,omitempty"`
	// Unique ID for the target account.
	TargetAccountID *string `json:"targetAccountId,omitempty"`
}

func (o *BankFeedAccountMappingResponse) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *BankFeedAccountMappingResponse) GetSourceAccountID() *string {
	if o == nil {
		return nil
	}
	return o.SourceAccountID
}

func (o *BankFeedAccountMappingResponse) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *BankFeedAccountMappingResponse) GetTargetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.TargetAccountID
}
