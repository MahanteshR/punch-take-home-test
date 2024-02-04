package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type config struct {
	Host     string
	Port     string
	UserName string
	Password string
	Database string
}

func Connect() (*sql.DB, error) {
	LoadEnv()

	config := config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		UserName: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}

	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		config.Host, config.Port, config.Database, config.UserName, config.Password)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = CreateTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

func CreateTables(db *sql.DB) error {
	// create person table
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS` + ` person(
		id uuid PRIMARY KEY,
		name TEXT NOT NULL
	);`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS` + ` person_uidx ON person (name);`)
	if err != nil {
		return err
	}

	// create relation table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS` + ` relation(
		id uuid PRIMARY KEY,
		name text NOT NULL,
		type text NOT NULL
	);`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS` + ` relation_uidx ON relation (name, type);`)
	if err != nil {
		return err
	}

	// create relationship table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS` + ` connection(
		id uuid PRIMARY KEY,
		first_person text NOT NULL,
		relationship text NOT NULL,
		second_person text NOT NULL
	);`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS connection_uidx` +
		` ON connection (first_person, relationship, second_person);`)
	if err != nil {
		return err
	}

	return err

}

func main() {
	// You can use the Connect function to get a database connection
	db, err := Connect()
	defer db.Close()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	fmt.Println("Tables created successfully")

}

func LoadEnv() {
	envPath := "../configs/.env"

	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("No .env file found")
	}
}
