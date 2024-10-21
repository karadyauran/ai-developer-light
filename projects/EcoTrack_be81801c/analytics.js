
export function generateInsights() {
    const data = fetchData();
    const totalActivities = data.activities.length;
    const averageImpact = totalActivities > 0 ? data.totalImpact / totalActivities : 0;
    return {
        totalActivities,
        averageImpact
    };