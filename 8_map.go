package main

import "fmt"

func main() {
	paises := make(map[string]string)
	fmt.Println(paises)
	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	fmt.Println(paises["Mexico"])

	campeonato := map[string]int{
		"Real Madrid": 9,
		"Barcelona":   10,
		"Mineiro":     9,
		"Boca":        8,
	}

	campeonato["River"] = 11
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)
	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo: %s con un puntaje de: %d\n", equipo, puntaje)
	}
	puntaje, existe := campeonato["Barcelona"]
	fmt.Println(puntaje, existe)
}
