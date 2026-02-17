package main

import "fmt"

func main() {
	var number int
	var number2 int
	var largest int
	var secondLargest int

	fmt.Print("Enter integer 1: ")
	fmt.Scan(&number)

	fmt.Print("Enter integer 2: ")
	fmt.Scan(&number2)

	if number >= number2 {
		largest = number
		secondLargest = number2
	} else {
		largest = number2
		secondLargest = number
	}

	for i := 3; i <= 10; i++ {
		fmt.Printf("Enter integer %d: ", i)
		fmt.Scan(&number)

		if number > largest {
			secondLargest = largest
			largest = number
		} else if number > secondLargest {
			secondLargest = number
		}
	}

	fmt.Println("The largest number is:", largest)
	fmt.Println("The second largest number is:", secondLargest)
}
