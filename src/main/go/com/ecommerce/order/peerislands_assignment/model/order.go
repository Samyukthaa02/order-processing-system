
package model

import (
    "time"
)

type Order struct {
    ID              uint      `gorm:"primaryKey"`
    CustomerName    string    `gorm:"not null"`
    OrderDate       time.Time `gorm:"not null"`
    Status          string    `gorm:"not null"`
    ShippingAddress string    `gorm:"not null"`
    TotalAmount     float64   `gorm:"not null"`
    OrderItems      []OrderItem
}
