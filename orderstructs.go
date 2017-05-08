package main

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

type JSONInput struct {
	Available_cookies int
	Orders            []Order
	Pagination        PageInfo
}

type JSONOutput struct {
	Remaining_cookies  int
	Unfulfilled_orders []int
}

func (order *Order) Contains(food string) bool {
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
