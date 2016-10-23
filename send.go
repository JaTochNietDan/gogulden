package gogulden

import "errors"

var ErrInvalidAddress = errors.New("Invalid Gulden address provided")

func (c *Client) SendToAddress(address string, amount float32, comment, commentTo string, subtractFee bool) (string, error) {
	validation, err := c.ValidateAddress(address)
	if err != nil {
		return "", err
	}

	if !validation.Valid {
		return "", ErrInvalidAddress
	}

	var transactionId string
	err = c.runCommand(&transactionId, "sendtoaddress", address, amount, comment, commentTo, subtractFee)
	return transactionId, nil
}
