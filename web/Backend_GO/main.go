package main

import (
	"github.com/csarD/Backend_GO/bd"
	"github.com/csarD/Backend_GO/handlers"
	"log"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Connection ERROR")
		return
	}
	handlers.Manejadores()
}
