package postgresql

import (
	"testing"

	"github.com/google/uuid"
)

func TestUserRepository_Create(t *testing.T) {
	user := &User{
		ID:          uuid.New(),
		Username:    "testuser",
		Email:       "testuser@example.com",
		Password:    "password",
		Nickname:    "Test User",
		Description: "test user",
	}

	repo := NewUserRepository()
	err := repo.Create(user)
	t.Logf("Create user: %+v\n", user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}
}
