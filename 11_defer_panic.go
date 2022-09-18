package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	archivo := "prueba.txt"
	f, err := os.Open(archivo)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		//os.Exit(1)
	}
	ejemploPanic()
}
func ejemploPanic() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("ERROR con PANIC: %v", r)
		}
	}()
	a := 1
	if a == 1 {
		panic("se encontro el valor de 1")
	}
}
