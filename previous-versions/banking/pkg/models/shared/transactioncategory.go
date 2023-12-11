// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TransactionCategory - Status of the bank transaction category.
type TransactionCategory struct {
	// A Boolean indicating whether there are other bank transaction categories beneath this one in the hierarchy.
	HasChildren *bool `json:"hasChildren,omitempty"`
	// The unique identifier of the bank transaction category.
	ID           string  `json:"id"`
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// The name of the bank transaction category.
	Name string `json:"name"`
	// The unique identifier of the parent bank transaction category.
	ParentID           *string `json:"parentId,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// The status of the transaction category.
	Status *TransactionCategoryStatus `json:"status,omitempty"`
}

func (o *TransactionCategory) GetHasChildren() *bool {
	if o == nil {
		return nil
	}
	return o.HasChildren
}

func (o *TransactionCategory) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TransactionCategory) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *TransactionCategory) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *TransactionCategory) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *TransactionCategory) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *TransactionCategory) GetStatus() *TransactionCategoryStatus {
	if o == nil {
		return nil
	}
	return o.Status
}
