// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CompanyGroupAssignment struct {
	GroupID *string `json:"groupId,omitempty"`
}

func (o *CompanyGroupAssignment) GetGroupID() *string {
	if o == nil {
		return nil
	}
	return o.GroupID
}
