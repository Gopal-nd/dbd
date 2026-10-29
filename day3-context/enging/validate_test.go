package enging

import (
	"errors"
	"fmt"
	"testing"
)

func ValidateError(o *Order) error {

	var err []error

	if o.Price <= 0 {
		err = append(err, fmt.Errorf("Price cannot less than 1 %w", ErrorInvalidOrder))
	}
	if o.Qty <= 0 {
		err = append(err, fmt.Errorf("Price cannot less than 1 %w", ErrorInvalidOrder))
	}
	return errors.Join(err...)

}

func TestValidate(t *testing.T) {

	tests := []struct {
		name   string
		order  Order
		errWar bool
	}{{"valid buy", Order{Side: "BUY", Price: 100, Qty: 1}, false},
		{"valid sell", Order{Side: "SELL", Price: 100, Qty: 1}, false},
		{"zero price", Order{Side: "BUY", Price: 0, Qty: 1}, true},
		{"negative qty", Order{Side: "SELL", Price: 100, Qty: -5}, true},
		{"bad side", Order{Side: "HOLD", Price: 100, Qty: 1}, true}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateError(&tt.order)

			if tt.errWar && !errors.Is(err, ErrorInvalidOrder) {
				t.Errorf("want ErrInvalidOrder, got %v", err)
			}

			if !tt.errWar && err != nil {
				t.Errorf("want ErrInvalidOrder, got %v", err)
			}

		})
	}

}
