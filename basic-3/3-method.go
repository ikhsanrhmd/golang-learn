//method merupakan fungsi yg melekat pada type (bisa menjadi struct atau tipe data yg lain)

//deklarasi
// func (receiver structType) MethodName(parameterList) (returnType){
	//block statement
//}

package main
import "fmt"

type Employee struct{
	fristName string
	LastName string
}

//function
func fullName(fristName string, LastName string) (fullName string){
	fullName = fristName + "-" + LastName
	return
}

//method
func (emp Employee) fullName() string{
	return emp.fristName + " " + emp.LastName
}
func main(){
	//function
	user1 := Employee{
		fristName: "ikhsan",
		LastName: "rahmad",
	}
	user2 := Employee{
		fristName: "ilham",
		LastName: "santoso",
	}
	fmt.Println(user1.fristName, user1.LastName)

	//method
	fmt.Println(user1.fullName())
	fmt.Println(user2.fullName())
}