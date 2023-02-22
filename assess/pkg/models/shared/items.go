package shared

type ItemsReportComponentMeasuresMeasures struct {
	Index              *int64   `json:"index,omitempty"`
	MeasureDisplayName *string  `json:"measureDisplayName,omitempty"`
	Value              *float64 `json:"value,omitempty"`
}

type ItemsReportComponentMeasures struct {
	Components           []Items                                `json:"components,omitempty"`
	Dimension            *int64                                 `json:"dimension,omitempty"`
	DimensionDisplayName *string                                `json:"dimensionDisplayName,omitempty"`
	Item                 *int64                                 `json:"item,omitempty"`
	ItemDisplayName      *string                                `json:"itemDisplayName,omitempty"`
	Measures             []ItemsReportComponentMeasuresMeasures `json:"measures,omitempty"`
}

type ItemsMeasures struct {
	Index              *int64   `json:"index,omitempty"`
	MeasureDisplayName *string  `json:"measureDisplayName,omitempty"`
	Value              *float64 `json:"value,omitempty"`
}

type Items struct {
	Components           []ItemsReportComponentMeasures `json:"components,omitempty"`
	Dimension            *int64                         `json:"dimension,omitempty"`
	DimensionDisplayName *string                        `json:"dimensionDisplayName,omitempty"`
	Item                 *int64                         `json:"item,omitempty"`
	ItemDisplayName      *string                        `json:"itemDisplayName,omitempty"`
	Measures             []ItemsMeasures                `json:"measures,omitempty"`
}
