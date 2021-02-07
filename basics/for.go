package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	cond := true
	cntr := 0

	// Классический цикл for
	for i := 1; i < len(args); i++ {
		fmt.Println(args[i])
	}

	// Классический цикл while
	for cond {
		if cntr == 10 { cond = false }
		fmt.Printf("%d из 10\n", cntr)
		cntr++
	}

	// Классичсекий бесконечный цикл
	for {
		if cntr > 20 { break }
		fmt.Printf("%d из 20\n", cntr)
		cntr++
	}

	// Классический цикл foreach
	for key, value := range args {
		fmt.Printf("Ключ - %d, значение - %s\n", key, value)
	}
}