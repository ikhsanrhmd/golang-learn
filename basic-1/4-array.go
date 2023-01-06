//array merupakan kumpalan dari element/data dengan tipe yang sama
package main
import "fmt"

func main(){
	var bilangan [5]int
	fmt.Println("banyak element", len(bilangan))
	bilangan[0] = 10
	bilangan[1] = 14
	bilangan[2] = 10
	bilangan[len(bilangan)-1] = 11


	fmt.Println(bilangan)
	fmt.Println(bilangan[1])

	var kota = [3]string{"bandung","jakarta"}
	var kota2 = [3]string{"padang","jambi"}
	var kota3 []string
	kota3 = append(kota[:], kota2[:]...)
	
	fmt.Println(kota, kota2)
	fmt.Println(kota3)

	//inisialisasi element array diawal
	var bilangans = [5]int {1,2,3,4,5}
	fmt.Println(bilangans[1])

	//array element(n) otomatis
    var data = [...]int {1:1 , 2:3, 5:4}
	fmt.Println(data)
}