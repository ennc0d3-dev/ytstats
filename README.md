# YouTube Stats

A comprehensive YouTube statistics platform featuring a Go-based HTTP API, professional CLI tool, and browser extension for real-time video stats.

## Features

### 🖥️ CLI Tool (NEW!)
- **Professional command-line interface** powered by Cobra + Viper
- Multiple commands: `serve`, `get`, `version`
- Configuration via files, environment variables, or flags
- Multiple output formats: table, JSON, YAML
- Beautiful formatted output with emojis and thousands separators

### 🌐 API Service
- 📊 Fetch YouTube video statistics (views, likes, comments, etc.)
- 📈 Prometheus metrics endpoint
- 🔭 OpenTelemetry tracing with stdout exporter
- 📝 Structured logging with zerolog
- 📄 **OpenAPI 3.0 specification** for API documentation
- 🐳 Docker support for easy deployment

### 🧪 Testing & Quality
- Comprehensive test suite with 42.3% coverage
- Unit + Integration tests
- Race detector enabled
- GitHub Actions CI/CD pipeline
- Security scanning (Gosec + CodeQL)
- Automated dependency updates

### 🌟 Chrome Extension
- Real-time stats overlay while watching YouTube videos
- Modern glassmorphism UI
- Auto-refresh with configurable intervals
- Customizable settings (API endpoint, refresh rate)

## Project Components

1. **CLI Tool**: Multi-command interface with flexible configuration
2. **API Service**: RESTful API for fetching YouTube statistics
3. **Chrome Extension**: Browser extension for in-video stats overlay
4. **CI/CD Pipeline**: Production-ready automated testing and deployment

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
./yt-stats serve  # Start API server
```

## CLI Usage

The CLI tool provides multiple commands for different use cases.

### Start API Server

```bash
# Start server (default port 8998)
yt-stats serve

# Custom port
yt-stats serve --port 9000

# Custom log level
yt-stats serve --log-level debug

# Using config file
yt-stats serve --config ~/.yt-stats.yaml
```

### Get Video Statistics

```bash
# Get stats with default table output
yt-stats get dQw4w9WgXcQ

# JSON output
yt-stats get dQw4w9WgXcQ --format json

# YAML output
yt-stats get dQw4w9WgXcQ --format yaml

# Custom fields
yt-stats get dQw4w9WgXcQ --fields views,likes

# All fields
yt-stats get dQw4w9WgXcQ --fields views,likes,comments,favorites
```

**Example output (table format):**
```
╔═══════════════════════════════════════════════════════════╗
║ YouTube Video Statistics                                  ║
╠═══════════════════════════════════════════════════════════╣
║ Title: Rick Astley - Never Gonna Give You Up             ║
╠═══════════════════════════════════════════════════════════╣
║ 👁️  Views:     1,234,567,890                             ║
║ 👍 Likes:     12,345,678                                 ║
║ 💬 Comments:  123,456                                    ║
╚═══════════════════════════════════════════════════════════╝
```

### Configuration

The CLI supports multiple configuration methods (in order of precedence):

1. **Command-line flags** (highest priority)
2. **Environment variables** (YTSTATS_*)
3. **Configuration file** (.yt-stats.yaml)
4. **Defaults** (lowest priority)

**Create a configuration file:**
```bash
cp .yt-stats.yaml.example ~/.yt-stats.yaml
# Edit with your settings
```

**Example .yt-stats.yaml:**
```yaml
apiKey: "your_youtube_api_key_here"
port: 8998
logLevel: "info"
```

**Environment variables:**
```bash
export YTSTATS_API_KEY="your_key"
export YTSTATS_PORT=8998
export YTSTATS_LOG_LEVEL="debug"
```

### Version Information

```bash
yt-stats version
# Output:
# yt-stats version 1.0.0
#   commit: dev
#   built:  unknown
```

### Help

```bash
# General help
yt-stats --help

# Command-specific help
yt-stats serve --help
yt-stats get --help
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

## API Documentation

### OpenAPI Specification

The API is fully documented using **OpenAPI 3.0** specification.

**Location**: [`openapi.yaml`](openapi.yaml)

**Features:**
- Complete endpoint documentation
- Request/response schemas
- Validation patterns
- Example requests and responses
- Ready for SDK generation

**View the API docs:**
```bash
# Online with Swagger Editor
# Visit: https://editor.swagger.io/
# Paste contents of openapi.yaml

# Generate client SDK
npm install -g @openapitools/openapi-generator-cli
openapi-generator-cli generate -i openapi.yaml -g python -o ./client

# Serve interactive docs locally
npx @stoplight/prism-cli mock openapi.yaml
```

