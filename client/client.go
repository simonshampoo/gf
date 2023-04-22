package client

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

func DialClientWS(apiKey string) (*ethclient.Client, error) {

	endPoint := fmt.Sprintf("wss://mainnet.infura.io/ws/v3/%s", apiKey)

	client, err := ethclient.Dial(endPoint)

	if err != nil {
		return nil, err
	}
	fmt.Println("connected")
	return client, nil
	// _ = client
	//
	// address := "0xc56334C9c54D06Dc73aD5576EA8A22c90806af8a"
	// account := common.HexToAddress(address)
	// balance, err := client.BalanceAt(context.Background(), account, nil)
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	// fmt.Println(balance)
	//
	// fbalance := new(big.Float)
	//
	// fbalance.SetString(balance.String())
	//
	// ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	//
	// fmt.Println(ethValue)
}
