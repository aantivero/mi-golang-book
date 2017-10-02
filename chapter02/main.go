package main

import "fmt"

func main() {
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("1.0 + 1.0 = ", 1.0+1.0)
	fmt.Println(len("Hola, mundo"))
	//imprime el valor entero, el caracter es representado por un byte y el byte es un entero
	fmt.Println("Hola, mundo"[1])
	fmt.Println("Hola, " + "mundo") //concatenación
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
