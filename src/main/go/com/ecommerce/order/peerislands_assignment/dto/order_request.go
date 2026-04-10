
package dto

import "time"

type OrderRequest struct {
    CustomerName    string    `json:"customerName"`
    OrderDate       time.Time `json:"orderDate"`
    ShippingAddress string    `json:"shippingAddress"`
    OrderItems      []OrderItemRequest `json:"orderItems"`
}
