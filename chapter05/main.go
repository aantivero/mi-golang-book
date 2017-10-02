package main

import "fmt"

func main() {
	arreglos()
	promedio()
	promedioRange()
	slicesArreglos()
	slicesAppendYCopy()
	mapas()
	mapas2()
	mapas3()
}

func arreglos() {
	//arrays es un secuencia numerada de elementos de tipos simples de un tamanio fijo
	//un array compuesto de cinco elementos
	var x [5]int
	//empiezan en la posicion cero
	x[4] = 100
	fmt.Println(x)
	fmt.Println(x[4])
}

func promedio() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	//len es una funcion
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	//hay que transformar int to float64
	fmt.Println(total / float64(len(x)))
}
func promedioRange() {
	x := [5]float64{
		98,
		93,
		77,
		82,
		83, //la coma extra es requerida por go
	}
	var total = 0.0
	//range sirve para hacer un loop de la variable
	//usamos el _ para indicar que no lo utilizaremos la variable iterator
	for _, valor := range x {
		total += valor
	}
	fmt.Println(total / float64(len(x)))
}

func slicesArreglos() {
	//son para poder manipular mejor los arrays, son indexables y tienen un tamanio
	//a diferencia de los arrays su tamanio puede variar
	//cuando se define no va su tamanio entre corchetes
	var x []float64
	fmt.Println(x)
	//crear un slice de tamanio 5
	a := make([]float64, 5)
	a[0] = 123
	fmt.Println(a)
	//crear un slice de tamanio 5 y capacidad 10
	b := make([]float64, 5, 10)
	fmt.Println(b)
	//otro forma de crear es usando [low:high]
	arr := [5]float64{1, 2, 3, 4, 5}
	z := arr[1:4]
	fmt.Println(z)
	/*
		se puede omitir el low y high
		arr[0:] = arr[0:len(arr)]
		arr[:5] = arr[0:5]
		arr[:] = arr[0:len(arr)]
	*/
}

func slicesAppendYCopy() {
	//append agrega elementos al final del slice
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println("append ", slice1, slice2)

	//copy recibe 2 argumento destino y origen
	//copia todo el origen en el destino dependiendo del tamanio del destino
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println("copy ", slice3, slice4)
}

func mapas() {
	//es una coleccion desordenada de pares key-value
	//var x map[string]int esto trae un error de ejecucion 'nil map'
	x := make(map[string]int)
	x["llave"] = 10
	fmt.Println(x)
	y := make(map[int]int)
	y[1] = 123
	fmt.Println(y)
}

func mapas2() {
	elementos := make(map[string]string)
	elementos["H"] = "Hydrogen"
	elementos["He"] = "Helium"
	elementos["Li"] = "Litium"
	elementos["Be"] = "Beryllium"
	elementos["B"] = "Boron"
	fmt.Println(elementos["Be"])
	//buscamos un elemento que no existe, imprime vacio, el map devuelve el valor cero que para string es vacio
	fmt.Println(elementos["No"])
	//podemos verificar con go valor cero
	//al acceder a un valor en un map podemos obtener 2 resultados, el valor y si existe
	nombre, existe := elementos["Un"]
	fmt.Println(nombre, existe)
	//lo podemos usar en condicionales
	if algo, existe := elementos["Algo"]; existe {
		//no imprime nada porque no se cumple la condicion
		fmt.Println(algo, existe)
	}
	//forma corta de crear maps como los arrays
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Litium",
		"Be": "Beryllium",
		"B":  "Boron",
	}
	fmt.Println(elements)
}

func mapas3() {
	//podemos crear mapas mas complejos
	elementos := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Litium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
	}
	fmt.Println(elementos)
	if elementoLi, existe := elementos["Li"]; existe {
		fmt.Println(elementoLi, elementoLi["name"], elementoLi["state"])
	}
}
