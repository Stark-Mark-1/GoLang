package main
import "fmt"
func main(){
	var a = "hi! how are you?"	//For declaring variable you only need to use "var" Go will infer the type of variable
	var b = 36
	var c = 33 
	var d = b + c 
	e := 89.0/232.0 // when inside functions := can be used as well to declare a vairable
	fmt.Println(a)
	fmt.Println(d)
	fmt.Println(e)
}