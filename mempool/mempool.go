package mempool

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
)

func RetrievePendingTx(client *gethclient.Client) {

	ch := make(types.Transaction)
	client.SubscribeFullPendingTransactions(ctx, ch)
}
