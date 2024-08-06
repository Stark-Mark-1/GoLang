package main
import "fmt"
func add(a,b int)int{
	return a+b
}

//multiple return values

func val()(int, string){
	return 22, "utkarsh"
}
//variadic functions

func sum(nums ...int){
	tot :=0
	for _, num := range nums{
		tot+=num
	}
	fmt.Printf("%d\n", tot)
}

//Closuers (I need it more than Go does lol)
func intseq() func() int{
	i := 0
	return func() int{
		i++
		return i
	}
}


func main(){
	res := add(22, 44)
	fmt.Println("22 + 44 = ", res)

	a,b := val()
	fmt.Println(a)
	fmt.Println(b)

	numbers := []int{1, 2, 3, 4, 1321, 56, 5665}
	sum(numbers...)

	newInt := intseq()

	fmt.Println("Closure Number: ", newInt())
}