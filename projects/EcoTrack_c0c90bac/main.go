
import (
	"fmt"
	"os"
)

func main() {
	user := getUserProfile()
	data := getEnvironmentalData()
	carbonFootprint := calculateCarbonFootprint(data)
	generateReport(user, carbonFootprint)
}

func getUserProfile() User {
	return User{Name: "Alice", Email: "alice@example.com"}
}

func getEnvironmentalData() EnvironmentalData {
	return EnvironmentalData{ElectricityUsage: 350, WaterUsage: 100, GasUsage: 50}
}

func calculateCarbonFootprint(data EnvironmentalData) float64 {
	return (data.ElectricityUsage*0.5 + data.WaterUsage*0.1 + data.GasUsage*0.3) * 0.001
}

func generateReport(user User, footprint float64) {
	report := fmt.Sprintf("User: %s\nCarbon Footprint: %.2f tons CO2", user.Name, footprint)
	fmt.Println(report)
	saveReport(report)
}

func saveReport(report string) {
	file, err := os.Create("report.txt")
	if err != nil {
		fmt.Println("Error creating report:", err)
		return
	}
	defer file.Close()
	file.WriteString(report)
}

type User struct {
	Name  string
	Email string
}

type EnvironmentalData struct {
	ElectricityUsage float64
	WaterUsage       float64
	GasUsage         float64