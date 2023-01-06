//interface merupakan kumpulan method yang harus di implement oleh sebuah object(behavior of object).

// type interface_name interface {
// 	method_name1 <return>
// }

package main
import "fmt"

type calculate interface{
	large() int
}
type square struct{
	side int
}

func (s square) large() int{
	return s.side * s.side
}
func main(){
	var dimResult calculate
	sqr := square{10}
	dimResult = sqr

	fmt.Println("Large square : ", dimResult.large())
}