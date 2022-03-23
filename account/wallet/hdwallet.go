package wallet

import (
	"fmt"

	"github.com/doitmagic/blockchain/util"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

// CreateHdWallet
func CreateHdWallet(mnemonic ...string) (err error) {

	var gMnemonic string
	if len(mnemonic) > 0 {
		gMnemonic = mnemonic[0]
	} else {
		gMnemonic, err = util.GenerateMnemonic(128)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Mnemonic: %s \n", mnemonic)

	wallet, err := hdwallet.NewFromMnemonic(gMnemonic)
	if err != nil {
		return err
	}

	// The root path for Ethereum is m/44'/60'/0'/0 according to the specification
	// from https://github.com/ethereum/EIPs/issues/84,
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		return err
	}
	fmt.Printf("root address %s \n", account.Address.Hex())

	privateKey, err := wallet.PrivateKeyHex(account)
	if err != nil {
		return err
	}
	fmt.Printf("PrivateKey %v for root \n", privateKey)

	publicKeyRoot, err := wallet.PublicKeyHex(account)
	if err != nil {
		return err
	}
	fmt.Printf("PublicKey %v for root \n", publicKeyRoot)

	//first child account
	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	account, err = wallet.Derive(path, false)
	if err != nil {
		return err
	}
	fmt.Printf("child derivated address %s \n", account.Address.Hex())

	privateKeyChild, err := wallet.PrivateKeyHex(account)
	if err != nil {
		return err
	}
	fmt.Printf("PrivateKey %v for child \n", privateKeyChild)

	publicKeyChild, err := wallet.PublicKeyHex(account)
	if err != nil {
		return err
	}
	fmt.Printf("PublicKey %v for child \n", publicKeyChild)

	return nil
}
