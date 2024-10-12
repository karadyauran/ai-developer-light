
import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func authenticateUser(username, password string) bool {
	db, err := sql.Open("sqlite3", "./userdb.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var storedPassword string
	err = db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
	if err != nil {
		return false
	}
	return storedPassword == password
}

func initDatabase() {
	db, err := sql.Open("sqlite3", "./userdb.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec(`
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

func main() {
	initDatabase()