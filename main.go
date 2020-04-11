package main

import (
	application "fiberapp/application"
	"fiberapp/config"
	"fiberapp/database"
	"log"
)

func init() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn := &database.Connection{}
	if err := conn.Connect(); err != nil {
		log.Fatal(err)
	}
	database.SetDbConnection(conn)
	app := application.Initialize()
	log.Fatal(app.Start())
}
