// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// RecordReference - Links the current record to the underlying record or data type that created it.
//
// For example, if a journal entry is generated based on an invoice, this property allows you to connect the journal entry to the underlying invoice in our data model.
type RecordReference struct {
	// Allowed name of the 'dataType'.
	DataType *string `json:"dataType,omitempty"`
	// 'id' of the underlying record or data type.
	ID *string `json:"id,omitempty"`
}

func (o *RecordReference) GetDataType() *string {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *RecordReference) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AccountsReceivableTracking - Categories, and a project and customer, against which the item is tracked.
type AccountsReceivableTracking struct {
	CategoryRefs []TrackingCategoryRef  `json:"categoryRefs"`
	CustomerRef  *AccountingCustomerRef `json:"customerRef,omitempty"`
	// Defines if the bill or bill credit note is billed/rebilled to a project.
	IsBilledTo BilledToType1 `json:"isBilledTo"`
	// Defines if the bill or bill credit note is billed/rebilled to a project.
	IsRebilledTo BilledToType1 `json:"isRebilledTo"`
	ProjectRef   *ProjectRef   `json:"projectRef,omitempty"`
	// Links the current record to the underlying record or data type that created it.
	//
	// For example, if a journal entry is generated based on an invoice, this property allows you to connect the journal entry to the underlying invoice in our data model.
	RecordRef *RecordReference `json:"recordRef,omitempty"`
}

func (o *AccountsReceivableTracking) GetCategoryRefs() []TrackingCategoryRef {
	if o == nil {
		return []TrackingCategoryRef{}
	}
	return o.CategoryRefs
}

func (o *AccountsReceivableTracking) GetCustomerRef() *AccountingCustomerRef {
	if o == nil {
		return nil
	}
	return o.CustomerRef
}

func (o *AccountsReceivableTracking) GetIsBilledTo() BilledToType1 {
	if o == nil {
		return BilledToType1("")
	}
	return o.IsBilledTo
}

func (o *AccountsReceivableTracking) GetIsRebilledTo() BilledToType1 {
	if o == nil {
		return BilledToType1("")
	}
	return o.IsRebilledTo
}

func (o *AccountsReceivableTracking) GetProjectRef() *ProjectRef {
	if o == nil {
		return nil
	}
	return o.ProjectRef
}

func (o *AccountsReceivableTracking) GetRecordRef() *RecordReference {
	if o == nil {
		return nil
	}
	return o.RecordRef
}
