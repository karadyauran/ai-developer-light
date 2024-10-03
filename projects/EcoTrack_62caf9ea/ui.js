import { getUserData, saveUserData } from './databaseHandler.js';
import { calculateCarbonFootprint } from './carbonCalculator.js';
import { trackEmissions } from './emissionTracker.js';

const userForm = document.getElementById('userForm');
const resultDisplay = document.getElementById('resultDisplay');
const tipsDisplay = document.getElementById('tipsDisplay');

userForm.addEventListener('submit', (event) => {
  event.preventDefault();
  const formData = new FormData(userForm);
  const userData = Object.fromEntries(formData.entries());
  saveUserData(userData);
  const emissions = trackEmissions(userData);
  const footprint = calculateCarbonFootprint(emissions);
  displayResults(footprint);
});

function displayResults(footprint) {
  resultDisplay.innerHTML = `Your carbon footprint is ${footprint} kg CO2 per year.`;
  const tips = getSustainabilityTips(footprint);
  tipsDisplay.innerHTML = `Tips: ${tips.join(', ')}`;
}

function getSustainabilityTips(footprint) {
  const tips = [];
  if (footprint > 10000) {
    tips.push('Consider using public transport');
  }
  if (footprint > 5000) {
    tips.push('Reduce energy consumption');
  }
  if (footprint > 2000) {
    tips.push('Adopt a plant-based diet');
  }
  return tips;
}

window.addEventListener('load', () => {
  const savedUserData = getUserData();
  if (savedUserData) {
    Object.keys(savedUserData).forEach(key => {
      const input = document.querySelector(`[name="${key}"]`);
      if (input) {
        input.value = savedUserData[key];
      }
    });
  }
});