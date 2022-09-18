package main

import "fmt"

func main() {
	numero, estado := uno(1)
	fmt.Println(numero)
	fmt.Println(estado)
	fmt.Println(dos(2))
	fmt.Println(Calcular(2, 4, 6, 8, 10))
}

func uno(numero int) (int, bool) {
	if numero == 1 {
		return 0, false
	} else {
		return numero, true
	}

}
func dos(numero int) (z int) {
	z = numero * 2
	return

}

func Calcular(numero ...int) (total int) {
	total = 0
	//el guion bajo es cuando no se necesita el valor
	for _, num := range numero {
		total += num
	}
	return
}

//no existe el override de funciones
