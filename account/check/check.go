package check

import (
	"context"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//AddressValid validate address
//address must be string like 0x323b5d4c32345ced77393b3530b1eed0f346429d
func AddressValid(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	return re.MatchString(address)
}

//IsAccount check if Address is an Account or a Smart Contract,
//if return false is a Smart Contract
func IsAccount(address string) (isAccount bool, err error) {
	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, "http://localhost:8545")
	if err != nil {
		return false, err
	}
	addr := common.HexToAddress(address)
	bytecode, err := client.CodeAt(context.Background(), addr, nil) // nil is latest block
	if err != nil {
		return false, err
	}

	isContract := len(bytecode) > 0

	return !isContract, nil
}
