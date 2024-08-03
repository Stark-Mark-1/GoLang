package main
import(
	"fmt"
	"slices"
)
func main(){
	var s []int
	fmt.Println(s)

	s = make ([]int,5)
	fmt.Println(s)

	s = append(s, 1, 2 ,3)
	fmt.Println(s)

	subSlice := s[2:5]
	fmt.Println(subSlice)

	arr := [5]int {1,2,3,4,5}
	sliceOfArray := arr[ 1:5 ]
	fmt.Println(sliceOfArray)

	fmt.Println(len(s))

	dest := make([]int, 5)
	dest = {1,2,3,4,5}
	src := make([]int, len(dest))
	copy(dest, src)
	fmt.Println(src)
}