
package dto

type OrderItemResponse struct {
    ID        uint    `json:"id"`
    ProductID uint    `json:"productId"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
}
