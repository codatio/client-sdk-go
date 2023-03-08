package operations

import (
	"net/http"
)

type RequestRecurringRevenueMetricsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type RequestRecurringRevenueMetricsRequest struct {
	PathParams RequestRecurringRevenueMetricsPathParams
}

type RequestRecurringRevenueMetrics200ApplicationJSONDimensionsItems struct {
	Index *int64 `json:"index,omitempty"`
}

type RequestRecurringRevenueMetrics200ApplicationJSONDimensions struct {
	DisplayName *string                                                           `json:"displayName,omitempty"`
	Index       *int64                                                            `json:"index,omitempty"`
	Items       []RequestRecurringRevenueMetrics200ApplicationJSONDimensionsItems `json:"items,omitempty"`
	Type        *string                                                           `json:"type,omitempty"`
}

type RequestRecurringRevenueMetrics200ApplicationJSONErrorsDetails struct {
	AdditionalProp1 []string `json:"additionalProp1,omitempty"`
	AdditionalProp2 []string `json:"additionalProp2,omitempty"`
	AdditionalProp3 []string `json:"additionalProp3,omitempty"`
}

type RequestRecurringRevenueMetrics200ApplicationJSONErrors struct {
	Details *RequestRecurringRevenueMetrics200ApplicationJSONErrorsDetails `json:"details,omitempty"`
	Message *string                                                        `json:"message,omitempty"`
	Type    *string                                                        `json:"type,omitempty"`
}

type RequestRecurringRevenueMetrics200ApplicationJSONMeasures struct {
	DisplayName *string `json:"displayName,omitempty"`
	Index       *int64  `json:"index,omitempty"`
	Type        *string `json:"type,omitempty"`
	Units       *string `json:"units,omitempty"`
}

type RequestRecurringRevenueMetrics200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresMeasures struct {
	Index              *int64   `json:"index,omitempty"`
	MeasureDisplayName *string  `json:"measureDisplayName,omitempty"`
	Value              *float64 `json:"value,omitempty"`
}

type RequestRecurringRevenueMetrics200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresReportComponentMeasures struct {
	Dimension            *int64                                                                                                                                                                      `json:"dimension,omitempty"`
	DimensionDisplayName *string                                                                                                                                                                     `json:"dimensionDisplayName,omitempty"`
	Item                 *int64                                                                                                                                                                      `json:"item,omitempty"`
	ItemDisplayName      *string                                                                                                                                                                     `json:"itemDisplayName,omitempty"`
	Measures             []RequestRecurringRevenueMetrics200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresMeasures `json:"measures,omitempty"`
}

type RequestRecurringRevenueMetrics200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresMeasures struct {
	Index              *int64   `json:"index,omitempty"`
	MeasureDisplayName *string  `json:"measureDisplayName,omitempty"`
	Value              *float64 `json:"value,omitempty"`
}

type RequestRecurringRevenueMetrics200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasures struct {
	Components           []RequestRecurringRevenueMetrics200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresReportComponentMeasures `json:"components,omitempty"`
	Dimension            *int64                                                                                                                                                              `json:"dimension,omitempty"`
	DimensionDisplayName *string                                                                                                                                                             `json:"dimensionDisplayName,omitempty"`
	Item                 *int64                                                                                                                                                              `json:"item,omitempty"`
	ItemDisplayName      *string                                                                                                                                                             `json:"itemDisplayName,omitempty"`
	Measures             []RequestRecurringRevenueMetrics200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresMeasures                `json:"measures,omitempty"`
}

type RequestRecurringRevenueMetrics200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresMeasures struct {
	Index              *int64   `json:"index,omitempty"`
	MeasureDisplayName *string  `json:"measureDisplayName,omitempty"`
	Value              *float64 `json:"value,omitempty"`
}

type RequestRecurringRevenueMetrics200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasures struct {
	Components           []RequestRecurringRevenueMetrics200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasures `json:"components,omitempty"`
	Dimension            *int64                                                                                                                                       `json:"dimension,omitempty"`
	DimensionDisplayName *string                                                                                                                                      `json:"dimensionDisplayName,omitempty"`
	Item                 *int64                                                                                                                                       `json:"item,omitempty"`
	ItemDisplayName      *string                                                                                                                                      `json:"itemDisplayName,omitempty"`
	Measures             []RequestRecurringRevenueMetrics200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresMeasures                `json:"measures,omitempty"`
}

type RequestRecurringRevenueMetrics200ApplicationJSONReportDimensionsItemsReportComponentMeasuresMeasures struct {
	Index              *int64   `json:"index,omitempty"`
	MeasureDisplayName *string  `json:"measureDisplayName,omitempty"`
	Value              *float64 `json:"value,omitempty"`
}

type RequestRecurringRevenueMetrics200ApplicationJSONReportDimensionsItemsReportComponentMeasures struct {
	Components           []RequestRecurringRevenueMetrics200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasures `json:"components,omitempty"`
	Dimension            *int64                                                                                                                `json:"dimension,omitempty"`
	DimensionDisplayName *string                                                                                                               `json:"dimensionDisplayName,omitempty"`
	Item                 *int64                                                                                                                `json:"item,omitempty"`
	ItemDisplayName      *string                                                                                                               `json:"itemDisplayName,omitempty"`
	Measures             []RequestRecurringRevenueMetrics200ApplicationJSONReportDimensionsItemsReportComponentMeasuresMeasures                `json:"measures,omitempty"`
}

type RequestRecurringRevenueMetrics200ApplicationJSONReportDimensionsItems struct {
	Components           []RequestRecurringRevenueMetrics200ApplicationJSONReportDimensionsItemsReportComponentMeasures `json:"components,omitempty"`
	Dimension            *int64                                                                                         `json:"dimension,omitempty"`
	DimensionDisplayName *string                                                                                        `json:"dimensionDisplayName,omitempty"`
	Item                 *int64                                                                                         `json:"item,omitempty"`
	ItemDisplayName      *string                                                                                        `json:"itemDisplayName,omitempty"`
}

type RequestRecurringRevenueMetrics200ApplicationJSONReportInfo struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty"`
}

// RequestRecurringRevenueMetrics200ApplicationJSON
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
type RequestRecurringRevenueMetrics200ApplicationJSON struct {
	Dimensions []RequestRecurringRevenueMetrics200ApplicationJSONDimensions            `json:"dimensions,omitempty"`
	Errors     []RequestRecurringRevenueMetrics200ApplicationJSONErrors                `json:"errors,omitempty"`
	Measures   []RequestRecurringRevenueMetrics200ApplicationJSONMeasures              `json:"measures,omitempty"`
	ReportData []RequestRecurringRevenueMetrics200ApplicationJSONReportDimensionsItems `json:"reportData,omitempty"`
	ReportInfo *RequestRecurringRevenueMetrics200ApplicationJSONReportInfo             `json:"reportInfo,omitempty"`
}

type RequestRecurringRevenueMetricsResponse struct {
	ContentType                                            string
	StatusCode                                             int
	RawResponse                                            *http.Response
	RequestRecurringRevenueMetrics200ApplicationJSONObject *RequestRecurringRevenueMetrics200ApplicationJSON
}
