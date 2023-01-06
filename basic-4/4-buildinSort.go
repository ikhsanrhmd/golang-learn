package main
import (
	"fmt"
	"sort"
)
func main(){
	var data = []int{1,21,3,43,7,9,10,13}
	fmt.Println("sebelum: ",data)
	sort.Ints(data)
	fmt.Println("sesudah: ",data)

	var dataString = []string{"i","k","hsan"}
	fmt.Println("\nsebelum: ",dataString)
	sort.Strings(dataString)
	fmt.Println("sesudah: ",dataString)
	fmt.Println("\n")
	datastatus := sort.IntsAreSorted(data)
	fmt.Println(datastatus)
	fmt.Println("\n")
	
	//ascending
	//return data2[i] > data2[j]
	//descending
	//return data2[i] < data2[j]
	
	var data2 = []int{1,21,3,43,7,9,10,13}
	sort.SliceStable(data2, func(i,j int) bool{
		return data2[i] > data2[j]
	})
	fmt.Println("hasil: ",data2)

}

