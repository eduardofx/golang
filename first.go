package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
    fmt.Println("Hello, World!")
	auxiliar.Write()

	erro := checkmail.ValidateFormat("eduardok.fx@gmail.com")
	fmt.Println(erro)
}