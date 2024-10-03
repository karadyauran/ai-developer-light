import { retrieveEmissionData } from './databaseHandler.js';

const emissionFactors = {
  transport: 2.3,
  energy: 0.5,
  food: 1.1
};

export function calculateCarbonFootprint(emissionData) {
  let totalFootprint = 0;
  Object.keys(emissionData).forEach(category => {
    const factor = emissionFactors[category] || 0;
    totalFootprint += emissionData[category] * factor;
  });
  return totalFootprint;
}

export function getAnnualCarbonFootprint() {
  const emissionData = retrieveEmissionData();
  return calculateCarbonFootprint(emissionData);
}

export function suggestReductionStrategies(footprint) {
  const strategies = [];
  if (footprint > 10000) {
    strategies.push('Carpool to reduce transport emissions');
  }
  if (footprint > 5000) {
    strategies.push('Install energy-efficient appliances');
  }
  if (footprint > 2000) {
    strategies.push('Reduce meat consumption');
  }
  return strategies;
}