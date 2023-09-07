package memory

import (
	"fmt"

	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/block"
)

// BlockRepository is the in-memory block repository implementation
type BlockRepository struct {
	block block.Block
}

// NewBlockRepository constructor
func NewBlockRepository() *BlockRepository {
	block := block.Block{}
	return &BlockRepository{
		block: block,
	}
}

// Get returns the current block in memory
func (mr *BlockRepository) Get() (block.Block, error) {
	if mr.block.Number == "" {
		return mr.block, fmt.Errorf("no block exists in memory")
	}

	return mr.block, nil
}

// Add adds given block as the current block in memory
func (mr *BlockRepository) Add(block block.Block) error {
	mr.block = block

	return nil
}
