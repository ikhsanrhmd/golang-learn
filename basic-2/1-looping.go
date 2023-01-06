//looping merupakan statement melakukan suatu hal secara berulang

package main
import "fmt"

//1,3,5,7,11
func main(){
	for i:= 0; i<14; i++{
		//fmt.Println(i*i)
	}
    
	
	var mapString = map[string]int{"satu":1,"dua":2}
	var sliceInt = []int{1,2,3,4,5,6}
	for key, value := range sliceInt{  //looping sebanyak data pada slice
		fmt.Printf("Key:%d- value:%v\n", key, value)
	}
	for key, value := range mapString{  //looping sebanyak data pada map
		fmt.Printf("Key:%s- value:%v\n", key, value)
	}
}