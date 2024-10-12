
import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func openDatabase(dbName string) *sql.DB {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func initUserDB() {
	db := openDatabase("./userdb.sqlite")
	defer db.Close()
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE,
		password TEXT
	);
	INSERT OR IGNORE INTO users (username, password) VALUES ('testuser', 'password123');
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func initDataDB() {
	db := openDatabase("./datadb.sqlite")
	defer db.Close()
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS data (id INTEGER PRIMARY KEY AUTOINCREMENT, entry TEXT)")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	initUserDB()
	initDataDB()