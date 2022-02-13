package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var num int
	var addrs string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter your Full Name - ")
	scanner.Scan() // For Long string with space and special character
	line := scanner.Text()

	fmt.Printf("Enter Your Contact - ")
	fmt.Scanf("%d", &num) // For Int
	fmt.Printf("Enter Your Address - ")
	fmt.Scanf("%s", &addrs) // For single String

	fmt.Println("==================")
	fmt.Println("== User Details ==")
	fmt.Println("==================")
	fmt.Println("Full Name:", line)
	fmt.Println("Mobile : ", num)
	fmt.Println("Address : ", addrs)

}
