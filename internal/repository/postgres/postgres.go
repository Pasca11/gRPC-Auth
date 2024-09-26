package postgres

import (
	"fmt"
	"github.com/Pasca11/gRPC-Auth/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

type Database struct {
	DB *sqlx.DB
}

func NewDatabase() (*Database, error) {
	dataBase := &Database{}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	connStr := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " host=" + dbHost + " port=" + dbPort + " sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	dataBase.DB = db
	err = dataBase.CreateNewUserTable()
	if err != nil {
		return nil, err
	}
	err = dataBase.CreateNewNoteTable()
	if err != nil {
		return nil, err
	}
	return dataBase, err
}

func (db *Database) CreateNewUserTable() error {
	newTableString := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(100) NOT NULL,
		role VARCHAR(20) NOT NULL
	);`

	_, err := db.DB.Exec(newTableString)
	return err
}

func (db *Database) CreateNewNoteTable() error {
	newTableString := `
		CREATE TABLE IF NOT EXISTS notes (
 	    id SERIAL PRIMARY KEY,
 	    text VARCHAR NOT NULL,
 		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
 		deadline TIMESTAMP,
 		user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE
	);`
	_, err := db.DB.Exec(newTableString)
	return err
}

func (db *Database) CreateUser(user *models.User) error {
	stmt := `INSERT INTO users (username, password, role) VALUES ($1, $2, $3)`
	_, err := db.DB.Exec(stmt, user.Username, user.Password, user.Role)
	if err != nil {
		return fmt.Errorf("repo.insert.user: %w", err)
	}
	return nil
}

func (db *Database) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := db.DB.Get(user, "SELECT * FROM users WHERE username=$1", username)
	if err != nil {
		return nil, fmt.Errorf("repo.get.user: %w", err)
	}
	return user, nil
}

func (db *Database) GetUserById(id int) (*models.User, error) {
	user := &models.User{}
	err := db.DB.Get(user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("repo.get.user: %w", err)
	}
	return user, nil
}
