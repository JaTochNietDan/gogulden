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
	Confirmations int       `json:"confirmations"`
	TransactionId string    `json:"txid"`
	Time          time.Time `json:"-"`
}

type Details struct {
	Account             string  `json:"account"`
	Address             string  `json:"address"`
	Category            string  `json:"category"`
	Amount              float32 `json:"amount"`
	Vout                int     `json:"vout"`
	SecuredByCheckpoint bool    `json:"-"`

	SecuredByCheckpointStr string `json:"secured_by_checkpoint"`
}

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

func (t *Transaction) UnmarshalJSON(data []byte) error {
	type Alias Transaction

	aux := &struct {
		*Alias

		Time int64 `json:"time"`
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	for idx := range t.Details {
		d := &t.Details[idx]
		d.SecuredByCheckpoint = d.SecuredByCheckpointStr == "yes"
	}

	t.Time = time.Unix(aux.Time, 0)
	return nil
}

func (t *TransactionItem) UnmarshalJSON(data []byte) error {
	type Alias TransactionItem

	aux := &struct {
		*Alias

		Time int64 `json:"time"`
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	t.Details.SecuredByCheckpoint = t.Details.SecuredByCheckpointStr == "yes"
	t.Time = time.Unix(aux.Time, 0)
	return nil
}
