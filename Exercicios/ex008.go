// Calcule o fatorial de um número.

package main

import "fmt"

func main(){
	n:= 5
	fatorial:= 1

	for i:=1; i <=n; i++{
		fatorial*=i


	}
	fmt.Println(fatorial)
}