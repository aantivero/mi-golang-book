package main

import "fmt"

func first() {
	fmt.Println("Call first")
}

func second() {
	fmt.Println("Call second")
}

func main() {
	//defer ejecuta al final de la función la llamada
	//osea se ejecuta una vez que todo haya finalizado, se ejecuta siempre aun ante un panic
	defer second()
	first()
	fmt.Println("----")
}
