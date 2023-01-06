package main
import "fmt"
import "strings"


func main(){
	var secret interface{}
	secret  = "hallo"
	var number = secret.(string) + "hello" //harus sama dengan type var yg dideklarasikan
	fmt.Println(secret," world : ", number)

	secret= []string{"apple", "manggo", "banana"}
	var gruits = strings.Join(secret.([]string), ", ")
	fmt.Println(gruits, "its favorite fruits")

	//map type interface

	var data = map[string]interface{}{}
	data["name"] = "ikhsan r"
	data["age"] = 24
	data["isMale"] = true

	fmt.Println(data)
}