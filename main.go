package main

import (
	"github.com/YAWAL/CryptoTrader/gdaxClient"
	"os"
	"fmt"
	"encoding/json"
)

type Product struct {
	Id             string `json:"id"`
	BaseCurrency   string `json:"base_currency"`
	QuoteCurrency  string `json:"quote_currency"`
	BaseMinSize    string `json:"base_min_size"`
	BaseMaxSize    string `json:"base_max_size"`
	QuoteIncrement string `json:"quote_increment"`
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

	data, err := json.Marshal(res.Body)
	if err != nil {
		fmt.Println("1,5: ", err)
		return
	}
	fmt.Println("data: ", data)

	res.Header.Set("Content-Type", "application/json")
	fmt.Println("resBody : ", res.Body)
	defer res.Body.Close()

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
	fmt.Println("4: ", products)

}
