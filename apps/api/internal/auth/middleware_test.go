package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAuthMiddleware_ExtractsBearerToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()

	r.Use(AuthMiddleware())

	r.GET("/test", func(c *gin.Context) {
		token, exists := c.Get("token")
		if !exists {
			t.Fatal("token was not found in context")
		}

		if token != "test-token" {
			t.Errorf("expected test-token, got %v", token)
		}

		c.Status(http.StatusOK)
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/test",
		nil,
	)

	req.Header.Set("Authorization", "Bearer test-token")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestAuthMiddleware_MissingAuthorizationHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.Use(AuthMiddleware())

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

	r := gin.New()
	r.Use(AuthMiddleware())

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

	r := gin.New()
	r.Use(AuthMiddleware())

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
