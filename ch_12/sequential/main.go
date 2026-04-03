package sequential

import (
	"fmt"
	"time"
)

func main() {

	orders := []Order{
		{
			TableNumber: 1,
			PrepTime:    2 * time.Second,
		},
		{
			TableNumber: 2,
			PrepTime:    3 * time.Second,
		},
	}
	// sequential management
	//for id, order := range orders {
	//	ProcessOrder(id,order)
	//}

	// go routine
	for id, order := range orders { // reason why despite running you do not get any output is because
		go ProcessOrder(id, order) // main thread spawns go routines , but does not wait for them and closes
		// before completion of children threads spawned , and since main thread (main func)
	} // is entry point , its end means program ends .

}

type Order struct {
	TableNumber int
	PrepTime    time.Duration
}

func ProcessOrder(waiterId int, order Order) {
	// simulate cooking time
	fmt.Printf("Preparing order for table %d...\n", order.TableNumber)

	time.Sleep(order.PrepTime)

	fmt.Printf("Order ready for table %d! \n\n", order.TableNumber)
}
