export const updateUI = (route) => {
  const routeDetails = document.getElementById('routeDetails');
  const routeList = document.createElement('ul');
  routeList.innerHTML = '';

  route.steps.forEach(step => {
    const listItem = document.createElement('li');
    listItem.textContent = `${step.instruction} (Distance: ${step.distance}km)`;
    routeList.appendChild(listItem);
  });

  routeDetails.innerHTML = '';
  routeDetails.appendChild(routeList);
};

export const displayMessage = (message) => {
  const messageBox = document.getElementById('messageBox');
  messageBox.textContent = message;
};

export const clearUI = () => {
  const routeDetails = document.getElementById('routeDetails');
  routeDetails.innerHTML = '';
  const messageBox = document.getElementById('messageBox');
  messageBox.textContent = '';
};