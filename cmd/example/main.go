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
	fmt.Printf("Your API key is a-ok!")

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
