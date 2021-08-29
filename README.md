# atomicassets-go

A simple API Wrapper for the AtomicHub API

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

## Currently Implemented Endpoints

- Assets
- Collections
- Schemas
- Templates

#### &copy; 2021 | World of Cryptopups
