/* A Keystore File (sometimes called a UTC file) in Ethereum 156 is an encrypted version of your private key */

package keystores

import (
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

//It's create a new Keystore and account
func CreateKeystore(path, passphrase string) (string, error) {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	// Create a new account with the specified encryption passphrase.
	account, err := ks.NewAccount(passphrase)
	if err != nil {
		return "", err
	}

	return account.Address.Hex(), nil
}

//ImportKeystore get the account updated above from the local keystore
func ImportKeystore(path, file, passphrase, newPassphrase string) (string, error) {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	account, err := ks.Import(jsonBytes, passphrase, newPassphrase)
	if err != nil {
		return "", err
	}

	if err := os.Remove(file); err != nil {
		return "", err
	}

	return account.Address.Hex(), nil
}

//DeleteAccountFromKeystore delete the account updated above from the local keystore
func DeleteAccountFromKeystore(path, file, passphrase, newPassphrase string) error {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	account, err := ks.Import(jsonBytes, passphrase, "temp passphrase")
	if err != nil {
		return err
	}
	err = ks.Delete(account, "temp passphrase")
	if err != nil {
		return err
	}

	return nil
}
