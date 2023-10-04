// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Propertiestracking1 - Categories, and a project and customer, against which the item is tracked.
type Propertiestracking1 struct {
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

func (o *Propertiestracking1) GetCategoryRefs() []TrackingCategoryRef {
	if o == nil {
		return []TrackingCategoryRef{}
	}
	return o.CategoryRefs
}

func (o *Propertiestracking1) GetCustomerRef() *AccountingCustomerRef {
	if o == nil {
		return nil
	}
	return o.CustomerRef
}

func (o *Propertiestracking1) GetIsBilledTo() BilledToType1 {
	if o == nil {
		return BilledToType1("")
	}
	return o.IsBilledTo
}

func (o *Propertiestracking1) GetIsRebilledTo() BilledToType1 {
	if o == nil {
		return BilledToType1("")
	}
	return o.IsRebilledTo
}

func (o *Propertiestracking1) GetProjectRef() *ProjectRef {
	if o == nil {
		return nil
	}
	return o.ProjectRef
}

func (o *Propertiestracking1) GetRecordRef() *InvoiceTo {
	if o == nil {
		return nil
	}
	return o.RecordRef
}
