# B2-GO-AA-UTILS

## Get Started

- Install the module:

```bash
go get github.com/b2network/b2-go-aa-utils
```

- Refer to [b2-account-infra](https://github.com/b2network/b2-account-infra) repo to collect deployed contract addresses

## Example

```go
package main

import (
	"log"

	"github.com/b2network/b2-go-aa-utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	B2DevRpcUrl   = "http://43.135.203.73:8123"
	SCARegistry   = "0x231aec684Ad0e63c2F4d176EddCE97A1B666247c"
	KernelFactory = "0x7516283Ff7090B8286E23a16f8b5b35B3ba541A2"
)

func main() {
	client, err := ethclient.Dial(B2DevRpcUrl)
	if err != nil {
		panic(err)
	}
	owner := "bc1qe40k9zyjyndl2t7f0fxws7h94pjrfz0zan5yak"
	target, err := b2aa.GetSCAAddress(client, common.HexToAddress(SCARegistry), common.HexToAddress(KernelFactory), owner)
	if err != nil {
		panic(err)
	}
	log.Println((target))
}
```