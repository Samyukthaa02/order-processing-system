
package dto

type OrderItemRequest struct {
    ProductID uint    `json:"productId"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
}
