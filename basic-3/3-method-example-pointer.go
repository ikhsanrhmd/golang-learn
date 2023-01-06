package main

import "fmt"

type Employee struct{
	nama string
	salary int
}

func (e *Employee) changeName(newName string, newSalary int){
	(*e).nama = newName
	(*e).salary = newSalary
}

func main(){
	e := Employee{
		nama: "ikhsan",
		salary: 999,
	}
	fmt.Println("e sebelum diganti = ", e)

	ep := &e

	ep.changeName("rahmad", 888)
	fmt.Println("setelah diganti= ",e)
}