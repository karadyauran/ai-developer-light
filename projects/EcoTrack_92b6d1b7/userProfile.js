  name: '',
  activities: [],
  emissions: 0,
};

export function getUserData() {
  return userData;
}

export function setUserData(data) {
  userData = { ...userData, ...data };
}

export function addActivity(activity) {
  userData.activities.push(activity);
  setUserData(userData);
}

export function clearActivities() {
  userData.activities = [];
  setUserData(userData);