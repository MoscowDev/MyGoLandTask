package main

import (
	"fmt"
)

func main() {	

	var miles int
	var gallons int

	var totalMiles int
	var totalGallons int

	fmt.Print("Enter miles driven (-1 to quit): ")
	fmt.Scan(&miles)

	for miles != -1 {

		fmt.Print("Enter gallons used: ")
		fmt.Scan(&gallons)

		if gallons == 0 {
			fmt.Println("Gallons must be greater than zero; skipping this trip.")
			fmt.Print("\nEnter miles driven (-1 to quit): ")
			fmt.Scan(&miles)
			continue
		}

		mpg := float64(miles) / float64(gallons)

		fmt.Printf("Miles per gallon for this trip: %.2f\n", mpg)

		totalMiles += miles
		totalGallons += gallons

		fmt.Print("\nEnter miles driven (-1 to quit): ")
		fmt.Scan(&miles)
	}

	if totalGallons != 0 {
		combinedMPG := float64(totalMiles) / float64(totalGallons)
		fmt.Printf("\nCombined miles per gallon: %.2f\n", combinedMPG)
	} else {
		fmt.Println("No trips were entered.")
	}
}

