package main

import (
	"net/http/httptest"
	"testing"

	"github.com/kallepan/go-backend/app/repository"
	"github.com/kallepan/go-backend/app/service"
	"github.com/kallepan/go-backend/test"
)

// Test Ping
func TestPing(t *testing.T) {
	// Fetch the service
	rep := repository.SystemRepositoryInit()
	svc := service.SystemServiceInit(rep)
	// Create a response recorder
	w := httptest.NewRecorder()
	ctx := test.GetGinTestContext(w)

	// Mock GET request with JSON
	test.GET(ctx, nil, nil)

	// Call the handler
	svc.GetPing(ctx)

	// Check the status code
	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}
}
