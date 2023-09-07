package subscriber

import "github.com/stretchr/testify/mock"

// MockRepository mocks the subscriber repository
type MockRepository struct {
	mock.Mock
}

// Add mock
func (m *MockRepository) Add(subscriber Subscriber) error {
	args := m.Called(subscriber)

	return args.Error(0)
}

// Update mock
func (m *MockRepository) Update(subscriber Subscriber) error {
	args := m.Called(subscriber)

	return args.Error(0)
}

// Exist mock
func (m *MockRepository) Exist(address string) bool {
	args := m.Called(address)

	return args.Get(0).(bool)
}

// GetByAddress mock
func (m *MockRepository) GetByAddress(address string) (*Subscriber, error) {
	args := m.Called(address)

	return args.Get(0).(*Subscriber), args.Error(1)
}
