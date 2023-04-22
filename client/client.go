package client

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var GethClient *gethclient.Client

func DialClientWS(apiKey string) {

	endPoint := fmt.Sprintf("wss://mainnet.infura.io/ws/v3/%s", apiKey)

	client, err := rpc.Dial(endPoint)
	if err != nil {
		panic(err)
	}
	newClient := gethclient.New(client)
	fmt.Println("connected")
	GethClient = newClient
}
