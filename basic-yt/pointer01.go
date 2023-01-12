// secara default digolang menggunakan pass by value, bukan by reference
// mkdsnya jika mengirim variable ke func, method atau var lain. yang dikirim adalah duplicasi value nya
package main

import "fmt"

type address struct {
	city, province, country string
}

func main() {
	var address1 address = address{"padang", "sumbar", "indo"}
	var address3 *address = &address1
	//pass by reference
	var address2 *address = &address1

	//pass by value
	//address2 := address1
	address2.city = "bandung"

	*address2 = address{"jakarta", "dki", "indo"}

	//function new -> membuat data dari struct menjadi kosong

	var address4 *address = new(address)
	address4.city = "jakarta"
	fmt.Println(address4)

	fmt.Println(address2)
	fmt.Println(address1)
	fmt.Println(address3) //mengikuti pointer *address2

}
