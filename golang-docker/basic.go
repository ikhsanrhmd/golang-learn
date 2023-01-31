package main

import (
	"fmt"
	"unsafe"

	geo "tes/geometry"

	"rsc.io/quote"
)

func rectProps(length, width float64) (float64, float64) {
	area := length * width
	perameter := 2 * (length + width)
	return area, perameter

}

// func main(){
func basic() {
	var x = 1
	name := "devops"
	fmt.Println("hello wordl")
	fmt.Println(quote.Go())
	fmt.Println(x)
	fmt.Printf("type of name %T and size is %d \n", name, unsafe.Sizeof(name))
	a, p := rectProps(1, 2)
	fmt.Printf("\narea: %f", a)
	fmt.Printf("parameter: %f\n", p)

	var dayOfMonth = map[string]int{"jan": 31}
	fmt.Println(dayOfMonth)

	area := geo.Area(1, 2)
	fmt.Println(area)

	diagonal := geo.Diagonal(5, 25)
	fmt.Println(diagonal)
}
