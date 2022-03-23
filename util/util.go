package util

import hdwallet "github.com/miguelmota/go-ethereum-hdwallet"

//GenerateMnemonic utility function to generate mnemonic words
//bitsNr can be 128 or 256
func GenerateMnemonic(bitsNr int) (string, error) {
	mnemonic, err := hdwallet.NewMnemonic(bitsNr)
	if err != nil {
		return "", err
	}

	return mnemonic, nil
}
