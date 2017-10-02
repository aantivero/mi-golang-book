package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

//de esta forma siempre se accede a una copia
func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

//si quiero modificar uno de sus propiedades debe acceder por puntero
func circleAreaP(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

//metodos
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

//Circle estructura circulo
type Circle struct {
	x, y, r float64
}

//Rectangle estructura para rectangulo
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	//var cx, cy, cr float64 = 0, 0, 5
	//se puede crear una instancia de diversas maneras
	//var c Circle se crea todo en cero 0 int, 0.0 floats, "" string, nil para pointers
	//c := new(Circle) aloja memoria para todas las propiedades y devuelve un puntero
	//la forma mas usual es que los parametros tenga valores iniciales
	c := Circle{x: 0, y: 0, r: 5}
	//c := Circle{0, 0, 5}
	//c := &Circle{0, 0, 5} igual al anterior con un puntero

	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	//fmt.Println(circleArea(cx, cy, cr))
	fmt.Println("acceder ", c.x, c.y)
	fmt.Println("Area circulo ", circleArea(c))
	//lamar por puntero
	fmt.Println("puntero ", circleAreaP(&c))
	//los metodos son un tipo especial de funciones
	fmt.Println("metodo area ", c.area())

	//ahora para el rectangulo
	r := Rectangle{0, 0, 10, 10}
	fmt.Println("Area rectangulo ", r.area())

	a := new(Android)
	a.Person.Talk()
	a.Talk()

}

//Person tiene un nombre
type Person struct {
	Name string
}

//Talk es una funcion
func (p *Person) Talk() {
	fmt.Println("Hi, my name is ", p.Name)
}

/* de esta forma la relacion "es" en lugar de "tiene"
type Androidp struct {
	Person Person
	Model  string
}*/

//Android tiene nombre y modelo
type Android struct {
	Person
	Model string
}
