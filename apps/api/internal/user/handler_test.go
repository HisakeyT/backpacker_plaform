package user

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/HisakeyT/backpacker-platform/internal/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type fakeRepository struct {
	user *User
	err  error
}

func (f *fakeRepository) Create(user *User) error {
	return nil
}

func (f *fakeRepository) FindByID(id uint) (*User, error) {
	return f.user, f.err
}

func (f *fakeRepository) FindByEmail(email string) (*User, error) {
	return f.user, f.err
}

func TestHandler_Register(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		request    string
		repository Repository
		wantStatus int
	}{
		{
			name:       "正常に登録できる",
			request:    `{"nickname":"test","email":"test@example.com","password":"password","password_confirmation":"password"}`,
			repository: &fakeRepository{},
			wantStatus: http.StatusCreated,
		},
		{
			name:    "メールアドレスが重複している",
			request: `{"nickname":"test","email":"test@example.com","password":"password","password_confirmation":"password"}`,
			repository: &fakeRepository{
				user: &User{
					ID:    1,
					Email: "test@example.com",
				},
			},
			wantStatus: http.StatusConflict,
		},
		{
			name:       "パスワードが一致しない",
			request:    `{"nickname":"test","email":"test@example.com","password":"password","password_confirmation":"wrong-password"}`,
			repository: &fakeRepository{},
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "不正なJSON",
			request:    `{"nickname":}`,
			repository: &fakeRepository{},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()

			jwtManager := auth.NewJWTManager("secret")
			useCase := NewUseCase(tt.repository, jwtManager)
			handler := NewHandler(useCase)

			router.POST("/register", handler.Register)

			req := httptest.NewRequest(
				http.MethodPost,
				"/register",
				strings.NewReader(tt.request),
			)
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf(
					"expected status %d, got %d",
					tt.wantStatus,
					rec.Code,
				)
			}
		})
	}
}

func TestHandler_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)

	email := "test@example.com"
	nickName := "test"

	password := "password"
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name       string
		request    string
		repository Repository
		wantStatus int
	}{
		{
			name:    "正常にログインできる",
			request: `{"email":"test@example.com","password":"password"}`,
			repository: &fakeRepository{
				user: &User{
					ID:           1,
					Nickname:     nickName,
					Email:        email,
					PasswordHash: string(passwordHash),
				},
			},
			wantStatus: http.StatusOK,
		},
		{
			name:    "認証情報が不正",
			request: `{"email":"test@example.com","password":"wrong-password"}`,
			repository: &fakeRepository{
				user: &User{
					ID:           1,
					Nickname:     nickName,
					Email:        email,
					PasswordHash: string(passwordHash),
				},
			},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "不正なJSON",
			request:    `{"email":}`,
			repository: &fakeRepository{},
			wantStatus: http.StatusBadRequest,
		},
		{
			name:    "ユーザーが存在しない",
			request: `{"email":"unknown@example.com","password":"password"}`,
			repository: &fakeRepository{
				err: ErrUserNotFound,
			},
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()

			jwtManager := auth.NewJWTManager("secret")
			useCase := NewUseCase(tt.repository, jwtManager)
			handler := NewHandler(useCase)

			router.POST("/login", handler.Login)

			req := httptest.NewRequest(
				http.MethodPost,
				"/login",
				strings.NewReader(tt.request),
			)
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf(
					"expected status %d, got %d",
					tt.wantStatus,
					rec.Code,
				)
			}
		})
	}
}
