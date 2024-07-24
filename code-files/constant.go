package main 
import(
	"fmt"
	"math"
)

	const s string = "Hello Everyone, this is a constant"  // for constant we can declare them same as variable also we cam either explicitly declare the data type or Go can declare it as well

func main(){
	const n = 4646659844  
	const d = 23e20/n
	fmt.Println(d)
	fmt.Println(int64(n))
	fmt.Println(math.Sin(d))
}