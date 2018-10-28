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

	// change the base URL if needed
	// briq.BaseURL = "https://www.givebriq.com/v0"

	// you can override default http client
	// briq.Client = http.Client{...}

	// override per_page value
	// 100 is already the default value
	briq.DefaultLimit = 100

	// three next line is exactly the same:
	// res, err := b.Transactions(briq.Page(1, 100))
	// res, err := b.Transactions(briq.Page(1))
	res := b.Transactions(briq.Page())
	if res.Error != nil {
		panic(res.Error)
	}

	for _, t := range res.Transactions {
		fmt.Printf("[%d] %s -> %s\n", t.Amount, t.From, t.To)
	}

	b.DoTransaction(briq.TransactionInput{
		From:    "kevin",
		To:      "elias",
		Amount:  1,
		Comment: "testing my sdk",
	})
}
