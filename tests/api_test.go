package tests

import (
	"net/http"
	"net/http/httptest"
	"product-management-system/api"
	"testing"
)

func TestGetProductByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/products/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.GetProductByID)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
