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
