package main
import(
	"fmt"
	"maps"
)
 func main(){
	m := make(map[string]int)

	m["k1"]= 1;
	m["k2"]= 2
	
	fmt.Println("map: ", m)

	v1 := m["k1"]
	v2 := m["k2"]
	fmt.Println("v1: ", v1,"v2: ", v2);

	delete(m, "k2")
	fmt.Println("map: ", m)

	clear(m)
	fmt.Println("map: ", m)

 }