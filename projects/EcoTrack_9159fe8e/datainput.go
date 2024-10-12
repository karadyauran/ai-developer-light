
import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func storeData(data string) error {
	db, err := sql.Open("sqlite3", "./datadb.sqlite")
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS data (id INTEGER PRIMARY KEY AUTOINCREMENT, entry TEXT)")
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO data (entry) VALUES (?)", data)
	return err
}

func initDataDB() {
	db, err := sql.Open("sqlite3", "./datadb.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS data (id INTEGER PRIMARY KEY AUTOINCREMENT, entry TEXT)")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	initDataDB()