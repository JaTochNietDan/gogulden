package gogulden

import (
	"bytes"
	"math/rand"
	"net/http"
	"time"

	gojson "encoding/json"

	"github.com/gorilla/rpc/json"
)

type Client struct {
	username string
	password string
	host     string
	client   *http.Client
}

type rpcRequest struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
	Id     uint64        `json:"id"`
}

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
	message, err := gojson.Marshal(&rpcRequest{
		Method: command,
		Params: args,
		Id:     uint64(rand.Int63()),
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

	err = json.DecodeClientResponse(resp.Body, result)
	if err != nil {
		return err
	}

	return nil
}
