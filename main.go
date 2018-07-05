package main

import (
	"github.com/YAWAL/CryptoTrader/gdaxClient"
	"os"
	"fmt"
	"encoding/json"
)

type Product struct {
	Id             int `json:"id"`
	BaseCurrency   string `json:"base_currency"`
	QuoteCurrency  string `json:"quote_currency"`
	BaseMinSize    int `json:"base_min_size"`
	BaseMaxSize    string `json:"base_max_size"`
	QuoteIncrement string `json:"quote_increment"`

	CancelOnly   bool   `json:"cancel_only"`
	DisplayName  string `json:"display_name"`
	LimitOnly    bool   `json:"limit_only"`
	MarginEnable bool   `json:"margin_enable"`
	//base_currency:	"BTC"
	//base_max_size:	"50"
	//base_min_size:	"0.001"
	//cancel_only:	false
	//display_name:	"BTC/EUR"
	//id:	"BTC-EUR"
	//limit_only:	false
	//margin_enabled:	false
	MaxMarketFunds string `json:"max_market_funds"`
	//max_market_funds:	"600000"
	MinMarketFunds string `json:"min_market_funds"`
	//min_market_funds:	"10"
	PostOnly bool `json:"post_only"`
	//post_only:	false
	//quote_currency:	"EUR"
	//quote_increment:	"0.01"
	Status string `json:"status"`
	//status:	"online"
	StatusMessage string `json:"status_message"`
	//status_message:	null
}

func main() {

	secret := os.Getenv("COINBASE_SECRET")
	key := os.Getenv("COINBASE_KEY")
	passphrase := os.Getenv("COINBASE_PASSPHRASE")

	client := gdaxClient.NewClient(secret, key, passphrase)
	//prod := Product{}

	products := Product{}
	res, err := client.Request("GET", "/products", nil, nil)
	if err != nil {
		fmt.Println("1 : ", err)
		return
	}
	res.Header.Set("Content-Type", "application/json")
	fmt.Println("res: ", res)
	fmt.Println("resBody : ", res.Body)
	defer res.Body.Close()

	data, err := json.Marshal(res.Body)
	if err != nil {
		fmt.Println("1,5: ", err)
		return
	}
	fmt.Println("data: ", data)

	//data, err := httputil.DumpResponse(res,false)
	//if err!= nil{
	//	fmt.Println("2: ", err)
	//	return
	//}

	err = json.Unmarshal(data, &products)
	if err != nil {
		fmt.Println("3: ", err)
		return
	}
	fmt.Println("products: ", products)

}
