package main

import (
	_ "nx-go-api/docs"
	"nx-go-api/infrastructure"
)

func main() {
	print("nx starting...")
	server := infrastructure.NewServer()
	server.Run()
	defer server.Close()

}
