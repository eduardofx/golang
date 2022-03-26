package main

import "fmt"

func main() {
	fmt.Println("struct")

	var u usuario
	u.name = "Eduardo"
	u.age = 18
	fmt.Println(u)

	u2 := usuario{name: "Eduardo"}
	fmt.Println(u2)

	u3 := estudante{curso: "Sistemas", usuario: usuario{name: "Eduardo", age: 18}}
	fmt.Println(u3)
}

type usuario struct {
	name string
	age  int
}

type estudante struct {
	usuario
	curso     string
	faculdade string
}
