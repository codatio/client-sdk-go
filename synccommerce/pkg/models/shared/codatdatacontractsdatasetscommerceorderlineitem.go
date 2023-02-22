package shared

type CodatDataContractsDatasetsCommerceOrderLineItem struct {
	DiscountAllocations []CodatDataContractsDatasetsCommerceDiscountAllocation `json:"discountAllocations,omitempty"`
	ID                  *string                                                `json:"id,omitempty"`
	ProductRef          *CodatDataContractsDatasetsCommerceProductRef          `json:"productRef,omitempty"`
	ProductVariantRef   *CodatDataContractsDatasetsCommerceProductRef          `json:"productVariantRef,omitempty"`
	Quantity            *float64                                               `json:"quantity,omitempty"`
	TaxPercentage       *float64                                               `json:"taxPercentage,omitempty"`
	TotalAmount         *float64                                               `json:"totalAmount,omitempty"`
	TotalTaxAmount      *float64                                               `json:"totalTaxAmount,omitempty"`
	UnitPrice           *float64                                               `json:"unitPrice,omitempty"`
}
