package figuras

import "fmt"

type FigurasGeo interface {
	area() float64
	perimetro() float64
}

func Calcular(f FigurasGeo) {
	fmt.Println("Area: ", f.area())
	fmt.Println("Permimetro: ", f.perimetro())
}
