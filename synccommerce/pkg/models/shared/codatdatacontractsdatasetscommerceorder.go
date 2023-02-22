package shared

import (
	"time"
)

type CodatDataContractsDatasetsCommerceOrder struct {
	ClosedDate         *time.Time                                        `json:"closedDate,omitempty"`
	Country            *string                                           `json:"country,omitempty"`
	CreatedDate        *time.Time                                        `json:"createdDate,omitempty"`
	Currency           *string                                           `json:"currency,omitempty"`
	CustomerRef        *CodatDataContractsDatasetsCommerceCustomerRef    `json:"customerRef,omitempty"`
	ID                 *string                                           `json:"id,omitempty"`
	LocationRef        *CodatDataContractsDatasetsCommerceLocationRef    `json:"locationRef,omitempty"`
	ModifiedDate       *time.Time                                        `json:"modifiedDate,omitempty"`
	OrderLineItems     []CodatDataContractsDatasetsCommerceOrderLineItem `json:"orderLineItems,omitempty"`
	OrderNumber        *string                                           `json:"orderNumber,omitempty"`
	Payments           []CodatDataContractsDatasetsCommercePaymentRef    `json:"payments,omitempty"`
	ServiceCharges     []CodatDataContractsDatasetsCommerceServiceCharge `json:"serviceCharges,omitempty"`
	SourceModifiedDate *time.Time                                        `json:"sourceModifiedDate,omitempty"`
	TotalAmount        *float64                                          `json:"totalAmount,omitempty"`
	TotalDiscount      *float64                                          `json:"totalDiscount,omitempty"`
	TotalGratuity      *float64                                          `json:"totalGratuity,omitempty"`
	TotalRefund        *float64                                          `json:"totalRefund,omitempty"`
	TotalTaxAmount     *float64                                          `json:"totalTaxAmount,omitempty"`
}
