package gogulden

import "errors"

// ErrInvalidAddress will be returned when a Gulden address is invalid.
var ErrInvalidAddress = errors.New("Invalid Gulden address provided")

// SendToAddress will send a specified amount of Gulden to an address with a
// comment if set to anything but a blank string.
//
// The parameter subtractFee can be used to define whether or not the fee
// should be subtracted from the amount specified. If it is false, it will be
// subtracted from the rest of your wallet balance. If it is true, then it will
// be subtracted from the amount that you specified so the address you sent to
// will receive slightly less than you sent.
func (c *Client) SendToAddress(address string, amount float32, comment, commentTo string, subtractFee bool) (string, error) {
	validation, err := c.ValidateAddress(address)
	if err != nil {
		return "", err
	}

	if !validation.Valid {
		return "", ErrInvalidAddress
	}

	var transactionID string
	err = c.runCommand(&transactionID, "sendtoaddress", address, amount, comment, commentTo, subtractFee)
	return transactionID, nil
}

// SendFrom will send a specified amount to an address from a specific account
// in the wallet.
func (c *Client) SendFrom(account, to string, amount float32, minconf int, comment, commentTo string) (string, error) {
	validation, err := c.ValidateAddress(to)
	if err != nil {
		return "", err
	}

	if !validation.Valid {
		return "", ErrInvalidAddress
	}

	var transactionID string
	err = c.runCommand(&transactionID, "sendfrom", account, to, amount, minconf, comment, commentTo)
	return transactionID, nil
}
