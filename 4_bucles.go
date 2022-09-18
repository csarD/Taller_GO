package main

import (
	"fmt"
)

func main() {

	//go_to()

	//continuar()

	//normal_for()
	//undefined_for()

}
func continuar() {
	var j = 0
	for j < 11 {
		fmt.Println("El valor de j:", j)
		if j == 5 {
			fmt.Println("Multiplicando por 2")
			j = j * 2
			continue
		}
		fmt.Println("Paso por aqui")
		j++
	}
}
func normal_for() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
func undefined_for() {
	numero := 0
	for {
		fmt.Println("Continuo")
		fmt.Println("Ingrese el numero secreto")
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}
	}
}
func go_to() {
	var k int = 0
RUTINA:
	fmt.Println("Empezando el for")
	for k < 10 {
		if k == 4 {
			k = k + 2
			fmt.Println("voy a rutina sumando 2 a k")
			goto RUTINA
		}
		fmt.Println("Valor de i:", k)
		k++

	}
}
