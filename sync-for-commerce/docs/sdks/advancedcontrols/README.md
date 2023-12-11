# AdvancedControls
(*AdvancedControls*)

## Overview

Advanced company management and sync preferences.

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
	"github.com/codatio/client-sdk-go/sync-for-commerce/v3/pkg/models/shared"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v3"
	"context"
	"log"
)

func main() {
    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AdvancedControls.CreateCompany(ctx, &shared.CreateCompany{
        Description: syncforcommerce.String("Requested early access to the new financing scheme."),
        Name: "Bank of Dave",
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
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 400,401,402,403,429,500,503 | application/json            |
| sdkerrors.SDKError          | 400-600                     | */*                         |

## GetConfiguration

Returns a company's commerce sync configuration'.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/sync-for-commerce/v3/pkg/models/shared"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v3"
	"context"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v3/pkg/models/operations"
	"log"
)

func main() {
    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
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
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 400-600                     | */*                         |

## ListCompanies

Returns a list of companies.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/sync-for-commerce/v3/pkg/models/shared"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v3"
	"context"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v3/pkg/models/operations"
	"log"
)

func main() {
    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AdvancedControls.ListCompanies(ctx, operations.ListCompaniesRequest{
        OrderBy: syncforcommerce.String("-modifiedDate"),
        Page: syncforcommerce.Int(1),
        PageSize: syncforcommerce.Int(100),
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
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |

## SetConfiguration

Sets a company's commerce sync configuration.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/sync-for-commerce/v3/pkg/models/shared"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v3"
	"context"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v3/pkg/models/operations"
	"log"
)

func main() {
    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
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
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.ErrorMessage              | 400,401,402,403,404,409,429,500,503 | application/json                    |
| sdkerrors.SDKError                  | 400-600                             | */*                                 |
