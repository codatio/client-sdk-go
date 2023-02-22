package shared

type ItemsTypeEnum1 string

const (
	ItemsTypeEnum1Unknown  ItemsTypeEnum1 = "Unknown"
	ItemsTypeEnum1Primary  ItemsTypeEnum1 = "Primary"
	ItemsTypeEnum1Landline ItemsTypeEnum1 = "Landline"
	ItemsTypeEnum1Mobile   ItemsTypeEnum1 = "Mobile"
	ItemsTypeEnum1Fax      ItemsTypeEnum1 = "Fax"
)

type PhoneNumbersitems struct {
	Number *string        `json:"number,omitempty"`
	Type   ItemsTypeEnum1 `json:"type"`
}
