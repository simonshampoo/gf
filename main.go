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
	c, err := client.DialClientWS(apiKey)
	defer c.Close()
	if err != nil {
		panic(err)
	}
	mempool.BlockNumber(apiKey)
	fmt.Println(c)
}
