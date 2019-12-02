package middlewares

import (
	"api/helpers"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func TestAuthMiddlewareMissingToken(t *testing.T) {
	req, _ := http.NewRequest("POST", "/get-commit-information", strings.NewReader(`{"repo":"admiral", "commit":"553aa1e036f04c414a84ad234586a35dd5845db1"}`))

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(HealthCheckHandler)

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AuthMiddleware(handler).ServeHTTP(w, r)
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "{\"error\":\"missing token\"}", rr.Body.String())
}
func TestAuthMiddlewareWrongToken(t *testing.T) {
	req, _ := http.NewRequest("POST", "/get-commit-information", strings.NewReader(`{"token":"test", "repo":"admiral", "commit":"553aa1e036f04c414a84ad234586a35dd5845db1"}`))

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(HealthCheckHandler)

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AuthMiddleware(handler).ServeHTTP(w, r)
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Equal(t, "{\"error\":\"Invalid token.\"}", rr.Body.String())
}

func TestAuthMiddlewareIntegration(t *testing.T) {

	//Generate JWT Token
	jwtClaims := &helpers.Claims{}
	_, token := jwtClaims.GenerateToken("TestUser", time.Now().Add(2*time.Minute).Unix())

	req, _ := http.NewRequest("POST", "/get-commit-information", strings.NewReader(fmt.Sprintf(`{"token":"%s", "repo":"admiral", "commit":"553aa1e036f04c414a84ad234586a35dd5845db1"}`, token)))
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(HealthCheckHandler)
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AuthMiddleware(handler).ServeHTTP(w, r)
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "{\"alive\": true}", rr.Body.String())
}
