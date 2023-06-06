package usecase

import (
	"context"
	"errors"
	"myapi/internal/app/domain"

	"testing"
)

type MockUserRepo struct{}

func (m *MockUserRepo) Save(ctx context.Context, user domain.User) error { return nil }
func (m *MockUserRepo) FindAll(ctx context.Context) ([]domain.User, error){
	users := []domain.User {
		{},
	}

	 return users, nil 
}
func (m *MockUserRepo) FindByUUID(ctx context.Context, uuid string) (domain.User, error){ return domain.User{}, nil }
func (m *MockUserRepo) FindByUsername(ctx context.Context, username string) (domain.User, error){ return domain.User{}, nil }
func (m *MockUserRepo) DeleteById(ctx context.Context, id int) error{ return nil }

func TestUser_Save(t *testing.T) {

	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Valid User",
			name:        "asdasdas",
			expectedErr: errors.New("asdasd"),
		},
	}

	for _, tc := range testCases {
		// Run Tests
		t.Run(tc.test, func(t *testing.T) {
			// Create a new customer
			user := domain.User{
				Username: "ciroazzi",
				Password: "123",
			}

			ucase := NewUserUsecase(&MockUserRepo{})
			err := ucase.Save(context.TODO(), user)
			// Check if the error matches the expected error
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}
}
