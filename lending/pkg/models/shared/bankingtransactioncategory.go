// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// BankingTransactionCategory - The Banking Transaction Categories data type provides a list of hierarchical categories associated with a transaction for greater contextual meaning to transaction activity.
//
// Responses are paged, so you should provide `page` and `pageSize` query parameters in your request.
type BankingTransactionCategory struct {
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

func (o *BankingTransactionCategory) GetHasChildren() *bool {
	if o == nil {
		return nil
	}
	return o.HasChildren
}

func (o *BankingTransactionCategory) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *BankingTransactionCategory) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *BankingTransactionCategory) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *BankingTransactionCategory) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *BankingTransactionCategory) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *BankingTransactionCategory) GetStatus() *TransactionCategoryStatus {
	if o == nil {
		return nil
	}
	return o.Status
}
