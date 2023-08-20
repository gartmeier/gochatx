package main

import "github.com/gartmeier/gochatx/pkg/server"

func main() {
	server := server.NewServer()
	server.Run()
}
