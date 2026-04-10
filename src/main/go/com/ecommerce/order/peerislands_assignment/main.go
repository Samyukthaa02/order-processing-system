
package main

import (
    "com/ecommerce/order/peerislands_assignment/config"
    "com/ecommerce/order/peerislands_assignment/controller"
    "com/ecommerce/order/peerislands_assignment/model"
    "com/ecommerce/order/peerislands_assignment/repository"
    "com/ecommerce/order/peerislands_assignment/service"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&model.Order{}, &model.OrderItem{})

    orderRepository := repository.OrderRepository{DB: db}
    orderService := service.OrderServiceImpl{OrderRepository: orderRepository}
    orderController := controller.OrderController{OrderService: &orderService}

    router := gin.Default()

    config.SetupSwagger(router)

    v1 := router.Group("/api/v1")
    {
        orders := v1.Group("/orders")
        {
            orders.POST("", orderController.CreateOrder)
            orders.GET("/:id", orderController.GetOrder)
            orders.GET("", orderController.GetAllOrders)
            orders.PUT("/:id", orderController.UpdateOrderStatus)
            orders.DELETE("/:id", orderController.DeleteOrder)
        }
    }

    router.Run(":8080")
}
