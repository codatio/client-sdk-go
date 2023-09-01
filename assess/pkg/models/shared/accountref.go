// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AccountRef - An account reference containing the account id and name.
type AccountRef struct {
	// The id of the account.
	ID *string `json:"id,omitempty"`
	// The name of the account.
	Name *string `json:"name,omitempty"`
}

func (o *AccountRef) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
