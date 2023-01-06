package main
import "fmt"

func isPalindrome(str string) bool{
    for i := 0; i < len(str); i++ {
        j := len(str)-1-i
        if str[i] != str[j] {
            return false   
        }
    }
    return true
}

func main(){
	// name := "ikhsan"
	// var reverseString string = ""
	// for i:=len(name)-1; i>0 ; i--{
	// 	reverseString += string(name[i])
	// 	fmt.Println(i)
	

	// }
	fmt.Println(isPalindrome("kupu-kupu"))
	fmt.Println(isPalindrome("katak"))
	fmt.Println(isPalindrome("civic"))
	
}