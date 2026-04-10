
package model

type OrderItem struct {
    ID        uint    `gorm:"primaryKey"`
    ProductID uint    `gorm:"not null"`
    Quantity  int     `gorm:"not null"`
    Price     float64 `gorm:"not null"`
    OrderID   uint
}
