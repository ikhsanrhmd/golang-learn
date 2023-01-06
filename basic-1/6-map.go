package main

import "fmt"
//map merupakan data structure berupak key dan value dengan tipe data apapun
func main(){
	var month = map[string]int{"januari":1, "desember":12 }
	fmt.Printf("%d\n", month["januari"])
	fmt.Printf("%v\n\n",month)
    
	var anyMap = map[string]interface{}{"ikhsan": 1, "hari": true} //key dengan value bernilai tipe apasaja
	fmt.Printf("%v",anyMap)

}