package main
import "fmt"

//sebuah fungsi untuk mencari sebuah bilangan berada di index berapa
//jika tidak ditemukan maka akan return -1
func linierSearch(elements []int, x int) int{
	for i := 0; i < len(elements); i++{
		if elements[i] == x{
			return i
		}
	} 
	return -1
}

func main(){
	var data = []int{1,2,3,4,5,6,7}
	fmt.Println(linierSearch(data,12))

}