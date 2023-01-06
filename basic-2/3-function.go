package main
import "fmt"

func VolBalok(alas, tinggi, lebar int) (luas int, vol int){ // (parameter1, parameter2, tipe data(int) (return(int)))
	luas = (2*alas*tinggi)+(2*alas*lebar)+(2*tinggi*lebar)
    vol = alas*tinggi*lebar
	return 
}


func main(){
	fmt.Println(VolBalok(10, 4, 20))
	resultLuas, resultVol := VolBalok(10, 4, 20)
	fmt.Printf("Luas: %d\n", resultLuas)
	fmt.Printf("Volume: %d\n", resultVol)

	resultLuas, _ = VolBalok(10, 4, 20)
	fmt.Printf("Luas: %d\n", resultLuas)
}
