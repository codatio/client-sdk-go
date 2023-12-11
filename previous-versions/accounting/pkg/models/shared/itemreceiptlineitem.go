// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// ItemReceiptLineItemDataType - Allowed name of the 'dataType'.
type ItemReceiptLineItemDataType string

const (
	ItemReceiptLineItemDataTypePurchaseOrders ItemReceiptLineItemDataType = "purchaseOrders"
)

func (e ItemReceiptLineItemDataType) ToPointer() *ItemReceiptLineItemDataType {
	return &e
}

func (e *ItemReceiptLineItemDataType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "purchaseOrders":
		*e = ItemReceiptLineItemDataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ItemReceiptLineItemDataType: %v", v)
	}
}

// RecordLineReference - Reference to the item receipt line this line was generated from.
type RecordLineReference struct {
	// Allowed name of the 'dataType'.
	DataType *ItemReceiptLineItemDataType `json:"dataType,omitempty"`
	// 'id' of the underlying record.
	ID *string `json:"id,omitempty"`
	// Line number of the underlying record.
	LineNumber *string `json:"lineNumber,omitempty"`
}

func (o *RecordLineReference) GetDataType() *ItemReceiptLineItemDataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *RecordLineReference) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RecordLineReference) GetLineNumber() *string {
	if o == nil {
		return nil
	}
	return o.LineNumber
}

type ItemReceiptLineItemTracking struct {
	RecordRefs []InvoiceTo `json:"recordRefs,omitempty"`
}

func (o *ItemReceiptLineItemTracking) GetRecordRefs() []InvoiceTo {
	if o == nil {
		return nil
	}
	return o.RecordRefs
}

type ItemReceiptLineItem struct {
	// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
	AccountRef *AccountRef `json:"accountRef,omitempty"`
	// Friendly name of the item or services received.
	Description *string  `json:"description,omitempty"`
	ItemRef     *ItemRef `json:"itemRef,omitempty"`
	// The item receipt line's number.
	LineNumber           *string              `json:"lineNumber,omitempty"`
	PurchaseOrderLineRef *RecordLineReference `json:"purchaseOrderLineRef,omitempty"`
	// Number of units of item or services received.
	Quantity *decimal.Big `decimal:"number" json:"quantity,omitempty"`
	// Amount of the line, inclusive of discounts but exclusive of tax.
	SubTotal *decimal.Big `decimal:"number" json:"subTotal,omitempty"`
	// Total amount of the line, including tax.
	TotalAmount *decimal.Big                 `decimal:"number" json:"totalAmount,omitempty"`
	Tracking    *ItemReceiptLineItemTracking `json:"tracking,omitempty"`
	// Price of each unit of item or services.
	UnitAmount *decimal.Big `decimal:"number" json:"unitAmount,omitempty"`
	// The measurement which defines a unit for this item (e.g. 'kilogram', 'litre').
	UnitOfMeasurement *string `json:"unitOfMeasurement,omitempty"`
}

func (i ItemReceiptLineItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *ItemReceiptLineItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ItemReceiptLineItem) GetAccountRef() *AccountRef {
	if o == nil {
		return nil
	}
	return o.AccountRef
}

func (o *ItemReceiptLineItem) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ItemReceiptLineItem) GetItemRef() *ItemRef {
	if o == nil {
		return nil
	}
	return o.ItemRef
}

func (o *ItemReceiptLineItem) GetLineNumber() *string {
	if o == nil {
		return nil
	}
	return o.LineNumber
}

func (o *ItemReceiptLineItem) GetPurchaseOrderLineRef() *RecordLineReference {
	if o == nil {
		return nil
	}
	return o.PurchaseOrderLineRef
}

func (o *ItemReceiptLineItem) GetQuantity() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *ItemReceiptLineItem) GetSubTotal() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.SubTotal
}

func (o *ItemReceiptLineItem) GetTotalAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *ItemReceiptLineItem) GetTracking() *ItemReceiptLineItemTracking {
	if o == nil {
		return nil
	}
	return o.Tracking
}

func (o *ItemReceiptLineItem) GetUnitAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.UnitAmount
}

func (o *ItemReceiptLineItem) GetUnitOfMeasurement() *string {
	if o == nil {
		return nil
	}
	return o.UnitOfMeasurement
}