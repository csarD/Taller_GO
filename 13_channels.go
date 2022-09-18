package main

import (
	"fmt"
	"time"
)

func main() {
	canal1 := make(chan bool)
	canal2 := make(chan bool)

	go bucle(canal1)

	go bucle2(canal2)
	fmt.Println("llegue hasta aqui")

Validar:

	msg1 := <-canal1
	msg2 := <-canal2
	if msg1 && msg2 {
		fmt.Println("fin de ambos hilos")
	} else {
		goto Validar
	}

}

func bucle(canal1 chan bool) {

	fmt.Println("Entrando al hilo 1")
	for i := 0; i < 1000000; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(i)
		fmt.Println("En hilo 1")
		if i == 10 {
			break
		}
	}

	canal1 <- true

}
func bucle2(canal2 chan bool) {

	fmt.Println("Entrando al hilo 2")
	for i := 0; i < 20; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(i)
		fmt.Println("En hilo 2")

	}

	canal2 <- true
}
