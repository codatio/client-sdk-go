package shared

type ReportDimensionsItems struct {
	Index *int64 `json:"index,omitempty"`
}

type ReportDimensions struct {
	DisplayName *string                 `json:"displayName,omitempty"`
	Index       *int64                  `json:"index,omitempty"`
	Items       []ReportDimensionsItems `json:"items,omitempty"`
	Type        *string                 `json:"type,omitempty"`
}

type ReportErrorsDetails struct {
	AdditionalProp1 []string `json:"additionalProp1,omitempty"`
	AdditionalProp2 []string `json:"additionalProp2,omitempty"`
	AdditionalProp3 []string `json:"additionalProp3,omitempty"`
}

type ReportErrors struct {
	Details *ReportErrorsDetails `json:"details,omitempty"`
	Message *string              `json:"message,omitempty"`
	Type    *string              `json:"type,omitempty"`
}

type ReportMeasures struct {
	DisplayName *string `json:"displayName,omitempty"`
	Index       *int64  `json:"index,omitempty"`
	Type        *string `json:"type,omitempty"`
	Units       *string `json:"units,omitempty"`
}

type ReportReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresMeasures struct {
	Index              *int64   `json:"index,omitempty"`
	MeasureDisplayName *string  `json:"measureDisplayName,omitempty"`
	Value              *float64 `json:"value,omitempty"`
}

type ReportReportDimensionsItemsReportComponentMeasuresReportComponentMeasures struct {
	Components           []Items                                                                             `json:"components,omitempty"`
	Dimension            *int64                                                                              `json:"dimension,omitempty"`
	DimensionDisplayName *string                                                                             `json:"dimensionDisplayName,omitempty"`
	Item                 *int64                                                                              `json:"item,omitempty"`
	ItemDisplayName      *string                                                                             `json:"itemDisplayName,omitempty"`
	Measures             []ReportReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresMeasures `json:"measures,omitempty"`
}

type ReportReportDimensionsItemsReportComponentMeasuresMeasures struct {
	Index              *int64   `json:"index,omitempty"`
	MeasureDisplayName *string  `json:"measureDisplayName,omitempty"`
	Value              *float64 `json:"value,omitempty"`
}

type ReportReportDimensionsItemsReportComponentMeasures struct {
	Components           []ReportReportDimensionsItemsReportComponentMeasuresReportComponentMeasures `json:"components,omitempty"`
	Dimension            *int64                                                                      `json:"dimension,omitempty"`
	DimensionDisplayName *string                                                                     `json:"dimensionDisplayName,omitempty"`
	Item                 *int64                                                                      `json:"item,omitempty"`
	ItemDisplayName      *string                                                                     `json:"itemDisplayName,omitempty"`
	Measures             []ReportReportDimensionsItemsReportComponentMeasuresMeasures                `json:"measures,omitempty"`
}

type ReportReportDimensionsItems struct {
	Components           []ReportReportDimensionsItemsReportComponentMeasures `json:"components,omitempty"`
	Dimension            *int64                                               `json:"dimension,omitempty"`
	DimensionDisplayName *string                                              `json:"dimensionDisplayName,omitempty"`
	Item                 *int64                                               `json:"item,omitempty"`
	ItemDisplayName      *string                                              `json:"itemDisplayName,omitempty"`
}

type ReportReportInfo struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty"`
}

// Report
// Assess reports follow a consistent structure. Reports contain four sections of information:
//
//  1. Report definition information such as:
//     a. The report info (e.g. enhanced_profit_and_loss).
//     b. The display name of the report (e.g. Enhanced Profit and Loss).
//  2. Information about the dimension contained in the reports such as:
//     a. The type of dimension (e.g. datetime, recordRef).
//     b. The display name of the dimension (e.g. Period, Category type, Category sub type).
//     c. The details about each item within the dimension (e.g. displayName:"Jan 2022", start:"...", end:"...", id:"...", name:"...").
//  3. Information about the measures contained in the report such as:
//     a. The display name of the measure (e.g. value of account, percentage change).
//     b. The type of the measure (e.g. currency, percentage).
//     c. The unit of the measure (e.g. %, GBP).
//  4. The data for the report. When the *includeDisplayName* parameter is set to *true*, it shows the *dimensionDisplayName* and *itemDisplayName* to make the data human-readable. The default setting for *includeDisplayName* is *false*.
//
// Reports can be rendered as follows (ordering is implicit rather than explicit):
//
// ![A table showing an example of how a report can be rendered](https://files.readme.io/1fa20ca-Report1.png)
//
// # Data model
//
// ## Dimensions
type Report struct {
	Dimensions []ReportDimensions            `json:"dimensions,omitempty"`
	Errors     []ReportErrors                `json:"errors,omitempty"`
	Measures   []ReportMeasures              `json:"measures,omitempty"`
	ReportData []ReportReportDimensionsItems `json:"reportData,omitempty"`
	ReportInfo *ReportReportInfo             `json:"reportInfo,omitempty"`
}
