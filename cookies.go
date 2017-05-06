// cookies
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

type JSONOutput struct {
	Remaining_cookies  int64
	Unfulfilled_orders []int64
}

func main() {

	// order info
	var available_cookies int64
	var orders []Order

Process:
	for page := 1; ; page++ {
		// Get current page
		// TODO functionify to pageData := getPage(page int)
		resp, _ := getPage(page)
		htmlData, _ := ioutil.ReadAll(resp.Body)

		var obj JSONInfo
		error := json.Unmarshal(htmlData, &obj)

		if error != nil {
			panic(error)
		}

		available_cookies = obj.Available_cookies

		for _, v := range obj.Orders {
			orders = append(orders, v)
		}

		fmt.Println(available_cookies)
		fmt.Println(orders)

		if obj.Pagination.Current_page*obj.Pagination.Per_page > obj.Pagination.Total {
			break Process
		}
	} // end Process
}

func getPage(pageNum int) (resp *http.Response, err error) {
	page := strconv.Itoa(pageNum)
	url := "https://backend-challenge-fall-2017.herokuapp.com/orders.json?page=" + page
	fmt.Println(url)
	resp, err = http.Get(url)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	return resp, err
}
