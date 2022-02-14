package main

import (
	"fmt"
)

func main() {
	var s, e int
	fmt.Println("(:: Enter two number, we will find odd-even b/w this series ::)")
	fmt.Printf("Start Number : ")
	fmt.Scanf("%d", &s)
	fmt.Printf("End Number : ")
	fmt.Scanf("%d", &e)
	for i := s; i <= e; i++ {
		if i % 2 == 0 {
		  fmt.Println(i, "- even")
		} else {
		  fmt.Println(i, "- odd")
		}
	  }
}