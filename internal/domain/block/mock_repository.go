package block

import "github.com/stretchr/testify/mock"

// MockRepository mocks the block repository
type MockRepository struct {
	mock.Mock
}

// Get mock
func (m *MockRepository) Get() (Block, error) {
	args := m.Called()

	return args.Get(0).(Block), args.Error(1)
}

// Add mock
func (m *MockRepository) Add(block Block) error {
	args := m.Called(block)

	return args.Error(0)
}
