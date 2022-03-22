package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, "http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we managed to connect")
	networkID, _ := client.NetworkID(ctx)
	fmt.Println(networkID)
}
