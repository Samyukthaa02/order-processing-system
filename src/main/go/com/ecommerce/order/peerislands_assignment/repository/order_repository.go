
package repository

import (
    "com/ecommerce/order/peerislands_assignment/model"
    "gorm.io/gorm"
)

type OrderRepository struct {
    DB *gorm.DB
}

func (r *OrderRepository) FindAll() ([]model.Order, error) {
    var orders []model.Order
    err := r.DB.Find(&orders).Error
    return orders, err
}

func (r *OrderRepository) FindById(id uint) (model.Order, error) {
    var order model.Order
    err := r.DB.First(&order, id).Error
    return order, err
}

func (r *OrderRepository) Save(order model.Order) (model.Order, error) {
    err := r.DB.Save(&order).Error
    return order, err
}

func (r *OrderRepository) DeleteById(id uint) error {
    return r.DB.Delete(&model.Order{}, id).Error
}
