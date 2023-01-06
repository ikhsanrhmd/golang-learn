package main

import "fmt"
/*
5! = 5 *4! --> 5*4*3*2*1 
   
*/
func factorialLoop(n int) int{
	var temp int = 1
	for i := n; i > 0 ; i--{
		temp *= i
	}
	return temp
}

func factorial(n int) int {
	//base case
	if n==1{
		return 1
	} else{ //recurrent relation
		return n*factorial(n-1) //fungsi memanggil dirinya sendiri
	}
}
func main(){
	fmt.Println(factorial(5))
	fmt.Println(factorialLoop(5))
}

