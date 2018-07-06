package handlers

import (
	"github.com/YAWAL/CryptoTrader/gdaxClient"
	"fmt"
	"path"
)

func GetProducts(cl *gdaxClient.Client) ([]gdaxClient.Product, error) {
	products := []gdaxClient.Product{}

	res, err := cl.Request("GET", "/products", nil, &products)
	if err != nil {
		fmt.Println("1 : ", err)
		return nil, err
	}
	res.Header.Set("Content-Type", "application/json")

	defer res.Body.Close()
	return products, nil
}

func GetBooks(cl *gdaxClient.Client, productId string) (gdaxClient.Book, error) {
	book := gdaxClient.Book{}

	url := path.Clean("/products/" + productId + "/book")
	fmt.Println("url : ", url)

	res, err := cl.Request("GET", "/products/BTC-USD/book", nil, &book)

	fmt.Println("book.Bids : ", book.Bids)
	fmt.Println("book.Asks : ", book.Asks)
	//fmt.Println("url : ", url)

	if err != nil {
		return book, err
	}
	res.Header.Set("Content-Type", "application/json")

	defer res.Body.Close()
	return book, nil
}

func GetBook(c *gdaxClient.Client, product string, level int) (gdaxClient.Book, error) {
	var book gdaxClient.Book

	requestURL := fmt.Sprintf("/products/%s/book?level=%d", product, level)
	_, err := c.Request("GET", requestURL, nil, &book)
	return book, err
}