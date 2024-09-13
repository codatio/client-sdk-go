// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PropertieTracking1 - Categories, and a project and customer, against which the item is tracked.
type PropertieTracking1 struct {
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
	RecordRef *InvoiceTo `json:"recordRef,omitempty"`
}

func (o *PropertieTracking1) GetCategoryRefs() []TrackingCategoryRef {
	if o == nil {
		return []TrackingCategoryRef{}
	}
	return o.CategoryRefs
}

func (o *PropertieTracking1) GetCustomerRef() *AccountingCustomerRef {
	if o == nil {
		return nil
	}
	return o.CustomerRef
}

func (o *PropertieTracking1) GetIsBilledTo() BilledToType1 {
	if o == nil {
		return BilledToType1("")
	}
	return o.IsBilledTo
}

func (o *PropertieTracking1) GetIsRebilledTo() BilledToType1 {
	if o == nil {
		return BilledToType1("")
	}
	return o.IsRebilledTo
}

func (o *PropertieTracking1) GetProjectRef() *ProjectRef {
	if o == nil {
		return nil
	}
	return o.ProjectRef
}

func (o *PropertieTracking1) GetRecordRef() *InvoiceTo {
	if o == nil {
		return nil
	}
	return o.RecordRef
}
