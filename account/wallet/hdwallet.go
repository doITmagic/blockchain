package wallet

import (
	"fmt"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

// CreateHdWallet
//mnemonic ex: "tag volcano eight thank tide danger coast health above argue embrace heavy"
func CreateHdWallet(accountsNr ...int) error {

	mnemonic, err := hdwallet.NewMnemonic(128)
	if err != nil {
		return err
	}

	fmt.Printf("Mnemonic: %s \n", mnemonic)

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
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

	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	account, err = wallet.Derive(path, false)
	if err != nil {
		return err
	}
	fmt.Printf("derivated address %s \n", account.Address.Hex())

	return nil
}
