package shared

type StatusEnum string

const (
	StatusEnumUnknown  StatusEnum = "Unknown"
	StatusEnumActive   StatusEnum = "Active"
	StatusEnumArchived StatusEnum = "Archived"
)
