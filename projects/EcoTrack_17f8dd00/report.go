package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateReport(totalCarbon float64) {
	rand.Seed(time.Now().UnixNano())
	tips := []string{
		"Consider using public transport.",
		"Try to bike or walk when possible.",
		"Turn off lights when not in use.",
		"Reduce, reuse, and recycle.",
		"Use energy-efficient appliances.",
	}
	fmt.Printf("Your total carbon footprint today is: %.2f kg CO2\n", totalCarbon)
	fmt.Println("Eco-friendly tips:")
	for i := 0; i < 3; i++ {
		fmt.Println("-", tips[rand.Intn(len(tips))])
	}
}