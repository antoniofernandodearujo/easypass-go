package services

import (
	"easypass-go/models"
	"sync"
)

type TicketService struct {
    ticketsAvailable int
    purchaseChannel  chan *models.TicketPurchaseRequest
    mu               sync.Mutex
}

func NewTicketService(totalTickets int) *TicketService {
    return &TicketService{
        ticketsAvailable: totalTickets,
        purchaseChannel:  make(chan *models.TicketPurchaseRequest),
    }
}

func (s *TicketService) PurchaseTicket(request *models.TicketPurchaseRequest) bool {
    s.mu.Lock()
    defer s.mu.Unlock()

    if s.ticketsAvailable > 0 {
        s.ticketsAvailable--
        return true
    }
    return false
}

func (s *TicketService) Run() {
    go func() {
        for request := range s.purchaseChannel {
            if s.PurchaseTicket(request) {
                request.Success <- true
            } else {
                request.Success <- false
            }
        }
    }()
}
