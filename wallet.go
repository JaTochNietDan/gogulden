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

func (c *Client) WalletInfo() (*WalletInfo, error) {
	var walletInfo WalletInfo
	err := c.runCommand(&walletInfo, "getwalletinfo")
	if err != nil {
		return nil, err
	}
	return &walletInfo, nil
}
