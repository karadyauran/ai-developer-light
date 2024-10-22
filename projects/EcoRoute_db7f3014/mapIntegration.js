export const displayMap = (mapElement) => {
  const map = new google.maps.Map(mapElement, {
    center: { lat: 40.7128, lng: -74.0060 },
    zoom: 12
  });

  return map;
};

export const plotRouteOnMap = (map, route) => {
  const pathCoordinates = route.steps.map(step => ({
    lat: step.location.lat,
    lng: step.location.lng
  }));

  const routePath = new google.maps.Polyline({
    path: pathCoordinates,
    geodesic: true,
    strokeColor: '#FF0000',
    strokeOpacity: 1.0,
    strokeWeight: 2
  });

  routePath.setMap(map);
};

export const clearMap = (map) => {
  for (let i = 0; i < mapOverlays.length; i++) {
    mapOverlays[i].setMap(null);
  }
  mapOverlays = [];
};

let mapOverlays = [];