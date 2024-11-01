// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Links struct {
	Current  HalRef  `json:"current"`
	Next     *HalRef `json:"next,omitempty"`
	Previous *HalRef `json:"previous,omitempty"`
	Self     HalRef  `json:"self"`
}

func (o *Links) GetCurrent() HalRef {
	if o == nil {
		return HalRef{}
	}
	return o.Current
}

func (o *Links) GetNext() *HalRef {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *Links) GetPrevious() *HalRef {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *Links) GetSelf() HalRef {
	if o == nil {
		return HalRef{}
	}
	return o.Self
}
