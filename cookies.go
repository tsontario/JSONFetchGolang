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

var URL = "https://backend-challenge-fall-2017.herokuapp.com/orders.json?page="
var FOOD = "Cookie"

// Structs for JSON unmarshalling
type Product struct {
	Title  string
	Amount int
}

type Order struct {
	Id        int
	Fulfilled bool
	Products  []Product
}

type PageInfo struct {
	Current_page int
	Per_page     int
	Total        int
}

// The entire JSON object
type JSONInfo struct {
	Available_cookies int
	Orders            []Order
	Pagination        PageInfo
}

type JSONOutput struct {
	Remaining_cookies  int
	Unfulfilled_orders []int
}

func (order Order) Contains(food string) bool {
	for _, product := range order.Products {
		if product.Title == food {
			return true
		}
	}
	return false
}

func (order *Order) NumCookies() int {
	numCookies := 0
	for _, item := range order.Products {
		if item.Title == FOOD {
			numCookies = item.Amount
		}
	}
	return numCookies
}

func main() {

	// order info
	var available_cookies int
	var orders []Order

Process:
	for page := 1; ; page++ {
		resp := getPage(page)
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

	output := JSONOutput{}
	output.Remaining_cookies = available_cookies
	for _, v := range orders {
		if !v.Fulfilled {
			output.Unfulfilled_orders = append(output.Unfulfilled_orders, v.Id)
		}
	}

	fmt.Println(output)
}

func getPage(pageNum int) *http.Response {
	page := strconv.Itoa(pageNum)
	url := URL + page

	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("%s\n", err)
	}
	return resp
}

func fulfillNoCookieOrders(orders []Order) {
	for i := 0; i < len(orders); i++ {
		if !orders[i].Contains(FOOD) {
			// If the order doesn't contain a cookie we set it to fulfilled
			orders[i].Fulfilled = true
		}
	}
}

//////////////////////////////
// Priority Queue
//////////////////////////////
type PriorityQueue []*Order

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Establish ordering. Ties for numCookies are ordered ascending by ID
	if pq[i].NumCookies() == pq[j].NumCookies() {
		return pq[i].Id < pq[j].Id
	}
	return (pq[i].NumCookies() > pq[j].NumCookies())

}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	order := x.(*Order)
	*pq = append(*pq, order)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	order := old[n-1]
	*pq = old[0 : n-1]
	return order
}
