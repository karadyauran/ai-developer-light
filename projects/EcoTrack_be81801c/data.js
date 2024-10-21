    activities: [],
    totalImpact: 0
};

export function fetchData() {
    const data = localStorage.getItem('ecoTrackData');
    return data ? JSON.parse(data) : defaultData;
}

export function saveData(data) {
    localStorage.setItem('ecoTrackData', JSON.stringify(data));