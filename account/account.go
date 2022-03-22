package main

import (
	"context"
	"fmt"
	"log"

	"github.com/doitmagic/blockchain/account/balance"
	"github.com/doitmagic/blockchain/block"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, "http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	//get an account address from ganache
	account := common.HexToAddress("0xFDDC6D56AC4c76F877Ea0eCD577Db487562575Aa")
	//get the account balance, if you enter as a third parameter the block number the resul
	// will reflect account balance at the time of that block
	myBalance, err := balance.GetBalanceForAccount(client, account)
	if err != nil {
		log.Fatal(err)
	}
	//print ballance in wei
	fmt.Println(myBalance)

	//obtain last block
	blockHeader, err := block.GetBlockByNumber(client, 0)
	if err == nil {
		//get the account balance at the time of the block using last block number
		balanceAt, err := client.BalanceAt(context.Background(), account, blockHeader.Number)
		if err != nil {
			log.Fatal(err)
		}
		//ballance in wei
		fmt.Printf("balance at the time of that block nr %d is %v ", blockHeader.Number, balanceAt)
	}

	//print balance in ethereum
	fmt.Println(balance.ConvertWeiBalanceToEthValue(myBalance))

	pendingBalance, err := balance.GetPendingBalanceForAccount(client, account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance.ConvertWeiBalanceToEthValue(pendingBalance))
}
