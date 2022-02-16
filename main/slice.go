package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Printf("Enert number of input you want to give : ")
	fmt.Scanf("%d", &a)
	fmt.Printf("Start to enter numbers - \n")
	var x = make([]int,a)
	for g := 0; g < a; g++ {
	    fmt.Scanf("%d",&x[g])
	}
	
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