package main

//////////////////////////////
// Priority Queue
//////////////////////////////
type PriorityQueue []*Order

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Establish ordering. Ties for numCookies are ordered ascending by ID
	if pq[i].NumItem(FOOD) == pq[j].NumItem(FOOD) {
		return pq[i].Id < pq[j].Id
	}
	return (pq[i].NumItem(FOOD) > pq[j].NumItem(FOOD))

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
