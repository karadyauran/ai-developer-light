import { trackActivity } from './activityTracker.js';
import { calculateEmissions } from './analytics.js';
import { getEcoTips } from './ecoTips.js';
import { loadData, saveData } from './dataStorage.js';

function initApp() {
  initUI();
  const userData = loadData();
  if (userData) {
    updateDashboard(userData);
  }
}

function updateDashboard(data) {
  const emissions = calculateEmissions(data.activities);
  displayEmissions(emissions);
  const tips = getEcoTips(emissions);
  displayTips(tips);
}

function displayEmissions(emissions) {
  const emissionsElement = document.getElementById('emissions');
  emissionsElement.textContent = `Your daily emissions: ${emissions} kg CO2`;
}

function displayTips(tips) {
  const tipsElement = document.getElementById('eco-tips');
  tipsElement.innerHTML = '';
  tips.forEach(tip => {
    const tipItem = document.createElement('li');
    tipItem.textContent = tip;
    tipsElement.appendChild(tipItem);
  });
}

document.getElementById('track-btn').addEventListener('click', () => {
  const activity = document.getElementById('activity-input').value;
  const activities = trackActivity(activity);
  saveData({ activities });
  updateDashboard({ activities });
});
