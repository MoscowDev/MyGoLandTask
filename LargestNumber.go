package main

import "fmt"

func main() {

	var counter int = 1
	var number int
	var largest int
	fmt.Print("Enter integer 1: ")
	fmt.Scan(&number)

	largest = number
	for counter < 10 {

		fmt.Print("Enter integer", counter+1, ": ")
		fmt.Scan(&number)

		if number > largest {
			largest = number
		}

		counter++
	}

	fmt.Println("The largest number is:", largest)
}
