package main

import "fmt"

func main(){
	var today int = 1
	// switch today {
	// case 1:
	// 	fmt.Println("pertama")
	// case 2:
	// 	fmt.Println("kedua")
	// default: //sama dengan else
	// 	fmt.Println("invalid")
	// }
	
	switch {
	case today == 1:
		fmt.Println("pertama")
		fallthrough
	case today >= 2:
		fmt.Println("kedua")
	default: //sama dengan else
		fmt.Println("invalid")
	}
}