package usecase

import (
	"testing"

	"Go-Dispatch-Bootcamp/mocks"
	"Go-Dispatch-Bootcamp/types"

	"github.com/stretchr/testify/assert"
)

var user = types.User{
	Id:         1,
	Username:   "test",
	Identifier: "test",
	FirstName:  "John",
	LastName:   "Doe",
}

type DemoService struct {
	mocks.DemoService
}

func (s *DemoService) GetUsers() (*[]types.User, error) {
	result := []types.User{user}
	return &result, nil
}

func TestDemoUsecase(t *testing.T) {
	uc := New(&DemoService{})

	t.Run("Fetch test", func(t *testing.T) {
		got, err := uc.Fetch()
		expected := user

		assert.Equal(t, (*got)[0], expected)
		assert.NoError(t, err)
	})
}
