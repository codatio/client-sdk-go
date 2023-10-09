// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/utils"
)

type Grouping struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	GroupingLevels       *GroupingLevels        `json:"groupingLevels,omitempty"`
	GroupingPeriod       *GroupingPeriod        `json:"groupingPeriod,omitempty"`
}

func (g Grouping) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *Grouping) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Grouping) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
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
