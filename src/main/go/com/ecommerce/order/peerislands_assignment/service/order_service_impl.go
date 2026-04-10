
package service

import (
    "com/ecommerce/order/peerislands_assignment/dto"
    "com/ecommerce/order/peerislands_assignment/mapper"
    "com/ecommerce/order/peerislands_assignment/model"
    "com/ecommerce/order/peerislands_assignment/repository"
    "errors"
)

type OrderServiceImpl struct {
    OrderRepository repository.OrderRepository
}

func (s *OrderServiceImpl) CreateOrder(orderRequest dto.OrderRequest) (dto.OrderResponse, error) {
    order := mapper.ToOrder(orderRequest)
    order.Status = model.Pending
    savedOrder, err := s.OrderRepository.Save(order)
    if err != nil {
        return dto.OrderResponse{}, err
    }
    return mapper.ToOrderResponse(savedOrder), nil
}

func (s *OrderServiceImpl) GetOrder(id uint) (dto.OrderResponse, error) {
    order, err := s.OrderRepository.FindById(id)
    if err != nil {
        return dto.OrderResponse{}, errors.New("Order not found")
    }
    return mapper.ToOrderResponse(order), nil
}

func (s *OrderServiceImpl) GetAllOrders() ([]dto.OrderResponse, error) {
    orders, err := s.OrderRepository.FindAll()
    if err != nil {
        return nil, err
    }
    var orderResponses []dto.OrderResponse
    for _, order := range orders {
        orderResponses = append(orderResponses, mapper.ToOrderResponse(order))
    }
    return orderResponses, nil
}

func (s *OrderServiceImpl) UpdateOrderStatus(id uint, status string) (dto.OrderResponse, error) {
    order, err := s.OrderRepository.FindById(id)
    if err != nil {
        return dto.OrderResponse{}, errors.New("Order not found")
    }
    order.Status = status
    updatedOrder, err := s.OrderRepository.Save(order)
    if err != nil {
        return dto.OrderResponse{}, err
    }
    return mapper.ToOrderResponse(updatedOrder), nil
}

func (s *OrderServiceImpl) DeleteOrder(id uint) error {
    return s.OrderRepository.DeleteById(id)
}
