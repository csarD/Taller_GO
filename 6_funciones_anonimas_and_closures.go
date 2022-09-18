package main

import "fmt"

var Calculo func(int, int) int = func(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(Calculo(1, 2))

	Calculo = func(x int, y int) int {
		return x - y
	}
	fmt.Println(Calculo(1, 2))

	Operaciones()

	//Closures
	tablaDel := 2
	MiTabla := Tabla(tablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}
}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())
}

//Closures

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	fmt.Println("inicializando")
	return func() int {
		fmt.Println("dentro del for")
		secuencia += 1
		return numero * secuencia
	}
}
