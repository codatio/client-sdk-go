// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// TrackingCategoryRef - References a category against which the item is tracked.
//
// Deprecated type: This will be removed in a future release, please migrate away from it as soon as possible.
type TrackingCategoryRef struct {
	// Unique identifier to the tracking category.
	ID string `json:"id"`
	// Name of tracking category.
	Name *string `json:"name,omitempty"`
}

func (o *TrackingCategoryRef) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TrackingCategoryRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
