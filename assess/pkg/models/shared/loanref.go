// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type LoanRef struct {
	// The dataConnectionId the object being referred to is associated with.
	DataConnectionID *string `json:"dataConnectionId,omitempty"`
	// The id of the object being referred to.
	ID *string `json:"id,omitempty"`
	// The object type data is referring to, e.g. Account.
	Type *string `json:"type,omitempty"`
}
