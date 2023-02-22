package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

type GetDataAssessAccountsCategoriesResponse struct {
	Categories  []shared.Categories
	ContentType string
	StatusCode  int
}
