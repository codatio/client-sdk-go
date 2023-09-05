// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateBankTransactions struct {
	AccountID    *string                        `json:"accountId,omitempty"`
	Transactions []CreateBankAccountTransaction `json:"transactions,omitempty"`
}

func (o *CreateBankTransactions) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *CreateBankTransactions) GetTransactions() []CreateBankAccountTransaction {
	if o == nil {
		return nil
	}
	return o.Transactions
}