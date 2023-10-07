package mocks

import (
	"github.com/kallepan/go-backend/app/domain/dao"
)

type MockUserRepository struct {
	data map[string]*dao.User
}

// CheckIfUserExists implements repository.UserRepository.
func (m MockUserRepository) CheckIfUserExists(username string) bool {
	user := m.data[username]

	return user != nil
}

// GetUserByUsername implements repository.UserRepository.
func (m MockUserRepository) GetUserByUsername(username string) (*dao.User, error) {
	return m.data[username], nil
}

// RegisterUser implements repository.UserRepository.
func (m MockUserRepository) RegisterUser(user *dao.User) (string, error) {
	m.data[user.Username] = user
	return user.Username, nil
}

func NewMockUserRepositoryInit() *MockUserRepository {
	return &MockUserRepository{
		data: make(map[string]*dao.User),
	}
}
