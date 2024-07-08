package main

import (
	"net/http"

	common "github.com/piyushbag/oms-common"
	pb "github.com/piyushbag/oms-common/api"
)

// handler handles requests
type handler struct {
	// gateway service
	client pb.OrderServiceClient
}

// NewHandler creates a new handler
func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client}
}

// registerRoutes registers routes
func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.createOrder)
	// register routes
}

// createOrder creates a new order
func (h *handler) createOrder(w http.ResponseWriter, r *http.Request) {
	// create order
	customerID := r.URL.Query().Get("customerID")
	var items []*pb.ItemsWithQuantity
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, "invalid request")
		return
	}

	// create order
	h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})
}
