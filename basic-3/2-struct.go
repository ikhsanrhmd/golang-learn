//struct mirip seperti class mirip seperti OOP
//struct dapat digunakan untuk formating response

package main
import "fmt"

type Person struct {
	fristName string
	LastName string
	Age int
}

//struct didalam struct
type User struct{
	fristName string
	LastName string
	Address Address
}

type Address struct{
	Street string
	City string
}


func main(){
	//log declaration with assigned value each name field
	var person1 = Person{
		fristName: "ikhsan",
		LastName: "rahmad",
		Age: 24,
	}
	fmt.Println(person1)

	//sort declaration
	person2 := Person{"ikhsan", "rahmad", 24}
	fmt.Println(person2)

	address1 := Address{"marapalam", "padang"}
	user1 := User{"\nikhsan", "rahmad", address1}
	fmt.Println(user1)
	
	//get value "padang"
	fmt.Println(user1.Address.City)
   

	//struct on map
	type Mobil struct{
		Type string
		CC int
	}
	var data = map[string]Mobil{}
	data["1"]=Mobil{
		Type: "SUV",
		CC: 2100,
	}
	data["2"]=Mobil{
		Type: "Sport",
		CC: 3000,
	}
	//fmt.Println(data)
    for k, v := range data{
		fmt.Println(k, v)
		fmt.Println(k, v.CC, v.Type)
		
	}
}