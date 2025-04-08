# AdvancedControls
(*AdvancedControls*)

## Overview

View and manage mapping configured for a company's commerce sync.

### Available Operations

* [CreateCompany](#createcompany) - Create company
* [GetConfiguration](#getconfiguration) - Get company configuration
* [ListCompanies](#listcompanies) - List companies
* [SetConfiguration](#setconfiguration) - Set configuration

## CreateCompany

Creates a Codat company

### Example Usage

```go
package main

import(
	"context"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v2"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"log"
)

func main() {
    ctx := context.Background()

    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.AdvancedControls.CreateCompany(ctx, &shared.CreateCompany{
        Name: "string",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Company != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [shared.CreateCompany](../../pkg/models/shared/createcompany.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |
| `opts`                                                           | [][operations.Option](../../pkg/models/operations/option.md)     | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.CreateCompanyResponse](../../pkg/models/operations/createcompanyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorMessage  | 400, 401, 402, 403, 429 | application/json        |
| sdkerrors.ErrorMessage  | 500, 503                | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## GetConfiguration

Returns a company's commerce sync configuration'.

### Example Usage

```go
package main

import(
	"context"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v2"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.AdvancedControls.GetConfiguration(ctx, operations.GetConfigurationRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Configuration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetConfigurationRequest](../../pkg/models/operations/getconfigurationrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetConfigurationResponse](../../pkg/models/operations/getconfigurationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorMessage  | 401, 402, 403, 404, 429 | application/json        |
| sdkerrors.ErrorMessage  | 500, 503                | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## ListCompanies

Returns a list of companies.

### Example Usage

```go
package main

import(
	"context"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v2"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.AdvancedControls.ListCompanies(ctx, operations.ListCompaniesRequest{
        OrderBy: syncforcommerce.String("-modifiedDate"),
        Query: syncforcommerce.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Companies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListCompaniesRequest](../../pkg/models/operations/listcompaniesrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListCompaniesResponse](../../pkg/models/operations/listcompaniesresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.ErrorMessage       | 400, 401, 402, 403, 404, 429 | application/json             |
| sdkerrors.ErrorMessage       | 500, 503                     | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## SetConfiguration

Sets a company's commerce sync configuration.

### Example Usage

```go
package main

import(
	"context"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v2"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.AdvancedControls.SetConfiguration(ctx, operations.SetConfigurationRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Configuration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.SetConfigurationRequest](../../pkg/models/operations/setconfigurationrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.SetConfigurationResponse](../../pkg/models/operations/setconfigurationresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorMessage            | 400, 401, 402, 403, 404, 409, 429 | application/json                  |
| sdkerrors.ErrorMessage            | 500, 503                          | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |