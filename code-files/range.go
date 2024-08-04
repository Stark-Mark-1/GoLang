package main

import (
	"fmt"
)
func main(){

	//Arrays Slice
	nums := []int{2,3,4}
	sum := 0
	for _, num := range nums{
		sum += num
	}
	fmt.Println("sum: ", sum)

	for i, num := range nums{
		if num==3 {
			fmt.Println("index: ", i)
		}
	}

	//maps

	ut := map[string]int{"a": 22, "b": 33, "c": 44}
	
	for k, v := range ut{
		fmt.Printf("%s is key for %d\n", k,v)
	}

	for i,c := range "utkarsh"{
		fmt.Println(i,c)
	}
}