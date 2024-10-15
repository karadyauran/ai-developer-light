
export function trackActivity(activity) {
  if (activity) {
    activities.push(activity);
  }
  return activities;