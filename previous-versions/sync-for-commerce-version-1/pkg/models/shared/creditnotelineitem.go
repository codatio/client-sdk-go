// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type CreditNoteLineItemAccountingProjectReference struct {
	// Unique identifier to the project reference.
	ID string `json:"id"`
	// The project's name.
	Name *string `json:"name,omitempty"`
}

func (o *CreditNoteLineItemAccountingProjectReference) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreditNoteLineItemAccountingProjectReference) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreditNoteLineItemTracking - Categories, and a project and customer, against which the item is tracked.
type CreditNoteLineItemTracking struct {
	CategoryRefs []TrackingCategoryRefItems `json:"categoryRefs"`
	CustomerRef  *AccountingCustomerRef     `json:"customerRef,omitempty"`
	// Defines if the bill or bill credit note is billed/rebilled to a project.
	IsBilledTo BilledToType `json:"isBilledTo"`
	// Defines if the bill or bill credit note is billed/rebilled to a project.
	IsRebilledTo BilledToType                                  `json:"isRebilledTo"`
	ProjectRef   *CreditNoteLineItemAccountingProjectReference `json:"projectRef,omitempty"`
	// Links the current record to the underlying record or data type that created it.
	//
	// For example, if a journal entry is generated based on an invoice, this property allows you to connect the journal entry to the underlying invoice in our data model.
	RecordRef *RecordRef `json:"recordRef,omitempty"`
}

func (o *CreditNoteLineItemTracking) GetCategoryRefs() []TrackingCategoryRefItems {
	if o == nil {
		return []TrackingCategoryRefItems{}
	}
	return o.CategoryRefs
}

func (o *CreditNoteLineItemTracking) GetCustomerRef() *AccountingCustomerRef {
	if o == nil {
		return nil
	}
	return o.CustomerRef
}

func (o *CreditNoteLineItemTracking) GetIsBilledTo() BilledToType {
	if o == nil {
		return BilledToType("")
	}
	return o.IsBilledTo
}

func (o *CreditNoteLineItemTracking) GetIsRebilledTo() BilledToType {
	if o == nil {
		return BilledToType("")
	}
	return o.IsRebilledTo
}

func (o *CreditNoteLineItemTracking) GetProjectRef() *CreditNoteLineItemAccountingProjectReference {
	if o == nil {
		return nil
	}
	return o.ProjectRef
}

func (o *CreditNoteLineItemTracking) GetRecordRef() *RecordRef {
	if o == nil {
		return nil
	}
	return o.RecordRef
}

type CreditNoteLineItem struct {
	// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
	AccountRef *AccountRef `json:"accountRef,omitempty"`
	// Friendly name of each line item. For example, the goods or service for which credit has been issued.
	Description *string `json:"description,omitempty"`
	// Value of any discounts applied.
	DiscountAmount *decimal.Big `decimal:"number" json:"discountAmount,omitempty"`
	// Percentage rate of any discount applied to the line item.
	DiscountPercentage *decimal.Big `decimal:"number" json:"discountPercentage,omitempty"`
	// The credit note is a direct income if `True`.
	IsDirectIncome *bool `json:"isDirectIncome,omitempty"`
	// Reference to the product, service type, or inventory item to which the direct cost is linked.
	ItemRef *ItemRef `json:"itemRef,omitempty"`
	// Number of units of the goods or service for which credit has been issued.
	Quantity *decimal.Big `decimal:"number" json:"quantity"`
	// Amount of credit associated with the line item, including discounts but excluding tax.
	SubTotal *decimal.Big `decimal:"number" json:"subTotal,omitempty"`
	// Amount of tax associated with the line item.
	TaxAmount *decimal.Big `decimal:"number" json:"taxAmount,omitempty"`
	// Reference to the tax rate to which the line item is linked.
	TaxRateRef *TaxRateRef `json:"taxRateRef,omitempty"`
	// Total amount of the line item, including discounts and tax.
	TotalAmount *decimal.Big `decimal:"number" json:"totalAmount,omitempty"`
	// Categories, and a project and customer, against which the item is tracked.
	Tracking *CreditNoteLineItemTracking `json:"tracking,omitempty"`
	// Reference to the tracking categories to which the line item is linked.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	TrackingCategoryRefs []TrackingCategoryRefItems `json:"trackingCategoryRefs,omitempty"`
	// Unit price of the goods or service.
	UnitAmount *decimal.Big `decimal:"number" json:"unitAmount"`
}

func (c CreditNoteLineItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreditNoteLineItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreditNoteLineItem) GetAccountRef() *AccountRef {
	if o == nil {
		return nil
	}
	return o.AccountRef
}

func (o *CreditNoteLineItem) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreditNoteLineItem) GetDiscountAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.DiscountAmount
}

func (o *CreditNoteLineItem) GetDiscountPercentage() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.DiscountPercentage
}

func (o *CreditNoteLineItem) GetIsDirectIncome() *bool {
	if o == nil {
		return nil
	}
	return o.IsDirectIncome
}

func (o *CreditNoteLineItem) GetItemRef() *ItemRef {
	if o == nil {
		return nil
	}
	return o.ItemRef
}

func (o *CreditNoteLineItem) GetQuantity() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.Quantity
}

func (o *CreditNoteLineItem) GetSubTotal() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.SubTotal
}

func (o *CreditNoteLineItem) GetTaxAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *CreditNoteLineItem) GetTaxRateRef() *TaxRateRef {
	if o == nil {
		return nil
	}
	return o.TaxRateRef
}

func (o *CreditNoteLineItem) GetTotalAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *CreditNoteLineItem) GetTracking() *CreditNoteLineItemTracking {
	if o == nil {
		return nil
	}
	return o.Tracking
}

func (o *CreditNoteLineItem) GetTrackingCategoryRefs() []TrackingCategoryRefItems {
	if o == nil {
		return nil
	}
	return o.TrackingCategoryRefs
}

func (o *CreditNoteLineItem) GetUnitAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.UnitAmount
}
