package main
import "fmt"
func luasSegitiga(alas float64, tinggi float64) float64{
	return (alas * tinggi) * 0.5
}

func main(){
	fmt.Println(luasSegitiga(20,25))
}