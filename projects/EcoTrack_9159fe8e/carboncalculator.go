
import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
)

func calculateCarbonFootprint() float64 {
	db, err := sql.Open("sqlite3", "./datadb.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT entry FROM data")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var total float64
	for rows.Next() {
		var entry string
		err = rows.Scan(&entry)
		if err != nil {
			log.Fatal(err)
		}
		value, err := strconv.ParseFloat(entry, 64)
		if err == nil {
			total += value
		}
	}
	return total * 0.5
}

func main() {
	total := calculateCarbonFootprint()
	log.Printf("Total Carbon Footprint: %f", total)