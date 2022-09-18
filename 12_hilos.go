package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go Nombre("Gabriel Lopez")
	go Nombre("Oscar Perez")
	fmt.Println("CESAR De la O")

}
func Nombre(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(letra)
	}
}
