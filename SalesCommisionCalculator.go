package main

import "fmt"

func main() {

	var totalSales float64 = 0

	for {
		var itemValue float64
		fmt.Print("Enter item value sold (-1 to end): ")
		fmt.Scan(&itemValue)

		if itemValue == -1 {
			break
		}

		if itemValue < 0 {
			fmt.Println("Invalid input. Enter a positive value.")
			continue
		}

		totalSales += itemValue
	}

	earnings := 200 + (0.09 * totalSales)

	fmt.Printf("Total sales: $%.2f\n", totalSales)
	fmt.Printf("Weekly earnings: $%.2f\n", earnings)
}
