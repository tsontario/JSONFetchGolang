// cookies
package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const URL = "https://backend-challenge-fall-2017.herokuapp.com/orders.json?page="
const FOOD = "Cookie"

func main() {

	// order info
	var available_cookies int
	var orders []Order

	// Get the data
Process:
	for page := 1; ; page++ {
		resp := getPage(page)
		htmlData, _ := ioutil.ReadAll(resp.Body)

		var jsonData JSONInput
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

	// Fulfill orders
	fulfillNoCookieOrders(orders)
	unfulfilledOrders := PriorityQueue{}

	for i := 0; i < len(orders); i++ {
		if orders[i].NumCookies() != 0 && !orders[i].Fulfilled {
			heap.Push(&unfulfilledOrders, &orders[i])
		}
	}

	for len(unfulfilledOrders) > 0 && available_cookies > 0 {
		item := heap.Pop(&unfulfilledOrders).(*Order)
		cookies := item.NumCookies()

		if cookies <= available_cookies {
			available_cookies -= cookies
			item.Fulfilled = true
		}
	}

	output JSONOutput
	output.Remaining_cookies = available_cookies
	for _, v := range orders {
		if !v.Fulfilled {
			output.Unfulfilled_orders = append(output.Unfulfilled_orders, v.Id)
		}
	}

	res, _ := json.Marshal(output)
	fmt.Println(string(res))
}

func getPage(pageNum int) *http.Response {
	page := strconv.Itoa(pageNum)
	url := URL + page

	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("%s\n", err)
	} else if resp.StatusCode != 200 {
		fmt.Printf("Response Error Code %s\n", resp.StatusCode)
	}
	return resp
}

func fulfillNoCookieOrders(orders []Order) {
	for i := 0; i < len(orders); i++ {
		if !orders[i].Contains(FOOD) {
			orders[i].Fulfilled = true
		}
	}
}
