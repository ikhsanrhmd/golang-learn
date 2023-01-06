package main
import "fmt"

//10 --> 10,5,2,1
// func faktorBilangan(bilangan int) int{
// 	//var value int

	
// 	return bilangan
// }

func main(){
	var bilangan int = 10
	for i:=1; i < bilangan; i++{
		if (bilangan % i == 0){
			fmt.Println(i)

		}
		
	}

}
	// fmt.Println(faktorBilangan(10))
