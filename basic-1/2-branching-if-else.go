package main
import "fmt"

func main()  {
	// var angka int = 10
	// if angka == 10 {
	// 	fmt.Println("Angka 10")
	// } else if angka == 100{
	//     fmt.Println("Angka 100")
	// } else {
	//     fmt.Println("Angka != 100")
	// }
	
	var data = 85
	if data >= 70 && data <= 90{
		fmt.Println("Data antara 70 s/d 90")
	} else if data >= 60 && data <= 70{
		fmt.Println("data antara 60 s/d 70")
	} else{
		fmt.Println("angka diluar 70 s/d 90")
	}
} 