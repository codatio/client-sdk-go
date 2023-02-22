package codatio

import (
	"context"
	"fmt"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/utils"
	"net/http"
)

type accountTransactions struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newAccountTransactions(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *accountTransactions {
	return &accountTransactions{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// GetAccountTransaction - Get account transaction
// Gets the account transactions for a given company.Gets the specified account transaction for a given company and connection.
func (s *accountTransactions) GetAccountTransaction(ctx context.Context, request operations.GetAccountTransactionRequest) (*operations.GetAccountTransactionResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/companies/{companyId}/connections/{connectionId}/data/accountTransactions/{accountTransactionId}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := utils.ConfigureSecurityClient(s.defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAccountTransactionResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.AccountTransaction
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.AccountTransaction = out
		}
	}

	return res, nil
}

// ListAccountTransactions - List account transactions
// Gets the account transactions for a given company.
func (s *accountTransactions) ListAccountTransactions(ctx context.Context, request operations.ListAccountTransactionsRequest) (*operations.ListAccountTransactionsResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/companies/{companyId}/connections/{connectionId}/data/accountTransactions", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request.QueryParams); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := utils.ConfigureSecurityClient(s.defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListAccountTransactionsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListAccountTransactionsLinks
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Links = out
		}
	}

	return res, nil
}
