package gogulden

import (
	"encoding/json"
	"strconv"
	"time"
)

type Transaction struct {
	Account             string    `json:"account"`
	Address             string    `json:"address"`
	Category            string    `json:"category"`
	Amount              float32   `json:"amount"`
	Vout                int       `json:"vout"`
	SecuredByCheckpoint bool      `json:"-"`
	Confirmations       int       `json:"confirmations"`
	TransactionId       string    `json:"txid"`
	Time                time.Time `json:"-"`
}

func (c *Client) Transactions(account string, count int, from string, includeWatchOnly bool) ([]*Transaction, error) {
	params := []string{}
	if account != "" {
		params = append(params, account)
	}

	if len(params) == 1 && count > 0 {
		params = append(params, strconv.Itoa(count))
	}

	if len(params) == 2 && from != "" {
		params = append(params, from)
	}

	if len(params) == 3 {
		params = append(params, strconv.FormatBool(includeWatchOnly))
	}

	transactions := []*Transaction{}
	err := c.runCommand(&transactions, "listtransactions", params...)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
	type Alias Transaction

	aux := &struct {
		*Alias

		SecuredByCheckpoint string `json:"secured_by_checkpoint"`
		Time                int64  `json:"time"`
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	t.SecuredByCheckpoint = aux.SecuredByCheckpoint == "yes"
	t.Time = time.Unix(aux.Time, 0)

	return nil
}
