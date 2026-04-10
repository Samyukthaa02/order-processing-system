
package dto

import "time"

type OrderResponse struct {
    ID              uint      `json:"id"`
    CustomerName    string    `json:"customerName"`
    OrderDate       time.Time `json:"orderDate"`
    Status          string    `json:"status"`
    ShippingAddress string    `json:"shippingAddress"`
    TotalAmount     float64   `json:"totalAmount"`
    OrderItems      []OrderItemResponse `json:"orderItems"`
}
