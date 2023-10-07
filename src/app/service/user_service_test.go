package service

import (
	"net/http/httptest"
	"testing"

	"github.com/kallepan/go-backend/app/domain/dao"
	"github.com/kallepan/go-backend/app/mocks"
	"github.com/kallepan/go-backend/test"
)

type RegisterUserTest struct {
	expectedStatus int
	request        dao.User
}

var registerUserTests = []RegisterUserTest{
	{
		// Valid user
		expectedStatus: 200,
		request: dao.User{
			Email:     "test@example.com",
			Firstname: "Test",
			Lastname:  "User",
			Username:  "testuser",
			Password:  "testpassword",
		},
	},
	{
		// Duplicate username
		expectedStatus: 400,
		request: dao.User{
			Email:     "test@example.com",
			Firstname: "Test",
			Lastname:  "User",
			Username:  "testuser",
			Password:  "testpassword",
		},
	},
	{
		// Invalid Email
		expectedStatus: 400,
		request: dao.User{
			Email:     "testexample.com",
			Firstname: "Test",
			Lastname:  "User",
			Username:  "testuser2",
			Password:  "testpassword",
		},
	},
}

func TestRegisterUser(t *testing.T) {
	userRepo := mocks.NewMockUserRepositoryInit()
	userService := UserServiceInit(userRepo)

	for i, testStep := range registerUserTests {
		t.Logf("Test %d", i)

		w := httptest.NewRecorder()
		ctx := test.GetGinTestContext(w)

		test.POST(ctx, nil, nil, testStep.request)

		userService.RegisterUser(ctx)

		if w.Code != testStep.expectedStatus {
			t.Errorf("Expected status %d but got %d", testStep.expectedStatus, w.Code)
		}

		if w.Code != 200 {
			continue
		}

		// Check password hash
		if user, err := userRepo.GetUserByUsername(testStep.request.Username); err != nil {
			t.Errorf("Error getting user: %s", err.Error())
		} else if user.Password == testStep.request.Password {
			t.Errorf("Password not hashed")
		}

	}
}
