package mempool

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
)

func RetrievePendingTx(client *gethclient.Client) {

	ctx := context.Background()
	txs := make(chan *types.Transaction)
	sub, err := client.SubscribeFullPendingTransactions(ctx, txs)

	if err != nil {
		panic(err)
	}
	print(sub)
	defer sub.Unsubscribe()

	for {
		select {
		case tx := <-txs:
			// check if transaction is profitable
			if tx.Data() != nil {
				fmt.Println("==========================================================\ncalldata: ", common.Bytes2Hex(tx.Data()), "\ngas: ", tx.Gas(), "\ngas price: ", tx.GasPrice())
			}
		case <-ctx.Done():
			return
		}
	}
}
