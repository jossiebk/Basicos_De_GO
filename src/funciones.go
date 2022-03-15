package main

import "fmt"

// Uso de una funcion basica con un parametro.
func normalFunction(message string) {
	fmt.Println(message)
}

// Uso de una funcion con multiples parametros de entrada.
func tripleArgumento(a int, b int, c string) {
	fmt.Println(a, b, c)
}

// Definicion de una funcion con un valor de retorno.
// En go se define el tipo de dato de salida unicamente.
func returnValue(a int) int {
	return a * 2
}

// Con Go se pueden retornar multiples valores de salida, definiendo el tipo de dato y nombre del valor de retorno.
// el valor retornado sera en el orden de asignacion es decir c y d retornaran a y  a*3 respectivamente
func dobleReturn(a int) (c, d int) {
	return a, a * 3
}

func dobleReturnTipoDistinto(a int) (c int, d string) {
	Saludo := "Hola mundo"
	return a, Saludo
}

// Posterior se pueden llamar las funciones dentro de un main segun sea necesario.
func main() {
	normalFunction("Hola mundo")
	tripleArgumento(1, 2, "hola")
	value := returnValue(4)
	fmt.Println("Value: ", value)
	//value1, value2 := dobleReturn(3)
	//fmt.Println("Value1 y Value2 ", value1, value2)
	value1, _ := dobleReturn(3)
	fmt.Println("Solo value1: ", value1)
	valor1, valor2 := dobleReturnTipoDistinto(10)
	fmt.Printf("El valor 1 es: %d y el valor 2 es: %s \n", valor1, valor2)
}
