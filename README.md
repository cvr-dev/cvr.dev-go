# Go API for cvr.dev

![Tests](https://github.com/cvr-dev/cvr.dev-go/actions/workflows/test.yml/badge.svg?branch=main)

The official [cvr.dev](https://cvr.dev/) Go client library.

[cvr.dev](https://cvr.dev/) is a web service that maintains an updated cache of the Danish CVR
registry.

We aim to provide a much simpler and more modern API compared to
CVR's own Elastic Search solution.
Our focus is on high availability and robustness, making it as easy and
reliable as possible to retrieve data about Danish companies from the CVR
database.

## Installation

Make sure that you have go installed, and then run the following in your
project folder:

```bash
go get -u github.com/cvr-dev/cvr.dev-go
```

Also: make sure that your project uses go modules.

## Docs

The Go API is available at [pkg.go.dev](https://pkg.go.dev/github.com/cvr-dev/cvr.dev-go).
The HTTP API is available at [docs.cvr.dev](https://docs.cvr.dev/).

## Usage

In the `cmd/example` dir there's a simple example program that verifies that
your API key works and fetches different data from the server.

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cvr-dev/cvr.dev-go"
)

func main() {
	c := cvr.NewClient(os.Getenv("CVR_DEV_TEST_API_KEY"))

	err := c.TestAPIKey()
	if err != nil {
		log.Printf("Failed to authorize: %s", err)
		return
	}
	log.Printf("Your API key is a-ok!")

	virksomheder, err := c.CVRVirksomhederByCVRNumre(10103940)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range virksomheder {
		fmt.Printf("Virksomhed: %s (%d)\n", v.VirksomhedMetadata.NyesteNavn.Navn, v.CVRNummer)
	}

	produktionsenheder, err := c.CVRProduktionsenhederByPNumre(1003388394)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range produktionsenheder {
		fmt.Printf("Produktionsenhed: %s (%d)\n", p.Metadata.NyesteNavn.Navn, p.PNummer)
	}

}
```

## Test

This project has two types of tests: live tests and local tests. The live tests
are run against our live servers and require that you set a valid API key in
the environment variable `CVR_DEV_TEST_API_KEY`. When at some point we start
charging money for using the service, these tests will count towards your
usage.

If you do not wish to run the live tests, you must add the `-short` flag:

```bash
$ go test ./... -short
```

If you wish to all of the tests, you should use:

```bash
$ go test ./...
```

## Alternatives

We want you to have the best experience possible; if for some reason didn't find
what you were looking for at cvr.dev, reach out to us at kontakt@cvr.dev.

If you just want to check out the market, these are potential alternatives:

- [Virk's official integration](https://datacvr.virk.dk/data/cvr-hj%C3%A6lp/indgange-til-cvr/system-til-system-adgang)
- [cvrapi.dk](https://cvrapi.dk)
- [risika.dk](https://risika.dk)
- [eanapi.dk](https://eanapi.dk)
