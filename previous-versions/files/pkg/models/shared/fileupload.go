// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type FileUpload struct {
	// The file to be uploaded as an attachment.
	File CodatFile `multipartForm:"file"`
}

func (o *FileUpload) GetFile() CodatFile {
	if o == nil {
		return CodatFile{}
	}
	return o.File
}
