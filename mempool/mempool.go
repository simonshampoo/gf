package mempool

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient/gethclient"
)

func RetrievePendingTx(client *gethclient.Client) {

	ctx := context.Background()
	txHashes := make(chan common.Hash)
	sub, err := client.SubscribePendingTransactions(ctx, txHashes)

	if err != nil {
		panic(err)
	}

	defer sub.Unsubscribe()

	for {
		select {
		case txHash := <-txHashes:
			// check if transaction is profitable
			fmt.Println(txHash)
		case <-ctx.Done():
			return
		}
	}
}
