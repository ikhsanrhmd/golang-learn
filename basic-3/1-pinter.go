// pointer merupakan variable yg disimpan dalam memory address(RAM)
// pointer mempunyai kemampuan untuk manipulasi data yg di pointing/ditunjuk
package main
import "fmt"

//pass by value
func ubahData(data string) string{
	data = "ilham"
	return data
}

//pass by refernce
func ubahDataPointer(data *string) string{
	*data = "pak ilham" //mengubah value dari alamat yg ada di variable
	return  *data
}

func main(){
	//contoh1
	var namaData = "ikhsan"
	namaResult := ubahData(namaData)
	fmt.Println("nama: ", namaData)
	fmt.Println("nama result: ", namaResult)

	namaResultPointer := ubahDataPointer(&namaData)
	fmt.Println("\nnama: ", namaData)
	fmt.Println("nama result: ", namaResultPointer)

	// //contoh2
	// var name string  = "isan"
	// var pointerName *string = &name

	// fmt.Println("name address: ", pointerName)
	// fmt.Println("name value", *pointerName)
	// fmt.Println("name address: ", &name)
	// fmt.Println("name value", name)

	// *pointerName = "rahmad" //manipulasi data
	// fmt.Println("\nchange value name: ", *pointerName)
	
	var a = 6
	var b = new(int)
	fmt.Println("\nB: ", *b)
	b = &a
	fmt.Println("B: ", *b)

}