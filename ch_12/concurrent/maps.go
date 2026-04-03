package main

import (
	"fmt"
	"sync"
)

func main() {
	db := database{users: make(map[int]string)}

	wg := sync.WaitGroup{}
	wg.Add(2)

	// two go routines writing to a map concurrently
	// classic race condition problems can arise
	// write -write or read-write is unsafe , only read-read is safe
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			db.add(i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			db.add(i)
		}
	}()

	wg.Wait()
	fmt.Println(db.users)
}

type database struct {
	// go maps are not thread safe
	users map[int]string
	sync.RWMutex
}

func (db *database) add(i int) {
	db.Lock()
	defer db.Unlock()
	db.users[i] = fmt.Sprintf("user-%d", i)
}
