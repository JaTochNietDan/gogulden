package gogulden

// Validation is used to store the result of a validation call ValidateAddress.
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

// ValidateAddress will validate a Gulden address to ensure that it is real. It
// will also return useful information such as whether or not the address is
// belonging to the wallet this is connected to.
//
// Error may still be nil, even when an address is invalid. Therefore you
// should always refer to the valid field in the returned struct for validation.
func (c *Client) ValidateAddress(address string) (*Validation, error) {
	var validation Validation
	err := c.runCommand(&validation, "validateaddress", address)
	return &validation, err
}
