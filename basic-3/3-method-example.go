package main
import "fmt"

type Lingkaran struct{
	jari int
}

type Persegi struct{
	sisi int
}

func (c Lingkaran) HitungLuas() float64{
	luas := 3.14 * float64(c.jari) * float64(c.jari)
	return luas
}
func (s Persegi) HitungLuas() float64{
	luas := s.sisi * s.sisi
	return float64(luas)
}

func main(){
	lingkaran := Lingkaran{
		jari: 5,
	}
	persegi := Persegi{
		sisi: 10,
	}

	fmt.Println(lingkaran.HitungLuas())
	
	fmt.Println(persegi.HitungLuas())
}
