package hex

import (
	"math/big"

	"github.com/stretchr/testify/mock"
)

// MockProvider mocks hex provider
type MockProvider struct {
	mock.Mock
}

// HexToInt mock
func (m *MockProvider) HexToInt(hex string) *big.Int {
	args := m.Called(hex)

	return args.Get(0).(*big.Int)
}

// IntToHex mock
func (m *MockProvider) IntToHex(num int64) string {
	args := m.Called()

	return args.Get(0).(string)
}
