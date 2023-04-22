package main

import (
	"fmt"
	"gf/client"
	"gf/mempool"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	// Access the value of an environment variable
	apiKey := os.Getenv("API_KEY")

	client.DialClientWS(apiKey)

	c := client.GethClient
	if err != nil {
		panic(err)
	}
	fmt.Println(c)

	mempool.RetrievePendingTx(c)
	/*

		so what do i need to do?

		first i gotta establish some ws connection to a node

		then i wanna fire a shitton of rpc at it

		one of these rpc will be subscribing to new transactions in the mempool

		then for every tx to a contract, i'll fork chain, send tx, then see if my balance increases

		if my balance increases, then

	*/
}
