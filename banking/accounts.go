package codatio

import (
	"context"
	"fmt"
	"github.com/codatio/client-sdk-go/banking/pkg/models/operations"
	"github.com/codatio/client-sdk-go/banking/pkg/utils"
	"net/http"
)

type accounts struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newAccounts(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *accounts {
	return &accounts{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// GetBankingAccount - Get account
// Gets a specified bank account for a given company
func (s *accounts) GetBankingAccount(ctx context.Context, request operations.GetBankingAccountRequest) (*operations.GetBankingAccountResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/companies/{companyId}/connections/{connectionId}/data/banking-accounts/{accountId}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetBankingAccountResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetBankingAccountSourceModifiedDate
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.SourceModifiedDate = out
		}
	}

	return res, nil
}

// ListBankingAccounts - List accounts
// Gets a list of all bank accounts of the SMB, with rich data like balances, account numbers and institutions holdingthe accounts.
func (s *accounts) ListBankingAccounts(ctx context.Context, request operations.ListBankingAccountsRequest) (*operations.ListBankingAccountsResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/companies/{companyId}/connections/{connectionId}/data/banking-accounts", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request.QueryParams); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListBankingAccountsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListBankingAccountsLinks
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Links = out
		}
	}

	return res, nil
}
