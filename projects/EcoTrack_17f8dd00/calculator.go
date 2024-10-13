package main

func CalculateTotalCarbon(selectedActivities []int, activities []Activity) float64 {
	var totalCarbon float64
	for _, index := range selectedActivities {
		if index >= 0 && index < len(activities) {
			totalCarbon += activities[index].CarbonCost
		}
	}
	return totalCarbon
}