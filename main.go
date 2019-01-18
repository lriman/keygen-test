package main

import (
	"log"
	"github.com/keygen-test/generator"
)

func main() {
	a := new(App)

	a.Initialize("127.0.0.1", "testuser", "mypass", "mydb")
	log.Println("Database success connection")

	go generator.KeyGen(a.DB)

	log.Println("Service is open on 8001 port")
	a.Run(":8001")
}

