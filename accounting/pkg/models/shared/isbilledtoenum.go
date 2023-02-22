package shared

type IsBilledToEnum string

const (
	IsBilledToEnumUnknown       IsBilledToEnum = "Unknown"
	IsBilledToEnumNotApplicable IsBilledToEnum = "NotApplicable"
	IsBilledToEnumCustomer      IsBilledToEnum = "Customer"
	IsBilledToEnumProject       IsBilledToEnum = "Project"
)
