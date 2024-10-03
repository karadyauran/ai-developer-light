const userDataKey = 'ecoTrackUserData';

export function getUserData() {
  const data = localStorage.getItem(userDataKey);
  return data ? JSON.parse(data) : null;
}

export function saveUserData(data) {
  localStorage.setItem(userDataKey, JSON.stringify(data));
}

export function retrieveEmissionData() {
  const data = localStorage.getItem('emissionData');
  return data ? JSON.parse(data) : initializeEmissionData();
}

function initializeEmissionData() {
  const initialData = {
    transport: 0,
    energy: 0,
    food: 0
  };
  localStorage.setItem('emissionData', JSON.stringify(initialData));
  return initialData;
}

export function saveEmissionData(emissionData) {
  localStorage.setItem('emissionData', JSON.stringify(emissionData));
}

export function updateEmissionData(category, value) {
  const emissionData = retrieveEmissionData();
  if (emissionData[category] !== undefined) {
    emissionData[category] += value;
    saveEmissionData(emissionData);
  }
}

export function resetUserData() {
  localStorage.removeItem(userDataKey);
  initializeEmissionData();
}