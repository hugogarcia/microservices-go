package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8080"

func main() {
	log.Printf("Starting broker on port %s\n", webPort)

	app := Config{}
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
