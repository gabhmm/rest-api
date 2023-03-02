package main

import "github.com/gabhmm/rest-api.git/server"

func main() {
	server := server.NewServer()

	server.Run()
}
