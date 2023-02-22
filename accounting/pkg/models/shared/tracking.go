package shared

type TrackingIsBilledToEnum string

const (
	TrackingIsBilledToEnumUnknown       TrackingIsBilledToEnum = "Unknown"
	TrackingIsBilledToEnumNotApplicable TrackingIsBilledToEnum = "NotApplicable"
	TrackingIsBilledToEnumCustomer      TrackingIsBilledToEnum = "Customer"
	TrackingIsBilledToEnumProject       TrackingIsBilledToEnum = "Project"
)

type Tracking struct {
	CategoryRefs []Items                `json:"categoryRefs"`
	CustomerRef  *CustomerRef           `json:"customerRef,omitempty"`
	IsBilledTo   TrackingIsBilledToEnum `json:"isBilledTo"`
	IsRebilledTo IsBilledToEnum         `json:"isRebilledTo"`
	ProjectRef   *ProjectRef            `json:"projectRef,omitempty"`
}
