package main

import "fmt"

var x string = "Hola mundo"

func main() {
	var y string
	y = "Esto es otra variable"
	fmt.Println(x)
	fmt.Println(y)
	var z = `Esto es una 
	cadena 
	de caracteres `
	fmt.Println(z)
	z += "mas caracteres extras"
	fmt.Println(z)
	var a = "hola"
	var b = "mundo"
	fmt.Println(a == b)
	b = "hola"
	fmt.Println(a == b)
	w := "Saludo"
	fmt.Println(w)
	f()
	const pi float64 = 3.142847
	fmt.Println(pi)
	var (
		nombre = "Ale"
		edad   = 39
		peso   = "100Kg"
	)
	fmt.Println("Persona ", nombre, edad, peso)
	calcularElDoble()
}

func f() {
	fmt.Println(x)
}

func calcularElDoble() {
	fmt.Println("Ingrese un número:")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}
