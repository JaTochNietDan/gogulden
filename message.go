package gogulden

func (c *Client) SignMessage(address string, message string) (string, error) {
	err := c.runCommand(&message, "signmessage", address, message)
	return message, err
}
