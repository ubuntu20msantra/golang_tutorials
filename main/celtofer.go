package main

import (
	"fmt"
)

func main() {

	var input float64
	fmt.Printf("Enter temp. in F : ")
	fmt.Scanf("%f", &input)

	var output float64

	output = ((input - 32) * 5/9)

	fmt.Println("================")
	fmt.Println("==== F to C ====")
	fmt.Println("================")

	fmt.Println("Temp. in Celcius : ", output)
}