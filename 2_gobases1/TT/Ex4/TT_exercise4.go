package main

import "fmt"

func main() {

	var monthNumber int
	fmt.Println("Enter the month number:")
	fmt.Scanln(&monthNumber)

	var months = map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December"}

	fmt.Println("The month is:", months[monthNumber])

	//Elegi resolverlo con un map porque es mucho mas resumido que un switch

}
