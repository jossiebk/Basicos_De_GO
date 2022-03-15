package main

import "fmt"

func main() {
	/*
		################################################ TIPOS DE DATOS PRIMITIVOS ################################################
	*/
	// Para la declaracion de variables se debe definir var + nombre + tipo de dato y su valor, el tipo de dato es opcional dado que al
	// Inicializar la variable con su valor Go identifica el tipo de dato al cual pertenece.
	// o unicamente como nombre := valor donde intepretara el tipo de dato.

	// Declaracion de constantes
	// Se utiliza la palabra reservada const
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//Declaracion de variables enteras
	base := 12          //Forma 1
	var altura int = 14 //Forma 2
	var area int        //Forma 3

	fmt.Println(base, altura, area)

	// Numeros enteros con signo positivo o negativo
	// int8 va desde -128 a 127 bits
	var entero8bit int8 = -16
	fmt.Println("El valor entero de 8 bits es: ", entero8bit)
	// int16 va desde -32768 a 32767 bits
	var entero16bit int16 = -32111
	fmt.Println("El valor entero de 16 bits es: ", entero16bit)
	// int32 va desde -2147483648 a 2147483647 bits
	var entero32bit int32 = 2144444444
	fmt.Println("El valor entero de 32 bits es: ", entero32bit)
	// int64 va desde -9223372036854775808 a 9223372036854775807 bits
	var entero64bit int64 = 4654321764764136315
	fmt.Println("El valor entero de 64 bits es: ", entero64bit)

	// Numeros enteros sin signo, es decir enteros positivos.
	// int8 va desde 0 a 255 bits
	var positivo8bit uint8 = 65
	fmt.Println("El valor entero de 8 bits es: ", positivo8bit)
	// int16 va desde 0 a 32767 bits
	var positivo16bit uint16 = 65535
	fmt.Println("El valor entero de 16 bits es: ", positivo16bit)
	// int32 va desde  a 4294967295 bits
	var positivo32bit uint32 = 2144444444
	fmt.Println("El valor entero de 32 bits es: ", positivo32bit)
	// int64 va desde 0 a 18446744073709551615 bits
	var positivo64bit uint64 = 4654321764764136315
	fmt.Println("El valor entero de 64 bits es: ", positivo64bit)

	// Valores decimales
	var decimal32Bits float32 = 123.222
	fmt.Println("El valor decimal de 32 bits es: ", decimal32Bits)
	var decimal64Bits float64 = 1.7e+308
	fmt.Println("El valor decimal de 64 bits es: ", decimal64Bits)

	// Adicionalmente Go nos permite manejar datos completos es decir un valor real+imaginario ya sea de 64 o 128 bits
	var variableCompleja = 2 + 16i
	fmt.Println("El valor complejo: ", variableCompleja)

	// Go no posee un tipo de dato Char como muchos otros lenguajes de alto nivel para representar carateres pero posee el tipo rune
	// para el manejo la representacion de caracteres ASCII y sus caracteres especiales en representacion ASCII.
	var caracterEspecial rune = 'ð'
	fmt.Println("El valor especial ASCII es: ", caracterEspecial)

	// Cadenas de caracteres.
	var cadenaNueva string = "Hola mundo"
	fmt.Println("El valor de la cadena es: ", cadenaNueva)
	//Zero values
	//Go nos asigna valores "vacios" a variables sin un valor asignado
	var a int     //asigna 0
	var b float64 //asigna 0
	var c string  //asigna cadena vacia
	var d bool    //asigna false
	// Go no maneja valores NULL, sino que auto asigna un valor para el manejo de sus variables cuando no se les inicializa un valor.

	fmt.Println(a, b, c, d)

	//Area de un cuadrado base*altura
	const BaseCuadrado = 10
	areaCuadrado := BaseCuadrado * BaseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)

	/*
		################################################ OPERADORES ARITMETICOS ################################################
	*/
	x := 10
	y := 50
	// Suma
	result := x + y
	fmt.Println("Suma: ", result)

	// Resta
	result = y - x
	fmt.Println("Resta: ", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion: ", result)

	// Division
	result = y / x
	fmt.Println("Division: ", result)

	// Residuo
	result = y % x
	fmt.Println("Residuo: ", result)

	// Incremental
	x++
	fmt.Println("Incremental: ", x)

	// Decremental
	x--
	fmt.Println("Decremental: ", x)

	// Reto
	// Calcular area del rectangulo, trapecio y cirulo
	BaseTriangulo := 5
	AlturaTriangulo := 15
	AreaTriangulo := (BaseTriangulo * AlturaTriangulo) / 2
	fmt.Println("Area de un triangulo: ", AreaTriangulo)

}
