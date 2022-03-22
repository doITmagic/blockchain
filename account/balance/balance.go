package balance

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//GetBalanceForAccount get the balance for an account, optional at the moment of block, in wei
func GetBalanceForAccount(client *ethclient.Client, address common.Address, blockNumber ...int64) (balance *big.Int, err error) {
	balance, err = client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

//GetPendingBalanceForAccount util ti find pending account balance is, after submitting or waiting for a transaction to be confirmed
func GetPendingBalanceForAccount(client *ethclient.Client, address common.Address) (pendingBalance *big.Int, err error) {
	pendingBalance, err = client.PendingBalanceAt(context.Background(), address)
	if err != nil {
		return nil, err
	}
	return pendingBalance, nil
}

//ConvertWeiBalanceToEthValue convert wei to ethereum
// with formula: ETH = wei / 10^18
func ConvertWeiBalanceToEthValue(weiBalance *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(weiBalance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)
	return ethValue
}
