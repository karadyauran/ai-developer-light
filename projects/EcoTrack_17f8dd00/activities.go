package main

type Activity struct {
	Name       string
	CarbonCost float64
}

func GetActivities() []Activity {
	return []Activity{
		{"Car Travel", 2.3},
		{"Public Transport", 0.5},
		{"Biking", 0.0},
		{"Home Electricity", 1.5},
		{"Recycling", -0.2},
	}
}