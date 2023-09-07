package queries

import (
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/block"
	"github.com/ydimitriou/eth-blockchain-parser/internal/pkg/hex"
)

type GetBlockResult struct {
	Number int64
}

type GetBlockHandler interface {
	Handle() (*GetBlockResult, error)
}

type getBlockHandler struct {
	repo        block.Repository
	hexProvider hex.Provider
}

func NewGetBlockHandler(repo block.Repository, hp hex.Provider) GetBlockHandler {
	return getBlockHandler{
		repo:        repo,
		hexProvider: hp,
	}
}

func (h getBlockHandler) Handle() (*GetBlockResult, error) {
	var b *GetBlockResult
	res, err := h.repo.Get()
	if err != nil || res.Number == "" {
		return b, err
	}
	b = &GetBlockResult{
		Number: h.hexProvider.HexToInt(res.Number).Int64(),
	}

	return b, nil
}
