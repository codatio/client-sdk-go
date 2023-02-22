package shared

import (
	"time"
)

// TrackingCategory
// Details of a category used for tracking transactions.
//
// :::note Language tip
//
// Parameters used to track types of spend in various parts of an organization can be called  **dimensions**, **projects**, **classes**, or **locations** in different accounting platforms. In Codat, we refer to these as tracking categories.
// :::
//
// Explore the <a className="external" href="https://api.codat.io/swagger/index.html#/TrackingCategories" target="_blank">Tracking Categories</a> endpoints in Swagger.
//
// View the coverage for tracking categories in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=trackingCategories" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Tracking categories are used to monitor cost centres and control budgets that sit outside the standard chart of accounts. Customers may use tracking categories to group together and track the income and costs of specific departments, projects, locations or customers.
//
// From their accounting system, customers can:
//
// - Create and maintain tracking categories and tracking category types.
// - View all tracking categories that are available for use.
// - View the relationships between the categories.
// - Assign invoices, bills, credit notes, or bill credit notes to one or more categories.
// - View the categories that a transaction belongs to.
// - View all transactions in a tracking category.
//
// From the [TrackingCategories](https://api.codat.io/swagger/index.html#/TrackingCategories) endpoints, you can retrieve:
//
// - A [list of the latest tracking categories](https://api.codat.io/swagger/index.html#/TrackingCategories/TrackingCategories_ListPaged) for a company.
// - The details of a [specific tracking category](https://api.codat.io/swagger/index.html#/TrackingCategories/TrackingCategories_Single).
type TrackingCategory struct {
	HasChildren        *bool       `json:"hasChildren,omitempty"`
	ID                 *string     `json:"id,omitempty"`
	ModifiedDate       *time.Time  `json:"modifiedDate,omitempty"`
	Name               *string     `json:"name,omitempty"`
	ParentID           *string     `json:"parentId,omitempty"`
	SourceModifiedDate *time.Time  `json:"sourceModifiedDate,omitempty"`
	Status             *StatusEnum `json:"status,omitempty"`
}
