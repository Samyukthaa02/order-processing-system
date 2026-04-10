
package mapper

import (
    "com/ecommerce/order/peerislands_assignment/dto"
    "com/ecommerce/order/peerislands_assignment/model"
)

func ToOrder(dto dto.OrderRequest) model.Order {
    return model.Order{
        CustomerName:    dto.CustomerName,
        OrderDate:       dto.OrderDate,
        ShippingAddress: dto.ShippingAddress,
        OrderItems:      ToOrderItems(dto.OrderItems),
    }
}

func ToOrderResponse(order model.Order) dto.OrderResponse {
    return dto.OrderResponse{
        ID:              order.ID,
        CustomerName:    order.CustomerName,
        OrderDate:       order.OrderDate,
        Status:          order.Status,
        ShippingAddress: order.ShippingAddress,
        TotalAmount:     order.TotalAmount,
        OrderItems:      ToOrderItemResponses(order.OrderItems),
    }
}

func ToOrderItems(dtos []dto.OrderItemRequest) []model.OrderItem {
    items := make([]model.OrderItem, len(dtos))
    for i, dto := range dtos {
        items[i] = ToOrderItem(dto)
    }
    return items
}

func ToOrderItem(dto dto.OrderItemRequest) model.OrderItem {
    return model.OrderItem{
        ProductID: dto.ProductID,
        Quantity:  dto.Quantity,
        Price:     dto.Price,
    }
}

func ToOrderItemResponses(items []model.OrderItem) []dto.OrderItemResponse {
    dtos := make([]dto.OrderItemResponse, len(items))
    for i, item := range items {
        dtos[i] = ToOrderItemResponse(item)
    }
    return dtos
}

func ToOrderItemResponse(item model.OrderItem) dto.OrderItemResponse {
    return dto.OrderItemResponse{
        ID:        item.ID,
        ProductID: item.ProductID,
        Quantity:  item.Quantity,
        Price:     item.Price,
    }
}
