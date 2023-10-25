package main

import "fmt"

func main(){
	array:= [5]int{1,2,3,4,5}
	fmt.Println(array)

	slice:=[]int{}
	slice = append(slice, 1, 2, 3, 4, 5,6)
	fmt.Println(slice)
}

