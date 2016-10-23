package gogulden

import (
	"encoding/json"
	"time"
)

type Transaction struct {
	TransactionData
	Details []Details `json:"details"`
}

type TransactionItem struct {
	Details
	TransactionData
}

type TransactionData struct {
	Confirmations int    `json:"confirmations"`
	TransactionId string `json:"txid"`

	UnixTime unixTime `json:"time"`
}

type Details struct {
	Account  string  `json:"account"`
	Address  string  `json:"address"`
	Category string  `json:"category"`
	Amount   float32 `json:"amount"`
	Vout     int     `json:"vout"`

	SecuredByCheckpoint securedByCheckpoint `json:"secured_by_checkpoint"`
}

type securedByCheckpoint bool
type unixTime time.Time

func (c *Client) GetTransactionsOnAccount(account string, count int, from int, includeWatchOnly bool) ([]*TransactionItem, error) {
	transactions := []*TransactionItem{}
	err := c.runCommand(&transactions, "listtransactions", account, count, from, includeWatchOnly)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (c *Client) GetAllTransactions() ([]*TransactionItem, error) {
	transactions := []*TransactionItem{}
	err := c.runCommand(&transactions, "listtransactions")
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (c *Client) GetTransaction(transactionId string) (*Transaction, error) {
	var transaction Transaction
	err := c.runCommand(&transaction, "gettransaction", transactionId)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (c *Client) SendToAddress(address string, amount float32, comment, commentTo string, subtractFee bool) (string, error) {
	var transactionId string
	err := c.runCommand(&transactionId, "sendtoaddress", address, amount, comment, commentTo, subtractFee)
	if err != nil {
		return "", err
	}

	return transactionId, nil
}

func (t *unixTime) UnmarshalJSON(data []byte) error {
	var timeSeconds int64
	if err := json.Unmarshal(data, &timeSeconds); err != nil {
		return err
	}

	*t = unixTime(time.Unix(timeSeconds, 0))
	return nil
}

func (s *securedByCheckpoint) UnmarshalJSON(data []byte) error {
	var secured string
	if err := json.Unmarshal(data, &secured); err != nil {
		return err
	}

	*s = securedByCheckpoint(secured == "yes")
	return nil
}

func (d *Details) Secured() bool {
	return bool(d.SecuredByCheckpoint)
}

func (td *TransactionData) Time() time.Time {
	return time.Time(td.UnixTime)
}
