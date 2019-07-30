package main

import "fmt"

func main() {
	var num1, num2 int
	var op string

	fmt.Println("Which operation do you want to perform? (+, -, *, /, %)")
	fmt.Scan(&op)

	if op == "+" || op == "-" || op == "*" || op == "/" || op == "%" {
		fmt.Println("Enter the two numbers: ")
		fmt.Scan(&num1, &num2)

		switch op {
		case "+":
			fmt.Println("Sum = ", num1+num2)
		case "-":
			fmt.Println("Differece = ", num1-num2)
		case "*":
			fmt.Println("Product = ", num1*num2)
		case "/":
			fmt.Println("Quotient = ", num1/num2)
		case "%":
			fmt.Println("Remainder = ", num1%num2)
		default:
			fmt.Println("invalid")
		}
	} else {
		fmt.Println("Invalid Operation")
	}

}
