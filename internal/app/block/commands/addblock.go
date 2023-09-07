package commands

import "github.com/ydimitriou/eth-blockchain-parser/internal/domain/block"

type AddBlockRequest struct {
	Number string
}

type AddBlockHandler interface {
	Handle(req AddBlockRequest) error
}

type addBlockHandler struct {
	repo block.Repository
}

func NewAddBlockHandler(repo block.Repository) AddBlockHandler {
	return addBlockHandler{repo: repo}
}

func (h addBlockHandler) Handle(req AddBlockRequest) error {
	b := block.Block{Number: req.Number}
	err := h.repo.Add(b)
	if err != nil {
		return err
	}

	return nil
}
