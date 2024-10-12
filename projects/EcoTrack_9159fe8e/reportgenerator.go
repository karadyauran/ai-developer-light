
import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
)

func generateReport() string {
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
	var entries []string
	for rows.Next() {
		var entry string
		err = rows.Scan(&entry)
		if err != nil {
			log.Fatal(err)
		}
		entries = append(entries, entry)
	}
	return strings.Join(entries, ", ")
}

func main() {
	report := generateReport()
	log.Println("Report: ", report)