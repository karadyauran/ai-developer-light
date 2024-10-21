import { calculateImpact } from './utils.js';

export function initApp() {
    const data = fetchData();
    displayFootprint(data.totalImpact);
}

export function updateFootprint() {
    const data = fetchData();
    data.totalImpact = data.activities.reduce((sum, activity) => sum + activity.impact, 0);
    saveData(data);
    displayFootprint(data.totalImpact);
}

function displayFootprint(totalImpact) {
    const footprintDisplay = document.getElementById('footprint-display');
    footprintDisplay.textContent = `Total Carbon Footprint: ${totalImpact.toFixed(2)} kg CO2`;