package main

import (
	"fmt"
)

func main() {
	fmt.Println("Archivo para primer commit.")
	fmt.Println("Segundo commit")

	fmt.Printf("Mi suma es: %d\n", fn_add(10, 11))
}

// fn_add verifica si los numeros son mayores a 0,
// en caso de no ser así, regresa -1
func fn_add(a, b int64) int64 {
	if a > 0 && b > 0 {
		return a + b
	}
	return -1
}
