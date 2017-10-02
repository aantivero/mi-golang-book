package main

import "fmt"

/*
func main() {
	panic("PANIC...") //causa un runtime error
	//podemos usar el recover para manejar el panic
	//recover para el panic y devuelve el valor pasado al panic
	str := recover() //de esta forma nunca se llama
	fmt.Println(str)
}
*/
func main() {
	//la forma correcta de implementar panic y recover con defer
	defer func() {
		str := recover()
		fmt.Println("-----> ", str)
	}()
	panic("PANIC")
}
