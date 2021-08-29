package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func getCurrencies() {
	resp, err := http.Get("https://api.coinbase.com/v2/currencies")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)
}

func getExchangeRates() {
	resp, err := http.Get("https://api.coinbase.com/v2/exchange-rates")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)
}

func getPrices() {
	resp, err := http.Get("https://api.coinbase.com/v2/prices/BTC-USD/spot")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)
}

func main() {
	getCurrencies()
	getExchangeRates()
	getPrices()
}
