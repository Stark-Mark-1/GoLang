package main
import(
	"fmt"
)
func main(){
	var a [5]int
	fmt.Println(a)
	
	fmt.Println("adding a value to the array")
	a[4] = 2525
	fmt.Println("A after insetion", a)

	// Now the multiple ways we can implement arrays

	b := [5]int{1,2,3,4,5}
	fmt.Println("b := [5]int{1,2,3,4,5} gives: ", b)

	b = [...]int{1,2,3,4,5}
	fmt.Println("b := [...]int{1,2,3,4,5} gives: ", b)

	b = [...]int{100, 3:400, 500}
	fmt.Println("b:= [...]int{100, 4:500, 600, 700} gives: ", b)

	// Two D Array
	var twoDArray [3][4]int
	
	for i := 0; i<3; i++{
		for j :=0; j<4; j++{
			twoDArray[i][j] =  i+j
			}
		}
	fmt.Println(twoDArray)
}