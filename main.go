package main

import (
	"nx-go-api/app/configuration"
)

func main() {
	print("nx starting...")
	server := configuration.NewServer()
	server.Run()
	defer server.Close()

}
