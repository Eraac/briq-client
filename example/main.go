package main

import (
	"fmt"
	"os"

	"github.com/Eraac/briq-client"
)

func main() {
	b := briq.NewBriq(
		os.Getenv("ORGANIZATION_KEY"),
		os.Getenv("API_KEY"),
	)

	// already the default value
	briq.DefaultLimit = 100

	// same as:
	// res, err := b.Transactions(briq.Page(1, 100))
	// res, err := b.Transactions(briq.Page(1))
	res, err := b.Transactions(briq.Page())
	if err != nil {
		panic(err)
	}

	for _, t := range res.Transactions {
		fmt.Printf("app: %s | from: %s | to: %s | amount: %d\n", t.App, t.From, t.To, t.Amount)
	}
}
