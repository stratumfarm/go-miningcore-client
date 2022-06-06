# go-miningcore-api

## About
This is a go wrapper around the miningcore api. It implements the miningcore api v2, which also returns meta object, on all available endpoints.

## Install
```
$ go get github.com/stratumfarm/go-miningcore-client
```

## Usage
```go
package main

import (
	"context"
	"fmt"

	"github.com/stratumfarm/go-miningcore-client"
)

func main() {
    c := miningcore.New(
        "https://localhost:8443",
        miningcore.WithTimeout(time.Second * 5),
        miningcore.WithoutTLSVerfiy(),
    )

    pools, statusCode, err := c.GetPools(context.Background())
    if err != nil {
        fmt.Println(statusCode, err)
    }

    for _, v := range pools {
        fmt.Println(v)
    }
}
```