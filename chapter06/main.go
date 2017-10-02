package main

import "fmt"

/*
funciones
*/

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
	fmt.Println(f1())
	fmt.Println(f3())
	fmt.Println(f4())
	a, b := f4()
	fmt.Println("simple ", a, "doble ", b)
	fmt.Println(add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	//puedo pasar por parametro un slice
	xxs := []int{1, 2, 3, 4, 5}
	fmt.Println(add(xxs...))

	//Closure funciones dentro de funciones
	suma := func(x, y int) int {
		return x + y
	}
	fmt.Println("closure: ", suma(3, 9))
	p := 0
	increment := func() int {
		p++
		return p
	}
	fmt.Println("increment: ", increment())
	fmt.Println("increment: ", increment())
	//una forma de usar closure es escribir funciones que devuelvan funciones
	//que al llamarse puedan generar una secuencia de numeros
	nextEven := makeEvenGenerator()
	fmt.Println("1 closure: ", nextEven())
	fmt.Println("2 closure: ", nextEven())
	fmt.Println("3 closure: ", nextEven())

	//Recursion
	//una funcion puede llamarse a si misma
	fmt.Println("Recursion factorial 18: ", factorial(18))
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

//call stack
func f1() int {
	return f2()
}
func f2() int {
	return 1
}

//return types con nombre
func f3() (r int) {
	r = 100
	return r
}

//multiples valores
func f4() (int, int) {
	r := 10
	x := r * 1
	y := r * 2
	return x, y
}

//parametros variables
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

//Closure: funcion que devuelve funcion
//cada vez que es llamada suma 2 a "i", que a diferencia de una variable local persiste entre llamadas
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

//Recursion, funciones que se llaman a si mismas
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
