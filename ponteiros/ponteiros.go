package main

import "fmt"

func main() {
	// Ponteiro é uma referência na memória para um valor

	var n int
	var pont *int

	n = 10
	pont = &n // &n = endereço de memória do n

	fmt.Println(n, pont)

	fmt.Println(*pont) // * é o operador de ponteiro [desreferenciar]
}
