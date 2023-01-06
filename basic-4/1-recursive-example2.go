package main

import "fmt"


func fibonaci(n int) int {
	if n<=1{
		return n
	} else{ //recurrent relation
		return fibonaci(n-1) + fibonaci(n-2) //fungsi memanggil dirinya sendiri
	}
}
func main(){
	fmt.Println(fibonaci(2))
}

