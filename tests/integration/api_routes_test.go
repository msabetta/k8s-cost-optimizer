package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"k8s-cost-optimizer/api"
)

func TestHealthRoute(t *testing.T) {
	handler := api.SetupRouter()
	req := httptest.NewRequest("GET", "/health", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", status)
	}
}
