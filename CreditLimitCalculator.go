package main

import (
	"fmt"
)

func main() {

	var creditLimit float64
	var balance float64
	var totalCharges float64
	var totalCredits float64	
	fmt.Print("Enter credit limit: ")
	fmt.Scan(&creditLimit)	;
	fmt.Print("Enter balance: ")
	fmt.Scan(&balance);
	fmt.Print("Enter total charges: ")
	fmt.Scan(&totalCharges);
	fmt.Print("Enter total credits: ")
	fmt.Scan(&totalCredits);	
	newBalance := balance + totalCharges - totalCredits
	fmt.Printf("New balance: %.2f\n", newBalance)
	if newBalance > creditLimit {
		fmt.Println("Credit limit exceeded.")
	} else {
		fmt.Println("Credit limit not exceeded.")
	}	
}