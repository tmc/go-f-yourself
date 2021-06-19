
package main

import (
	alias "github.com/tmc/go-f-yourself"
	"github.com/inlined/go-functions/support/emulator"
)

func main() {
	emulator.Serve(map[string]interface{}{
		"NotACloudFunction": alias.NotACloudFunction,
		"NotAFunction": alias.NotAFunction,
		"PubSubListener": alias.PubSubListener,
		"Webhook": alias.Webhook,
	})
}

