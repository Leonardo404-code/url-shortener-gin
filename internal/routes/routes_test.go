package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupRoutes(t *testing.T) {
	r := SetupRoutes()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Hello", w.Body.String())
}
