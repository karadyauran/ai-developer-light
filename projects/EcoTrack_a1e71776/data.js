  driving: { impact: 0.21 },
  cycling: { impact: 0 },
  walking: { impact: 0 },
  publicTransport: { impact: 0.05 },
  waste: { impact: 2.1 },
  electricity: { impact: 0.5 }
};

export function getActivityImpact(activity) {
  return activities[activity] ? activities[activity].impact : 0;