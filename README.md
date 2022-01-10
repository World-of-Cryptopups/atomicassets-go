# atomicassets-go

A simple API Wrapper for the AtomicHub API

### Note

All functions are not tested or guaranteed to work correctly yet.

## Install

```
go get -u github.com/World-of-Cryptopups/atomicassets-go
```

## Usage

```go
package main

import (
    "fmt"
    "log"

	"github.com/World-of-Cryptopups/atomicassets-go"
)

func main () {
	a := atomicassets.New()

	q, err := a.GetAssets(&atomicassets.GetAssetsQuery{
		CollectionName: "cryptopuppie",
		Limit: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q.Data[0].AssetID)
	fmt.Println(q.Data[0].Name)
}
```

- You can also create a new instance with your a custom endpoint and http client.

```go
package main

import (
    "fmt"
    "log"

	"github.com/World-of-Cryptopups/atomicassets-go"
)

func main () {
	a := atomicassets.NewCustom("your-custom-endpoint", &http.Client{})
}
```

## Currently Implemented Endpoints

Only `/atomicassets/` endpoints are implemented since the name implies.

- Assets
- Collections
- Schemas
- Templates
- Offers
- Transfers
- Accounts
- Burns

#### &copy; 2021 | World of Cryptopups
