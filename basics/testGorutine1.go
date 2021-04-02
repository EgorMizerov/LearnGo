package main

/*
Вывод:
Все горутины умирают после того, как умирает процесс main
Если горутина запущена внутри второстепенной функии и эта функция
закончила свою работу, то запущенная ею горутина не умирает!
*/

import (
	"fmt"
	"time"
)

func main() {
	test()
	time.Sleep(time.Second * 5)
}

func test() {
	defer func() {
		fmt.Println("Test закончилась")
	}()
	go func() {
		for {
			time.Sleep(time.Second * 1)
			fmt.Println("Хай")
		}
	}()
}
