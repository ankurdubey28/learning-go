package main

import (
	"fmt"
	s "learning-go/ch_12/sequential"
	"sync"
	"time"
)

func main() {
	orders := []s.Order{
		{
			TableNumber: 1,
			PrepTime:    2 * time.Second,
		},
		{
			TableNumber: 2,
			PrepTime:    3 * time.Second,
		},
		{TableNumber: 3,
			PrepTime: 3 * time.Second},
	}

	wg := sync.WaitGroup{}

	for waiterId, order := range orders {
		wg.Add(1)
		go func(waiterId int, order s.Order) {
			defer wg.Done()
			s.ProcessOrder(waiterId, order)
		}(waiterId, order)
	}
	wg.Wait()
	fmt.Println("all orders done")
}
