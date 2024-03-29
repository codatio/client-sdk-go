// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GroupingPeriod struct {
	// Array of grouping period options.
	GroupingPeriodOptions []string `json:"groupingPeriodOptions,omitempty"`
	// Grouping period i.e. Daily sales.
	SelectedGroupingPeriod *string `json:"selectedGroupingPeriod,omitempty"`
}

func (o *GroupingPeriod) GetGroupingPeriodOptions() []string {
	if o == nil {
		return nil
	}
	return o.GroupingPeriodOptions
}

func (o *GroupingPeriod) GetSelectedGroupingPeriod() *string {
	if o == nil {
		return nil
	}
	return o.SelectedGroupingPeriod
}
