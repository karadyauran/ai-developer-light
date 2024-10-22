import { fetchRouteData } from './dataService.js';

export const getOptimizedRoute = async (start, end, preferences) => {
  const routeData = await fetchRouteData(start, end);
  return optimizeRoute(routeData, preferences);
};

const optimizeRoute = (routeData, preferences) => {
  const filteredRoutes = routeData.routes.filter(route => {
    if (preferences.transportMode) {
      return route.transportMode === preferences.transportMode;
    }
    return true;
  });

  filteredRoutes.sort((a, b) => a.emissions - b.emissions);
  return filteredRoutes[0];
};