package commands

import (
	"fmt"

	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/subscriber"
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/transaction"
)

// UpdateSubscriberRequest Model of UpdateSubscriberHandler
type UpdateSubscriberRequest struct {
	Address     string
	Transaction TransactionsRequest
}

type TransactionsRequest struct {
	From string
	To   string
}

// UpdateSubscriberHandler interface for updating subscribers transactions
type UpdateSubscriberHandler interface {
	Handle(req UpdateSubscriberRequest) error
}

type updateSubscriberHandler struct {
	repo subscriber.Repository
}

// NewUpdateSubscriberHandler constructor
func NewUpdateSubscriberHandler(repo subscriber.Repository) UpdateSubscriberHandler {
	return updateSubscriberHandler{repo: repo}
}

// Handle handles subscriber updates requests
func (h updateSubscriberHandler) Handle(req UpdateSubscriberRequest) error {
	s, err := h.repo.GetByAddress(req.Address)
	if err != nil {
		return err
	}
	if s == nil {
		return fmt.Errorf("update failed, subscriber with adress %v does not exist", req.Address)
	}
	tx := transaction.Transaction{From: req.Transaction.From, To: req.Transaction.To}
	s.Transactions = append(s.Transactions, tx)

	return h.repo.Update(*s)
}
