package main

import (
	"fmt"
	"os"
)

/*
Ejercicio 3 - Impuestos de salario #3

Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de
error reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible
(el mensaje mostrado por consola deberá decir: “error: el mínimo imponible es de 150.000 y
el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado
por parámetro).
*/

func checkSalary(salary int) error {

	if salary < 150000 {
		return fmt.Errorf("error: el minimo imponible es de 150000 y el salario ingresado es de: %d", salary)
	}

	return nil

}

func main() {

	salary := 100000

	err := checkSalary(salary)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}
