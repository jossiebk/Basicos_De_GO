package main

import "fmt"

func main() {
	// Arrays
	var arreglo [4]int //si lo declaramos de esta forma, todos sus valores seran "zero" values hasta que se le asignen valores
	// Como en muchos lenguajes el conteo inicia en 0
	arreglo[0] = 5
	arreglo[1] = 10
	fmt.Println(arreglo)
	fmt.Println(len(arreglo)) // nos indicara su tamaño
	fmt.Println(cap(arreglo)) // nos indicara la capacidad maxima de almacenaje de este array

	// Slices
	// A un slice no se le indica la cantidad de elementos que podra tener.
	slice := []int{0, 2, 4, 6, 8, 10}
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// Algunos metodos de un slice
	fmt.Println(slice[0])   // imprime el indice 0
	fmt.Println(slice[:3])  // imprime los indices 0,1 y 2, excluyendo el tercer indice
	fmt.Println(slice[2:4]) //imprime del indice 2 al 4 inclusivamente
	fmt.Println(slice[4:])  // imprime del indice 4 al final inclusivamente

	// Append para agregar datoss
	slice = append(slice, 12)
	fmt.Println(slice)

	nuevoSlice := []int{15, 16, 17}
	slice = append(slice, nuevoSlice...) // Los 3 puntos al final indican que tomara el nuevo slice, extraera sus valoresy los agregara independientemente sin tener una lista anidada.
	fmt.Println(slice)

	/*
		################################################ RECORRER SLICES ################################################
	*/
	sliceDeCadena := []string{"hola", "mundo", "mi", "primer", "slice"}
	for i, valor := range sliceDeCadena {
		fmt.Println("Indice: ", i, " Valor: ", valor)
	}

	/*
		################################################ maps ################################################
	*/
	// Es una estructura de datos tipo clave/valor
	m := make(map[string]int)
	m["jossie"] = 15
	m["pepito"] = 22
	fmt.Println(m)

	// Para recorrer un map, esto al tener muchos valores, su recorrido ocurre de forma concurrente por lo que los valores pueden no mostrarse en el orden ingresado.
	// GO ya implementa de forma nativa concurrencia para su ejecucion, una de las cosas que lo hace un lenguaje tan potente.
	for i, valor := range m {
		fmt.Println("Indice: ", i, " Valor: ", valor)
	}

	// Encontrar un valor especifico
	// la variable ejemplo contendra el valor, y la variable codigo devolvera un codigo true o false dependiendo si el valor existe o no en el map
	ejemplo, codigo := m["jossie"]
	fmt.Println("Valor: ", ejemplo, " con codigo: ", codigo)
}
