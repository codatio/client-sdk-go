package operations

import (
	"net/http"
)

type GetEnhancedBalanceSheetPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetEnhancedBalanceSheetQueryParams struct {
	IncludeDisplayNames *bool  `queryParam:"style=form,explode=true,name=includeDisplayNames"`
	NumberOfPeriods     int64  `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	PeriodLength        int64  `queryParam:"style=form,explode=true,name=periodLength"`
	ReportDate          string `queryParam:"style=form,explode=true,name=reportDate"`
}

type GetEnhancedBalanceSheetRequest struct {
	PathParams  GetEnhancedBalanceSheetPathParams
	QueryParams GetEnhancedBalanceSheetQueryParams
}

type GetEnhancedBalanceSheet200ApplicationJSONDimensionsItems struct {
	Index *int64 `json:"index,omitempty"`
}

type GetEnhancedBalanceSheet200ApplicationJSONDimensions struct {
	DisplayName *string                                                    `json:"displayName,omitempty"`
	Index       *int64                                                     `json:"index,omitempty"`
	Items       []GetEnhancedBalanceSheet200ApplicationJSONDimensionsItems `json:"items,omitempty"`
	Type        *string                                                    `json:"type,omitempty"`
}

type GetEnhancedBalanceSheet200ApplicationJSONErrorsDetails struct {
	AdditionalProp1 []string `json:"additionalProp1,omitempty"`
	AdditionalProp2 []string `json:"additionalProp2,omitempty"`
	AdditionalProp3 []string `json:"additionalProp3,omitempty"`
}

type GetEnhancedBalanceSheet200ApplicationJSONErrors struct {
	Details *GetEnhancedBalanceSheet200ApplicationJSONErrorsDetails `json:"details,omitempty"`
	Message *string                                                 `json:"message,omitempty"`
	Type    *string                                                 `json:"type,omitempty"`
}

type GetEnhancedBalanceSheet200ApplicationJSONMeasures struct {
	DisplayName *string `json:"displayName,omitempty"`
	Index       *int64  `json:"index,omitempty"`
	Type        *string `json:"type,omitempty"`
	Units       *string `json:"units,omitempty"`
}

type GetEnhancedBalanceSheet200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresMeasures struct {
	Index              *int64   `json:"index,omitempty"`
	MeasureDisplayName *string  `json:"measureDisplayName,omitempty"`
	Value              *float64 `json:"value,omitempty"`
}

type GetEnhancedBalanceSheet200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresReportComponentMeasures struct {
	Dimension            *int64                                                                                                                                                               `json:"dimension,omitempty"`
	DimensionDisplayName *string                                                                                                                                                              `json:"dimensionDisplayName,omitempty"`
	Item                 *int64                                                                                                                                                               `json:"item,omitempty"`
	ItemDisplayName      *string                                                                                                                                                              `json:"itemDisplayName,omitempty"`
	Measures             []GetEnhancedBalanceSheet200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresMeasures `json:"measures,omitempty"`
}

type GetEnhancedBalanceSheet200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresMeasures struct {
	Index              *int64   `json:"index,omitempty"`
	MeasureDisplayName *string  `json:"measureDisplayName,omitempty"`
	Value              *float64 `json:"value,omitempty"`
}

type GetEnhancedBalanceSheet200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasures struct {
	Components           []GetEnhancedBalanceSheet200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresReportComponentMeasures `json:"components,omitempty"`
	Dimension            *int64                                                                                                                                                       `json:"dimension,omitempty"`
	DimensionDisplayName *string                                                                                                                                                      `json:"dimensionDisplayName,omitempty"`
	Item                 *int64                                                                                                                                                       `json:"item,omitempty"`
	ItemDisplayName      *string                                                                                                                                                      `json:"itemDisplayName,omitempty"`
	Measures             []GetEnhancedBalanceSheet200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresMeasures                `json:"measures,omitempty"`
}

type GetEnhancedBalanceSheet200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresMeasures struct {
	Index              *int64   `json:"index,omitempty"`
	MeasureDisplayName *string  `json:"measureDisplayName,omitempty"`
	Value              *float64 `json:"value,omitempty"`
}

type GetEnhancedBalanceSheet200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasures struct {
	Components           []GetEnhancedBalanceSheet200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasures `json:"components,omitempty"`
	Dimension            *int64                                                                                                                                `json:"dimension,omitempty"`
	DimensionDisplayName *string                                                                                                                               `json:"dimensionDisplayName,omitempty"`
	Item                 *int64                                                                                                                                `json:"item,omitempty"`
	ItemDisplayName      *string                                                                                                                               `json:"itemDisplayName,omitempty"`
	Measures             []GetEnhancedBalanceSheet200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresMeasures                `json:"measures,omitempty"`
}

type GetEnhancedBalanceSheet200ApplicationJSONReportDimensionsItemsReportComponentMeasuresMeasures struct {
	Index              *int64   `json:"index,omitempty"`
	MeasureDisplayName *string  `json:"measureDisplayName,omitempty"`
	Value              *float64 `json:"value,omitempty"`
}

type GetEnhancedBalanceSheet200ApplicationJSONReportDimensionsItemsReportComponentMeasures struct {
	Components           []GetEnhancedBalanceSheet200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasures `json:"components,omitempty"`
	Dimension            *int64                                                                                                         `json:"dimension,omitempty"`
	DimensionDisplayName *string                                                                                                        `json:"dimensionDisplayName,omitempty"`
	Item                 *int64                                                                                                         `json:"item,omitempty"`
	ItemDisplayName      *string                                                                                                        `json:"itemDisplayName,omitempty"`
	Measures             []GetEnhancedBalanceSheet200ApplicationJSONReportDimensionsItemsReportComponentMeasuresMeasures                `json:"measures,omitempty"`
}

type GetEnhancedBalanceSheet200ApplicationJSONReportDimensionsItems struct {
	Components           []GetEnhancedBalanceSheet200ApplicationJSONReportDimensionsItemsReportComponentMeasures `json:"components,omitempty"`
	Dimension            *int64                                                                                  `json:"dimension,omitempty"`
	DimensionDisplayName *string                                                                                 `json:"dimensionDisplayName,omitempty"`
	Item                 *int64                                                                                  `json:"item,omitempty"`
	ItemDisplayName      *string                                                                                 `json:"itemDisplayName,omitempty"`
}

type GetEnhancedBalanceSheet200ApplicationJSONReportInfo struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty"`
}

// GetEnhancedBalanceSheet200ApplicationJSON
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
type GetEnhancedBalanceSheet200ApplicationJSON struct {
	Dimensions []GetEnhancedBalanceSheet200ApplicationJSONDimensions            `json:"dimensions,omitempty"`
	Errors     []GetEnhancedBalanceSheet200ApplicationJSONErrors                `json:"errors,omitempty"`
	Measures   []GetEnhancedBalanceSheet200ApplicationJSONMeasures              `json:"measures,omitempty"`
	ReportData []GetEnhancedBalanceSheet200ApplicationJSONReportDimensionsItems `json:"reportData,omitempty"`
	ReportInfo *GetEnhancedBalanceSheet200ApplicationJSONReportInfo             `json:"reportInfo,omitempty"`
}

type GetEnhancedBalanceSheetResponse struct {
	ContentType                                     string
	StatusCode                                      int
	RawResponse                                     *http.Response
	GetEnhancedBalanceSheet200ApplicationJSONObject *GetEnhancedBalanceSheet200ApplicationJSON
}
