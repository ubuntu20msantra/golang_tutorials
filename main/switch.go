package main

import (
	"fmt"
)

func main() {
	fmt.Println("===================================")
	fmt.Println("==== Welcome to SBX Restrurant ====")
	fmt.Println("===================================")

	fmt.Println(" (:: See Menu to Order ::) ")
		fmt.Println("---------------------------")
		fmt.Println("1. Tea - Rs.10.00/- [No.1]")
		fmt.Println("2. Special Tea - Rs.15.00/- [No.2]")
		fmt.Println("3. Tandoor Tea - Rs.25.00/- [No.3]")
	fmt.Println("-----------------------------------")

	var num int 

		fmt.Printf("Enter your Order Number - ")
		fmt.Scanf("%d", &num)
		switch num {
			case 1: fmt.Println("You Ordered - Tea")
			case 2: fmt.Println("You Ordered - Special Tea")
			case 3: fmt.Println("You Ordered - Tandoor Tea")
			default: fmt.Println("Invalid Order Number")
			}
}