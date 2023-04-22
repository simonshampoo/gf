package mempool

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
)

//func Subscribe() {
//	payload := map[string]interface{}{
//		"jsonrpc": "2.0",
//		"method":  "eth_blockNumber",
//		"params":  []interface{}{},
//		"id":      1,
//	}
//	requestBody, err := json.Marshal(payload)
//	if err != nil {
//		fmt.Println("Failed to serialize JSON payload:", err)
//		return
//	}
//
//	req, err := http.NewRequest("POST", rpcURL, bytes.NewBuffer(requestBody))
//	if err != nil {
//		fmt.Println("Failed to create HTTP request:", err)
//		return
//	}
//}

func BlockNumber(apiKey string) {
	rpcURL := fmt.Sprintf("wss://mainnet.infura.io/ws/v3/%s", apiKey)

	client, err := rpc.DialWebsocket(context.Background(), rpcURL, "")
	if err != nil {
		fmt.Println("Failed to connect to WebSocket server:", err)
		return
	}

	defer client.Close()

	var blockNumber string
	err = client.Call(&blockNumber, "eth_blockNumber")
	if err != nil {
		fmt.Println("Failed to retrieve block number:", err)
		return
	}

	fmt.Println("Block number:", blockNumber)
}
