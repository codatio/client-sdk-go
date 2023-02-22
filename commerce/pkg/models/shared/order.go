package shared

import (
	"time"
)

// Order
// Orders contain the transaction details for all products sold by the company, and include details of any payments, service charges, or refunds related to each order.
//
// From the Orders endpoints you can retrieve:
//
// A list of all the orders for a commerce company:
// `GET /companies/{companyId}/connections/{connectionId}/data/commerce-orders`.
// The details of an individual order:
// `GET /companies/{companyId}/connections/{connectionId}/data/commerce-orders/{orderId}`.
// Note that for refunds `quantity` is a negative value and `unitPrice` is a positive value.
type Order struct {
	ClosedDate         *time.Time      `json:"closedDate,omitempty"`
	Country            *string         `json:"country,omitempty"`
	CreatedDate        *time.Time      `json:"createdDate,omitempty"`
	Currency           *string         `json:"currency,omitempty"`
	CustomerRef        *Zero           `json:"customerRef,omitempty"`
	ID                 string          `json:"id"`
	LocationRef        *Zero           `json:"locationRef,omitempty"`
	ModifiedDate       *time.Time      `json:"modifiedDate,omitempty"`
	OrderLineItems     []LineItem      `json:"orderLineItems,omitempty"`
	OrderNumber        *string         `json:"orderNumber,omitempty"`
	Payments           []PaymentRef    `json:"payments,omitempty"`
	ServiceCharges     []ServiceCharge `json:"serviceCharges,omitempty"`
	SourceModifiedDate *time.Time      `json:"sourceModifiedDate,omitempty"`
	TotalAmount        *float64        `json:"totalAmount,omitempty"`
	TotalDiscount      *float64        `json:"totalDiscount,omitempty"`
	TotalGratuity      *float64        `json:"totalGratuity,omitempty"`
	TotalRefund        *float64        `json:"totalRefund,omitempty"`
	TotalTaxAmount     *float64        `json:"totalTaxAmount,omitempty"`
}
