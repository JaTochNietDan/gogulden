package gogulden

import (
	"fmt"
	"os"
)

type WalletInfo struct {
	Version            int     `json:"walletversion"`
	Balance            float32 `json:"balance"`
	UnconfirmedBalance float32 `json:"unconfirmed_balance"`
	ImmatureBalance    float32 `json:"immature_balance"`
	TransactionCount   int     `json:"txcount"`
	KeyPoolOldest      int     `json:"keypoololdest"`
	KeyPoolSize        int     `json:"keypoolsize"`
}

type Account struct {
	Name    string
	Balance float32
}

func (c *Client) WalletInfo() (*WalletInfo, error) {
	var walletInfo WalletInfo
	err := c.runCommand(&walletInfo, "getwalletinfo")
	if err != nil {
		return nil, err
	}
	return &walletInfo, nil
}

func (c *Client) Accounts() ([]*Account, error) {
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

func (c *Client) GetBalance(account string, minConf int, includeWatchOnly bool) (float32, error) {
	var balance float32
	err := c.runCommand(&balance, "getbalance", account, minConf, includeWatchOnly)
	return balance, err
}

func (c *Client) GetAccount(address string) (string, error) {
	var account string
	err := c.runCommand(&account, "getaccount", address)
	return account, err
}

func (c *Client) GetAddresses(account string) ([]string, error) {
	var addresses []string
	err := c.runCommand(&addresses, "getaddressesbyaccount", account)
	return addresses, err
}

func (c *Client) GetUnconfirmedBalance() (float32, error) {
	var balance float32
	err := c.runCommand(&balance, "getunconfirmedbalance")
	return balance, err
}

func (c *Client) BackupWallet(path string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	err = c.runCommand(nil, "backupwallet", fmt.Sprintf("%s/%s", wd, path))
	if err.Error() != "result is null" {
		return err
	}

	return nil
}

func (c *Client) DumpWallet(path string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	err = c.runCommand(nil, "dumpwallet", fmt.Sprintf("%s/%s", wd, path))
	if err.Error() != "result is null" {
		return err
	}

	return nil
}

func (c *Client) GetPrivateKey(address string) (string, error) {
	var key string
	err := c.runCommand(&key, "dumpprivkey", address)
	return key, err
}
