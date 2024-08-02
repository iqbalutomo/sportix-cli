package testing

import (
	"sportix-cli/internal/entity"
	"sportix-cli/internal/handler"
	"testing"

	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type MockRepo struct {
	Mock mock.Mock
}

func (m *MockRepo) CreateUser(user *entity.User) error {
	args := m.Mock.Called(user)

	return args.Error(0)
}

func (m *MockRepo) FindUserByEmail(email string) (*entity.User, error) {
	args := m.Mock.Called(email)

	if user, ok := args.Get(0).(*entity.User); ok {
		return user, args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *MockRepo) ValidateUser(email, password string) (*entity.User, error) {
	return &entity.User{}, nil
}

func (m *MockRepo) FindBalanceByEmail(email string) (float64, error) {
	return 0, nil
}

func (m *MockRepo) UpdateBalance(userID uint, deposit float64) error {
	return nil
}

func TestRegister(t *testing.T) {
	mockRepo := new(MockRepo)
	h := handler.NewUserHandler(mockRepo)

	password := "fancybear"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	tests := []struct {
		name       string
		email      string
		password   string
		role       string
		mockReq    func()
		exprectErr bool
	}{
		{
			name:     "testing baleeee",
			email:    "bale@testing.com",
			password: string(hashedPassword),
			role:     "user",
			mockReq: func() {
				user := &entity.User{
					Username: "testing baleee",
					Email:    "bale@testing.com",
					Password: password,
					Role:     "user",
				}

				mockRepo.Mock.On("FindUserByEmail", "bale@testing.com").Return(nil, nil)
				mockRepo.Mock.On("CreateUser", mock.MatchedBy(func(u *entity.User) bool {
					return u.Email == user.Email
				})).Return(nil)
			},
			exprectErr: false,
		},
		{
			name:     "testing baleeee",
			email:    "gaadaboskuh@email.com",
			password: string(hashedPassword),
			role:     "user",
			mockReq: func() {
				user := &entity.User{
					Username: "testing baleee",
					Email:    "gaadaboskuh@email.com",
					Password: password,
					Role:     "user",
				}

				mockRepo.Mock.On("FindUserByEmail", "gaadaboskuh@email.com").Return(user, nil)
				mockRepo.Mock.On("CreateUser", mock.MatchedBy(func(u *entity.User) bool {
					return u.Email == user.Email
				})).Return(nil)
			},
			exprectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.email, func(t *testing.T) {
			test.mockReq()
			err := h.Register("testing baleee", test.email, test.password, test.role)
			if err != nil != test.exprectErr {
				t.Errorf("Register() error = %v, expectErr %v", err, test.exprectErr)
			}
		})
	}
}
