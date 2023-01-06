package main
import(
	"fmt"
)

func BinarySort(data []int, x int) int {
	var kiri = 0
	var kanan = len(data) - 1
	
	for kiri <= kanan{
		var tengah = (kiri + kanan) / 2
		if data[tengah] == x {
			return tengah
		} else if x < data[tengah]{
			kanan = tengah -1
		} else if x > data[tengah]{
			kiri = tengah + 1
		}
 
	}
	return -1
}

func main(){
	var data = []int{1,3,7,9,10,13}
	fmt.Println(BinarySort(data,10))
}
 