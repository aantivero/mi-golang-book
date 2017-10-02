package main

import (
	"fmt"
)

/**
Estructuras de control
*/
func main() {
	cicloFor()
	cicloForYCondicional()
	ahoraConSwitch(9)
}

func cicloFor() {
	/*i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}*/
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func cicloForYCondicional() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " es par")
		} else {
			fmt.Println(i, " es impar")
		}
	}
}

func ahoraConSwitch(i int) {
	switch i {
	case 0:
		fmt.Println("Cero")
	case 1:
		fmt.Println("Uno")
	case 2:
		fmt.Println("Dos")
	case 3:
		fmt.Println("Tres")
	case 4:
		fmt.Println("Cuatro")
	case 5:
		fmt.Println("Cinco")
	case 6:
		fmt.Println("Seis")
	case 7:
		fmt.Println("Siete")
	case 8:
		fmt.Println("Ocho")
	case 9:
		fmt.Println("Nueve")
	default:
		fmt.Println("Desconocido")
	}
}
