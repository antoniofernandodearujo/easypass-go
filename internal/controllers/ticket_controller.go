package controllers

import (
	"easypass-go/models"
	"easypass-go/services"
	"encoding/json"
	"net/http"
)

type TicketController struct {
	service *services.TicketService
}

// NewTicketController cria um novo TicketController.
func NewTicketController(service *services.TicketService) *TicketController {
	return &TicketController{service: service}
}

// PurchaseTicketHandler lida com as requisições de compra de ingressos.
func (c *TicketController) PurchaseTicketHandler(w http.ResponseWriter, r *http.Request) {
	var request models.TicketPurchaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// Canal para receber o status da compra
	request.Success = make(chan bool)
	c.service.purchaseChannel <- &request
	success := <-request.Success

	if success {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Ticket purchased successfully"))
	} else {
		http.Error(w, "Ticket purchase failed", http.StatusConflict)
	}
}
