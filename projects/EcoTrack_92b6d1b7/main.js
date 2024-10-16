import { calculateEmissions } from './emissionsCalculator.js';
import { getTips } from './tips.js';
import { updateProgress } from './progressTracker.js';
import { renderUI } from './ui.js';

function initApp() {
  const userData = getUserData();
  renderUI(userData, handleUserAction);
}

function handleUserAction(action, data) {
  switch (action) {
    case 'calculate':
      const emissions = calculateEmissions(data);
      setUserData({ ...data, emissions });
      updateProgress(emissions);
      const tips = getTips(emissions);
      renderUI(getUserData(), handleUserAction, tips);
      break;
    case 'updateProfile':
      setUserData(data);
      renderUI(getUserData(), handleUserAction);
      break;
    default:
      break;
  }
}
