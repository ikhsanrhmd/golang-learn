//Slice
//mirip seperti array(element yg hanya bisa satu tipe data) tapi size dynamic

//Array
//array merupakan kumpalan dari element/data dengan tipe yang sama

//kapan menggunakan array & slice?
// kondisi menggunakan slice jika data yg akan di masukan tidak diketahui jumlah nya, spt: data pada db
// menggunakan array jika data yg ingin digunakan sudah fix 

package main
import "fmt"
import "reflect"

func main(){
  var bilangan = [5]int{1,2,3,4,5}

  // create slice from array
  var bagian_bilangan []int = bilangan[1:5] //mengambil array dari element 1 s/d 3
  bagian_bilangan = append(bagian_bilangan, 10000) //menambahkan data ke slice akan menambah data di array juga

  fmt.Println(reflect.ValueOf(bagian_bilangan).Kind())
  fmt.Println(bagian_bilangan)

  var datas = []int{1,2,4,5,3}
  fmt.Println(datas)
  datas = append(datas, 22)
  fmt.Println(datas)

  var datas2 []int = bilangan[:4]
  fmt.Println(datas2)
  
  //long declaration
  var even_number []int
  fmt.Printf("element = %v, len = %d, cap= %d\n", even_number, len(even_number), cap(even_number))
  //long declaration with value
  var odd_number = []int{1, 3, 5, 7, 9}
  fmt.Printf("element = %v, len = %d, cap= %d\n", odd_number, len(odd_number), cap(odd_number))
  //using make function
  var prime = make([]int, 5, 10) //make digunakan untuk menentukan banyak (len, capacity )
  fmt.Printf("element = %v, len = %d, cap= %d\n", prime, len(prime), cap(prime))

  copy(datas2, odd_number) //copy (dest, source)
  fmt.Printf("element = %v, len = %d, cap= %d\n", datas2, len(datas2), cap(datas2))
}