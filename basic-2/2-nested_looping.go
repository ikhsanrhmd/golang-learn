package main
import "fmt"

func main(){
	// for i:= 10; i>= 0; i--{
	// 	for j:= 1; j<=i; j++{
	// 		fmt.Printf("*")
	// 	}
	// 	fmt.Println("")
	// }
	for i:= 1; i<= 10; i++{
		for x:= (10-i); x>=0; x--{
			fmt.Printf(" ")
		}
		for j:= 1; j<=i; j++{
			fmt.Printf("*")
		}
		fmt.Println("")
	}

	//looping string
	var name = "IKHSAN"
	for k:=0; k < len(name); k++{
		for l:= 0; l<=k; l++{
			fmt.Printf(""+ string(name[l]))
		}
		fmt.Println()
    }

	for _, value := range name{
		fmt.Printf(string(value))
	}
}