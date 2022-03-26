package main

import "fmt"

func main() {
	soma := somar(2, 2)
	fmt.Println(soma)

	rsoma, rsub, _ := calcula(2, 2)
	fmt.Println(rsoma, rsub)
}

func calcula(n1, n2 int8) (int8, int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao, soma - subtracao
}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}
