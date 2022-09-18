package main

import "fmt"

type SerVivo interface {
	estaVivo() bool
}
type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}
type animal interface {
	respirar()

	comer()
	EsCarnivoro() bool
}
type vegetal interface {
	ClasificacionVegetal() string
}

type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	esVivo     bool
}
type mujer struct {
	hombre
}

func (h *hombre) respirar()      { h.respirando = true }
func (h *hombre) comer()         { h.comiendo = true }
func (h *hombre) pensar()        { h.pensando = true }
func (h *hombre) estaVivo() bool { return h.esVivo }
func (h *hombre) sexo() string {
	if h.esHombre {
		return "Hombre"
	} else {
		return "Mujer"
	}
}

func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Println(hu.sexo())
}

type perro struct {
	respirando bool

	comiendo  bool
	carnivoro bool
	esVivo    bool
}

func (h *perro) respirar()         { h.respirando = true }
func (h *perro) comer()            { h.comiendo = true }
func (h *perro) EsCarnivoro() bool { return h.carnivoro }
func (h *perro) estaVivo() bool    { return h.esVivo }
func AninamesRespirar(an animal) {
	an.respirar()

}
func AnimalesCarnivoros(an animal) int {
	if an.EsCarnivoro() {
		return 1
	}
	return 0
}
func estoyVivo(v SerVivo) bool { return v.estaVivo() }
func main() {
	fmt.Println("Starting")
	Pedro := new(hombre)
	Pedro.esHombre = true
	Maria := new(mujer)
	Maria.esHombre = false
	HumanosRespirando(Pedro)
	HumanosRespirando(Maria)
	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	Dogo.esVivo = true
	AninamesRespirar(Dogo)
	totalCarnivoros = +AnimalesCarnivoros(Dogo)
	fmt.Println(estoyVivo(Dogo))
	fmt.Println(totalCarnivoros)
}
