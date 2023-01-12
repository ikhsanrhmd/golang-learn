package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) namaPria() {
	man.Name = "mr. " + man.Name
}

func main() {
	pria := Man{"isan"}
	pria.namaPria()

	fmt.Println(pria.Name)
}
