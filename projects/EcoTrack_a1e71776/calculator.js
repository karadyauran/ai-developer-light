
export function calculateImpact(activity, duration) {
  const impactPerUnit = getActivityImpact(activity);
  return impactPerUnit * duration;