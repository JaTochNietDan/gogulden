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

func (c *Client) Transactions(account string, count int, from string, includeWatchOnly bool) ([]*TransactionItem, error) {
	params := []interface{}{}
	if account != "" {
		params = append(params, account)
	}

	if len(params) == 1 && count > 0 {
		params = append(params, count)
	}

	if len(params) == 2 && from != "" {
		params = append(params, from)
	}

	if len(params) == 3 {
		params = append(params, includeWatchOnly)
	}

	transactions := []*TransactionItem{}
	err := c.runCommand(&transactions, "listtransactions", params...)
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
