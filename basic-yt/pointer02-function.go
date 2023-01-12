//data di function secara default pass by vaule
//untuk mengubah data asli di parameter dapat menggunakan pointer di function (*)

package main

import "fmt"

type Address struct {
	city, province, country string
}

func changeCountry(address *Address) { //agar data di function berubah perlu menggunakan pointer (*)
	//func changeCountry(address Address) {
	address.country = "indonesia"

}

func main() {
	alamat := Address{"manchester", "south-england", "england"}
	changeCountry(&alamat)
	fmt.Println(alamat)
}
