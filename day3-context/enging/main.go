package enging

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, World!")

	book := &Book{
		Orders: make([]Order, 0),
		Trades: make([]Trade, 0),
	}

	order := &Order{
		Entity: Entity{
			Id:        1,
			CreatedAt: time.Now(),
		},
		Symbol: "TATA",
		Side:   BUY,
		Price:  100,
		Qty:    100,
		Filled: 0,
	}

	order2 := &Order{
		Entity: Entity{
			Id:        2,
			CreatedAt: time.Now(),
		},
		Symbol: "TATA",
		Side:   SELL,
		Price:  100,
		Qty:    100,
		Filled: 0,
	}

	book.AddOrder(order)
	book.CancelOrder(order.Entity.Id)
	book.AddOrder(order2)
}
