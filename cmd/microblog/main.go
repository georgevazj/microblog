package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/georgevazj/microblog/internal/data"
	"github.com/georgevazj/microblog/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("PORT")
	serv, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}

	// connection to the database.
	// d := data.New()
	// if err := d.DB.Ping(); err != nil {
	//	log.Fatal(err)
	// }

	// Start the server
	go serv.Start()

	// Wait for an interrupt signal to gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attempt to close the server gracefully
	serv.Close()
	data.Close()
}
