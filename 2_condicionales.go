package main

import "fmt"

var estado bool

func main() {
	//estado = true
	if !estado {
		fmt.Println("Hello World")
	} else {
		fmt.Println(estado)
	}
	numero := 1

	switch numero {
	case 1:

		fmt.Println(1)

	case 2:
		fmt.Println(2)

	case 3:
		fmt.Println(3)

	case 4:
		fmt.Println(4)

	default:
		fmt.Println("Mayor a 4")

	}

}
