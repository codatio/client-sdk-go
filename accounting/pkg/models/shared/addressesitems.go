package shared

type ItemsTypeEnum string

const (
	ItemsTypeEnumUnknown  ItemsTypeEnum = "Unknown"
	ItemsTypeEnumBilling  ItemsTypeEnum = "Billing"
	ItemsTypeEnumDelivery ItemsTypeEnum = "Delivery"
)

type Addressesitems struct {
	City       *string       `json:"city,omitempty"`
	Country    *string       `json:"country,omitempty"`
	Line1      *string       `json:"line1,omitempty"`
	Line2      *string       `json:"line2,omitempty"`
	PostalCode *string       `json:"postalCode,omitempty"`
	Region     *string       `json:"region,omitempty"`
	Type       ItemsTypeEnum `json:"type"`
}
