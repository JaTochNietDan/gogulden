# gogulden ![godoc Reference](https://godoc.org/github.com/JaTochNietDan/gogulden?status.svg)
Go API interface for the Gulden cryptocurrency https://gulden.com/

## Usage
You must be running the Gulden RPC server to use this library. If you are using the downloaded wallet, then you can do this by simply running ./Gulden -server. Gulden will require you to set up authentication information for this RPC server before running it. It explains how to do this in a popup when you try to run the wallet with the -server flag.

```go
package main

import (
	"fmt"
	"log"

	"github.com/JaTochNietDan/gogulden"
)

func main() {
	client, err := gogulden.NewClient("username", "password", "http://127.0.0.1:9232")
	if err != nil {
		log.Fatalln("Couldn't initialize client:", err)
	}

	walletInfo, err := client.WalletInfo()
	if err != nil {
		log.Fatalln("Couldn't get wallet info:", err)
	}

	fmt.Printf("Wallet: %#v\n", walletInfo)
	// Got wallet info: &gogulden.WalletInfo{Version:60000, Balance:0, UnconfirmedBalance:0, ImmatureBalance:0, TransactionCount:0, KeyPoolOldest:1477213584, KeyPoolSize:101}
}

```
