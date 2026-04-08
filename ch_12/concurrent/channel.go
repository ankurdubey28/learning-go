package main

import (
	"fmt"
	"sync"
)

func main() {
	myChan := make(chan int, 1)
	defer close(myChan)
	wg := &sync.WaitGroup{}

	//myChan <- 5
	//fmt.Print(<-myChan) // will deadlock cause no go routine to listen

	// makes this send only
	wg.Add(2)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 5
		wg.Done()
	}(myChan, wg)

	// makes this receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isOpen := <-ch
		if isOpen {
			fmt.Print(val)
		}
		wg.Done()
		//close(ch)   // stops you from closing the channel
	}(myChan, wg)

	wg.Wait()
}
