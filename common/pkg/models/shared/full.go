package shared

// FullImage
// Image reference.
type FullImage struct {
	Alt *string `json:"alt,omitempty"`
	Src *string `json:"src,omitempty"`
}

type Full struct {
	Image *FullImage `json:"image,omitempty"`
}
