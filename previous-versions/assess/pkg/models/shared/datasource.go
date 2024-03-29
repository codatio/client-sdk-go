// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DataSource struct {
	// An array containing bank account data for each connected banking data source that have the following data types enabled: `banking-accounts`, `banking-transactions`.
	Accounts []Accounts `json:"accounts,omitempty"`
}

func (o *DataSource) GetAccounts() []Accounts {
	if o == nil {
		return nil
	}
	return o.Accounts
}
