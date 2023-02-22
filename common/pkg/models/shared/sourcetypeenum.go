package shared

type SourceTypeEnum string

const (
	SourceTypeEnumAccounting SourceTypeEnum = "Accounting"
	SourceTypeEnumBanking    SourceTypeEnum = "Banking"
	SourceTypeEnumCommerce   SourceTypeEnum = "Commerce"
	SourceTypeEnumOther      SourceTypeEnum = "Other"
	SourceTypeEnumUnknown    SourceTypeEnum = "Unknown"
)
