package main

import "fmt"

func add(input int) chan int {
	result := make(chan int)
	go func() {
		result <- input + 10
		close(result)
	}()
	return result
}

func main() {
	var input int
	fmt.Println("Enter your number : ")
	fmt.Scan(&input)
	final := add(input)
	result := <-final
	fmt.Println("result : ", result)
}
