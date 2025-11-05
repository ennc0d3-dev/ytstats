// +build integration

package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/api/youtube/v3"
)

// TestIntegration_VideoStats tests the full flow with a real API call
// This test requires YTSTATS_API_KEY to be set and will make real API calls
func TestIntegration_VideoStats(t *testing.T) {
	apiKey := os.Getenv("YTSTATS_API_KEY")
	if apiKey == "" {
		t.Skip("Skipping integration test: YTSTATS_API_KEY not set")
	}

	router := mux.NewRouter()
	SetupRoutes(router)

	// Test with a well-known video ID (Rick Astley - Never Gonna Give You Up)
	req, err := http.NewRequest("GET", "/stats?video_id=dQw4w9WgXcQ", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var stats youtube.VideoStatistics
	err = json.Unmarshal(rr.Body.Bytes(), &stats)
	require.NoError(t, err)

	// Verify we got real statistics
	assert.Greater(t, stats.ViewCount, uint64(0), "View count should be greater than 0")
	t.Logf("Video stats: Views=%d, Likes=%d, Comments=%d",
		stats.ViewCount, stats.LikeCount, stats.CommentCount)
}

// TestIntegration_InvalidVideoID tests error handling with invalid video ID
func TestIntegration_InvalidVideoID(t *testing.T) {
	apiKey := os.Getenv("YTSTATS_API_KEY")
	if apiKey == "" {
		t.Skip("Skipping integration test: YTSTATS_API_KEY not set")
	}

	router := mux.NewRouter()
	SetupRoutes(router)

	req, err := http.NewRequest("GET", "/stats?video_id=invalid-video-id", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Should return error for invalid video ID
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
