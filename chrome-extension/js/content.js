// YouTube Stats Overlay - Content Script
// Runs on YouTube watch pages and injects stats overlay

console.log('YouTube Stats Overlay - Content script loaded');

class StatsOverlay {
  constructor() {
    this.videoId = null;
    this.overlay = null;
    this.apiEndpoint = 'http://localhost:8998';
    this.refreshInterval = null;
    this.isEnabled = true;
    
    this.init();
  }

  async init() {
    const settings = await chrome.storage.sync.get(['apiEndpoint', 'isEnabled', 'refreshRate']);
    this.apiEndpoint = settings.apiEndpoint || 'http://localhost:8998';
    this.isEnabled = settings.isEnabled !== false;
    const refreshRate = settings.refreshRate || 30000;

    if (!this.isEnabled) {
      console.log('Stats overlay is disabled');
      return;
    }

    this.waitForPlayer();
    this.observeVideoChanges();
    
    if (refreshRate > 0) {
      this.refreshInterval = setInterval(() => this.updateStats(), refreshRate);
    }
  }

  waitForPlayer() {
    const checkPlayer = setInterval(() => {
      const player = document.querySelector('#movie_player');
      if (player) {
        clearInterval(checkPlayer);
        this.extractVideoId();
        if (this.videoId) {
          this.createOverlay();
          this.updateStats();
        }
      }
    }, 500);
  }

  extractVideoId() {
    const urlParams = new URLSearchParams(window.location.search);
    this.videoId = urlParams.get('v');
    console.log('Video ID:', this.videoId);
  }

  createOverlay() {
    const existing = document.getElementById('yt-stats-overlay');
    if (existing) existing.remove();

    this.overlay = document.createElement('div');
    this.overlay.id = 'yt-stats-overlay';
    this.overlay.className = 'yt-stats-overlay';
    
    const header = document.createElement('div');
    header.className = 'yt-stats-header';
    header.innerHTML = '<span class="yt-stats-title">📊 Video Stats</span><button class="yt-stats-toggle" title="Toggle stats">−</button><button class="yt-stats-close" title="Close">×</button>';
    
    const content = document.createElement('div');
    content.className = 'yt-stats-content';
    content.innerHTML = '<div class="yt-stats-loading">Loading stats...</div>';
    
    this.overlay.appendChild(header);
    this.overlay.appendChild(content);

    const player = document.querySelector('#movie_player');
    if (player) {
      player.appendChild(this.overlay);
      
      this.overlay.querySelector('.yt-stats-close').addEventListener('click', () => {
        this.overlay.remove();
      });
      
      this.overlay.querySelector('.yt-stats-toggle').addEventListener('click', (e) => {
        const content = this.overlay.querySelector('.yt-stats-content');
        const button = e.target;
        if (content.style.display === 'none') {
          content.style.display = 'block';
          button.textContent = '−';
        } else {
          content.style.display = 'none';
          button.textContent = '+';
        }
      });
    }
  }

  async updateStats() {
    if (!this.videoId || !this.overlay) return;

    try {
      const response = await fetch(this.apiEndpoint + '/stats?video_id=' + this.videoId);
      
      if (!response.ok) {
        throw new Error('HTTP error! status: ' + response.status);
      }
      
      const stats = await response.json();
      this.displayStats(stats);
    } catch (error) {
      console.error('Error fetching stats:', error);
      this.displayError(error.message);
    }
  }

  displayStats(stats) {
    const content = this.overlay.querySelector('.yt-stats-content');
    const viewCount = parseInt(stats.viewCount || 0).toLocaleString();
    const likeCount = parseInt(stats.likeCount || 0).toLocaleString();
    const commentCount = parseInt(stats.commentCount || 0).toLocaleString();
    
    content.innerHTML = '<div class="yt-stats-item"><span class="yt-stats-label">👁️ Views</span><span class="yt-stats-value">' + viewCount + '</span></div><div class="yt-stats-item"><span class="yt-stats-label">👍 Likes</span><span class="yt-stats-value">' + likeCount + '</span></div><div class="yt-stats-item"><span class="yt-stats-label">💬 Comments</span><span class="yt-stats-value">' + commentCount + '</span></div><div class="yt-stats-footer"><small>Last updated: ' + new Date().toLocaleTimeString() + '</small></div>';
  }

  displayError(message) {
    const content = this.overlay.querySelector('.yt-stats-content');
    content.innerHTML = '<div class="yt-stats-error">⚠️ Error loading stats<br><small>' + message + '</small><br><small>Make sure yt-stats API is running on ' + this.apiEndpoint + '</small></div>';
  }

  observeVideoChanges() {
    let lastUrl = location.href;
    new MutationObserver(() => {
      const url = location.href;
      if (url !== lastUrl) {
        lastUrl = url;
        if (url.includes('/watch?v=')) {
          console.log('Video changed, updating overlay');
          this.extractVideoId();
          if (this.videoId) {
            this.createOverlay();
            this.updateStats();
          }
        } else {
          const existing = document.getElementById('yt-stats-overlay');
          if (existing) existing.remove();
        }
      }
    }).observe(document.body, { childList: true, subtree: true });
  }

  destroy() {
    if (this.refreshInterval) {
      clearInterval(this.refreshInterval);
    }
    if (this.overlay) {
      this.overlay.remove();
    }
  }
}

if (document.readyState === 'loading') {
  document.addEventListener('DOMContentLoaded', () => new StatsOverlay());
} else {
  new StatsOverlay();
}
