package gogulden

// SignMessage allows you to sign a message using an address. This means you
// can write a message and provide the signing key along with it for individual
// clients to evaluate and confirm that the message came from the owner of this
// address.
func (c *Client) SignMessage(address, message string) (string, error) {
	err := c.runCommand(&message, "signmessage", address, message)
	return message, err
}

// VerifyMessage will verify a message from an address with the provided
// signature. Error may be nil so the bool that is returned should always be
// checked after the error is checked.
func (c *Client) VerifyMessage(address, signature, message string) (bool, error) {
	var valid bool
	err := c.runCommand(&valid, "verifymessage", address, signature, message)
	return valid, err
}
