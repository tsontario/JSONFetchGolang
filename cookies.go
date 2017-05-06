// cookies
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	//"strings"
)

var URL = "https://backend-challenge-fall-2017.herokuapp.com/orders.json?page="

// Structs for JSON unmarshalling
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

type JSONOutput struct {
	Remaining_cookies  int64
	Unfulfilled_orders []int64
}

func (order Order) Contains(food string) bool {
	for _, product := range order.Products {
		if product.Title == food {
			return true
		}
	}
	return false
}

func main() {

	// order info
	var available_cookies int64
	var orders []Order

Process:
	for page := 1; ; page++ {
		resp, _ := getPage(page)
		htmlData, _ := ioutil.ReadAll(resp.Body)

		var jsonData JSONInfo
		error := json.Unmarshal(htmlData, &jsonData)

		if error != nil {
			panic(error)
		}

		for _, v := range jsonData.Orders {
			orders = append(orders, v)
		}

		if jsonData.Pagination.Current_page*jsonData.Pagination.Per_page > jsonData.Pagination.Total {
			available_cookies = jsonData.Available_cookies
			break Process
		}
	} // end Process

	fulfillNoCookieOrders(orders)
	fmt.Println(available_cookies)
	fmt.Println(orders)
}

func getPage(pageNum int) (resp *http.Response, err error) {
	page := strconv.Itoa(pageNum)
	url := URL + page

	resp, err = http.Get(url)

	if err != nil {
		fmt.Printf("%s\n", err)
	}
	return resp, err
}

func fulfillNoCookieOrders(orders []Order) {
	for i := 0; i < len(orders); i++ {
		if !orders[i].Contains("Cookie") {
			// If the order doesn't contain a cookie we set it to fulfilled
			orders[i].Fulfilled = true
			fmt.Printf("Fulfilled order %d.\n", orders[i].Id)
		}
	}
}
