package main

import (
	"os"
	"fmt"
	"github.com/preichenberger/go-gdax"
)

func main() {

	secret := os.Getenv("COINBASE_SECRET")
	key := os.Getenv("COINBASE_KEY")
	passphrase := os.Getenv("COINBASE_PASSPHRASE")

	//client := gdaxClient.NewClient(secret, key, passphrase)

	cl := gdax.NewClient(secret,key,passphrase)

	//products, err := handlers.GetProducts(client)
	products, err := cl.GetProducts()

	if err != nil {
		fmt.Println("1 : ", err)
	}

	for _, prod := range products {
		fmt.Println(prod.Id)

	}

	//book, err := handlers.GetBook(client, "BTC-USD",1)
	book, err := cl.GetBook("BCH-EUR",1)

	if err != nil {
		fmt.Println("2 : ", err)
	}
		fmt.Println(book)

}


