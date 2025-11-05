package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHandleVideoInfo_MissingVideoID(t *testing.T) {
	req, err := http.NewRequest("GET", "/stats", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/stats", handleVideoInfo).Methods(http.MethodGet)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "video_id parameter is missing")
}

func TestHandleVideoInfo_EmptyVideoID(t *testing.T) {
	req, err := http.NewRequest("GET", "/stats?video_id=", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/stats", handleVideoInfo).Methods(http.MethodGet)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "video_id parameter is missing")
}

func TestHandleVideoInfo_InvalidVideoID_WithoutAPIKey(t *testing.T) {
	// This test verifies error handling when API key is not set
	req, err := http.NewRequest("GET", "/stats?video_id=invalid-id", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/stats", handleVideoInfo).Methods(http.MethodGet)

	router.ServeHTTP(rr, req)

	// Should return internal server error when API key is missing or invalid
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Contains(t, rr.Body.String(), "failed to retrieve video statistics")
}
