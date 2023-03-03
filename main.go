package main

import (
	"github.com/gabhmm/rest-api.git/database"
	"github.com/gabhmm/rest-api.git/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
