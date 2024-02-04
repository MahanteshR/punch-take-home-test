package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/punch-test/pkg/model"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
	connectionService "github.com/punch-test/pkg/service/connection"
	personService "github.com/punch-test/pkg/service/person"
	relationService "github.com/punch-test/pkg/service/relation"
	"github.com/punch-test/pkg/store/connection"
	"github.com/punch-test/pkg/store/person"
	"github.com/punch-test/pkg/store/relation"
)

func main() {
	db, err := initializeDB()
	if err != nil {
		log.Fatal(err)
	}

	personStore := person.New(db)
	relationStore := relation.New(db)
	connectionStore := connection.New(db)

	personSvc := personService.New(personStore)
	relationSvc := relationService.New(relationStore)
	connectionSvc := connectionService.New(connectionStore)

	for {
		reader := bufio.NewReader(os.Stdin)

		input, er := reader.ReadString('\n')
		if er != nil {
			inputError()
		}

		inputList := strings.Split(strings.TrimRight(input, "\n"), " ")
		if inputList[0] != "family-tree" {
			log.Fatalf("Invalid command line input, cmd should start with prefix 'family-tree'")
		}

		switch inputList[1] {
		case "add":
			if len(inputList) != 4 && len(inputList) != 6 {
				log.Fatalf("Invalid command line input for adding a person or relation")
			}

			switch inputList[2] {
			case "person":
				name := inputList[3]

				if e := personSvc.AddPerson(&model.Person{Name: name}); e != nil {
					log.Fatalf("Unable to add person")
				}

			case "relationship":
				name := inputList[3]
				relationship := inputList[5]

				if e := relationSvc.AddRelation(&model.Relationship{
					Name:     name,
					Relation: relationship,
				}); e != nil {
					log.Fatalf("Unable to add relationship for a person")
				}
			default:
				inputError()
			}

		case "connect":
			if len(inputList) != 7 {
				inputError()
			}

			name1, relationship, name2 := inputList[2], inputList[4], inputList[6]
			if e := connectionSvc.AddConnection(&model.Connect{
				FirstPerson:  name1,
				Connection:   relationship,
				SecondPerson: name2,
			}); e != nil {
				log.Fatalf("Unable to add connection")
			}

		case "count":
			if len(inputList) != 5 {
				inputError()
			}

			relationship, name := inputList[2], inputList[4]
			switch relationship {
			case "sons", "daughters", "wives":
				count, err := connectionSvc.CountConnections(relationship, name)
				if err != nil {
					log.Fatalf("Unable to add connection")
				}

				log.Printf("%v has %v %v", name, *count, relationship)
			default:
				inputError()
			}

		case "father":
			if len(inputList) != 4 {
				inputError()
			}

			name := inputList[3]
			if e := connectionSvc.GetConnection(inputList[1], name); e != nil {
				log.Fatalf("Unable to get relation")
			}

		default:
			inputError()
		}
	}
}

func initializeDB() (*sql.DB, error) {
	LoadEnv()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()

		return nil, err
	}

	return db, nil
}

func inputError() {
	log.Fatalf("Invalid command line input")
}

func LoadEnv() {
	envPath := "./configs/.env"

	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("No .env file found")
	}
}
