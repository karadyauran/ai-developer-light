import { updateEmissionData, retrieveEmissionData } from './databaseHandler.js';

export function trackEmissions(userData) {
  const emissionData = retrieveEmissionData();
  const transportEmission = calculateTransportEmission(userData.transportMode, userData.transportDistance);
  const energyEmission = calculateEnergyEmission(userData.energyUsage);
  const foodEmission = calculateFoodEmission(userData.foodType, userData.foodQuantity);
  updateEmissionData('transport', transportEmission);
  updateEmissionData('energy', energyEmission);
  updateEmissionData('food', foodEmission);
  return emissionData;
}

function calculateTransportEmission(mode, distance) {
  const factors = {
    car: 0.21,
    bus: 0.11,
    bike: 0.01
  };
  return distance * (factors[mode] || 0);
}

function calculateEnergyEmission(usage) {
  const factor = 0.3;
  return usage * factor;
}

function calculateFoodEmission(type, quantity) {
  const factors = {
    meat: 5,
    vegetarian: 2,
    vegan: 1
  };
  return quantity * (factors[type] || 0);
}