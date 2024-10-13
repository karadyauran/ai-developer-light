package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Activity struct {
	Name       string
	CarbonCost float64
}

func main() {
	activities := []Activity{
		{"Car Travel", 2.3},
		{"Public Transport", 0.5},
		{"Biking", 0.0},
		{"Home Electricity", 1.5},
		{"Recycling", -0.2},
	}
	fmt.Println("Welcome to EcoTrack")
	var totalCarbon float64
	for i, activity := range activities {
		fmt.Printf("%d. %s\n", i+1, activity.Name)
	}
	fmt.Println("Enter the number of activities you performed today (e.g. 1,3):")
	var input string
	fmt.Scanln(&input)
	selectedActivities := parseInput(input)
	for _, index := range selectedActivities {
		if index >= 0 && index < len(activities) {
			totalCarbon += activities[index].CarbonCost
		}
	}
	fmt.Printf("Your total carbon footprint today is: %.2f kg CO2\n", totalCarbon)
	generateReport(totalCarbon)
}

func parseInput(input string) []int {
	var indices []int
	for _, char := range input {
		index := char - '1'
		indices = append(indices, int(index))
	}
	return indices
}

func generateReport(totalCarbon float64) {
	rand.Seed(time.Now().UnixNano())
	tips := []string{
		"Consider using public transport.",
		"Try to bike or walk when possible.",
		"Turn off lights when not in use.",
		"Reduce, reuse, and recycle.",
		"Use energy-efficient appliances.",
	}
	fmt.Println("Eco-friendly tips:")
	for i := 0; i < 3; i++ {
		fmt.Println("-", tips[rand.Intn(len(tips))])
	}
}