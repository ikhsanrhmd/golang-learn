package main
import "fmt"

var firstname, lastname string

func main()  {
  fmt.Println("enter yourname")
  fmt.Scanln(&firstname, &lastname)
  fmt.Printf("Hi %s and  %s!\n", firstname, lastname)
  
  //long var
  var angka1 int = 2
  //short var
  angka2  := 1
  fmt.Println(angka1)
  fmt.Println(angka2)
}
