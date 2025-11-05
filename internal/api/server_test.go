package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestSetupRoutes(t *testing.T) {
	router := mux.NewRouter()
	SetupRoutes(router)

	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
	}{
		{
			name:           "Stats endpoint with missing video_id",
			method:         http.MethodGet,
			path:           "/stats",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid HTTP method",
			method:         http.MethodPost,
			path:           "/stats",
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.path, nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
		})
	}
}

// Note: Metrics endpoint is tested in integration tests as it's set up in StartServer
