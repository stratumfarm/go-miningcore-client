# go-miningcore-api
[![testing](https://github.com/stratumfarm/go-miningcore-client/actions/workflows/testing.yml/badge.svg)](https://github.com/stratumfarm/go-miningcore-client/actions/workflows/testing.yml)
[![lint](https://github.com/stratumfarm/go-miningcore-client/actions/workflows/lint.yml/badge.svg)](https://github.com/stratumfarm/go-miningcore-client/actions/workflows/lint.yml)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/stratumfarm/go-miningcore-client)
[![Go Report Card](https://goreportcard.com/badge/github.com/stratumfarm/go-miningcore-client)](https://goreportcard.com/report/github.com/stratumfarm/go-miningcore-client)
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