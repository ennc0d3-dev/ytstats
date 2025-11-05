# YouTube Stats Overlay - Chrome Extension

A Chrome extension that displays real-time YouTube video statistics as an overlay while you watch videos.

## Features

- 📊 Real-time video statistics overlay
- 👁️ View count display
- 👍 Like count display
- 💬 Comment count display
- ⚙️ Configurable API endpoint
- 🔄 Auto-refresh with customizable intervals
- 🎨 Clean, modern UI with glassmorphism design
- ⏸️ Collapsible overlay to minimize distraction
- 🚀 Lightweight and fast

## Installation

### Prerequisites

1. **yt-stats API Server**: The extension requires the yt-stats Go API server to be running
   ```bash
   # From the main project directory
   export YTSTATS_API_KEY=your_youtube_api_key
   go run cmd/yt-stats/main.go
   # Or using Docker
   docker-compose up -d
   ```

### Load Extension in Chrome

1. Open Chrome and navigate to `chrome://extensions/`
2. Enable "Developer mode" (toggle in top right)
3. Click "Load unpacked"
4. Select the `chrome-extension` directory from this project
5. The extension icon should appear in your browser toolbar

## Usage

### First Time Setup

1. Click the extension icon in your browser toolbar
2. Configure settings:
   - **Enable Stats Overlay**: Toggle on/off
   - **API Endpoint**: Default is `http://localhost:8998` (change if your API runs elsewhere)
   - **Refresh Rate**: How often to update stats (in seconds, 0 = never)
3. Click "Save Settings"

### Watching Videos

1. Navigate to any YouTube video
2. The stats overlay will automatically appear in the top-right corner
3. The overlay shows:
   - 👁️ Current view count
   - 👍 Current like count
   - 💬 Current comment count
   - Last update timestamp
4. Use the buttons:
   - **−** (minimize): Collapse the stats panel
   - **×** (close): Remove the overlay

### Customization

- **Move the overlay**: The overlay is positioned in the top-right by default. Future versions will support drag-and-drop positioning.
- **Refresh rate**: Set to 0 to disable auto-refresh, or up to 300 seconds (5 minutes)

## Troubleshooting

### "Error loading stats"

**Cause**: Cannot connect to the API server

**Solutions**:
1. Make sure the yt-stats API server is running:
   ```bash
   curl http://localhost:8998/stats?video_id=dQw4w9WgXcQ
   ```
2. Check the API endpoint in extension settings matches where your server is running
3. Ensure `YTSTATS_API_KEY` environment variable is set

### Overlay not appearing

**Causes**:
1. Extension might be disabled in settings
2. Not on a YouTube watch page
3. YouTube player hasn't loaded yet

**Solutions**:
1. Check extension settings (click icon) - ensure "Enable Stats Overlay" is checked
2. Make sure you're on a URL like `youtube.com/watch?v=...`
3. Refresh the page

### Stats show as 0 or incorrect

**Cause**: API returned empty or invalid data

**Solutions**:
1. Verify the video ID is valid
2. Check API server logs for errors
3. Ensure your YouTube API key has sufficient quota

## Technical Details

### Files

- **manifest.json**: Extension configuration
- **js/content.js**: Main content script that runs on YouTube pages
- **js/background.js**: Service worker for settings management
- **js/popup.js**: Settings UI logic
- **popup.html**: Settings page UI
- **css/overlay.css**: Styling for the stats overlay

### Permissions

- `storage`: Save user settings
- `activeTab`: Access current YouTube tab
- Host permissions:
  - `https://www.youtube.com/*`: Inject overlay on YouTube
  - `http://localhost:8998/*`: Fetch stats from local API

### Architecture

1. Content script (`content.js`) injects overlay into YouTube video player
2. Observes URL changes (YouTube is a Single Page Application)
3. Fetches stats from API endpoint periodically
4. Updates overlay UI with formatted data
5. Settings stored in Chrome sync storage

## Privacy

This extension:
- ✅ Only runs on YouTube watch pages
- ✅ Only communicates with your local API server (or configured endpoint)
- ✅ Does NOT collect or transmit any personal data
- ✅ Does NOT track your viewing history
- ✅ Does NOT inject ads or modify page content (except adding the overlay)
- ✅ All settings are stored locally in Chrome sync storage

## Future Enhancements

- [ ] Drag-and-drop positioning
- [ ] Theme customization (dark/light, colors)
- [ ] Additional stats (subscribers, upload date, etc.)
- [ ] Historical data charts
- [ ] Keyboard shortcuts
- [ ] Multiple overlay layouts
- [ ] Export stats data

## Development

### Testing Locally

1. Make changes to extension files
2. Go to `chrome://extensions/`
3. Click the reload icon on the extension card
4. Refresh any YouTube tabs to see changes

### Debugging

- **Console logs**: Open DevTools on YouTube page, content script logs appear there
- **Background logs**: Go to `chrome://extensions/`, click "service worker" under extension
- **Popup logs**: Right-click extension icon → "Inspect popup"

## License

Same as the main yt-stats project. See LICENSE file.
