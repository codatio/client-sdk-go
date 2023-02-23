package operations

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenuePathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenuePeriodUnitEnum string

const (
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenuePeriodUnitEnumDay   GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenuePeriodUnitEnum = "Day"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenuePeriodUnitEnumWeek  GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenuePeriodUnitEnum = "Week"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenuePeriodUnitEnumMonth GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenuePeriodUnitEnum = "Month"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenuePeriodUnitEnumYear  GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenuePeriodUnitEnum = "Year"
)

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenueQueryParams struct {
	IncludeDisplayNames *bool                                                                                      `queryParam:"style=form,explode=true,name=includeDisplayNames"`
	NumberOfPeriods     int64                                                                                      `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	PeriodLength        int64                                                                                      `queryParam:"style=form,explode=true,name=periodLength"`
	PeriodUnit          GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenuePeriodUnitEnum `queryParam:"style=form,explode=true,name=periodUnit"`
	ReportDate          string                                                                                     `queryParam:"style=form,explode=true,name=reportDate"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenueRequest struct {
	PathParams  GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenuePathParams
	QueryParams GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenueQueryParams
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONDimensionsItems struct {
	Index *int64 `json:"index,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONDimensions struct {
	DisplayName *string                                                                                                         `json:"displayName,omitempty"`
	Index       *int64                                                                                                          `json:"index,omitempty"`
	Items       []GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONDimensionsItems `json:"items,omitempty"`
	Type        *string                                                                                                         `json:"type,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONErrorsDetails struct {
	AdditionalProp1 []string `json:"additionalProp1,omitempty"`
	AdditionalProp2 []string `json:"additionalProp2,omitempty"`
	AdditionalProp3 []string `json:"additionalProp3,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONErrors struct {
	Details *GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONErrorsDetails `json:"details,omitempty"`
	Message *string                                                                                                      `json:"message,omitempty"`
	Type    *string                                                                                                      `json:"type,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONMeasures struct {
	DisplayName *string `json:"displayName,omitempty"`
	Index       *int64  `json:"index,omitempty"`
	Type        *string `json:"type,omitempty"`
	Units       *string `json:"units,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresMeasures struct {
	Index              *int64   `json:"index,omitempty"`
	MeasureDisplayName *string  `json:"measureDisplayName,omitempty"`
	Value              *float64 `json:"value,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresReportComponentMeasures struct {
	Dimension            *int64                                                                                                                                                                                                                    `json:"dimension,omitempty"`
	DimensionDisplayName *string                                                                                                                                                                                                                   `json:"dimensionDisplayName,omitempty"`
	Item                 *int64                                                                                                                                                                                                                    `json:"item,omitempty"`
	ItemDisplayName      *string                                                                                                                                                                                                                   `json:"itemDisplayName,omitempty"`
	Measures             []GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresMeasures `json:"measures,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresMeasures struct {
	Index              *int64   `json:"index,omitempty"`
	MeasureDisplayName *string  `json:"measureDisplayName,omitempty"`
	Value              *float64 `json:"value,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasures struct {
	Components           []GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresReportComponentMeasures `json:"components,omitempty"`
	Dimension            *int64                                                                                                                                                                                                            `json:"dimension,omitempty"`
	DimensionDisplayName *string                                                                                                                                                                                                           `json:"dimensionDisplayName,omitempty"`
	Item                 *int64                                                                                                                                                                                                            `json:"item,omitempty"`
	ItemDisplayName      *string                                                                                                                                                                                                           `json:"itemDisplayName,omitempty"`
	Measures             []GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasuresMeasures                `json:"measures,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresMeasures struct {
	Index              *int64   `json:"index,omitempty"`
	MeasureDisplayName *string  `json:"measureDisplayName,omitempty"`
	Value              *float64 `json:"value,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasures struct {
	Components           []GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresReportComponentMeasures `json:"components,omitempty"`
	Dimension            *int64                                                                                                                                                                                     `json:"dimension,omitempty"`
	DimensionDisplayName *string                                                                                                                                                                                    `json:"dimensionDisplayName,omitempty"`
	Item                 *int64                                                                                                                                                                                     `json:"item,omitempty"`
	ItemDisplayName      *string                                                                                                                                                                                    `json:"itemDisplayName,omitempty"`
	Measures             []GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasuresMeasures                `json:"measures,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONReportDimensionsItemsReportComponentMeasuresMeasures struct {
	Index              *int64   `json:"index,omitempty"`
	MeasureDisplayName *string  `json:"measureDisplayName,omitempty"`
	Value              *float64 `json:"value,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONReportDimensionsItemsReportComponentMeasures struct {
	Components           []GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONReportDimensionsItemsReportComponentMeasuresReportComponentMeasures `json:"components,omitempty"`
	Dimension            *int64                                                                                                                                                              `json:"dimension,omitempty"`
	DimensionDisplayName *string                                                                                                                                                             `json:"dimensionDisplayName,omitempty"`
	Item                 *int64                                                                                                                                                              `json:"item,omitempty"`
	ItemDisplayName      *string                                                                                                                                                             `json:"itemDisplayName,omitempty"`
	Measures             []GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONReportDimensionsItemsReportComponentMeasuresMeasures                `json:"measures,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONReportDimensionsItems struct {
	Components           []GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONReportDimensionsItemsReportComponentMeasures `json:"components,omitempty"`
	Dimension            *int64                                                                                                                                       `json:"dimension,omitempty"`
	DimensionDisplayName *string                                                                                                                                      `json:"dimensionDisplayName,omitempty"`
	Item                 *int64                                                                                                                                       `json:"item,omitempty"`
	ItemDisplayName      *string                                                                                                                                      `json:"itemDisplayName,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONReportInfo struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty"`
}

// GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSON
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
type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSON struct {
	Dimensions []GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONDimensions            `json:"dimensions,omitempty"`
	Errors     []GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONErrors                `json:"errors,omitempty"`
	Measures   []GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONMeasures              `json:"measures,omitempty"`
	ReportData []GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONReportDimensionsItems `json:"reportData,omitempty"`
	ReportInfo *GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONReportInfo             `json:"reportInfo,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenueResponse struct {
	ContentType                                                                                          string
	StatusCode                                                                                           int
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSONObject *GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue200ApplicationJSON
}
