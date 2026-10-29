package enging

import (
	"errors"
	"fmt"
	"slices"
	"sync"
	"time"
)

type Side string

const (
	BUY  Side = "buy"
	SELL Side = "sell"
)

type Entity struct {
	Id        int
	CreatedAt time.Time
}

func (e *Entity) Age() time.Duration {
	return time.Since(e.CreatedAt)
}

type Order struct {
	Entity Entity
	Symbol string
	Side   Side
	Price  int64
	Qty    int64
	Filled int64
}

type Trade struct {
	Order       Entity
	BuyOrderId  int64
	SellOrderId int64
	Price       int64
	Qty         int64
}

func (o *Order) Remaining() int64 {
	return o.Qty - o.Filled
}

type Book struct {
	sync.Mutex
	Orders []Order
	Trades []Trade
}

func (b *Book) AddOrder(order *Order) error {
	b.Lock()
	defer b.Unlock()
	if err := b.ValidateError(order); err != nil {
		return err
	}
	b.Orders = append(b.Orders, *order)
	return nil
}

func (b *Book) CancelOrder(id int) error {
	b.Lock()
	defer b.Unlock()
	var exist bool = false
	index := -1

	for i, v := range b.Orders {
		if v.Entity.Id == id {
			exist = true
			index = i
		}
	}

	if exist && index >= 0 {
		b.Orders = slices.Delete(b.Orders, index, index+1)
	} else {
		return fmt.Errorf("order with id %d does not exist %w", id, ErrorOrderNotFound)
	}

	return nil
}

func (b *Book) ValidateError(o *Order) error {

	var err []error

	if o.Price <= 0 {
		err = append(err, fmt.Errorf("Price cannot less than 1 %w", ErrorInvalidOrder))
	}
	if o.Qty <= 0 {
		err = append(err, fmt.Errorf("Price cannot less than 1 %w", ErrorInvalidOrder))
	}
	return errors.Join(err...)

}
