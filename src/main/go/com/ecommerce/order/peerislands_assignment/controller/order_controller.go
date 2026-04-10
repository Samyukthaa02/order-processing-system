
package controller

import (
    "com/ecommerce/order/peerislands_assignment/dto"
    "com/ecommerce/order/peerislands_assignment/service"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type OrderController struct {
    OrderService service.OrderService
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
    var orderRequest dto.OrderRequest
    if err := ctx.ShouldBindJSON(&orderRequest); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    orderResponse, err := c.OrderService.CreateOrder(orderRequest)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, orderResponse)
}

func (c *OrderController) GetOrder(ctx *gin.Context) {
    id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
        return
    }
    orderResponse, err := c.OrderService.GetOrder(uint(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, orderResponse)
}

func (c *OrderController) GetAllOrders(ctx *gin.Context) {
    orderResponses, err := c.OrderService.GetAllOrders()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, orderResponses)
}

func (c *OrderController) UpdateOrderStatus(ctx *gin.Context) {
    id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
        return
    }
    var updateOrderStatusRequest dto.UpdateOrderStatusRequest
    if err := ctx.ShouldBindJSON(&updateOrderStatusRequest); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    orderResponse, err := c.OrderService.UpdateOrderStatus(uint(id), updateOrderStatusRequest.Status)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, orderResponse)
}

func (c *OrderController) DeleteOrder(ctx *gin.Context) {
    id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
        return
    }
    err = c.OrderService.DeleteOrder(uint(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusNoContent, nil)
}
