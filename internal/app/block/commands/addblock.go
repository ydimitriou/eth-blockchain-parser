package commands

import "github.com/ydimitriou/eth-blockchain-parser/internal/domain/block"

// AddBlockRequest Model of AddBlockHandler
type AddBlockRequest struct {
	Number string
}

// AddBlockHandler interface for adding a block
type AddBlockHandler interface {
	Handle(req AddBlockRequest) error
}

type addBlockHandler struct {
	repo block.Repository
}

// NewAddBlockHandler constructor
func NewAddBlockHandler(repo block.Repository) AddBlockHandler {
	return addBlockHandler{repo: repo}
}

// Handle handles add block requests
func (h addBlockHandler) Handle(req AddBlockRequest) error {
	b := block.Block{Number: req.Number}
	err := h.repo.Add(b)
	if err != nil {
		return err
	}

	return nil
}
