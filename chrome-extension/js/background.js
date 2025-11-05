// Background Service Worker for YouTube Stats Overlay

console.log('YouTube Stats Overlay - Background service worker loaded');

// Set default settings on install
chrome.runtime.onInstalled.addListener(() => {
  chrome.storage.sync.set({
    apiEndpoint: 'http://localhost:8998',
    isEnabled: true,
    refreshRate: 30000
  }, () => {
    console.log('Default settings initialized');
  });
});

// Handle messages from content script (if needed for future features)
chrome.runtime.onMessage.addListener((request, sender, sendResponse) => {
  if (request.action === 'getSettings') {
    chrome.storage.sync.get(['apiEndpoint', 'isEnabled', 'refreshRate'], (settings) => {
      sendResponse(settings);
    });
    return true; // Keep channel open for async response
  }
});
