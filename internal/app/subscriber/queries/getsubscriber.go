package queries

import (
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/subscriber"
)

// GetSubscriberRequest Model of GetSubscriberHandler
type GetSubscriberRequest struct {
	Address string
}

// GetSubscriberResult represents the result of get subscriber request
type GetSubscriberResult struct {
	Address      string
	Transactions []TransactionResult
}

type TransactionResult struct {
	From string
	To   string
}

// GetSubscriberHandler interface for getting subscriber details
type GetSubscriberHandler interface {
	Handle(req GetSubscriberRequest) (*GetSubscriberResult, error)
}

type getSubscriberHandler struct {
	repo subscriber.Repository
}

// NewGetSubscriberHandler constructor
func NewGetSubscriberHandler(repo subscriber.Repository) GetSubscriberHandler {
	return getSubscriberHandler{repo: repo}
}

// Handle handles get subscriber requests
func (h getSubscriberHandler) Handle(req GetSubscriberRequest) (*GetSubscriberResult, error) {
	var sub *GetSubscriberResult
	res, err := h.repo.GetByAddress(req.Address)
	if err != nil || res == nil {
		return sub, err
	}

	sub = &GetSubscriberResult{
		Address:      res.Address,
		Transactions: []TransactionResult{},
	}

	for _, inTx := range res.Transactions {
		tx := TransactionResult{
			From: inTx.From,
			To:   inTx.To,
		}
		sub.Transactions = append(sub.Transactions, tx)
	}

	return sub, nil
}
