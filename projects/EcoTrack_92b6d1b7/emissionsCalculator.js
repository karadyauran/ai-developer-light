  car: 2.31,
  bus: 0.89,
  bike: 0,
  walk: 0,
  electricity: 0.92,
};

export function calculateEmissions(activities) {
  return activities.reduce((total, activity) => {
    const factor = emissionFactors[activity.type] || 0;
    return total + activity.amount * factor;
  }, 0);