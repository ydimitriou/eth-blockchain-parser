package memory

import (
	"fmt"

	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/block"
)

type BlockRepository struct {
	block block.Block
}

func NewBlockRepository() *BlockRepository {
	block := block.Block{}
	return &BlockRepository{
		block: block,
	}
}

func (mr *BlockRepository) Get() (block.Block, error) {
	if mr.block.Number == "" {
		return mr.block, fmt.Errorf("no block exists in memory")
	}

	return mr.block, nil
}

func (mr *BlockRepository) Add(block block.Block) error {
	mr.block = block

	return nil
}
