import { displayMap } from './mapIntegration.js';
import { getOptimizedRoute } from './routeOptimization.js';
import { getUserPreferences } from './userPreferences.js';
import { updateUI } from './ui.js';

const startApp = () => {
  const mapElement = document.getElementById('map');
  const startLocation = document.getElementById('startLocation');
  const endLocation = document.getElementById('endLocation');
  const routeButton = document.getElementById('routeButton');

  displayMap(mapElement);

  routeButton.addEventListener('click', async () => {
    const start = startLocation.value;
    const end = endLocation.value;
    const preferences = getUserPreferences();

    if (start && end) {
      try {
        const optimizedRoute = await getOptimizedRoute(start, end, preferences);
        updateUI(optimizedRoute);
      } catch (error) {
        console.error('Error finding route:', error);
      }
    }
  });
};

document.addEventListener('DOMContentLoaded', startApp);