package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1 int
var numero2 int
var resultado int
var leyenda string

func main() {
	fmt.Println("Ingrese numero 1:")

	//fmt.Scanf("%d", &numero1)
	fmt.Scanln(&numero1)
	fmt.Println("Ingrese numero 2:")
	//fmt.Scanf("%d", &numero2)
	fmt.Scanln(&numero2)
	fmt.Println("Descripcion")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		leyenda = scanner.Text()

	}

	//fmt.Println("El resultado")
	resultado = numero1 + numero2
	fmt.Println(leyenda, resultado)

}
