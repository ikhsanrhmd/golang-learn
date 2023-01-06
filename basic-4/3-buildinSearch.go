package main
import "fmt"
import "sort"

func main(){
	var data = []int{1,3,5,6,7}
	value := 4
	var findIndex = sort.SearchInts(data, value)

	if value == data[findIndex]{
		fmt.Println("data ditemukan", findIndex)
	}else{
		fmt.Println("data tdk ditemukan")
	}
}