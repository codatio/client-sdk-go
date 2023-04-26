# CompanyManagement

## Overview

Create new and manage existing Sync for Commerce companies.

### Available Operations

* [CreateCompany](#createcompany) - Create Sync for Commerce company
* [CreateConnection](#createconnection) - Create connection
* [ListCompanies](#listcompanies) - List companies
* [ListConnections](#listconnections) - List data connections
* [UpdateConnection](#updateconnection) - Update data connection

## CreateCompany

Creates a Codat company with a commerce partner data connection.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/synccommerce"
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := shared.CreateCompany{
        Name: "Bob's Burgers",
    }

    res, err := s.CompanyManagement.CreateCompany(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Company != nil {
        // handle response
    }
}
```

## CreateConnection

Create a data connection for company.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/synccommerce"
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/operations"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CreateConnectionRequest{
        RequestBody: codatsynccommerce.String("corrupti"),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.CompanyManagement.CreateConnection(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

## ListCompanies

Retrieve a list of all companies the client has created.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/synccommerce"
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/operations"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListCompaniesRequest{
        OrderBy: codatsynccommerce.String("-modifiedDate"),
        Page: 1,
        PageSize: codatsynccommerce.Int(100),
        Query: codatsynccommerce.String("provident"),
    }

    res, err := s.CompanyManagement.ListCompanies(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Companies != nil {
        // handle response
    }
}
```

## ListConnections

Retrieve previously created data connections.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/synccommerce"
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/operations"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListConnectionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codatsynccommerce.String("-modifiedDate"),
        Page: 1,
        PageSize: codatsynccommerce.Int(100),
        Query: codatsynccommerce.String("distinctio"),
    }

    res, err := s.CompanyManagement.ListConnections(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connections != nil {
        // handle response
    }
}
```

## UpdateConnection

Update a data connection

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/synccommerce"
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/operations"
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.UpdateConnectionRequest{
        UpdateConnection: &shared.UpdateConnection{
            Status: codatsynccommerce.String("Linked"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.CompanyManagement.UpdateConnection(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```
