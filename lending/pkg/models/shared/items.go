// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Items struct {
	// Unique identifier for the group.
	ID *string `json:"id,omitempty"`
}

func (o *Items) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
