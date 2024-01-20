package main

import (
	"fmt"
)

func main() {
	fmt.Println("Archivo para primer commit.")
	fmt.Println("Segundo commit")
	fmt.Println("Tercer commit en la rama main")
	fmt.Println("Cuarto commit de main")

	fmt.Printf("Mi suma es: %d\n", fn_add(10, 11))
	fmt.Printf("Suma no válida: %d\n", fn_add(-1, 22))
}

// fn_add verifica si los numeros son mayores a 0,
// en caso de no ser así, regresa -1
func fn_add(a, b int64) int64 {
	if a > 0 && b > 0 {
		return a + b
	}
	return -1
}
