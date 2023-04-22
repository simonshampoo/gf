package main

import (
	"fmt"
	"gf/client"
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
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
}
