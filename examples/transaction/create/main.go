package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cpurta/bluzelle-db-go/types"
	blz "github.com/cpurta/bluzelle-db-go/v1"
)

func main() {
	blzAddress := "bluzelle1u2g3qtrq67maslw0lnj3gv7xy6sq7cgn0llrwq"
	uuid := "2d280e84-dad1-453c-af56-7c16719e72e7"

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

	_, err = blzClient.Transaction().Create(ctx, &types.MsgCreate{
		Creator: blzAddress,
		Uuid:    uuid,
		Key:     "exampleKey",
		Value:   []byte("exampleValue"),
		Lease: &types.Lease{
			Hours: uint32(1),
		},
	})
	if err != nil {
		fmt.Println("unable to get count for provided uuid:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Successfully created a key value pair!")
}
