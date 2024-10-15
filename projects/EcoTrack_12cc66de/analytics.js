  driving: 2.31,
  publicTransport: 0.1,
  cycling: 0,
  walking: 0
};

export function calculateEmissions(activities) {
  return activities.reduce((total, activity) => {
    return total + (emissionFactors[activity] || 0);
  }, 0);