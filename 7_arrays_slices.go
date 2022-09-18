package main

import "fmt"

var matriz [5][7]int
var slice = []int{1, 2, 3}

func main() {
	tabla := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	matriz[3][5] = 1
	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}
	fmt.Println(matriz)
	fmt.Println(slice)
	variante2()
	variante3()
	variante4()
}
func variante2() {
	elementos := [5]int{10, 20, 30, 40, 5}
	porcion := elementos[2:]
	fmt.Println("porcion")
	fmt.Println(porcion)
}
func variante3() {
	elementos := make([]int, 5, 20)
	fmt.Println(len(elementos))
	fmt.Println(cap(elementos))
}
func variante4() {
	nums := make([]int, 0, 0)
	fmt.Println("variante 4")
	fmt.Println(len(nums))
	fmt.Println(cap(nums))

	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
}
