// cookies
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Custom structs for nested info
type Product struct {
	Title  string
	Amount int64
}

type Order struct {
	Id        int64
	Fulfilled bool
	Products  []Product
}

type PageInfo struct {
	Current_page int64
	Per_page     int64
	Total        int64
}

// The entire JSON object
type JSONInfo struct {
	Available_cookies int64
	Orders            []Order
	Pagination        PageInfo
}

func main() {

	// order info
	var available_cookies int64
	var orders []Order

	resp, err := http.Get("https://backend-challenge-fall-2017.herokuapp.com/orders.json")
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	htmlData, err := ioutil.ReadAll(resp.Body)

	var obj JSONInfo
	error := json.Unmarshal(htmlData, &obj)

	if error != nil {
		panic(error)
	}

	available_cookies = obj.Available_cookies

	for pagesToProcess {

		for _, v := range obj.Orders {
			orders = append(orders, v)
		}
		fmt.Println(available_cookies)
		fmt.Println(orders)
	}
}
