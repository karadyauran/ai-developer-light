  initialEmissions: 0,
  currentEmissions: 0,
  reduction: 0,
};

export function updateProgress(currentEmissions) {
  if (progress.initialEmissions === 0) {
    progress.initialEmissions = currentEmissions;
  }
  progress.currentEmissions = currentEmissions;
  progress.reduction = progress.initialEmissions - progress.currentEmissions;
}

export function getProgress() {
  return progress;