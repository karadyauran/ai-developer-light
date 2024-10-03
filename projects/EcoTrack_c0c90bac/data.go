
import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func loadEnvironmentalData(filename string) EnvironmentalData {
	dataBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read data file: %v", err)
	}
	var data EnvironmentalData
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		log.Fatalf("Failed to parse data: %v", err)
	}
	return data
}

func saveEnvironmentalData(filename string, data EnvironmentalData) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to encode data: %v", err)
	}
	if err := ioutil.WriteFile(filename, dataBytes, 0644); err != nil {
		log.Fatalf("Failed to write data file: %v", err)
	}
}

func calculateAverageUsage(data []EnvironmentalData) EnvironmentalData {
	var total ElectricityUsage, totalWaterUsage, totalGasUsage float64
	for _, d := range data {
		totalElectricityUsage += d.ElectricityUsage
		totalWaterUsage += d.WaterUsage
		totalGasUsage += d.GasUsage
	}
	count := float64(len(data))
	return EnvironmentalData{
		ElectricityUsage: totalElectricityUsage / count,
		WaterUsage:       totalWaterUsage / count,
		GasUsage:         totalGasUsage / count,
	}