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

		one of these rpc will be subscribing to new transactions in the mempool DONE

		now, we need to get the calldata for transactions that aren't simple transfers

		simply plug the calldata into my own private environment to see if it increases my balance

		if balance increases, then we send it to the chain with a higher gas fee

		no need to partake in PGAs... idk gas escalation techniques just yet

	*/
}
