package gogulden

func (c *Client) SendToAddress(address string, amount float32, comment, commentTo string, subtractFee bool) (string, error) {
	var transactionId string
	err := c.runCommand(&transactionId, "sendtoaddress", address, amount, comment, commentTo, subtractFee)
	if err != nil {
		return "", err
	}

	return transactionId, nil
}
