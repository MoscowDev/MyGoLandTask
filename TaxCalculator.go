package main

import (
	"fmt"
)	
func main() {

	for i := 0; i < 3; i++ {


		fmt.Print("Enter name: ")
		var name string
		fmt.Scan(&name)
		
		var income float64
		fmt.Print("Enter annual income: ")
		fmt.Scan(&income)
		var taxRate float64

		if income <= 30000 {
			taxRate = 0.15
		
		} else {
			taxRate = 0.20
		}

		taxAmount := income * taxRate
		fmt.Printf("Income: $%.2f, Tax Rate: %.2f%%, Tax Amount: $%.2f\n", income, taxRate*100, taxAmount)		
	}
}