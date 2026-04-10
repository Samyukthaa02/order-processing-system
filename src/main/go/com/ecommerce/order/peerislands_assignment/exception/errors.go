
package exception

import "errors"

var (
    ErrInvalidOrderRequest = errors.New("Invalid order request")
    ErrInvalidOrderState   = errors.New("Invalid order state")
    ErrOrderNotFound       = errors.New("Order not found")
)
