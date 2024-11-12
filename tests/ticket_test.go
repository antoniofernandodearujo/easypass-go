package tests

import (
	"easypass-go/models"
	"easypass-go/services"
	"testing"
)

func TestTicketConcurrency(t *testing.T) {
    service := services.NewTicketService(10)
    service.Run()

    successCount := 0
    failedCount := 0
    totalRequests := 20

    for i := 0; i < totalRequests; i++ {
        go func() {
            request := &models.TicketPurchaseRequest{Success: make(chan bool)}
            service.PurchaseTicket(request)

            if <-request.Success {
                successCount++
            } else {
                failedCount++
            }
        }()
    }

    if successCount != 10 || failedCount != 10 {
        t.Errorf("Expected 10 successful purchases and 10 failures, got %d successes and %d failures", successCount, failedCount)
    }
}