**Documented Endpoints:**
- `GET /stats?video_id={VIDEO_ID}` - Fetch video statistics
- `GET /metrics` - Prometheus metrics

## CI/CD Pipeline

The project has a **production-ready CI/CD pipeline** with GitHub Actions.

**Documentation**: See [`PIPELINE.md`](PIPELINE.md) for complete details.

**Workflows:**
1. **CI/CD Workflow** (`.github/workflows/ci.yml`)
   - Lint (golangci-lint v6)
   - Test (unit + integration with race detector)
   - Build (binary artifacts)
   - Security (Gosec scanner)
   - Docker (BuildKit with caching)

2. **Build Workflow** (`.github/workflows/go.yml`)
   - Production deployment
   - Docker Hub publishing
   - Codecov integration

3. **Security Workflow** (`.github/workflows/codeql.yml`)
   - CodeQL analysis
   - Weekly scans
   - Go + JavaScript coverage

**Pipeline Features:**
- ✅ Parallel job execution
- ✅ Smart caching (Go modules, Docker layers)
- ✅ Security scanning (Gosec + CodeQL)
- ✅ 42.3% test coverage
- ✅ Race detector enabled
- ✅ Automated dependency updates

**Status Badges:**
```markdown
![CI/CD](https://github.com/ennc0d3-dev/ytstats/workflows/CI/CD/badge.svg)
![CodeQL](https://github.com/ennc0d3-dev/ytstats/workflows/CodeQL/badge.svg)
```

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
│   ├── main.go             # Main CLI entry
│   └── cmd/                # CLI commands (NEW!)
│       ├── root.go         # Root command with viper config
│       ├── serve.go        # API server command
│       ├── get.go          # Stats fetcher command
│       └── version.go      # Version command
├── internal/api/           # API handlers and server logic
│   ├── server.go           # HTTP server setup
│   ├── handler.go          # Request handlers
│   ├── ytutil.go           # YouTube API integration
│   ├── handler_test.go     # Handler unit tests
│   ├── server_test.go      # Server route tests
│   └── integration_test.go # Integration tests
├── chrome-extension/       # Chrome browser extension
│   ├── manifest.json       # Extension configuration (Manifest V3)
│   ├── js/                 # JavaScript files
│   │   ├── content.js      # YouTube page overlay injector
│   │   ├── background.js   # Service worker
│   │   └── popup.js        # Settings UI logic
│   ├── css/                # Styling
│   │   └── overlay.css     # Stats overlay styles (glassmorphism)
│   ├── popup.html          # Settings page
│   └── README.md           # Extension documentation
├── .github/workflows/      # GitHub Actions CI/CD
│   ├── ci.yml              # Comprehensive CI pipeline (NEW!)
│   ├── go.yml              # Build and deployment
│   └── codeql.yml          # Security scanning
├── openapi.yaml            # OpenAPI 3.0 specification (NEW!)
├── PIPELINE.md             # CI/CD documentation (NEW!)
├── .yt-stats.yaml.example  # Config file example (NEW!)
├── Dockerfile              # Multi-stage Docker build
├── docker-compose.yml      # Docker Compose setup
├── Makefile                # Development tasks
├── test.sh                 # Test runner script
└── README.md               # This file
```

## Quick Reference

### Common Commands

```bash
# CLI Tool
yt-stats serve                      # Start API server
yt-stats get VIDEO_ID               # Get stats (table)
yt-stats get VIDEO_ID -f json       # Get stats (JSON)
yt-stats version                    # Show version

# API Server
curl "http://localhost:8998/stats?video_id=VIDEO_ID"
curl http://localhost:8998/metrics

# Development
make build                          # Build binary
make test                           # Run all tests with coverage
make test-unit                      # Run unit tests only
make docker-build                   # Build Docker image

# Docker
docker-compose up -d                # Start services
docker-compose logs -f              # View logs
docker-compose down                 # Stop services
```

### Configuration Priority

1. Command-line flags (--api-key, --port)
2. Environment variables (YTSTATS_API_KEY, YTSTATS_PORT)
3. Config file (~/.yt-stats.yaml)
4. Defaults

### Project Links

- **OpenAPI Spec**: [openapi.yaml](openapi.yaml)
- **Pipeline Docs**: [PIPELINE.md](PIPELINE.md)
- **Chrome Extension**: [chrome-extension/README.md](chrome-extension/README.md)
- **GitHub Repository**: https://github.com/ennc0d3-dev/ytstats

## License

See LICENSE file for details.
