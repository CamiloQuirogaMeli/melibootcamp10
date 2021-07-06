package main

import "fmt"

func main() {
	num := 3
	isPair(num)
	fmt.Println("Ejecución completada!")
}

func isPair(num int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	if (num % 2) != 0 {
		panic("no es un número par")
	}
	fmt.Println(num, " es un número par!")
}
