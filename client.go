package gogulden

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"encoding/json"
)

// Client should be initialised using NewClient.
type Client struct {
	username string
	password string
	host     string
	client   *http.Client
}

type rpcRequest struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
	ID     uint64        `json:"id"`
}

type rpcResponse struct {
	Result *json.RawMessage `json:"result"`
	Error  interface{}      `json:"error"`
	ID     uint64           `json:"id"`
}

// NewClient initialises the gogulden RPC client. The host should be the in the
// following format: http://127.0.0.1:9232.
func NewClient(username, password, host string) (*Client, error) {
	client := &Client{
		username: username,
		password: password,
		host:     host,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	_, err := client.Help("")
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) runCommand(result interface{}, command string, args ...interface{}) error {
	message, err := json.Marshal(&rpcRequest{
		Method: command,
		Params: args,
		ID:     uint64(rand.Int63()),
	})

	req, err := http.NewRequest("POST", c.host, bytes.NewBuffer(message))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.username, c.password)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var rpcResp rpcResponse
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return err
	}
	if rpcResp.Error != nil {
		if m, ok := rpcResp.Error.(map[string]interface{}); ok {
			if message, ok := m["message"]; ok {
				return fmt.Errorf("%v", message)
			}
		}
		return fmt.Errorf("%v", rpcResp.Error)
	}
	if rpcResp.Result == nil {
		return errors.New("result is null")
	}

	return json.Unmarshal(*rpcResp.Result, result)
}
