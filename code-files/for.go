package main
import (
	"fmt"
)

func main(){
	var i = 1
	for i<= 3{
		fmt.Println(i)
		i= i +1;
		for j:=1; j<4; j++{
			fmt.Println(j)
		}
		for k:=0; k<=6 ; k++{
			if k % 2 ==0{
				continue
			}
			fmt.Println(k)
		}
	}
}