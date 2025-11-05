# YouTube Stats Terminal App

A Go-based HTTP API service that fetches YouTube video statistics using the YouTube Data API v3. The project includes observability features with OpenTelemetry and Prometheus metrics.

## Features

- Fetch YouTube video statistics (views, likes, comments, etc.)
- OpenTelemetry tracing with stdout exporter
- Prometheus metrics endpoint
- Structured logging with zerolog
- Docker support for easy deployment

## Project Vision

This project has two main goals:

1. **API Service**: Fetch YouTube statistics for specific videos via HTTP API
2. **Browser Extension** (planned): Display video stats as an overlay/watermark while watching YouTube videos

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

## Development

### Building
```bash
go build -v ./cmd/yt-stats
```

### Running Tests
```bash
./test.sh
```

### Building with Docker
```bash
docker build -t yt-stats .
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
├── cmd/yt-stats/       # Application entry point
├── internal/api/       # API handlers and server logic
│   ├── server.go       # HTTP server setup
│   ├── handler.go      # Request handlers
│   └── ytutil.go       # YouTube API integration
├── Dockerfile          # Docker build configuration
├── docker-compose.yml  # Docker Compose setup
└── README.md          # This file
```

## License

See LICENSE file for details.
