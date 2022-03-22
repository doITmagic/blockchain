package block

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

//GetBlockByNumber return header information about a block
func GetBlockByNumber(client *ethclient.Client, nr int64) (header *types.Header, err error) {
	blockNumber := big.NewInt(nr)
	header, err = client.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		return nil, err
	}
	//prin block number
	//fmt.Println(header.Number.String())
	return header, nil
}
