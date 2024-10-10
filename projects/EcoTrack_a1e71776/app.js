import { calculateImpact } from './calculator.js';
import { generateReport } from './report.js';
import { activities } from './data.js';

const activityForm = document.getElementById('activity-form');
const reportSection = document.getElementById('report-section');

async function init() {
  const latestData = await fetchData();
  if (latestData) {
    updateActivities(latestData);
  }
}

function updateActivities(data) {
  Object.keys(activities).forEach(activity => {
    if (data[activity]) {
      activities[activity].impact = data[activity].impact;
    }
  });
}

function handleFormSubmit(event) {
  event.preventDefault();
  const selectedActivity = document.getElementById('activity').value;
  const duration = parseFloat(document.getElementById('duration').value);
  const impact = calculateImpact(selectedActivity, duration);
  displayReport(selectedActivity, duration, impact);
}

function displayReport(activity, duration, impact) {
  const report = generateReport(activity, duration, impact);
  reportSection.innerHTML = report;
}

activityForm.addEventListener('submit', handleFormSubmit);