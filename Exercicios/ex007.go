//Calcule a soma dos números de 1 a 100.

package main

import "fmt"

func main(){
	x := 0
	// tem que declarar a variável dentro da função main

	for i:=0; i <= 100; i++{
		x = x+i
		
	}
	fmt.Printf("A soma total foi %d", x)
}