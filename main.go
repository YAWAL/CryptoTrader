package main

import (
	"github.com/YAWAL/CryptoTrader/gdaxClient"
	"os"
	"fmt"
)

type Product struct {
	Id             string `json:"id"`
	BaseCurrency   string `json:"base_currency"`
	QuoteCurrency  string `json:"quote_currency"`
	BaseMinSize    string `json:"base_min_size"`
	BaseMaxSize    string `json:"base_max_size"`
	QuoteIncrement string `json:"quote_increment"`

	CancelOnly     bool   `json:"cancel_only"`
	DisplayName    string `json:"display_name"`
	LimitOnly      bool   `json:"limit_only"`
	MarginEnable   string   `json:"margin_enable"`
	MaxMarketFunds string `json:"max_market_funds"`
	MinMarketFunds string `json:"min_market_funds"`
	PostOnly       bool   `json:"post_only"`
	Status         string `json:"status"`
	StatusMessage  string `json:"status_message"`
}

//type prod []Product

func main() {

	secret := os.Getenv("COINBASE_SECRET")
	key := os.Getenv("COINBASE_KEY")
	passphrase := os.Getenv("COINBASE_PASSPHRASE")

	client := gdaxClient.NewClient(secret, key, passphrase)

	products := []Product{}

	res, err := client.Request("GET", "/products", nil, &products)
	if err != nil {
		fmt.Println("1 : ", err)
		return
	}
	res.Header.Set("Content-Type", "application/json")

	defer res.Body.Close()
	for _, prod := range products {
		fmt.Println(prod.Id)

	}

}
