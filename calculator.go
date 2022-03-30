package main

import "fmt"

func add(a, b int) int {

	output := a + b
	return output
}

func subtract(a, b int) int {
	output := a - b
	return output

}

func multiply(a, b int) int {
	output := a * b
	return output

}

func divide(a, b int) int {
	output := a / b
	return output

}

func main() {
	var a, b int
	var process string
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&a)
	fmt.Println("Enter process: (add, sub, mul, div)")
	fmt.Scanln(&process)
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&b)

	if process == "add" {
		res := add(a, b)
		fmt.Println(res)
	} else if process == "sub" {
		res := subtract(a, b)
		fmt.Println(res)
	} else if process == "mul" {
		res := multiply(a, b)
		fmt.Println(res)
	} else if process == "div" {
		res := divide(a, b)
		fmt.Println(res)
	} else {
		fmt.Println("Something went wrong, please check your input!")
	}
}
