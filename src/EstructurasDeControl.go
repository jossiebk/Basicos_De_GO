package main

import (
	"fmt"
)

func main() {
	/*
	 ******************************************************** CICLOS ********************************************************
	 */

	/*
		En el caso de GO el unico ciclo que posee es "FOR" junto con sus derivados.
	*/

	//For condicional
	for i := 0; i < 10; i++ {
		fmt.Println("Iteracion For tradicional: ", i)
	}

	//For While
	//  a diferencia del for tradicional lleva una condicion para permanecer en el ciclo o salir, como el while.
	contador := 0
	for contador < 10 {
		fmt.Println("Contador For While: ", contador)
		contador++
	}

	//For Forever
	//Esta variacion de FOR se ejecuta permanentemente hasta ser detenido manualmente, por lo que no se le agrega ninguna condicion de inicio o fin hasta que lo detengamos.
	/*contadorForever := 0
	for {
		fmt.Println("Contador For forever: ", contadorForever)
		contadorForever++
		// de no ser agregado correctamente puede saturar la memoria, por lo que se puede agregar internamente alguna forma para detenerlo ya sea con if u otro condicional
	}*/

	// For range
	listaNumerica := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	total := 0
	for _, numero := range listaNumerica {
		total += numero
	}
	fmt.Println("Total de for range: ", total)

	/*
	******************************************************** CONDICIONALES ********************************************************
	 */
	val1 := 1
	//val2:=2

	// Condificional If-Else
	if val1 == 1 {
		fmt.Println("La condicion se cumple, es 1")
	} else {
		fmt.Println("La condicion no se cumple, el valor no es igual a 1")
	}

	// Switch
	//modulo := 5 % 2
	//modulo:=3%2
	modulo := 20 % 2
	switch modulo {
	case 0:
		fmt.Println("Si es par")
	default:
		fmt.Println("Es impar")
	}

	//Switch sin condicionales
	// Ideal cuando son muchas condiciones internas.
	//valor3 := 6
	valor3 := 750
	//valor3 := 300
	switch {
	case valor3 < 100:
		fmt.Println("Es menor a 100")
	case valor3 > 100 && valor3 < 500:
		fmt.Println("Es mayor a 100")
	default:
		fmt.Println("No aplica en las condiciones")
	}
}
