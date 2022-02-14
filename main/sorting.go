package main

import (
	"fmt"
)

func main() {
	x := []int{
		1,96,86,68,
		57,82,63,0,
		37,34,83,27,
		19,97,4,17,
	  }
	//  fmt.Println(x[len(x)-1])
	  for i :=0; i<len(x)-2; i++ {
		for j :=i+1; j<=len(x)-1; j++ {
			if x[i]>x[j] {
				temp := x[j]
				x[j] = x[i]
				x[i] = temp
			}
		}
	  }

	  fmt.Println("---  smallest value ---")
      fmt.Println(x[0])

	  fmt.Println("---  largest value ---")
	  fmt.Println(x[len(x)-1])

	  fmt.Println("---  sorting ---")
		for i :=0; i<len(x);i++ {
			fmt.Printf("%d, ",x[i])
		}
	fmt.Println("\n---  Thank You ---")
}