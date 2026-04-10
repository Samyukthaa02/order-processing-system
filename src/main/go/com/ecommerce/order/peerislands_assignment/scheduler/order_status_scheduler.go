
package scheduler

import (
    "com/ecommerce/order/peerislands_assignment/service"
    "github.com/robfig/cron/v3"
    "log"
)

func StartOrderStatusScheduler(orderService service.OrderService) {
    c := cron.New()
    c.AddFunc("@hourly", func() {
        log.Println("Checking for pending orders...")
        orders, err := orderService.GetAllOrders()
        if err != nil {
            log.Printf("Error getting orders: %v", err)
            return
        }

        for _, order := range orders {
            if order.Status == "PENDING" {
                log.Printf("Processing order %d", order.ID)
                _, err := orderService.UpdateOrderStatus(order.ID, "PROCESSING")
                if err != nil {
                    log.Printf("Error updating order status: %v", err)
                }
            }
        }
    })
    c.Start()
}
