package main

import (
	"fmt"
)

func verificarSalario(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("el minimo imponible es de %d y el salario ingresado es de %d", 150000, salary)
	}
	return nil
}

func main() {
	var salary int
	fmt.Println("Ingrese su salario")
	fmt.Scanln(&salary)
	e := verificarSalario(salary)
	if e != nil {
		fmt.Println("Error:", e)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
