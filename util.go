package gogulden

type Validation struct {
	Valid        bool   `json:"isvalid"`
	Address      string `json:"address"`
	ScriptPubKey string `json:"scriptPubKey"`
	Mine         bool   `json:"ismine"`
	WatchOnly    bool   `json:"iswatchonly"`
	Script       bool   `json:"isscript"`
	PubKey       string `json:"pubkey"`
	Compressed   bool   `json:"iscompressed"`
}

func (c *Client) ValidateAddress(address string) (*Validation, error) {
	var validation Validation
	err := c.runCommand(&validation, "validateaddress", address)
	return &validation, err
}
