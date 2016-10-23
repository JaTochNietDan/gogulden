package gogulden

import (
	"encoding/json"
	"time"
)

// Transaction stores an entire piece of transaction information.
type Transaction struct {
	data
	Details []details `json:"details"`
}

// TransactionItem is a transaction returned when getting a list of
// transactions rather than a singular one.
type TransactionItem struct {
	details
	data
}

type data struct {
	Confirmations int    `json:"confirmations"`
	TransactionID string `json:"txid"`

	UnixTime unixTime `json:"time"`
}

type details struct {
	Account  string  `json:"account"`
	Address  string  `json:"address"`
	Category string  `json:"category"`
	Amount   float32 `json:"amount"`
	Vout     int     `json:"vout"`

	SecuredByCheckpoint securedByCheckpoint `json:"secured_by_checkpoint"`
}

type securedByCheckpoint bool
type unixTime time.Time

// GetTransactionsOnAccount will get all of the transactions on the specified
// account. This may also be limited by the count and from parameter.
func (c *Client) GetTransactionsOnAccount(account string, count int, from int, includeWatchOnly bool) ([]*TransactionItem, error) {
	transactions := []*TransactionItem{}
	err := c.runCommand(&transactions, "listtransactions", account, count, from, includeWatchOnly)
	return transactions, err
}

// GetAllTransactions will return all transactions known to this wallet.
func (c *Client) GetAllTransactions() ([]*TransactionItem, error) {
	transactions := []*TransactionItem{}
	err := c.runCommand(&transactions, "listtransactions")
	return transactions, err
}

// GetTransaction will return all of the information available about a given
// transaction id.
func (c *Client) GetTransaction(transactionID string) (*Transaction, error) {
	var transaction Transaction
	err := c.runCommand(&transaction, "gettransaction", transactionID)
	return &transaction, err
}

// SetTransactionFee will set the fee for all future transactions on this
// wallet to use.
func (c *Client) SetTransactionFee(amount float32) (bool, error) {
	var feeSet bool
	err := c.runCommand(&feeSet, "settxfee", amount)
	return feeSet, err
}

// UnmarshalJSON will turn the returned unix time in seconds from the RPC to a
// Go time object.
func (t *unixTime) UnmarshalJSON(data []byte) error {
	var timeSeconds int64
	if err := json.Unmarshal(data, &timeSeconds); err != nil {
		return err
	}

	*t = unixTime(time.Unix(timeSeconds, 0))
	return nil
}

// UnmarshalJSON will turn the "yes" or "no" return from the RPC into a bool.
func (s *securedByCheckpoint) UnmarshalJSON(data []byte) error {
	var secured string
	if err := json.Unmarshal(data, &secured); err != nil {
		return err
	}

	*s = securedByCheckpoint(secured == "yes")
	return nil
}

// Secured will return whether or not the transaction has been secured by the
// blockchain.
func (d *details) Secured() bool {
	return bool(d.SecuredByCheckpoint)
}

// Time will return the time at which the transaction was executed.
func (td *data) Time() time.Time {
	return time.Time(td.UnixTime)
}
