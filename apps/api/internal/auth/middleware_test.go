package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAuthMiddleware_ExtractsBearerToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userID := uint(123)
	jwtManager := NewJWTManager("test-secret")

	token, err := jwtManager.GenerateToken(userID)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	r := gin.New()

	r.Use(AuthMiddleware(jwtManager))

	r.GET("/test", func(c *gin.Context) {
		getUserID, exists := c.Get("userID")
		if !exists {
			t.Fatal("token was not found in context")
		}

		if getUserID != userID {
			t.Errorf("expected userID 123, got %v", userID)
		}

		c.Status(http.StatusOK)
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/test",
		nil,
	)

	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestAuthMiddleware_ValidToken(t *testing.T) {
	userID := uint(123)

	gin.SetMode(gin.TestMode)

	jwtManager := NewJWTManager("test-secret")

	token, err := jwtManager.GenerateToken(userID)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	r := gin.New()

	r.Use(AuthMiddleware(jwtManager))

	r.GET("/test", func(c *gin.Context) {
		getUserID, exists := c.Get("userID")
		if !exists {
			t.Fatal("user_id was not found in context")
		}

		if getUserID != userID {
			t.Errorf("expected userID 123, got %v", userID)
		}

		c.Status(http.StatusOK)
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/test",
		nil,
	)

	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestAuthMiddleware_MissingAuthorizationHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userID := uint(123)
	jwtManager := NewJWTManager("test-secret")

	_, err := jwtManager.GenerateToken(userID)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	r := gin.New()
	r.Use(AuthMiddleware(jwtManager))

	r.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/test",
		nil,
	)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected status 401, got %d", w.Code)
	}
}

func TestAuthMiddleware_InvalidAuthorizationScheme(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userID := uint(123)
	jwtManager := NewJWTManager("test-secret")

	_, err := jwtManager.GenerateToken(userID)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	r := gin.New()
	r.Use(AuthMiddleware(jwtManager))

	r.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/test",
		nil,
	)

	req.Header.Set("Authorization", "Basic test-token")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected status 401, got %d", w.Code)
	}
}

func TestAuthMiddleware_MissingToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userID := uint(123)
	jwtManager := NewJWTManager("test-secret")

	_, err := jwtManager.GenerateToken(userID)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	r := gin.New()
	r.Use(AuthMiddleware(jwtManager))

	r.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/test",
		nil,
	)

	req.Header.Set("Authorization", "Bearer")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected status 401, got %d", w.Code)
	}
}
