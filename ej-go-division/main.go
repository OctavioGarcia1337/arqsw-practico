package main

import (
	"fmt"

	"github.com/OctavioGarcia1337/arq-software/ej-go-division/div"
	"github.com/OctavioGarcia1337/arq-software/ej-go-division/sum"
)

func main() {
	// defers the execution of a function to the end
	defer func() {
		fmt.Println("Bye!")
	}()

	// inputs
	a, b := float64(20), float64(10)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// division
	resultDiv, err := div.Division(a, b)
	if err != nil {
		fmt.Println("Error in division: ", err.Error())
		return
	}
	fmt.Println("Div: ", resultDiv)

	// sum
	resultSum := sum.Sum(a, b)
	fmt.Println("Sum: ", resultSum)
}
