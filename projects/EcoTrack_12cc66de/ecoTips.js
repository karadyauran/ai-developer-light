  'Consider carpooling or using public transport',
  'Reduce energy consumption at home',
  'Switch to renewable energy sources',
  'Use energy-efficient appliances',
  'Reduce, reuse, and recycle materials'
];

export function getEcoTips(emissions) {
  if (emissions > 5) {
    return tips;
  }
  return ['Great job! Keep maintaining your low emissions.'];