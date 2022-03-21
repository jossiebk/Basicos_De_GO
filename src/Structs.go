package main

import "fmt"

/*
	################################################ Structs ################################################
*/

// Si creamos los nombres de parametros en minuscula Go lo tomara como variables privadas, si los escribimos con la inicial mayuscula seran publicos.
// Esto de parte del lenguaje GO
type carro struct {
	marca string
	anio  int
}

/*
	################################################ Structs y punteros ################################################
*/

func main() {

	miCarro := carro{marca: "patito", anio: 2022}
	fmt.Println(miCarro)

	// Para punteros
	a := 10 // Variable que contendra el dato
	b := &a // Variable que al agregarle & se le almacenara la posicion de memoria donde esta almacenado a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b) // al agregar el asterisco a la variable que apunta a una direccion de memoria se le indara que muestra el valor que esta en dicha posicion de memoria
}
