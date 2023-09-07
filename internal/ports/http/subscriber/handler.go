package subscriber

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ydimitriou/eth-blockchain-parser/internal/app"
	"github.com/ydimitriou/eth-blockchain-parser/internal/app/subscriber/commands"
	"github.com/ydimitriou/eth-blockchain-parser/internal/app/subscriber/queries"
)

// Handler is a subscriber http request handler
type Handler struct {
	subServices app.SubscriberServices
}

// NewHandler return an http request handler
func NewHandler(ss app.SubscriberServices) Handler {
	return Handler{subServices: ss}
}

// CreateSubscriberRequest represents the expected model for subscribe requests
type CreateSubscriberRequest struct {
	Address string `json:"address"`
}

// Create adds a new subscriber in storage
func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateSubscriberRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	subscriber := commands.AddSubscriberRequest{
		Address: req.Address,
	}
	err = h.subServices.Commands.AddSubscriberHandler.Handle(subscriber)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}
	w.WriteHeader(http.StatusCreated)
}

const GetTransactionsURLParam = "address"

// GetTransactions returns all inbound and outbound transactions for a subscribers address
func (h Handler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	address := v[GetTransactionsURLParam]
	req := queries.GetSubscriberRequest{Address: address}
	subTx, err := h.subServices.Queries.GetSubscriberHandler.Handle(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	if subTx == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "This Address is not a suscriber. Please subscribe to get your transactions")
		return
	}
	err = json.NewEncoder(w).Encode(subTx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
}
