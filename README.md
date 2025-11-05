# YouTube Stats Terminal App

A Go-based HTTP API service that fetches YouTube video statistics using the YouTube Data API v3. The project includes observability features with OpenTelemetry and Prometheus metrics.

## Features

- 📊 Fetch YouTube video statistics (views, likes, comments, etc.)
- 🔭 OpenTelemetry tracing with stdout exporter
- 📈 Prometheus metrics endpoint
- 📝 Structured logging with zerolog
- 🐳 Docker support for easy deployment
- 🧪 Comprehensive test suite with 42%+ coverage
- 🌐 **Chrome Extension**: Real-time stats overlay while watching YouTube videos

## Project Vision

This project has two main components:

1. **API Service**: Go-based HTTP API that fetches YouTube statistics
2. **Chrome Extension**: Browser extension that displays video stats as an overlay while you watch

## Prerequisites

- Go 1.24 or later
- YouTube Data API v3 key ([Get one here](https://console.cloud.google.com/apis/credentials))
- Docker and Docker Compose (optional, for containerized deployment)

## Quick Start

### Using Docker Compose (Recommended)

1. Clone the repository:
```bash
git clone https://github.com/ennc0d3/yt-stats.git
cd yt-stats
```

2. Create a `.env` file from the example:
```bash
cp .env.example .env
```

3. Edit `.env` and add your YouTube API key:
```bash
YTSTATS_API_KEY=your_actual_api_key_here
```

4. Start the service:
```bash
docker-compose up -d
```

5. Test the API:
```bash
# Replace VIDEO_ID with an actual YouTube video ID
curl "http://localhost:8998/stats?video_id=dQw4w9WgXcQ"
```

### Using Go Directly

1. Set your API key:
```bash
export YTSTATS_API_KEY=your_youtube_api_key
```

2. Build and run:
```bash
go build -o yt-stats ./cmd/yt-stats
./yt-stats
```

## API Endpoints

### Get Video Statistics
```
GET /stats?video_id={VIDEO_ID}
```

Returns JSON with video statistics including views, likes, comments, etc.

**Example:**
```bash
curl "http://localhost:8998/stats?video_id=dQw4w9WgXcQ"
```

### Prometheus Metrics
```
GET /metrics
```

Returns Prometheus-formatted metrics for monitoring.

## Chrome Extension

The project includes a Chrome extension that displays real-time YouTube statistics as an overlay while you watch videos.

### Installation

1. **Start the API server** (see Quick Start above)

2. **Load the extension:**
   - Open Chrome and go to `chrome://extensions/`
   - Enable "Developer mode" (toggle in top right)
   - Click "Load unpacked"
   - Select the `chrome-extension` directory from this project

3. **Configure settings:**
   - Click the extension icon in Chrome toolbar
   - Set API endpoint (default: `http://localhost:8998`)
   - Set refresh rate (default: 30 seconds)
   - Click "Save Settings"

4. **Watch YouTube videos** - The stats overlay will automatically appear!

### Features

- 👁️ Real-time view count
- 👍 Real-time like count
- 💬 Real-time comment count
- 🔄 Auto-refresh with configurable intervals
- 🎨 Modern glassmorphism UI
- ⏸️ Collapsible/closeable overlay
- ⚙️ Customizable settings

For detailed documentation, see [`chrome-extension/README.md`](chrome-extension/README.md)

## Testing

### Testing the API Service

#### 1. Quick API Test

**Start the server:**
```bash
# Using Docker Compose
docker-compose up -d

# OR using Go directly
export YTSTATS_API_KEY=your_youtube_api_key
go run cmd/yt-stats/main.go
```

**Test with curl:**
```bash
# Test with a well-known video (Rick Astley - Never Gonna Give You Up)
curl "http://localhost:8998/stats?video_id=dQw4w9WgXcQ"

# Expected response (example):
# {
#   "viewCount": "1234567890",
#   "likeCount": "12345678",
#   "commentCount": "123456",
#   "favoriteCount": "0"
# }
```

**Test the metrics endpoint:**
```bash
curl http://localhost:8998/metrics
```

#### 2. Test Different Scenarios

**Missing video_id parameter:**
```bash
curl "http://localhost:8998/stats"
# Expected: 400 Bad Request - "video_id parameter is missing"
```

**Invalid video_id:**
```bash
curl "http://localhost:8998/stats?video_id=invalid"
# Expected: 500 Internal Server Error - "failed to retrieve video statistics"
```

**Valid video_id:**
```bash
# Use any real YouTube video ID from youtube.com/watch?v=VIDEO_ID
curl "http://localhost:8998/stats?video_id=YOUR_VIDEO_ID"
# Expected: 200 OK with JSON stats
```

### Testing the Chrome Extension

#### 1. Initial Setup Test

**Load the extension:**
1. Open Chrome: `chrome://extensions/`
2. Enable "Developer mode"
3. Click "Load unpacked"
4. Select `chrome-extension/` folder
5. ✅ Extension should appear with "YouTube Stats Overlay" title

**Configure settings:**
1. Click extension icon in toolbar
2. Verify default settings:
   - ✅ "Enable Stats Overlay" is checked
   - ✅ API Endpoint shows `http://localhost:8998`
   - ✅ Refresh Rate shows `30` seconds
3. Click "Save Settings"
4. ✅ You should see "Settings saved!" message

#### 2. Test on YouTube

**Navigate to any YouTube video:**
```
https://www.youtube.com/watch?v=dQw4w9WgXcQ
```

**Verify overlay appears:**
- ✅ Stats overlay appears in top-right corner
- ✅ Shows "📊 Video Stats" header
- ✅ Shows three stats: 👁️ Views, 👍 Likes, 💬 Comments
- ✅ Numbers are formatted with thousands separators (e.g., "1,234,567")
- ✅ Shows "Last updated" timestamp at bottom

**Test overlay controls:**
- ✅ Click **−** button → Panel collapses (only header visible)
- ✅ Click **+** button → Panel expands again
- ✅ Click **×** button → Overlay disappears

**Test video navigation:**
- ✅ Click another video in sidebar
- ✅ Overlay should update with new video's stats
- ✅ Stats should refresh automatically (wait 30 seconds)

#### 3. Test Error Handling

**Stop the API server:**
```bash
docker-compose down
# OR press Ctrl+C if running with Go
```

**Reload YouTube page:**
- ✅ Overlay should show error message:
  - "⚠️ Error loading stats"
  - Shows connection error
  - Shows helpful hint about API server

**Restart API server:**
```bash
docker-compose up -d
```

**Reload YouTube page:**
- ✅ Overlay should work again with stats

#### 4. Test Settings Changes

**Open extension popup:**
1. Change refresh rate to `10` seconds
2. Click "Save Settings"
3. ✅ Settings saved message appears

**Reload YouTube page:**
- ✅ Stats should now refresh every 10 seconds
- ✅ Check "Last updated" timestamp changes

**Disable overlay:**
1. Open extension popup
2. Uncheck "Enable Stats Overlay"
3. Click "Save Settings"
4. Reload YouTube page
5. ✅ Overlay should NOT appear

### Running Automated Tests

#### Unit Tests (No API Key Required)

```bash
# Run all unit tests
go test -v -short ./...

# OR using Make
make test-unit

# Expected output:
# === RUN   TestHandleVideoInfo_MissingVideoID
# --- PASS: TestHandleVideoInfo_MissingVideoID (0.00s)
# === RUN   TestHandleVideoInfo_EmptyVideoID
# --- PASS: TestHandleVideoInfo_EmptyVideoID (0.00s)
# ...
# PASS
# ok      github.com/ennc0d3/yt-stats/internal/api
```

#### Integration Tests (Requires API Key)

```bash
# Set your YouTube API key
export YTSTATS_API_KEY=your_actual_api_key

# Run integration tests
go test -v -tags=integration ./...

# OR using test script
./test.sh

# Expected output:
# === RUN   TestIntegration_VideoStats
# --- PASS: TestIntegration_VideoStats (1.50s)
#     integration_test.go:XX: Video stats: Views=1234567890, Likes=12345678, Comments=123456
# PASS
```

#### Full Test Suite with Coverage

```bash
# Run all tests with coverage report
./test.sh

# OR using Make
make test

# Expected output:
# Running yt-stats tests...
#
# === Running Unit Tests ===
# ...tests pass...
#
# === Running Integration Tests ===
# ...tests pass...
#
# === Generating Coverage Report ===
# github.com/ennc0d3/yt-stats/internal/api/handler.go:13:    handleVideoInfo     57.1%
# github.com/ennc0d3/yt-stats/internal/api/server.go:18:     SetupRoutes         100.0%
# ...
# total:                                                      42.3%
#
# ✅ All tests passed!
```

#### Test with Race Detector

```bash
# Detect race conditions
go test -race ./...

# Expected: No race conditions detected
```

### Docker Build Test

```bash
# Test Docker build
make docker-build

# OR
docker build -t yt-stats:test .

# Expected: Build completes successfully
# ✅ Successfully tagged yt-stats:test
```

### Complete End-to-End Test

**1. Start everything:**
```bash
# Terminal 1: Start API
export YTSTATS_API_KEY=your_key
docker-compose up
```

**2. Test API:**
```bash
# Terminal 2: Test API
curl "http://localhost:8998/stats?video_id=dQw4w9WgXcQ" | jq
# ✅ Should return formatted JSON with stats
```

**3. Test Chrome Extension:**
- Load extension in Chrome
- Navigate to `https://www.youtube.com/watch?v=dQw4w9WgXcQ`
- ✅ Overlay should show matching stats from step 2

**4. Verify auto-refresh:**
- Wait 30 seconds (or your configured refresh rate)
- ✅ "Last updated" timestamp should change
- ✅ Stats might update if video stats changed

### Troubleshooting Tests

#### API Tests Failing?

**Check API key:**
```bash
echo $YTSTATS_API_KEY
# Should show your API key, not empty
```

**Check server is running:**
```bash
curl http://localhost:8998/metrics
# Should return Prometheus metrics
```

**Check logs:**
```bash
# Docker logs
docker-compose logs -f

# OR if running with Go, check terminal output
```

#### Chrome Extension Not Working?

**Open Chrome DevTools:**
1. On YouTube page, press F12
2. Check Console tab for errors
3. Look for "YouTube Stats Overlay - Content script loaded"

**Check extension background service worker:**
1. Go to `chrome://extensions/`
2. Find "YouTube Stats Overlay"
3. Click "service worker" link
4. Check for error messages

**Verify API is accessible:**
```bash
# From browser console (F12), run:
fetch('http://localhost:8998/stats?video_id=dQw4w9WgXcQ')
  .then(r => r.json())
  .then(console.log)
```

#### Tests Taking Too Long?

**Run only unit tests (faster):**
```bash
make test-unit
# Skips integration tests that make real API calls
```

**Increase test timeout:**
```bash
go test -timeout 5m ./...
```

## Development

### Building
```bash
# Using Go
go build -v ./cmd/yt-stats

# Using Make
make build
```

### Running Tests
```bash
# Run all tests with coverage
./test.sh

# Or using Make
make test              # All tests with coverage
make test-unit         # Unit tests only
make test-integration  # Integration tests (requires YTSTATS_API_KEY)
```

**Test Coverage**: 42.3% on API package with comprehensive unit and integration tests

### Building with Docker
```bash
# Using Docker directly
docker build -t yt-stats .

# Using Make
make docker-build

# Using Docker Compose
make docker-up
```

## Configuration

The application is configured via environment variables:

- `YTSTATS_API_KEY` (required): Your YouTube Data API v3 key

## Architecture

- **Port**: 8998
- **Logging**: Structured JSON logging via zerolog
- **Tracing**: OpenTelemetry with stdout exporter
- **Metrics**: Prometheus metrics exposed on `/metrics`
- **Router**: Gorilla Mux with OpenTelemetry middleware

## Project Structure

```
.
├── cmd/yt-stats/           # Application entry point
│   └── main.go             # Main application
├── internal/api/           # API handlers and server logic
│   ├── server.go           # HTTP server setup
│   ├── handler.go          # Request handlers
│   ├── ytutil.go           # YouTube API integration
│   ├── *_test.go           # Unit and integration tests
├── chrome-extension/       # Chrome browser extension
│   ├── manifest.json       # Extension configuration
│   ├── js/                 # JavaScript files
│   │   ├── content.js      # YouTube page overlay injector
│   │   ├── background.js   # Service worker
│   │   └── popup.js        # Settings UI logic
│   ├── css/                # Styling
│   │   └── overlay.css     # Stats overlay styles
│   ├── popup.html          # Settings page
│   └── README.md           # Extension documentation
├── Dockerfile              # Docker build configuration
├── docker-compose.yml      # Docker Compose setup
├── Makefile                # Development tasks
├── test.sh                 # Test runner script
└── README.md               # This file
```

## License

See LICENSE file for details.
