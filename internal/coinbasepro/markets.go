package main

import (
	"fmt"
	"github.com/preichenberger/go-coinbasepro/v2"
)

func client() {
	client := coinbasepro.NewClient()

	// optional, configuration can be updated with ClientConfig
	client.UpdateConfig(&coinbasepro.ClientConfig{
		BaseURL:    "https://api.pro.coinbase.com",
		Key:        "",
		Passphrase: "",
		Secret:     "",
	})

	stats, err := client.GetStats("BTC-USD")

	if err != nil {
		println(err.Error())
	}

	fmt.Printf("%+v\n", stats)

	rates, err := client.GetHistoricRates("BTC-USD")

	if err != nil {
		println(err.Error())
	}

	fmt.Printf("%+v\n", rates)

}

func main() {
	client()
}
