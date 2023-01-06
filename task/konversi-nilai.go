package main

import "fmt"

func main(){
	var score int
	fmt.Printf("enter your score: ")
	fmt.Scanln(&score)
	switch{
	case score >= 80 && score <= 100:
	  fmt.Println("A")
	case score >= 65 && score <= 79:
		fmt.Println("B+")
	case score >= 50 && score <= 64:
		fmt.Println("B")
	case score >= 35 && score <= 49:
		fmt.Println("C")
	case score >= 0 && score <= 34:
		fmt.Println("D")
	default:
		fmt.Println("Invalid")
	}
}