// Popup script for settings management

// Load saved settings
document.addEventListener('DOMContentLoaded', () => {
  chrome.storage.sync.get(['apiEndpoint', 'isEnabled', 'refreshRate'], (settings) => {
    document.getElementById('apiEndpoint').value = settings.apiEndpoint || 'http://localhost:8998';
    document.getElementById('isEnabled').checked = settings.isEnabled !== false;
    document.getElementById('refreshRate').value = (settings.refreshRate || 30000) / 1000;
  });
});

// Save settings
document.getElementById('save').addEventListener('click', () => {
  const apiEndpoint = document.getElementById('apiEndpoint').value;
  const isEnabled = document.getElementById('isEnabled').checked;
  const refreshRate = parseInt(document.getElementById('refreshRate').value) * 1000;
  
  // Validate inputs
  if (!apiEndpoint || !apiEndpoint.startsWith('http')) {
    showStatus('Please enter a valid API endpoint URL', 'error');
    return;
  }
  
  if (refreshRate < 0 || refreshRate > 300000) {
    showStatus('Refresh rate must be between 0 and 300 seconds', 'error');
    return;
  }
  
  // Save to storage
  chrome.storage.sync.set({
    apiEndpoint,
    isEnabled,
    refreshRate
  }, () => {
    showStatus('Settings saved! Reload YouTube pages to apply changes.', 'success');
    
    // Notify content scripts to reload settings
    chrome.tabs.query({url: 'https://www.youtube.com/*'}, (tabs) => {
      tabs.forEach(tab => {
        chrome.tabs.sendMessage(tab.id, {action: 'settingsUpdated'}).catch(() => {
          // Ignore errors for tabs without content script
        });
      });
    });
  });
});

function showStatus(message, type) {
  const status = document.getElementById('status');
  status.textContent = message;
  status.className = 'status ' + type;
  
  setTimeout(() => {
    status.className = 'status';
  }, 4000);
}
