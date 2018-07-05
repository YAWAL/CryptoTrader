package main

import (
	"github.com/YAWAL/CryptoTrader/gdaxClient"
	"os"
	"fmt"
	"encoding/json"
	"net/http/httputil"
)
type Product struct {
	Id             string `json:"id"`
	BaseCurrency   string `json:"base_currency"`
	QuoteCurrency  string `json:"quote_currency"`
	BaseMinSize    string `json:"base_min_size"`
	BaseMaxSize    string `json:"base_max_size"`
	QuoteIncrement string `json:"quote_increment"`
}

func main()  {

	secret := os.Getenv("COINBASE_SECRET")
	key := os.Getenv("COINBASE_KEY")
	passphrase := os.Getenv("COINBASE_PASSPHRASE")


	client := gdaxClient.NewClient(secret, key, passphrase)
	prod := Product{}
	res, err:= client.Request("GET","/products",nil,nil)
	if err!= nil{
		fmt.Println("1 : ",err)
		return
	}

	//body := res.Body

	fmt.Println("res : ",res)


	data, err := httputil.DumpResponse(res,false)
	if err!= nil{
		fmt.Println("2: ", err)
		return
	}
	fmt.Println("data: ", string(data))

	err = json.Unmarshal(data, &prod)
	if err!= nil{
		fmt.Println("3: ",err)
		return
	}
	fmt.Println("4: ",prod)

}
