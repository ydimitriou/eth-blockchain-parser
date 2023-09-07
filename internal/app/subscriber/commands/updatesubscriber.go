package commands

import (
	"fmt"

	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/subscriber"
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/transaction"
)

type UpdateSubscriberRequest struct {
	Address     string
	Transaction TransactionsRequest
}

type TransactionsRequest struct {
	From string
	To   string
}

type UpdateSubscriberHandler interface {
	Handle(req UpdateSubscriberRequest) error
}

type updateSubscriberHandler struct {
	repo subscriber.Repository
}

func NewUpdateSubscriberHandler(repo subscriber.Repository) UpdateSubscriberHandler {
	return updateSubscriberHandler{repo: repo}
}

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
