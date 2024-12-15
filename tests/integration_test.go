package tests

import (
	"net/http"
	"net/http/httptest"
	"product-management-system/main"
	"testing"
)

func TestIntegrationGetProduct(t *testing.T) {
	server := httptest.NewServer(main.SetupRoutes())
	defer server.Close()

	resp, err := http.Get(server.URL + "/products/1")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}
}
