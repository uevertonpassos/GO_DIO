// Imprima os números ímpares de 1 a 15.

package main

import "fmt"

func main(){

	for i:=0; i <=15; i++{
		if i % 2 == 1{
			fmt.Printf("O número %d é Ímpar!\n",i)
		}
	}
}