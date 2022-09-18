package main

import "fmt"

type Person struct {
	name string
	age  int
}
type Profesion struct {
	Person
	miProfesion string
}

func main() {
	myProf := Profesion{
		Person: Person{
			name: "Oscar",
			age:  27,
		},
		miProfesion: "Ingeniero",
	}
	fmt.Println(myProf)

	personas := make(map[string]Profesion)
	personas[myProf.name] = myProf
	myProf2 := Profesion{
		Person: Person{
			name: "Pablo",
			age:  30,
		},
		miProfesion: "Ingeniero",
	}
	personas[myProf2.name] = myProf2
	fmt.Println(personas)
}
