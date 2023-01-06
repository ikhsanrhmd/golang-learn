package main
import "fmt"
import "math"

/*
deklarasi variable int --> counter, bilangan =1
for jika counter kurang dari number maka bilangan++, if checkprime(bilangan) maka counter++, return  bilangan
*/

func checkPrime(number int) bool{
	if number < 2 {
		return false
	}
	sqrtNumber := int(math.Sqrt(float64(number)))
	for i := 2 ;  i<= sqrtNumber; i++{
		if number%i == 0{
			return false
		}
	}
	return true
}
func primeX(number int) int {
	var counter int
	var bilangan int = 1
	for counter < number{
		bilangan++
		if checkPrime(bilangan){
			counter++
		}
		
		fmt.Println("count:",counter, "bil:",bilangan)
	}
	return bilangan	
}
func main(){
	fmt.Println(checkPrime(4))
	fmt.Println(primeX(5))
}