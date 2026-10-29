package enging

import (
	"errors"
)

var (
	ErrorOrderNotFound = errors.New("order not found")
	ErrorInvalidOrder  = errors.New("invalid order")
	ErrorCancelOrder   = errors.New("cancel order")
)

// error.Is
// error.As
