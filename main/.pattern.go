package main

import (
	"fmt"
)

func main() {
	var op rune = '♥'
	fmt.Printf("♥♥♥♥♥♥♥    ♥♥        ♥♥♥♥♥♥  ♥♥\n")
	for i :=1; i<=6; i++ {
		fmt.Printf("   %c       %c%c       %c      %c   " , op,op,op,op,op)
		for j :=3; j<=2*i; j++ {
			fmt.Printf(" ")	
		}
		fmt.Printf("♥♥")
		fmt.Printf("\n")
			
		
	}
	fmt.Printf("♥♥♥♥♥♥♥    ♥♥♥♥♥♥♥♥  ♥♥♥♥♥♥                ♥♥\n")
}


    // fmt.Printf("   %c       %c%c       %c      %c   ♥♥        ♥♥   ♥♥           ♥♥      ♥♥   %c      %c  %c      %c\n" , op,op,op,op,op,op,op,op,op)
	// fmt.Printf("   %c       %c%c       %c      %c    ♥♥      ♥♥    ♥♥            ♥♥    ♥♥    %c      %c  %c      %c\n" , op,op,op,op,op,op,op,op,op)
	// fmt.Printf("   %c       %c%c       %c      %c     ♥♥    ♥♥     ♥♥♥♥           ♥♥  ♥♥     %c      %c  %c      %c\n" , op,op,op,op,op,op,op,op,op)
	// fmt.Printf("   %c       %c%c       %c      %c      ♥♥  ♥♥      ♥♥              ♥♥♥♥      %c      %c  %c      %c\n" , op,op,op,op,op,op,op,op,op)
	// fmt.Printf("   %c       %c%c       %c      %c       ♥♥♥♥       ♥♥               ♥♥       %c      %c  %c      %c\n" , op,op,op,op,op,op,op,op,op)