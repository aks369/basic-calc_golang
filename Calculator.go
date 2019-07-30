package main

import "fmt"

func main() {
	var num1, num2 float64
	var op string

	fmt.Println("Which operation do you want to perform? (+, -, *, /, %)")
	fmt.Scan(&op)

	fmt.Println("Enter the two numbers: ")
	fmt.Scan(&num1, &num2)

	switch op {
	case "+":
		fmt.Println(num1 + num2)
		break
	case "-":
		fmt.Println(num1 - num2)
		break
	case "*":
		fmt.Println(num1 * num2)
		break
	case "/":
		fmt.Println(num1 / num2)
		break
	default:
		fmt.Println("wrong operation")

	}
}
