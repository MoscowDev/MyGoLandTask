package main

import "fmt"

func main() {
	var palindrome int
	fmt.Print("Enter a number: ")
	fmt.Scan(&palindrome)
	original := palindrome
	var reversed int	
	for palindrome > 0 {
		digit := palindrome % 10
		reversed = reversed*10 + digit
		palindrome /= 10
	}	
	if original == reversed {
		fmt.Printf("%d is a palindrome.\n", original)
	} else {
		fmt.Printf("%d is not a palindrome.\n", original)
	}	
	

}