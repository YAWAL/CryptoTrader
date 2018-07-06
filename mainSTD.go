package main

import (
	"net/http"
	"fmt"
	"io"
	"os"
	"log"
	"encoding/json"
)

type Prod struct {
	Id             string `json:"id"`
	BaseCurrency   string `json:"base_currency"`
	QuoteCurrency  string `json:"quote_currency"`
	BaseMinSize    string `json:"base_min_size"`
	BaseMaxSize    string `json:"base_max_size"`
	QuoteIncrement string `json:"quote_increment"`


}
var prod json.RawMessage

func main()  {
	res, err := http.Get("https://api-public.sandbox.pro.coinbase.com/products")
	if err != nil {
		fmt.Println("1 : ", err)
		return
	}
	defer res.Body.Close()

	n, err := io.Copy(os.Stdout, res.Body)
	if err != nil {
		fmt.Println("2 : ", err)
		return
	}

	log.Println("n: ", n)
	log.Println("es.Body: ", res.Body)


	data, err := json.Marshal(res.Body)

	if err != nil {
		fmt.Println("3: ", err)
		return
	}
	log.Println("data: ", data)

	//
	err = json.Unmarshal(data, prod)
	//
	if err != nil {
		fmt.Println("3: ", err)
		return
	}
	log.Println("products: ", prod)


	}
