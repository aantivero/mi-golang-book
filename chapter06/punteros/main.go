package main

import "fmt"

func main() {
	x := 5
	cero(x)
	fmt.Println(x) //sigue valiendo 5
	//si queremos modificar x tenemos que utilizar punteros
	//& para ubicar la direccion de la variable la misma direccion de memoria
	zero(&x)
	fmt.Println(x) //ahora si vale 0

	//otra forma es utilizar el new function
	//new recibe el tipo como argumento y reserva memoria para ella y devuelve un puntero a este
	y := new(int)
	uno(y)
	fmt.Println(*y)
}

func cero(u int) {
	u = 0
}

//* representa un puntero, u es un puntero a int
func zero(u *int) {
	*u = 0
}

func uno(x *int) {
	*x = 1
}
