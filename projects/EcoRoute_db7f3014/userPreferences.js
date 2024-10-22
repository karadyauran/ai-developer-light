export const getUserPreferences = () => {
  const transportMode = document.querySelector('input[name="transportMode"]:checked').value;
  return { transportMode };
};

export const saveUserPreferences = (preferences) => {
  localStorage.setItem('userPreferences', JSON.stringify(preferences));
};

export const loadUserPreferences = () => {
  const savedPreferences = localStorage.getItem('userPreferences');
  return savedPreferences ? JSON.parse(savedPreferences) : { transportMode: 'driving' };
};

export const applyUserPreferences = () => {
  const preferences = loadUserPreferences();
  const transportModeInput = document.querySelector(`input[name="transportMode"][value="${preferences.transportMode}"]`);
  if (transportModeInput) {
    transportModeInput.checked = true;
  }
};