package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	leoArchivo()
	leoArchivo2()
	graboArchivo()
	graboArchivo2()
}

func leoArchivo() {
	//agregar ruta antes de correr
	archivo, err := ioutil.ReadFile("")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(archivo))
	}
}

func leoArchivo2() {
	//agregar ruta antes de correr
	archivo, err := os.Open("")
	if err != nil {
		fmt.Println(err)
	} else {
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {
			registro := scanner.Text()
			fmt.Println(registro)
		}
	}
	archivo.Close()
}

func graboArchivo() {
	//agregar ruta antes de correr
	archivo, err := os.Create("")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(archivo, "HOLAMUNDO NUEVA")
	archivo.Close()
}
func graboArchivo2() {
	//agregar ruta antes de correr
	fileName := ""
	if Append(fileName, "\n Holaaaaaaaaaaaaaaaaaaa") == false {
		fmt.Println("Error")
	}
}
func Append(fileName string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
