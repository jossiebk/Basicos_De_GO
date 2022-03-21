package main

import "fmt"

// Podemos definir diferentes structs
type cuadrado struct {
	base float32
}

type rectangulo struct {
	base   float32
	altura float32
}

//Las funciones asociadas a cada struct
func (cuadro cuadrado) area() float32 {
	return cuadro.base * cuadro.base
}

func (rect rectangulo) area() float32 {
	return rect.base * rect.altura
}

type figura interface {
	area() float32
}

// Y una interfaz que puede aplicar a distintos structs
func calcularAreaFigura(f figura) {
	fmt.Println("El area de la figura es: ", f.area())
}

func main() {
	FiguraCuadrado := cuadrado{base: 5}
	FiguraRectangulo := rectangulo{base: 5, altura: 10}

	calcularAreaFigura(FiguraCuadrado)
	calcularAreaFigura(FiguraRectangulo)
}
