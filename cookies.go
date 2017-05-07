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

func (order Order) NumCookies() int {
	numCookies := 0

	for _, item := range order.Products {
		if item.Title == "Cookie" {
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

	// TODO implement Priority Queue
	pq := make(PriorityQueue, len(orders))

	for i := 0; i < len(orders); i++ {

		pq[i] = &orders[i]
	}

	fmt.Println("BEFORE HEAPINIT")
	for _, v := range pq {

		fmt.Println(v)

	}
	fmt.Println()
	heap.Init(&pq)
	fmt.Println("AFTER HEAPINIT")
	for pq.Len() > 0 {
		fmt.Println(pq.Pop())
	}
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
			// If the order doesn't contain a cookie we fulfill the order
			orders[i].Fulfilled = true
		}
	}
}

// Priority Queue implementation
type PriorityQueue []*Order

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].NumCookies() == pq[j].NumCookies() {
		return pq[i].Id < pq[j].Id
	} else {
		return pq[i].NumCookies() > pq[j].NumCookies()
	}
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(order interface{}) {
	*pq = append(*pq, order.(*Order))
}

func (pq *PriorityQueue) Pop() interface{} {
	element := (*pq)[0]
	*pq = (*pq)[0:len(*pq)]

	return element
}
