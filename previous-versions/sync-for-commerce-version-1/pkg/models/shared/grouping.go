// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Grouping struct {
	GroupingLevels *GroupingLevels `json:"groupingLevels,omitempty"`
	GroupingPeriod *GroupingPeriod `json:"groupingPeriod,omitempty"`
}

func (o *Grouping) GetGroupingLevels() *GroupingLevels {
	if o == nil {
		return nil
	}
	return o.GroupingLevels
}

func (o *Grouping) GetGroupingPeriod() *GroupingPeriod {
	if o == nil {
		return nil
	}
	return o.GroupingPeriod
}
