package main

import "fmt"

func main() {

	fmt.Println("N\t10*N\t100*N\t1000*N")

	for number := 1; number <= 5; number++ {

		fmt.Println(
			number, "\t",
			number*10, "\t",
			number*100, "\t",
			number*1000,
		)
	}
}

