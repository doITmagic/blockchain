package wallet

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func GenerateWalletAddress() (string, error) {
	//generating a random private key for signing transactions
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", err
	}

	//convert it to bytes
	privateKeyBytes := crypto.FromECDSA(privateKey)
	//convert it to a hexadecimal string and strip off the 0x after it's hex encoded
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	//derived public key from private key
	publicKey := privateKey.Public()

	//convert it to a hexadecimal string and strip off the 0x and the first 2 characters
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	//generate the public address which is what you're used to seeing
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)

	// public address is Keccak-256 hash of the public key
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	adress := hexutil.Encode(hash.Sum(nil)[12:])
	fmt.Println(adress)

	return address, nil
}
