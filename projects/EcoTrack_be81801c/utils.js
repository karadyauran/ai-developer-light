    'driving': 2.31,
    'flying': 0.15,
    'cycling': 0.02,
    'walking': 0.01,
    'public_transport': 0.1
};

export function calculateImpact(activity) {
    return activityImpactMapping[activity] || 0;