
package service

import (
    "com/ecommerce/order/peerislands_assignment/dto"
)

type OrderService interface {
    CreateOrder(order dto.OrderRequest) (dto.OrderResponse, error)
    GetOrder(id uint) (dto.OrderResponse, error)
    GetAllOrders() ([]dto.OrderResponse, error)
    UpdateOrderStatus(id uint, status string) (dto.OrderResponse, error)
    DeleteOrder(id uint) error
}
