package main

import (
	"errors"
	"fmt"
)

func verificationMessageTest(salary int) error {
	if salary < 150000 {
		return errors.New("error: the salary entered does not reach the taxable minimum")
	}
	return nil
}

func main() {
	var salary int

	fmt.Println("Enter salary: ")
	fmt.Scanln(&salary)

	err := verificationMessageTest(salary)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Must pay tax")
	}
}
