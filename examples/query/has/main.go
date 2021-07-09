package main

import (
	"context"
	"fmt"
	"os"
	"time"

	blz "github.com/cpurta/bluzelle-db-go/v1"
)

func main() {
	uuid := "2d280e84-dad1-453c-af56-7c16719e72e7"
	key := "exampleKey"

	config := &blz.Config{
		Mnemonic:    "your mnemonic goes here",
		APIEndpoint: "https://client.sentry.testnet.private.bluzelle.com:26657",
		MaxGas:      10000,
		GasPrice:    0.02,
		UUID:        uuid,
	}

	blzClient, err := blz.NewBluzelleClient(config)
	if err != nil {
		fmt.Println("unable to set up new bluzelle client:", err.Error())
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	defer cancel()

	hasResponse, err := blzClient.Query().Has(ctx, uuid, key)
	if err != nil {
		fmt.Println("unable to determine if key exists for provided uuid:", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Has: %t\n", hasResponse.Has)
}
