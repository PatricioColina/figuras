package figuras

type Cuadrado struct {
	Alto  float64
	Ancho float64
}

func (c *Cuadrado) area() float64 {
	return c.Ancho * c.Alto
}

func (c *Cuadrado) perimetro() float64 {
	return 2*c.Ancho + 2*c.Alto
}
