package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kushagra/olx-api/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// migration means :- jo maine sql file bnayi hai usse apne databse ke upr apply krna.

func main() {
	if len(os.Args) < 2 {
		log.Fatal("please provide migration name <up | down>")
		return
	}

	cnf := config.MustLoad()

	// database se connect krna hai
	m, err := migrate.New(
		"file://migrations", // migration folder ka path
		cnf.DatabaseUrl,
	)

	if err != nil {
		log.Fatalf("error while creating migration obj: %v", err)
	}

	// switch case for up and down
	switch os.Args[1] {
	case "up":
		fmt.Println("Running migration....!")
		err = m.Up()
		if err != nil {
			log.Fatalf("error while running migration: %v", err)
		}
	case "down":
		fmt.Println("undoing migration....!")
		err = m.Steps(-1) // this steps used for down migration to previous one.
		if err != nil {
			log.Fatalf("error while running migration: %v", err)
		}
	default:
		log.Fatal("please provide migration name <up | down>")
		return
	}

}
