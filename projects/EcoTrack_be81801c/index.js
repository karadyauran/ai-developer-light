import { fetchData, saveData } from './data.js';
import { calculateImpact } from './utils.js';
import { generateInsights } from './analytics.js';

document.addEventListener('DOMContentLoaded', () => {
    initApp();

    const activityForm = document.getElementById('activity-form');
    const activityInput = document.getElementById('activity-input');
    const resultDisplay = document.getElementById('result-display');

    activityForm.addEventListener('submit', (event) => {
        event.preventDefault();
        const activity = activityInput.value.trim();
        if (activity) {
            const data = fetchData();
            const impact = calculateImpact(activity);
            data.activities.push({ activity, impact });
            saveData(data);
            updateFootprint();
            displayResult(impact);
        }
    });

    function displayResult(impact) {
        const insights = generateInsights();
        resultDisplay.innerHTML = `
            <p>Carbon Impact: ${impact} kg CO2</p>
            <p>Total Activities: ${insights.totalActivities}</p>
            <p>Average Impact: ${insights.averageImpact.toFixed(2)} kg CO2</p>
        `;
        activityInput.value = '';
    }