package gogulden

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
