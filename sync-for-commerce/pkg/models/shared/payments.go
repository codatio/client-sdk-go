// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Payments struct {
	Accounts map[string]ConfigAccount `json:"accounts,omitempty"`
	// Boolean indicator for syncing sales.
	SyncPayments *bool `json:"syncPayments,omitempty"`
}

func (o *Payments) GetAccounts() map[string]ConfigAccount {
	if o == nil {
		return nil
	}
	return o.Accounts
}

func (o *Payments) GetSyncPayments() *bool {
	if o == nil {
		return nil
	}
	return o.SyncPayments
}