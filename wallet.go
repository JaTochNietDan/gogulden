package gogulden

import (
	"fmt"
	"os"
)

// WalletInfo will contain the various information about the wallet.
type WalletInfo struct {
	Version            int     `json:"walletversion"`
	Balance            float32 `json:"balance"`
	UnconfirmedBalance float32 `json:"unconfirmed_balance"`
	ImmatureBalance    float32 `json:"immature_balance"`
	TransactionCount   int     `json:"txcount"`
	KeyPoolOldest      int     `json:"keypoololdest"`
	KeyPoolSize        int     `json:"keypoolsize"`
}

// Account will contain the name of an account and its balance.
type Account struct {
	Name    string
	Balance float32
}

// GetWalletInfo will return the information about the wallet to which this RPC
// is connected to.
func (c *Client) GetWalletInfo() (*WalletInfo, error) {
	var walletInfo WalletInfo
	err := c.runCommand(&walletInfo, "getwalletinfo")
	if err != nil {
		return nil, err
	}
	return &walletInfo, nil
}

// GetAccounts will return all of the accounts on the wallet.
func (c *Client) GetAccounts() ([]*Account, error) {
	data := map[string]float32{}
	err := c.runCommand(&data, "listaccounts")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for name, balance := range data {
		accounts = append(accounts, &Account{
			Name:    name,
			Balance: balance,
		})
	}

	return accounts, nil
}

// GetBalance will return the specific balance of an account specified. Use
// GetWalletInfo to get the overall balance of the wallet rather than this
// function.
func (c *Client) GetBalance(account string, minConf int, includeWatchOnly bool) (float32, error) {
	var balance float32
	err := c.runCommand(&balance, "getbalance", account, minConf, includeWatchOnly)
	return balance, err
}

// GetAccount will return the account associated with a specified address.
func (c *Client) GetAccount(address string) (string, error) {
	var account string
	err := c.runCommand(&account, "getaccount", address)
	return account, err
}

// GetAddress will return the main address associated with the specified
// account.
func (c *Client) GetAddress(account string) (string, error) {
	var address string
	err := c.runCommand(&address, "getaccountaddress", account)
	return address, err
}

// GetAddresses will return all of the addresses associated with the specified
// account.
func (c *Client) GetAddresses(account string) ([]string, error) {
	var addresses []string
	err := c.runCommand(&addresses, "getaddressesbyaccount", account)
	return addresses, err
}

// GetUnconfirmedBalance will get the total amount of Guldens coming into your
// wallet that are not yet confirmed fully.
func (c *Client) GetUnconfirmedBalance() (float32, error) {
	var balance float32
	err := c.runCommand(&balance, "getunconfirmedbalance")
	return balance, err
}

// BackupWallet can be used to write a backup of the wallet to a relative path.
func (c *Client) BackupWallet(path string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	err = c.runCommand(nil, "backupwallet", fmt.Sprintf("%s/%s", wd, path))
	if err != nil && err.Error() != "result is null" {
		return err
	}

	return nil
}

// DumpWallet can be used to write a human-readable dump of the wallet keys to
// a relative path.
func (c *Client) DumpWallet(path string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	err = c.runCommand(nil, "dumpwallet", fmt.Sprintf("%s/%s", wd, path))
	if err != nil && err.Error() != "result is null" {
		return err
	}

	return nil
}

// GetPrivateKey will return the private key associated with the specified
// address.
func (c *Client) GetPrivateKey(address string) (string, error) {
	var key string
	err := c.runCommand(&key, "dumpprivkey", address)
	return key, err
}

// NewAddress will generate a new address under the specified account.
func (c *Client) NewAddress(account string) (string, error) {
	var address string
	err := c.runCommand(&address, "getnewaddress", account)
	return address, err
}

// SetAccount will change the account that the specified address is associated
// with.
func (c *Client) SetAccount(address, account string) error {
	err := c.runCommand(nil, "setaccount", address, account)
	if err.Error() != "result is null" {
		return err
	}

	return nil
}
