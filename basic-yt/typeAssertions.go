//kemampuan merubah tipe data sesuai dengan yg kita inginkan
//seringkali digunakan dengan data interface kosong

//lebih aman menggunakan swicth assertion agar tidak ada panic err

package main

import "fmt"

func random() interface{} {
	return true

}

func main() {
	result := random()
	//resultString := result.(string)
	//fmt.Println(resultString)

	// resultInt := result.(int)
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("value", value, "is string")
	case int:
		fmt.Println("value", value, "is integer")
	default:
		fmt.Println("value unknown")
	}
}
