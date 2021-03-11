package main

import "fmt"

func main() {
	x := 10
	px := &x
	fmt.Println("Ссылка: ", px)
	fmt.Println("Значение: ", *px)
	fmt.Println("Ссылка переменной сслык: ", &px)
}
