package block

// Repository interface for block
type Repository interface {
	Get() (Block, error)
	Add(block Block) error
}
