package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hugogarcia/microservices/authentication-service/cmd/data"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "8181"

var counts int64

type Config struct {
	Repo   data.Repository
	Client *http.Client
}

func main() {
	log.Printf("Starting authentication service")

	conn, err := connectToDB()
	if err != nil {
		log.Panic(err)
	}

	app := Config{
		Client: &http.Client{},
	}
	app.setupRepo(conn)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func (app *Config) setupRepo(conn *sql.DB) {
	app.Repo = data.NewPostgresRepository(conn)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func connectToDB() (*sql.DB, error) {
	dsn := os.Getenv("DB_DSN")
	for counts < 10 {
		db, err := openDB(dsn)
		if err != nil {
			counts++
			log.Println("postgres not yet ready...")
			log.Println("backing off for 2 seconds")
			time.Sleep(time.Second * 2)
			continue
		}

		log.Println("connected to postres!")

		return db, nil
	}

	return nil, fmt.Errorf("could not connect to postgres")
}
