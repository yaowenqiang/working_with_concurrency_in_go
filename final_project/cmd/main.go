package main

import (
	"database/sql"
	"log"
	"os"
	"time"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"

)

const webPort = "80"


func main() {
	db := initDB()
	db.Ping()
	// connect to the database
	// create sessions
	// crate channels
	// create waitinggroup
	// set up the application config
	// set up mail
	// listen for web connections

}


func initDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to database")
	}
}

func connectToDB() *sql.DB {
	counts := 0
	dsn := os.Getenv('DSN')
		for {
			connection, err := openDB(dsn)
			if err != nil {
				log.Println("postgres not yet read...")

			} else {
				log.Println("connected to database.")
				return connection
			}


			if counts > 10 {
				return nil
			} else {
				log.Println("Backing off for 1 second")
				time.Sleep(1 * time.Second)
				continue
			}
		}

}

func openDB(dsn) (*sql.DB, error) {
	db, err := sql.OPen("pgx",dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil

}
